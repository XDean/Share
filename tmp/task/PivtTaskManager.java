package com.asml.brion.pivt.server.service.task;

import com.asml.brion.pivt.server.dao.TaskLogRepository;
import com.asml.brion.pivt.server.data.error.PivtException;
import com.asml.brion.pivt.server.model.task.TaskLogEntity;
import com.asml.brion.pivt.server.model.task.TaskRunInfo;
import io.reactivex.disposables.Disposable;
import io.reactivex.subjects.Subject;
import io.reactivex.subjects.UnicastSubject;
import lombok.Value;
import lombok.experimental.NonFinal;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang.exception.ExceptionUtils;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.scheduling.TaskScheduler;
import org.springframework.scheduling.annotation.SchedulingConfigurer;
import org.springframework.scheduling.config.ScheduledTaskRegistrar;
import org.springframework.stereotype.Service;

import java.util.*;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.stream.Collectors;

@Service
@Slf4j
public class PivtTaskManager implements SchedulingConfigurer, PivtTaskService, DisposableBean {

  @Autowired(required = false) List<PivtTask> tasks = Collections.emptyList();
  @Autowired TaskLogRepository logRepository;
  @Autowired TaskScheduler scheduler;

  List<TaskRunImpl> runs = new ArrayList<>();

  @Override
  public void configureTasks(ScheduledTaskRegistrar taskRegistrar) {
    tasks.forEach(t -> {
      if (!t.cron().isEmpty()) {
        taskRegistrar.addCronTask(() -> {
          TaskRunImpl r = new TaskRunImpl(t, "_system");
          r.run();
        }, t.cron());
      }
    });
  }

  @Override
  public void destroy() {
    runs.forEach(r -> r.stop("_system"));
  }

  @Override
  public List<PivtTask> getAll() {
    return Collections.unmodifiableList(tasks);
  }

  @Override
  public Optional<PivtTask> find(String id) {
    return tasks.stream().filter(t -> t.id().equals(id)).findAny();
  }

  @Override
  public int run(PivtTask task, String who) {
    TaskRunImpl r = new TaskRunImpl(task, who);
    scheduler.schedule(() -> r.run(), new Date());
    return r.runId;
  }

  @Override
  public Optional<Boolean> stop(PivtTask task, int runId, String who) {
    Optional<Boolean> memory = runs.stream()
            .filter(e -> e.getTaskId().equals(task.id()))
            .filter(e -> e.getRunId() == runId)
            .findAny()
            .map(r -> r.stop(who));
    if (memory.isPresent()) {
      return memory;
    } else if (logRepository.existsByTaskIdAndRunId(task.id(), runId)) {
      return Optional.of(false);
    } else {
      return Optional.empty();
    }
  }

  @Override
  public List<TaskRunInfo> getRecentRunInfo(PivtTask task, int limit) {
    return logRepository.findAllByTaskIdAndTypeOrderByIdAsc(
            task.id(), TaskLogEntity.Type.START, limit > 0 ? PageRequest.of(0, limit) : Pageable.unpaged())
            .stream()
            .map(e -> getInfo(task, e.getRunId(), false))
            .filter(Optional::isPresent)
            .map(Optional::get)
            .collect(Collectors.toList());
  }

  @Override
  public Optional<TaskRunInfo> getInfo(PivtTask task, int runId, boolean withLog) {
    List<TaskLogEntity> sysLogs = logRepository.findAllByTaskIdAndRunIdAndTypeInOrderByIdAsc(task.id(), runId, TaskLogEntity.Type.systemTypes());
    int sysSize = sysLogs.size();
    TaskLogEntity startLog;
    TaskLogEntity endLog;
    if (sysSize == 0) {
      return Optional.empty();
    } else if (sysSize == 1) {
      startLog = sysLogs.get(0);
      endLog = null;
    } else if (sysSize == 2) {
      startLog = sysLogs.stream()
              .filter(e -> e.getType() == TaskLogEntity.Type.START)
              .findAny()
              .orElseThrow(() -> PivtException.internal().message("A started task must have type=START log").build());
      endLog = sysLogs.stream()
              .filter(e -> e.getType() != TaskLogEntity.Type.START)
              .findAny()
              .orElseThrow(() -> PivtException.internal().message("A ended task must have type!=START log").build());
    } else {
      throw PivtException.internal().message("Any run must have at most 2 system logs").build();
    }
    return Optional.of(TaskRunInfo.builder()
            .taskId(task.id())
            .runId(runId)
            .startTime(startLog.getTime())
            .stopTime(endLog == null ? null : endLog.getTime())
            .startBy(startLog.getMessage())
            .stopBy(endLog == null ? null : endLog.getMessage())
            .status(endLog == null ? TaskRunInfo.Status.RUNNING : TaskRunInfo.Status.from(endLog.getType()))
            .logs(withLog ?
                    logRepository.findAllByTaskIdAndRunIdAndTypeInOrderByIdAsc(task.id(), runId, TaskLogEntity.Type.userTypes()) :
                    Collections.emptyList())
            .build());
  }

  @Value
  private class TaskRunImpl implements PivtTaskLogger {
    int runId;
    String startTrigger;
    PivtTask task;
    Subject<TaskLogEntity> subject = UnicastSubject.create();
    String taskId;
    Disposable disposable;
    @NonFinal
    AtomicBoolean stopped = new AtomicBoolean(false);

    TaskRunImpl(PivtTask task, String startTrigger) {
      this.task = task;
      this.taskId = task.id();
      this.runId = newTaskRun(taskId, startTrigger);
      this.startTrigger = startTrigger;
      runs.add(this);
      this.disposable = subject
              .doFinally(() -> {
                runs.remove(this);
              })
              .doOnError(e -> stopped.set(true))
              .onErrorReturn(e -> newRow(TaskLogEntity.Type.DONE_ERROR, ExceptionUtils.getStackTrace(e)))
              .buffer(1, TimeUnit.SECONDS, 10)
              .subscribe(rows -> {
                logRepository.saveAll(rows);
              }, error -> {
                logRepository.save(newRow(TaskLogEntity.Type.ERROR,
                        "Server internal error happened: " + ExceptionUtils.getStackTrace(error)));
              }, () -> {
                if (stopped.compareAndSet(false, true)) {
                  logRepository.save(newRow(TaskLogEntity.Type.DONE, ""));
                }
              });
    }

    @Override
    public void error(String message) {
      subject.onNext(newRow(TaskLogEntity.Type.ERROR, message));
    }

    @Override
    public void warn(String message) {
      subject.onNext(newRow(TaskLogEntity.Type.WARN, message));
    }

    @Override
    public void info(String message) {
      subject.onNext(newRow(TaskLogEntity.Type.INFO, message));
    }

    int newTaskRun(String taskId, String trigger) {
      int id = logRepository.findFirstByTaskIdOrderByRunIdDesc(taskId).map(i -> i.getRunId() + 1).orElse(0);
      TaskLogEntity e = TaskLogEntity.builder()
              .taskId(taskId)
              .runId(id)
              .type(TaskLogEntity.Type.START)
              .time(System.currentTimeMillis())
              .message(trigger)
              .build();
      logRepository.save(e);
      return id;
    }

    void run() {
      try {
        this.task.run(this);
        done();
      } catch (Exception e) {
        errorStop(e);
        log.debug(String.format("Task Run Error, %s#%d", taskId, runId), e);
      }
    }

    boolean stop(String who) {
      if (stopped.compareAndSet(false, true)) {
        this.disposable.dispose();
        logRepository.save(newRow(TaskLogEntity.Type.STOP, who));
        return true;
      } else {
        return false;
      }
    }

    void done() {
      if (!subject.hasComplete() && !subject.hasThrowable()) {
        subject.onComplete();
      }
    }

    void errorStop(Exception e) {
      if (!subject.hasComplete() && !subject.hasThrowable()) {
        subject.onError(e);
      }
    }

    TaskLogEntity newRow(TaskLogEntity.Type type, String message) {
      return TaskLogEntity.builder()
              .taskId(taskId)
              .runId(runId)
              .time(System.currentTimeMillis())
              .type(type)
              .message(message)
              .build();
    }
  }
}

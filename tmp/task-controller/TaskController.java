package com.asml.brion.pivt.server.controller.v2.task;

import com.asml.brion.pivt.server.annotation.PreAuthAdmin;
import com.asml.brion.pivt.server.controller.v2.SimpleMessageDTO;
import com.asml.brion.pivt.server.data.error.PivtErrorCode;
import com.asml.brion.pivt.server.model.task.TaskLogEntity;
import com.asml.brion.pivt.server.model.task.TaskRunInfo;
import com.asml.brion.pivt.server.service.task.PivtTask;
import com.asml.brion.pivt.server.service.task.PivtTaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.convert.converter.Converter;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.*;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/pivt/api/v2")
public class TaskController implements Converter<String, PivtTask> {
  private static final PivtErrorCode NO_SUCH_TASK = PivtErrorCode.of(HttpStatus.NOT_FOUND, "NO_SUCH_TASK");
  private static final PivtErrorCode NO_SUCH_TASK_RUN = PivtErrorCode.of(HttpStatus.NOT_FOUND, "NO_SUCH_TASK_RUN");

  @Autowired
  PivtTaskService taskService;

  @PreAuthAdmin
  @GetMapping("/tasks")
  public List<String> getAll() {
    return taskService.getAll().stream().map(e -> e.id()).collect(Collectors.toList());
  }

  @PreAuthAdmin
  @GetMapping("/task/{taskId}")
  public TaskInfoDTO get(@PathVariable("taskId") PivtTask task,
                         @RequestParam(name = "limit", required = false, defaultValue = "10") int limit) {
    return TaskInfoDTO.builder()
            .id(task.id())
            .cron(task.cron())
            .runs(taskService.getRecentRunInfo(task, limit).stream()
                    .map(e -> TaskRunInfoDTO.from(e))
                    .collect(Collectors.toList()))
            .build();
  }

  @PreAuthAdmin
  @PostMapping("/task/{taskId}")
  public TaskRunDTO run(@PathVariable("taskId") PivtTask task, Authentication auth) {
    int id = taskService.run(task, auth.getName());
    return TaskRunDTO.builder().id(id).build();
  }

  @PreAuthAdmin
  @GetMapping("/task/{taskId}/{runId}/stop")
  public ResponseEntity<SimpleMessageDTO> stop(@PathVariable("taskId") PivtTask task,
                                               @PathVariable("runId") int runId,
                                               Authentication auth) {
    Optional<Boolean> stopped = taskService.stop(task, runId, auth.getName());
    if (stopped.isPresent()) {
      if (stopped.get()) {
        return ResponseEntity.ok(SimpleMessageDTO.of("Stop Success"));
      } else {
        return ResponseEntity.badRequest().body(SimpleMessageDTO.of("Task had been stopped"));
      }
    } else {
      throw NO_SUCH_TASK_RUN.error();
    }
  }

  @PreAuthAdmin
  @GetMapping("/task/{taskId}/{runId}")
  public TaskRunInfoDTO getInfo(@PathVariable("taskId") PivtTask task,
                                @PathVariable("runId") int runId,
                                @RequestParam(name = "log", required = false) boolean log) {
    Optional<TaskRunInfo> info = taskService.getInfo(task, runId, log);
    if (info.isPresent()) {
      return TaskRunInfoDTO.from(info.get());
    } else {
      throw NO_SUCH_TASK_RUN.error();
    }
  }

  @PreAuthAdmin
  @GetMapping("/task/{taskId}/{runId}/log")
  public String getLog(@PathVariable("taskId") PivtTask task,
                       @PathVariable("runId") int runId) {
    Optional<TaskRunInfo> infoOpt = taskService.getInfo(task, runId, true);
    if (infoOpt.isPresent()) {
      SimpleDateFormat f = new SimpleDateFormat("yy-MM-dd HH:mm:ss");
      TaskRunInfo info = infoOpt.get();
      List<TaskLogEntity> logs = new ArrayList<>(info.getLogs());
      if (info.getStatus() == TaskRunInfo.Status.ERROR && info.getStopTime() != null) {
        logs.add(TaskLogEntity.builder()
                .time(info.getStopTime())
                .type(TaskLogEntity.Type.ERROR)
                .message(info.getStopBy())
                .build());
      }
      return logs.stream()
              .map(e -> String.format("%s [%5s] %s", f.format(e.getTime()), e.getType().toString(), e.getMessage()))
              .collect(Collectors.joining("\n"));
    } else {
      throw NO_SUCH_TASK_RUN.error();
    }
  }

  @Override
  public PivtTask convert(String source) {
    Optional<PivtTask> task = taskService.find(source);
    if (task.isPresent()) {
      return task.get();
    } else {
      throw NO_SUCH_TASK.error();
    }
  }
}

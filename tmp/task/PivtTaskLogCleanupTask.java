package com.asml.brion.pivt.server.service.task;

import com.asml.brion.pivt.server.configuration.schedule.PivtTaskProperties;
import com.asml.brion.pivt.server.dao.TaskLogRepository;
import com.asml.brion.pivt.server.model.task.TaskLogEntity;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import java.time.Duration;
import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Component
public class PivtTaskLogCleanupTask implements PivtTask {

  @Autowired TaskLogRepository logRepository;
  @Autowired PivtTaskProperties properties;

  @Override
  public String id() {
    return "task-log-cleanup";
  }

  @Override
  public String cron() {
    // clean up per day
    return "0 0 0 * * *";
  }

  @Override
  @Transactional
  public void run(PivtTaskLogger logger) throws Exception {
    Run run = new Run(logger);
    run.cleanupByTime();
    run.cleanupByCount();
  }

  private class Run {
    PivtTaskLogger logger;

    Run(PivtTaskLogger logger) {
      this.logger = logger;
    }

    private void cleanupByTime() {
      Duration maxDuration = properties.getManager().getCleanup().getMaxDuration();
      if (maxDuration == null) {
        return;
      }
      logger.info("Start Cleanup Task Log By Time");
      long liveMillis = maxDuration.getSeconds() * 1000;
      logger.info("Max Duration: " + maxDuration);
      List<TaskLogEntity> startLogs = logRepository.findAllByTimeBeforeAndTypeIs(System.currentTimeMillis() - liveMillis, TaskLogEntity.Type.START);
      logger.info(String.format("Find %d Expired", startLogs.size()));
      for (int i = 0; i < startLogs.size(); i++) {
        TaskLogEntity startLog = startLogs.get(i);
        logger.info(String.format("Delete Logs (%d/%d): %s#%d", i + 1, startLogs.size(), startLog.getTaskId(), startLog.getRunId()));
        logRepository.deleteAllByTaskIdAndRunId(startLog.getTaskId(), startLog.getRunId());
      }
      logger.info("Cleanup By Time Done.");
    }

    private void cleanupByCount() {
      logger.info("Start Cleanup Task Log By Count");
      int max = properties.getManager().getCleanup().getMaxCountPerTask();
      logger.info("Max Count: " + max);
      Map<String, List<TaskLogEntity>> startLogs = logRepository.findAllByTypeIs(TaskLogEntity.Type.START)
              .stream()
              .collect(Collectors.groupingBy(e -> e.getTaskId()));
      startLogs.values().removeIf(e -> e.size() <= max);
      logger.info(String.format("Find %s Tasks Need Cleanup", startLogs.size()));
      int index = 0;
      for (Map.Entry<String, List<TaskLogEntity>> entry : startLogs.entrySet()) {
        String taskId = entry.getKey();
        List<TaskLogEntity> logs = entry.getValue();
        logger.info(String.format("Delete Logs (%d/%d): %s (%d Runs)", index + 1, startLogs.size(), taskId, logs.size() - max));
        logs.stream()
                .sorted(Comparator.comparing(e -> e.getTime()))
                .limit(logs.size() - max)
                .forEach(e -> {
                  logRepository.deleteAllByTaskIdAndRunId(e.getTaskId(), e.getRunId());
                });
      }
      logger.info("Cleanup By Count Done.");
    }
  }
}

package com.asml.brion.pivt.server.model.task;

import lombok.Builder;
import lombok.Value;

import javax.annotation.Nullable;
import java.util.List;

@Value
@Builder
public class TaskRunInfo {
  public enum Status {
    RUNNING, STOP, DONE, ERROR;

    public static Status from(TaskLogEntity.Type type) {
      switch (type) {
        case DONE:
          return Status.DONE;
        case DONE_ERROR:
          return Status.ERROR;
        case STOP:
          return Status.STOP;
        default:
          return Status.RUNNING;
      }
    }
  }

  String taskId;

  int runId;

  long startTime;

  @Nullable
  Long stopTime;

  Status status;

  String startBy;

  @Nullable
  String stopBy;

  List<TaskLogEntity> logs;
}

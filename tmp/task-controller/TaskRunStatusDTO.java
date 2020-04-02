package com.asml.brion.pivt.server.controller.v2.task;

import com.asml.brion.pivt.server.model.task.TaskRunInfo;
import com.fasterxml.jackson.annotation.JsonValue;

public enum TaskRunStatusDTO {
  RUNNING("RUNNING", TaskRunInfo.Status.RUNNING),
  STOP("STOP", TaskRunInfo.Status.STOP),
  DONE("DONE", TaskRunInfo.Status.DONE),
  ERROR("ERROR", TaskRunInfo.Status.ERROR);


  public final String label;
  public final TaskRunInfo.Status value;

  TaskRunStatusDTO(String label, TaskRunInfo.Status value) {
    this.label = label;
    this.value = value;
  }

  @Override
  @JsonValue
  public String toString() {
    return label;
  }

  public static TaskRunStatusDTO fromValue(TaskRunInfo.Status value) {
    for (TaskRunStatusDTO b : TaskRunStatusDTO.values()) {
      if (value == b.value) {
        return b;
      }
    }
    return null;
  }
}

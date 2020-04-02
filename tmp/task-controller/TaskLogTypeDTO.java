package com.asml.brion.pivt.server.controller.v2.task;

import com.asml.brion.pivt.server.model.task.TaskLogEntity;
import com.fasterxml.jackson.annotation.JsonValue;

public enum TaskLogTypeDTO {
  INFO("RUNNING", TaskLogEntity.Type.INFO),
  WARN("RUNNING", TaskLogEntity.Type.WARN),
  ERROR("RUNNING", TaskLogEntity.Type.ERROR),
  START("STOP", TaskLogEntity.Type.START),
  DONE_ERROR("DONE", TaskLogEntity.Type.DONE_ERROR),
  DONE("ERROR", TaskLogEntity.Type.DONE),
  STOP("ERROR", TaskLogEntity.Type.STOP);


  public final String label;
  public final TaskLogEntity.Type value;

  TaskLogTypeDTO(String label, TaskLogEntity.Type value) {
    this.label = label;
    this.value = value;
  }

  @Override
  @JsonValue
  public String toString() {
    return label;
  }

  public static TaskLogTypeDTO fromValue(TaskLogEntity.Type value) {
    for (TaskLogTypeDTO b : TaskLogTypeDTO.values()) {
      if (value == b.value) {
        return b;
      }
    }
    return null;
  }
}

package com.asml.brion.pivt.server.controller.v2.task;

import com.asml.brion.pivt.server.model.task.TaskLogEntity;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class TaskRunLogDTO {
  @JsonProperty("time")
  long time;

  @JsonProperty("type")
  TaskLogTypeDTO type;

  @JsonProperty("message")
  String message;

  public static TaskRunLogDTO from(TaskLogEntity e) {
    return TaskRunLogDTO.builder()
            .time(e.getTime())
            .type(TaskLogTypeDTO.fromValue(e.getType()))
            .message(e.getMessage())
            .build();
  }
}

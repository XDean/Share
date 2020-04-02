package com.asml.brion.pivt.server.controller.v2.task;

import com.asml.brion.pivt.server.model.task.TaskRunInfo;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Data;

import java.util.Date;
import java.util.List;
import java.util.stream.Collectors;

@Data
@Builder
public class TaskRunInfoDTO {

  @JsonProperty("id")
  int runId;

  @JsonProperty("status")
  TaskRunStatusDTO status;

  @JsonProperty("start-time")
  long startTime;

  @JsonProperty("user-start-time")
  @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")
  Date userStartTime;

  @JsonProperty("stop-time")
  Long stopTime;

  @JsonProperty("user-stop-time")
  @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")
  Date userStopTime;

  @JsonProperty("start-by")
  String startBy;

  @JsonProperty("stop-by")
  String stopBy;

  @JsonProperty("logs")
  List<TaskRunLogDTO> logs;

  public static TaskRunInfoDTO from(TaskRunInfo e) {
    return TaskRunInfoDTO.builder()
            .runId(e.getRunId())
            .startTime(e.getStartTime())
            .userStartTime(new Date(e.getStartTime()))
            .stopTime(e.getStopTime())
            .userStopTime(e.getStopTime() == null ? null : new Date(e.getStopTime()))
            .status(TaskRunStatusDTO.fromValue(e.getStatus()))
            .startBy(e.getStartBy())
            .stopBy(e.getStopBy())
            .logs(e.getLogs().isEmpty() ? null : e.getLogs().stream()
                    .map(l -> TaskRunLogDTO.from(l))
                    .collect(Collectors.toList()))
            .build();
  }
}

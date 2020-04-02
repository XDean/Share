package com.asml.brion.pivt.server.controller.v2.task;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Data;

import java.util.List;

@Data
@Builder
public class TaskInfoDTO {
  @JsonProperty("id")
  String id;

  @JsonProperty("cron")
  String cron;

  @JsonProperty("runs")
  List<TaskRunInfoDTO> runs;
}

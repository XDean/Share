package com.asml.brion.pivt.server.controller.v2.task;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class TaskRunDTO {
  @JsonProperty("id")
  int id;
}

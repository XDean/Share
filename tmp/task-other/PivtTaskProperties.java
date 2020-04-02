package com.asml.brion.pivt.server.configuration.schedule;

import lombok.Data;

import java.time.Duration;

@Data
public class PivtTaskProperties {

  Manager manager;

  @Data
  public static class Manager {

    Cleanup cleanup;

    @Data
    public static class Cleanup {
      Duration maxDuration = null;
      int maxCountPerTask = 10;
    }
  }
}

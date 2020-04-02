package com.asml.brion.pivt.server.service.task;

public interface PivtTask {
  String id();

  void run(PivtTaskLogger logger) throws Exception;

  default String cron() {
    return "";
  }
}

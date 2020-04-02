package com.asml.brion.pivt.server.service.task;

public interface PivtTaskLogger {
  void error(String message);

  void warn(String message);

  void info(String message);
}

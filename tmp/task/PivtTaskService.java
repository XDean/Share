package com.asml.brion.pivt.server.service.task;

import com.asml.brion.pivt.server.model.task.TaskRunInfo;

import java.util.List;
import java.util.Optional;

public interface PivtTaskService {
  List<PivtTask> getAll();

  Optional<PivtTask> find(String id);

  int run(PivtTask task, String who);

  Optional<Boolean> stop(PivtTask task, int runId, String who);

  List<TaskRunInfo> getRecentRunInfo(PivtTask task, int limit);

  Optional<TaskRunInfo> getInfo(PivtTask task, int runId, boolean withLog);
}

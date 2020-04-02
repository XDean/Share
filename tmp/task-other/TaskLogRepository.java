package com.asml.brion.pivt.server.dao;

import com.asml.brion.pivt.server.model.task.TaskLogEntity;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Collection;
import java.util.List;
import java.util.Optional;

public interface TaskLogRepository extends JpaRepository<TaskLogEntity, Integer> {
  Optional<TaskLogEntity> findFirstByTaskIdOrderByRunIdDesc(String taskId);

  List<TaskLogEntity> findAllByTaskIdAndTypeOrderByIdAsc(String taskId, TaskLogEntity.Type type, Pageable pageable);

  List<TaskLogEntity> findAllByTaskIdAndRunIdAndTypeInOrderByIdAsc(String taskId, int runId, Collection<TaskLogEntity.Type> type);

  boolean existsByTaskIdAndRunId(String taskId, int runId);

  List<TaskLogEntity> findAllByTimeBeforeAndTypeIs(long time, TaskLogEntity.Type type);

  List<TaskLogEntity> findAllByTypeIs(TaskLogEntity.Type type);

  void deleteAllByTaskIdAndRunId(String taskId, int runId);
}

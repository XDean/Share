package com.asml.brion.pivt.server.configuration.schedule;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.task.TaskSchedulerBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.concurrent.ThreadPoolTaskScheduler;

@Configuration
@EnableScheduling
public class PivtTaskConfig {

  @Bean
  @ConfigurationProperties("pivt.task")
  public PivtTaskProperties pivtScheduleProperties() {
    return new PivtTaskProperties();
  }

  @Bean
  public ThreadPoolTaskScheduler taskScheduler(TaskSchedulerBuilder builder) {
    return builder.build();
  }
}

package xdean.share.spring.async;

import org.springframework.aop.interceptor.AsyncUncaughtExceptionHandler;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Primary;
import org.springframework.core.task.TaskExecutor;
import org.springframework.scheduling.annotation.AsyncConfigurer;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.concurrent.ConcurrentTaskExecutor;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

@SpringBootApplication
@EnableAsync
public class Application {
    public static void main(String[] args) throws Exception {
        ConfigurableApplicationContext ctx = SpringApplication.run(Application.class, args);
        TaskService service = ctx.getBean(TaskService.class);

        service.runVoid(0);
        service.runVoid(1);

        service.runFuture(2);
        service.runFuture(3);

        service.runCompletableFuture(4);
        service.runCompletableFuture(5);

        service.runListenableFuture(6);
        service.runListenableFuture(7);

        service.runQualifier(8);
        service.runQualifier(9);

        service.runError(10);
        service.runError(11);

        Thread.sleep(1000);
        System.exit(0);
    }

    @Bean
    @Primary
    public TaskExecutor executor() {
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        executor.setCorePoolSize(2);
        executor.setMaxPoolSize(2);
        executor.setQueueCapacity(500);
        executor.setThreadNamePrefix("MyPool");
        executor.initialize();
        return executor;
    }

    @Bean
    @Qualifier("io")
    public TaskExecutor iolExecutor() {
        return new ConcurrentTaskExecutor();
    }

    @Bean
    public AsyncConfigurer exceptionHandler() {
        return new AsyncConfigurer() {
            @Override
            public AsyncUncaughtExceptionHandler getAsyncUncaughtExceptionHandler() {
                return (ex, method, params) -> System.out.printf("Error: %s %s\n", method, ex);
            }
        };
    }
}

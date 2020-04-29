package xdean.share.spring.async;

import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.AsyncResult;
import org.springframework.stereotype.Component;
import org.springframework.util.concurrent.ListenableFuture;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Future;

@Component
public class TaskService {

    @Async
    public void runVoid(int i) {
        System.out.printf("%s run-%d\n", Thread.currentThread(), i);
    }

    @Async
    public Future<String> runFuture(int i) {
        System.out.printf("%s runFuture-%d\n", Thread.currentThread(), i);
        return new AsyncResult<>("runFuture-" + i);
    }

    @Async
    public CompletableFuture<String> runCompletableFuture(int i) {
        System.out.printf("%s runCompletableFuture-%d\n", Thread.currentThread(), i);
        return CompletableFuture.completedFuture("runCompletableFuture-" + i);
    }

    @Async
    public ListenableFuture<String> runListenableFuture(int i) {
        System.out.printf("%s runListenableFuture-%d\n", Thread.currentThread(), i);
        return new AsyncResult<>("runListenableFuture-" + i);
    }

    @Async("io")
    public void runQualifier(int i) {
        System.out.printf("%s runQualifier-%d\n", Thread.currentThread(), i);
    }

    @Async
    public void runError(int i) {
        System.out.printf("%s runError-%d\n", Thread.currentThread(), i);
        throw new IllegalArgumentException();
    }
}

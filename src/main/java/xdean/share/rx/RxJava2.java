package xdean.share.rx;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import io.reactivex.Flowable;
import io.reactivex.Observable;
import io.reactivex.rxjavafx.schedulers.JavaFxScheduler;
import io.reactivex.schedulers.Schedulers;
import javafx.application.Platform;

public class RxJava2 {
  public static void main(String[] args) {
    Flowable.range(1, 10)
        .filter(i -> i % 2 == 0)
        .map(i -> i * i)
        .reduce((a, b) -> a + b)
        .subscribe(e -> System.out.println(e));
  }
  
  public static void schedule() {
    new Thread(() -> {
      int value = calc();
      Platform.runLater(() -> {
        // UI work
        System.out.println(value);
      });
    }).start();
    Observable.fromCallable(() -> calc())
        .subscribeOn(Schedulers.computation())
        .observeOn(JavaFxScheduler.platform())
        .subscribe(e -> System.out.println(e));
  }

  public static int calc() {
    uncheck(() -> Thread.sleep(1000));
    return 1;
  }
}

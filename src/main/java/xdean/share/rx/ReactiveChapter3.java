package xdean.share.rx;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import java.util.concurrent.TimeUnit;

import org.junit.Test;

import io.reactivex.Flowable;
import io.reactivex.Observable;
import io.reactivex.rxjavafx.schedulers.JavaFxScheduler;
import io.reactivex.schedulers.Schedulers;
import io.reactivex.subscribers.DisposableSubscriber;
import javafx.application.Platform;
import xdean.jex.util.reflect.ReflectUtil;

public class ReactiveChapter3 {

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

  @Test
  public void testSchedule() throws Exception {
    splitor();
    Flowable.range(1, 3)
        .doOnNext(e -> System.out.println("doOnNext 1: \t" + e + " \t" + Thread.currentThread()))
        .doOnSubscribe(s -> System.out.println("doOnSubscribe 1: \t" + Thread.currentThread()))
        .doOnRequest(e -> System.out.println("doOnRequest 1: \t" + e + " \t" + Thread.currentThread()))
        .subscribeOn(Schedulers.newThread())
        .doOnNext(e -> System.out.println("doOnNext 2: \t" + e + " \t" + Thread.currentThread()))
        .doOnSubscribe(s -> System.out.println("doOnSubscribe 2: \t" + Thread.currentThread()))
        .doOnRequest(e -> System.out.println("doOnRequest 2: \t" + e + " \t" + Thread.currentThread()))
        .observeOn(Schedulers.io(), true, 1)
        .doOnNext(e -> System.out.println("doOnNext 3: \t" + e + " \t" + Thread.currentThread()))
        .doOnSubscribe(s -> System.out.println("doOnSubscribe 3: \t" + Thread.currentThread()))
        .doOnRequest(e -> System.out.println("doOnRequest 3: \t" + e + " \t" + Thread.currentThread()))
        .subscribeOn(Schedulers.computation())
        .doOnNext(e -> System.out.println("doOnNext 4: \t" + e + " \t" + Thread.currentThread()))
        .doOnSubscribe(s -> System.out.println("doOnSubscribe 4: \t" + Thread.currentThread()))
        .doOnRequest(e -> System.out.println("doOnRequest 4: \t" + e + " \t" + Thread.currentThread()))
        .subscribe(new DisposableSubscriber<Integer>() {
          @Override
          public void onStart() {
            request(1);
          };

          @Override
          public void onNext(Integer t) {
            request(1);
          }

          @Override
          public void onError(Throwable t) {

          }

          @Override
          public void onComplete() {

          }
        });
    Thread.sleep(1000);
  }

  @Test
  public void testDelayError() throws Exception {
    splitor(false);
    delayError(false);
    splitor(true);
    delayError(true);
  }

  public void delayError(boolean delay) {
    Flowable.interval(150, TimeUnit.MILLISECONDS)
        .doOnNext(i -> {
          if (i == 3) {
            throw new Exception();
          }
        })
        .doOnNext(e -> System.out.println("before obsereOn: " + e))
        .observeOn(Schedulers.newThread(), delay)
        .doOnNext(e -> System.out.println("after obsereOn: " + e))
        .doOnNext(e -> Thread.sleep(250))
        .blockingSubscribe(
            e -> System.out.println("onNext: " + e),
            e -> System.out.println("onError: " + e),
            () -> System.out.println("onComplete"));
  }

  public void splitor() {
    System.out.println("--------------------------------------------------------");
    System.out.println("Run " + ReflectUtil.getCaller(1, false).getMethodName());
    System.out.println("--------------------------------------------------------");
  }

  public void splitor(Object s) {
    System.out.println("--------------------------------------------------------");
    System.out.println("Run " + ReflectUtil.getCaller(1, false).getMethodName() + "(" + s + ")");
    System.out.println("--------------------------------------------------------");
  }
}

package xdean.share.rx;

import java.util.concurrent.TimeUnit;

import org.junit.Test;

import io.reactivex.Flowable;
import io.reactivex.schedulers.Schedulers;
import xdean.jex.util.reflect.ReflectUtil;

public class ReactiveChapter3 {
  @Test
  public void testSchedule() throws Exception {
    splitor();
    Flowable.range(1, 5)
        .doOnNext(e -> System.out.println("doOnNext 1: " + e + " on" + Thread.currentThread()))
        .doOnSubscribe(s -> System.out.println("doOnSubscribe 1: " + Thread.currentThread()))
        .subscribeOn(Schedulers.newThread())
        .doOnNext(e -> System.out.println("doOnNext 2: " + e + " on" + Thread.currentThread()))
        .doOnSubscribe(s -> System.out.println("doOnSubscribe 2: " + Thread.currentThread()))
        .observeOn(Schedulers.newThread())
        .doOnNext(e -> System.out.println("doOnNext 3: " + e + " on" + Thread.currentThread()))
        .doOnSubscribe(s -> System.out.println("doOnSubscribe 3: " + Thread.currentThread()))
        .blockingSubscribe();
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

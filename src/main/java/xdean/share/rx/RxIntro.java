package xdean.share.rx;

import org.reactivestreams.Publisher;
import org.reactivestreams.Subscriber;
import org.reactivestreams.Subscription;

public class RxIntro {
  public static void main(String[] args) {
    Publisher<Integer> publisher = s -> s.onSubscribe(new Subscription() {
      int i = 0;
      boolean cancel = false;

      @Override
      public void request(long n) {
        while (!cancel && i < 10 && n-- > 0) {
          s.onNext(i++);
        }
        if (i >= 10) {
          s.onComplete();
        }
      }

      @Override
      public void cancel() {
        cancel = true;
      }
    });
    MySubscriber<Integer> s = new MySubscriber<>();
    publisher.subscribe(s);
    s.request(1);
    s.request(2);
    s.request(10);
  }
}

class MySubscriber<T> implements Subscriber<T>, Subscription {
  Subscription s;

  @Override
  public void onSubscribe(Subscription s) {
    this.s = s;
  }

  @Override
  public void onNext(T t) {
    System.out.println("onNext: " + t);
  }

  @Override
  public void onError(Throwable t) {
    System.out.println("oNerror: " + t);
    cancel();
  }

  @Override
  public void onComplete() {
    System.out.println("onComplete");
  }

  @Override
  public void request(long n) {
    System.out.println("request: " + n);
    s.request(n);
  }

  @Override
  public void cancel() {
    System.out.println("cancel");
    s.cancel();
  }
}
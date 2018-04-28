package xdean.share.rx;

import java.util.function.Supplier;

import org.reactivestreams.Publisher;
import org.reactivestreams.Subscriber;
import org.reactivestreams.Subscription;

import com.google.common.base.MoreObjects;

import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.ToString;

public class ReactiveChapter1 {
  public static void main(String[] args) {
    Eater<Chicken> dean = new Eater<>("Dean");
    Publisher<Chicken> kfc = new KFC<>(Chicken::new);
    kfc.subscribe(dean);
    dean.getOrder().request(2);
    dean.getOrder().request(1);
    dean.getOrder().cancel();
    dean.getOrder().request(3);
  }

  @RequiredArgsConstructor
  public static class KFC<T> implements Publisher<T> {
    public final Supplier<T> factory;

    @Override
    public void subscribe(Subscriber<? super T> s) {
      s.onSubscribe(new ChikenOrder(s));
    }

    @RequiredArgsConstructor
    public class ChikenOrder implements Subscription {
      final Subscriber<? super T> subscriber;
      boolean cancel = false;
      int left = 5;

      @Override
      public void request(long n) {
        System.out.printf("Request %d chicken\n", n);
        if (cancel) {
          System.out.println("But the order is finished");
          return;
        }
        while (!cancel && n-- > 0 && left-- > 0) {
          subscriber.onNext(factory.get());
        }
        if (left <= 0) {
          cancel = true;
          subscriber.onComplete();
        }
      }

      @Override
      public void cancel() {
        System.out.println("Cancel chicken order");
        cancel = true;
      }

      @Override
      public String toString() {
        return MoreObjects.toStringHelper(this)
            .add("left", left)
            .add("cancel", cancel)
            .toString();
      }
    }
  }

  @RequiredArgsConstructor
  @Getter
  public static class Eater<T> implements Subscriber<T> {
    final String name;
    Subscription order;

    @Override
    public void onSubscribe(Subscription s) {
      System.out.println(name + " order chickens success: " + s);
      this.order = s;
    }

    @Override
    public void onNext(T t) {
      System.out.printf("%s eat chicken: %s\n", name, t);
    }

    @Override
    public void onError(Throwable t) {
      System.out.println("Error: " + t.getMessage());
    }

    @Override
    public void onComplete() {
      System.out.println("Clean! All chiken eated");
    }
  }

  @ToString
  public static class Chicken {
    private static int counter = 0;
    public final int id = ++counter;
  }
}

package xdean.share.rx;

import org.reactivestreams.Publisher;
import org.reactivestreams.Subscriber;
import org.reactivestreams.Subscription;

import com.google.common.base.MoreObjects;

import lombok.Getter;
import lombok.RequiredArgsConstructor;
import lombok.Value;

public class ReactiveChapter1 {
  public static void main(String[] args) {
    Eater dean = new Eater("Dean");
    Publisher<Chicken> kfc = new KFC();
    kfc.subscribe(dean);
    dean.getOrder().request(2);
    dean.getOrder().request(1);
    dean.getOrder().cancel();
    dean.getOrder().request(3);
  }

  public static class KFC implements Publisher<Chicken> {
    @Override
    public void subscribe(Subscriber<? super Chicken> s) {
      s.onSubscribe(new ChikenOrder(s));
    }

    @RequiredArgsConstructor
    private static class ChikenOrder implements Subscription {
      final Subscriber<? super Chicken> subscriber;
      boolean cancel = false;
      int left = 5;

      @Override
      public void request(long n) {
        System.out.printf("Request %d chicken\n", n);
        if (cancel) {
          System.out.println("But the order is canceled");
          return;
        }
        while (!cancel && n-- > 0 && left-- > 0) {
          subscriber.onNext(new Chicken(5 - left));
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
  public static class Eater implements Subscriber<Chicken> {
    final String name;
    Subscription order;

    @Override
    public void onSubscribe(Subscription s) {
      System.out.println(name + " order chickens success: " + s);
      this.order = s;
    }

    @Override
    public void onNext(Chicken t) {
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

  @Value
  static class Chicken {
    int id;
  }
}

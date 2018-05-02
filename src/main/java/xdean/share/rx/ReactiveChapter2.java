package xdean.share.rx;

import java.util.function.Supplier;

import org.reactivestreams.Publisher;
import org.reactivestreams.Subscriber;
import org.reactivestreams.Subscription;

import lombok.RequiredArgsConstructor;
import lombok.ToString;
import xdean.share.rx.ReactiveChapter1.Eater;
import xdean.share.rx.ReactiveChapter1.KFC;
import xdean.share.rx.ReactiveChapter2.Chicken.Type;

public class ReactiveChapter2 {
  public static void main(String[] args) {
    Eater<Chicken> dean = new Eater<>("Dean");
    KFC2<Chicken> kfc = new KFC2<>(Chicken::new);
    kfc.filter(Type.RIB).subscribe(dean);
    dean.getOrder().request(2);
    dean.getOrder().request(1);
    dean.getOrder().request(3);
  }

  public static class KFC2<T extends Chicken> extends KFC<T> {
    public KFC2(Supplier<T> factory) {
      super(factory);
    }

    public Publisher<Chicken> filter(Type type) {
      return s -> subscribe(new ChickenTypeWaiter(s, type));
    }
  }

  @RequiredArgsConstructor
  public static class ChickenTypeWaiter implements Subscriber<Chicken> {
    final Subscriber<? super Chicken> actual;
    final Type not;
    Subscription s;

    @Override
    public void onSubscribe(Subscription s) {
      actual.onSubscribe(s);
      this.s = s;
    }

    @Override
    public void onNext(Chicken t) {
      if (t.type == not) {
        System.out.println("Drop rib and request one more.");
        s.request(1);
      } else {
        actual.onNext(t);
      }
    }

    @Override
    public void onError(Throwable t) {
      actual.onError(t);
    }

    @Override
    public void onComplete() {
      actual.onComplete();
    }
  }

  @ToString
  public static class Chicken {
    public enum Type {
      LEG, BREAST, RIB;
    }

    private static int counter = 0;
    public final int id = ++counter;
    public final Type type;

    public Chicken() {
      switch (id % 3) {
      case 0:
        type = Type.LEG;
        break;
      case 1:
        type = Type.BREAST;
        break;
      case 2:
        type = Type.RIB;
        break;
      default:
        throw new InstantiationError();
      }
    }
  }
}

package xdean.share.rx;

import io.reactivex.Flowable;

public class RxJava2 {
  public static void main(String[] args) {
    Flowable.range(1, 10)
        .filter(i -> i % 2 == 0)
        .map(i -> i * i)
        .reduce((a, b) -> a + b)
        .subscribe(e -> System.out.println(e));
  }
}

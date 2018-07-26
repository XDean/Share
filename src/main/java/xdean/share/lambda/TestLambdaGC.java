package xdean.share.lambda;

import static xdean.jex.util.lang.ExceptionUtil.uncheck;

import xdean.jex.util.lang.FinalizeSupport;

public class TestLambdaGC {
  int i;

  public static void main(String[] args) throws Exception {
    sayBye(new TestLambdaGC(), 1).sayHelloLaterAnonymous(2000);
    gc();
    sayBye(new TestLambdaGC(), 2).sayHelloLaterLambda(2000);
    gc();
  }

  public void sayHelloLaterAnonymous(int millis) {
    new Thread(new Runnable() {
      @Override
      public void run() {
        uncheck(() -> Thread.sleep(millis));
        System.out.println("Hello");
      }
    }).start();
  }

  public void sayHelloLaterLambda(int millis) {
    new Thread(() -> {
      uncheck(() -> Thread.sleep(millis));
      System.out.println("Hello");
    }).start();
  }

  private static TestLambdaGC sayBye(TestLambdaGC t, int i) {
    t.i = i;
    System.out.println("To say bye " + i);
    FinalizeSupport.finalize(t, () -> System.out.println("Bye: " + i));
    return t;
  }

  static void gc() {
    for (int i = 0; i < 20; i++) {
      System.gc();
      uncheck(() -> Thread.sleep(1));
    }
  }
}

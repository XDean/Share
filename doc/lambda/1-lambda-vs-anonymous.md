# Lambda VS Anonymous

## Anonymous force capture `outer.this`

```
new Test().sayHelloLaterAnonymous(100000);
new Test().sayHelloLaterLambda(100000);

class Test {
  public void sayHelloLaterAnonymous(int millis) {
    new Thread(new Runnable(){
      @Override
      public void run() {
        Thread.sleep(millis);
        System.out.println("Hello");
      }
    }).start();
  }
  
  public void sayHelloLaterLambda(int millis) {
    new Thread(() -> {
      Thread.sleep(millis);
      System.out.println("Hello");
    }).start();
  }
}
```

## Anonymous shadowing fields and methods

```java
public class class ShadowTest {
  public int x = 0;

  public IntConsumer get() {
    int x = 1;
    return new IntConsumer() {
      int x = 2;

      @Override
      public void accept(int x) {
        System.out.println("x = " + x);
        // System.out.println("local.x = " + ?);
        System.out.println("this.x = " + this.x);
        System.out.println("ShadowTest.this.x = " + ShadowTest.this.x);
      }
    };
  }
}

//Lambda expression's parameter x cannot redeclare another local variable defined in an enclosing scope. 
int x = 1;
return x -> System.out.println(x);
```

## Lambda has no `this`

```java
new ChangeListenerEX<T>() {
  @Override
  public void changed(ObservableValue<? extends T> ob, T o, T n) {
    if (shouldRemove()) {
      ob.removeListener(this);
    } else {
      handleEvent(ob, o, n);
    }
  }
};

(ob, o, n) -> {
  if (shouldRemove()) {
    ob.removeListener(this); // here this is outer instance not the lambda itself
  } else {
    handleEvent(ob, o, n);
  }
}
```
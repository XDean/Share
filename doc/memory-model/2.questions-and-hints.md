# 尽可能Immutable

*Do not communicate by sharing memory; instead, share memory by communicating*





# 臭名昭著的DCL

只要是了解过设计模式的同学肯定都知道单例模式(singleton)。而单例模式中最著名的实现非二次检查锁定模式(Double-Check Locking)莫属，下面的代码块就是最早最原始的DCL实现。

```java
class SomeClass {
  private static Resource resource = null;
  public static Resource getResource() {
    if (resource == null) {
      synchronized {
        if (resource == null) {
          resource = new Resource();
        }
      }
    }
    return resource;
  }
}
```

先下一个结论，这一段代码不work。

如果是经常改sonarqube issue的同学可能也见过关于这一问题的issue。

在学习了java内存模型之后，你是否能发现bug所在的呢？我们该如何解决呢？

# This Escape

final有着特殊的语义，总能保持着可见性。但是前提是在正确构造完成后。下面一段代码就演示的final域的错误可见性。

```
public class ThisEscape {
  public static void main(String[] args) {
    new ThisEscape();
  }

  public static void print(ThisEscape t) {
    System.out.println(t.a);
  }

  private final int a;

  public ThisEscape() {
    print(this); // print 0
    this.a = 100;
    print(this); // print 100
  }
}
```

构造过程中应尽量避免传出this引用。
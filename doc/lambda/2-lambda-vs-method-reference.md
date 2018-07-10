# Lambda VS Method Reference

## Method-ref only capture left value

```java
public class Hello {
  String name;
  
  public void say() {
    Thread.sleep(100000);
    System.out.println("Hello " + name);
  }
}

public class Test {
  Hello hello;
  
  public void sayHelloLater(){
    new Thread(() -> hello.say()).start();
    new Thread(hello::say).start;
  }
}
```

## Method-ref has more clear type

```java
list.stream()
  .map
  .flatMap
  .map
  ...
  .reduce((a, b) -> a + b); // which type is this?
  

list.stream()
  .map
  .flatMap
  .map
  ...
  .reduce(Integer::sum);
  .reduce(Long::sum);
  .reduce(String::concat);
```
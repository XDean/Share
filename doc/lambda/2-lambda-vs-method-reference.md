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
  Hello hello = new Hello("a");
  
  public void sayHelloLater(){
    new Thread(() -> hello.say()).start(); // not sure
    new Thread(hello::say).start(); // print "a"
    hello = new Hello("b");
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


| Previous | Next |
| --- | --- |
| [Lambda VS Anonymous](1-lambda-vs-anonymous.md) | [Samples and Hints](3-samples-and-hints.md) |
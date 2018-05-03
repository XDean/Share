# Operator

## What is Operator

继续刚才KFC的例子。
我们都知道吮指原味鸡来自于鸡的不同部位，分别有鸡腿，鸡胸和鸡肋。

```java
class Chicken {
  enum Type {
    LEG, BREAST, RIB;
  }

  int id;
  Type type;
}
```

现在我这个人很挑，不吃鸡肋，该怎么办呢。

最直接的想法，我们当然可以很直接在`KFC`(`Publisher`)内做判断。

```java
Chicken chicken;
do {
  chicken = new Chicken();
} while (chicken.type != Chicken.Type.RIB);
subscriber.onNext(chicken);
```

但是很快我们就发现了问题
- 从现实的角度，KFC从此再也生产鸡肋了，很不科学
- 从代码的角度，我们的改动侵入了`KFC`类，难以拓展

想象一条流水线，KFC是起点，我是终点，我们不应该让KFC再也不生产鸡肋，而是应该在流水线上安排一个工人(KFC服务员)，由服务员来帮我们过滤鸡肋。KFC对服务员负责，服务员对我负责。

我们定义服务员类，服务员把我(`actual`)给保护起来了，要吃鸡先走服务员走一趟。

```java
class ChickenTypeWaiter implements Subscriber<Chicken>{
  Subscriber<? super Chicken> actual;
  Type not;
  
  @Override
  public void onNext(Chicken t) {
    if (t.type == not) {
      s.request(1); //①
    } else {
      actual.onNext(t);
    }
  }
  
  ... // other methods
}

kfc.subscribe(new ChickenTypeWaiter(dean, Type.RIB));
```

注意标记①处，遇到鸡肋的时候我们不能简单的跳过，而是要再`request`一块新的。否则遇到鸡肋就没有鸡吃了。

至此我们就解决了不吃鸡肋的问题，这一实现就是流中经典的`filter`操作符。

## But it's hard to use

很快我们又有了新的麻烦
- 一方面，并不是每个人都能找到`ChickenTypeWaiter`，这理应是`KFC`提供给我的服务而不是我自己去哪儿找来的路人甲
- 另一方面，聚合式API非常难用，如果多几个operator，我们就会陷入无穷无尽的嵌套括号当中

```java
kfc.subscribe(new Op1(new Op2(new Op3(new ChickenTypeWaiter(dean, Type.RIB)))));
```

既然这一功能应该是`KFC`的一部分，何不将其聚合进`KFC`。

```java
public static class KFC2<T extends Chicken> extends KFC<T> {
  public Publisher<Chicken> filter(Type type) {
    return s -> subscribe(new ChickenTypeWaiter(s, type));
  }
}

kfc.filter(Type.RIB).subscribe(dean);
```

这样我们就可以只通过短短的`filter(Type.RIB)`来进行过滤操作。`RxJava`的操作符也都是以这种[流式API](https://en.wikipedia.org/wiki/Fluent_interface)提供的，不仅可以简化代码展示更清晰的结构，也可以方便地提示上下文中的可用操作符。

## RxJava Operators

RxJava为我们提供了非常全面的操作符，大致分为几类

- 变换(Transform)，如`map`, `scan`
- 过滤(Filter)，如`filter`, `debounce`, `distinct`
- 聚合(Aggregate)，如`switch`, `reduce`
- 调度(Schedule)，如`subscribeOn`, `observeOn`

具体的各个操作符用法参考[Reactive 中文文档](https://mcxiaoke.gitbooks.io/rxdocs/content/Operators.html) (部分过时，但对于初学者已经足够)。

## Sample

```java
Flowable.range(1, 10)
    .filter(i -> i % 2 == 0)
    .map(i -> i * i)
    .reduce((a, b) -> a + b)
    .subscribe(e -> System.out.println(e));
```

![Reactive-Model.svg](Reactive-Model.svg)


[Sample Code](/src/main/java/xdean/share/rx/ReactiveChapter2.java)

| Previous | Next |
| --- | --- |
|[Reactive API](1-Reactive-API.md)   | [Scheduler](3-Scheduler.md) |

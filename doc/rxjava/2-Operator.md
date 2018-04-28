# Operator

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
} while (chicken.type != Chicken.Type.BREAST);
subscriber.onNext(chicken);
```

但是很快我们就发现了问题
- 从现实的角度，KFC从此再也生产鸡肋了，很不科学
- 从代码的角度，我们的改动侵入了`KFC`类，难以拓展

想象一条流水线，KFC是起点，我是终点，我们不应该让KFC再也不生产鸡肋，而是应该在流水线上安排一个工人(KFC服务员)，由服务员来帮我们过滤鸡肋。KFC对服务员负责，服务员对我负责。

```java

```

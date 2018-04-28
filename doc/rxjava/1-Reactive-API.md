# Reactive Stream API

`Reactive-streams` 是 `RxJava`的唯一依赖，它是一项响应式编程API标准，已被java标准库收录。

`Reactive-streams`只定义了4个接口共计7个方法，这7个方法构成了整个`Reactive`世界。

*Publisher<T>*

```java
void subscribe(Subscriber<? super T> subscriber)
```

*Subscriber<T>*

```java
void onSubscribe(Subscription s);
void onNext(T t);
void onError(Throwable t);
void onComplete();
```

*Subscription*

```java
void request(long n);
void cancel();
```

由于这些方法都太直观了没什么好说的，我就简单举个例子

- 我(`Subscriber`)向KFC(`Publisher`)订购(`subscribe(Dean)`)了30块吮指原味鸡。
- KFC把订单连同兑换券(`Subscription`)发送(`onSubscribe(30块鸡)`)给我。
- 当我想吃鸡的时候，我就向KFC要一块鸡(`request(1)`)，KFC就会给我一块鸡(`onNext(鸡)`)。
- 如果有一天KFC倒闭了，KFC得通知我(`onError(倒闭Exception)`)。
- 如果我30块鸡吃完了，KFC也得通知我(`onCompelete()`)。
- 如果有一天我不想吃了，我就告诉KFC不要了(`cancel()`)。

Previous | Next
--|--
 | [Where data from](2-Where-data-from.md)
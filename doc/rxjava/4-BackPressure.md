# Back Pressure

1. [Back Pressure](https://github.com/ReactiveX/RxJava/wiki/Backpressure-(2.0))
2. [Observable VS Flowable / Pull VS push](https://github.com/ReactiveX/RxJava/wiki/What's-different-in-2.0#which-type-to-use)

## Pull or Push

我们还是以吮指原味鸡为例，KFC分为后厨和前台。后厨产鸡(producer)，前台卖鸡(consumer)，由不同的人(thread)来负责。

- `Observable`采用的是推模型(push)。不管你前台你卖不卖得过来，后厨都一直生产。
- `Flowable`采用的拉模型(pull)。每当前台要(request)鸡了，后厨就生产一个。

两种策略各有优劣，

- 推模型下，前台根本卖不出去的鸡生产出来，容易产生物理上的浪费，鸡凉了就得倒了。
- 拉模型下，前台要鸡后台不一定生产的过来，容易产生时间上的浪费，前台和顾客就都得等着。

在实践中应该选用合适的模型。

## 拉模型的改进

为了防止顾客久等，KFC在前台和后厨之间放置了一个货架(buffer)，用来暂时存放多于的鸡。这是一种最典型最直接的处理方式来避免人员(thread)等待。事实上`RxJava`就是这么做的。`observeOn`操作符就带有一个`bufferSize`参数可以设置货架的大小。

## 当意外发生

虽然拉模型看上去很美，但是我们做服务业就得把各种意外考虑进去。

- KFC来了一群小孩，一共点了100个鸡
- 前台通知(request)后厨，做100个鸡，后厨开工
- 小孩突然打架了，前台服务员前去劝阻(thread block)
- 后厨鸡陆续产出，然而货架只能放50个鸡，后厨懵了

这就是所谓的背压(Back Pressure)问题，当消费者跟不上生产者的进度，问题就发生了。那么怎么解决背压问题呢，`RxJava`提供了几种方案

- `onBackPressureBuffer`，原理和上面的`observeOn`的buffer一样，在上例中我们只要把货架扩容，问题就暂时解决了。
- `onBackPressureDrop`，把放不下的鸡直接扔掉
- `onBackPressureLatest`，把最早的鸡扔掉
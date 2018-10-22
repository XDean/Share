# 理解Java内存模型

## 从一个问题开始

几天前我在stackoverflow上遇到了一个[有趣的问题](https://stackoverflow.com/questions/51647337/java-volatile-reordering-prevention-scope/51647983#51647983)，我们从这个问题开谈起

@Nik Kotovski:

> 听说`volatile`可以阻止重排序，那么他的作用范围是多大呢？一行，一个大括号还是多少？举例：
> 
>     int i,j,k;
>     volatile int v;
>     boolean flag = true;
>     
>     void someMethod() {
>         int i = 1;
>         if (flag) {
>             j = 2;
>         }
>         if (flag) {
>             k = 3;
>             v = 4;
>         }
>     }
>    
> 可以肯定的是k和v肯定不会重排序，那么i,j,flag和v呢？i,j,k之间呢？

我的回答很长，浓缩成一句即是

> The `volatile` only guarantee the `happens-before` relation.

题主给我的回复是

> Is there any association?

这个问题先按下不表

## 从单线程来

*Note that all variable declaration is `int a = 0;`, and all `r1` like variable is local.*

两行简单的代码

```java
a = 1;
b = 2;
```

我们理所应当地认为`b = 2`在`a = 1`之后执行，因为代码本身是有序的(`program order`)。

当然，大家都知道有那么个玩意叫重排序(`reordering`)。现代CPU为了防止CPU时钟浪费，允许提前执行之后的指令。
同时编译器以及JIT都有权利做重排序以提高执行效率。所以实际运行的时候可能就变成了(`execution order`):

```java
b = 2;
a = 1;
```

问题是谁给编译器勇气乱搞我们的代码的呢？答案是连续一致性`sequentially consistent`

> A set of actions is *sequentially consistent* if all actions occur in a total order (the execution order) 
that is consistent with program order, and furthermore, each read r of a variable v sees the value written 
by the write w to v such that:
> - w comes before r in the execution order, and
> - there is no other write w' such that w comes before w' and w' comes before r in the execution order.

当重排序不破坏连续一致性时，重排序就会被允许。简单的来说就是没有人能观测到重排序的影响，那不就可以想怎样怎样。

## 到多线程去

现在我们加入一个新的线程

```
+----------+----------+
| thread 1 | thread 2 |
+----------+----------+
| r1 = a   | r2 = b   |
+----------+----------+
| b = 2    | a = 1    |
+----------+----------+
```

可以预料到的(r1, r2)值，诸如(0, 0), (1, 0), (0, 2)。
但是我们刚刚说了重排序的问题，我们可以猜到一种结果是(1, 2)，这显然不是我们想看到的。
无知的编译器为了提高单单一个核心的运行效率就自以为是地重排了我们严谨的代码，我们必须给它一些警示来阻止他这么做。

## Synchronization Order

前面提到`program order`，我们的操作在每条线程都应以这一顺序来执行，即使事实上并非如此。
在多线程间，为了组织操作的顺序，Java定义了一系列同步操作(`Synchronization actions`)，
他们在整个程序中拥有确定的顺序(`Synchronization Order`)。

> Synchronization actions, which are:
> - Volatile read. A volatile read of a variable.
> - Volatile write. A volatile write of a variable.
> - Lock. Locking a monitor
> - Unlock. Unlocking a monitor.
> - The (synthetic) first and last action of a thread.
> - Actions that start a thread or detect that a thread has terminated (§17.4.4).

> Synchronization actions induce the synchronized-with relation on actions, defined as follows:
> - An unlock action on monitor m synchronizes-with all subsequent lock actions on m (where "subsequent" is defined according to the synchronization order).
> - A write to a volatile variable v (§8.3.1.4) synchronizes-with all subsequent reads of v by any thread (where "subsequent" is defined according to the synchronization order).
> - An action that starts a thread synchronizes-with the first action in the thread it starts.
> - The write of the default value (zero, false, or null) to each variable synchronizes-with the first action in every thread.
> - The final action in a thread T1 synchronizes-with any action in another thread T2 that detects that T1 has terminated.
> - If thread T1 interrupts thread T2, the interrupt by T1 synchronizes-with any point where any other thread (including T2) determines that T2 has been interrupted (by having an InterruptedException thrown or by invoking Thread.interrupted or Thread.isInterrupted).

## Happens-before

结合线程内的`program order`和线程间的`synchronized order`便组成了最重要的一个概念`happens-before`。

> Two actions can be ordered by a happens-before relationship. 
If one action happens-before another, then the first is visible to and ordered before the second.
> 
> If we have two actions x and y, we write hb(x, y) to indicate that x happens-before y.
> - If x and y are actions of the same thread and x comes before y in program order, then hb(x, y).
> - There is a happens-before edge from the end of a constructor of an object to the start of a finalizer (§12.6) for that object.
> - If an action x synchronizes-with a following action y, then we also have hb(x, y).
> - If hb(x, y) and hb(y, z), then hb(x, z).

`happens-before`其实就是保证了我们常说的有序性和可见性。

### final语义

一个特殊的情况是final句柄，final有着特殊的语义，final在构造完成后总是可见的。

```java
class FinalFieldExample { 
    final int x;
    int y; 
    static FinalFieldExample f;

    public FinalFieldExample() {
        x = 3; 
        y = 4; 
    } 

    static void writer() {
        f = new FinalFieldExample();
    } 

    static void reader() {
        if (f != null) {
            int i = f.x;  // guaranteed to see 3  
            int j = f.y;  // could see 0
        } 
    } 
}
```

## 回到起点

现在我们回到原来的问题，大家是不是都能回答@Nik Kotovski的问题了呢？

另一方面，谈到并发谈到内存模型，常常被提起的几个性质，现在是否有了更深的理解了呢？

- 原子性
- 可见性
- 有序性

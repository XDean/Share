# Machine Learning: Gradient Descent and Back Propagation

## 世界的组成

世界上只有两种东西：数据（值）和过程（函数）

$$Y=f(X)$$

一种常见的情况是，已知$X$和$f$，求$Y$。比如说给定日期返回星期几。

另一种情况是，已知一组$(X_i,Y_i) \in T, Y=f(X)$，给定新的$X'$，求$Y'$。比如给定一张图识别一个数字。

对于第一种情况，我们很容易用数字符号来构造一个过程来计算。
然而对于第二种情形，虽然每个正常的人脑都是这一过程的实现，但是却无法用数字符号准确描述。（尽管脑科学家在努力解构人脑，但是在可预见的未来不会有重大突破）

虽然无法准确描述这个过程，但我们可以退而求其次，找到一个可描述的过程$f' \approx f$

## Temp

$$
\begin{aligned}
A_0 &= (a_{0,0}, a_{0,1}, \cdots , a_{0,n_0})^T \\\\
A_l &= (a_{l,0}, a_{l,1}, \cdots , a_{l,n_l})^T \\\\
&= Sigmoid(NET_l) \\\\
&= Sigmoid( \Omega_l \times A_{l-1} + b_l ) \\\\
a_{l,i} &= Sigmoid(net_{l,i}) \\\\
&= Sigmoid(\sum_{k=0}^{n_{l-1}}\omega_{l,k,i} + b_l) \\\\
E &= \sum_{k=0}^{n_L}(a_{L,k}-t_k)^2 / 2
\end{aligned}
$$

$$
\begin{aligned}
\Delta_{l,i,j}=\frac{\partial E}{\partial \omega_{l,i,j}} &= \frac{\partial E}{\partial a_{l,j}} \frac{\partial a_{l,j}}{\partial net_{l,j}} \frac{\partial net_{l,j}}{\partial \omega_{l,i,j}} \\\\
&= \frac{\partial E}{\partial a_{l,j}} \frac{\partial Sigmoid(a_{l,j})}{\partial a_{l,j}} a_{l-1,i} \\\\
&= \frac{\partial E}{\partial a_{l,j}} a_{l,j} (1 - a_{l,j}) a_{l-1,i} \\\\
{\rm If} \quad l = L \\\\
\frac{\partial E}{\partial a_{L,j}} &= a_{L,j} - t_j \\\\
{\rm If} \quad l \not= L \\\\
\frac{\partial E}{\partial a_{l,j}} &= \frac{\partial E(net_{l+1,0}, net_{l+1,1}, \cdots , net_{l+1,n_{l+1}})}{\partial a_{l,j}} \\\\
&= \sum_{k=0}^{n{l+1}}\frac{\partial E}{\partial net_{l+1,k}}\frac{\partial net_{l+1,k}}{\partial a_{l,j}} \\\\
&= \sum_{k=0}^{n{l+1}}\frac{\partial E}{\partial net_{l+1,k}}\omega_{l,j,k} \\\\
{\rm Let} \quad  \Delta_{l,i,j}=\delta_{l,j}a_{l-1,i} \\\\
\delta_{l,j}&=\begin{cases}
(a_{L,j} - t_j)a_{L,j}(1-a_{L,j}),\quad l=L \\\\
\sum_{k=0}^{n{l+1}}\delta_{l+1,k}\omega_{l,j,k}a_{l,j}(1-a_{l,j}), \quad l \not= L
\end{cases} \\\\
\end{aligned}
$$

## Reference

1. [Principles of training multi-layer neural network using backpropagation](http://galaxy.agh.edu.pl/~vlsi/AI/backp_t_en/backprop.html)
2. [Wiki-Backpropagation](https://en.wikipedia.org/wiki/Backpropagation)
3. [一文弄懂神经网络中的反向传播法——BackPropagation](https://www.cnblogs.com/charlotte77/p/5629865.html)

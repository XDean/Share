# Gradient Descent and Back Propagation

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
\frac{\partial E}{\partial \omega_{l,i,j}} &= \frac{\partial E}{\partial a_{l,j}} \frac{\partial a_{l,j}}{\partial net_{l,j}} \frac{\partial net_{l,j}}{\partial \omega_{l,i,j}} \\\\
&= \frac{\partial E}{\partial a_{l,j}} \frac{\partial Sigmoid(a_{l,j})}{\partial a_{l,j}} a_{l-1,i} \\\\
&= \frac{\partial E}{\partial a_{l,j}} a_{l,j} (1 - a_{l,j}) a_{l-1,i} \\\\
{\rm If} \quad l = L \\\\
\frac{\partial E}{\partial a_{L,j}} &= a_{L,j} - t_j \\\\
{\rm Else} \\\\
\frac{\partial E}{\partial a_{l,j}} &= \frac{\partial E(net_{l+1,0}, net_{l+1,1}, \cdots , net_{l+1,n_{l+1}})}{\partial a_{l,j}} \\\\
&= \sum_{k=0}^{n{l+1}}\frac{\partial E}{\partial net_{l+1,k}}\frac{\partial net_{l+1,k}}{\partial a_{l,j}} \\\\
&= \sum_{k=0}^{n{l+1}}\frac{\partial E}{\partial net_{l+1,k}}\omega_{l,j,k}
\end{aligned}
$$

## Reference

1. [Principles of training multi-layer neural network using backpropagation](http://galaxy.agh.edu.pl/~vlsi/AI/backp_t_en/backprop.html)
2. [Wiki-Backpropagation](https://en.wikipedia.org/wiki/Backpropagation)
3. [一文弄懂神经网络中的反向传播法——BackPropagation](https://www.cnblogs.com/charlotte77/p/5629865.html)

# Gradient Descent and Back Propagation

## Temp


$$
\begin{aligned}
A_0 &= (a_{0,0}, a_{0,1}, \cdots , a_{0,n_0})^T \\\\
A_l &= (a_{l,0}, a_{l,1}, \cdots , a_{l,n_l})^T \\\\
&= Sigmoid( \Omega_l \times A_{l-1} + b_l ) \\\\
a_{l,i} &= Sigmoid(net_{l,i}) \\\\
&= Sigmoid(\sum_{k=0}^{n_{l-1}}\omega_{l,k,i} + b_l) \\\\
E &= \sum_{k=0}^{n_L}(a_{L,k}-t_k)^2 / 2
\end{aligned}
$$

## Reference

1. [Principles of training multi-layer neural network using backpropagation](http://galaxy.agh.edu.pl/~vlsi/AI/backp_t_en/backprop.html)
2. [Wiki-Backpropagation](https://en.wikipedia.org/wiki/Backpropagation)
3. [一文弄懂神经网络中的反向传播法——BackPropagation](https://www.cnblogs.com/charlotte77/p/5629865.html)

$$
% Simulating hand-drawn lines with TikZ
% Author: percusse
\documentclass{article}
\usepackage{tikz}
\usetikzlibrary{calc,decorations.pathmorphing,patterns}
\pgfdeclaredecoration{penciline}{initial}{
    \state{initial}[width=+\pgfdecoratedinputsegmentremainingdistance,
    auto corner on length=1mm,]{
        \pgfpathcurveto%
        {% From
            \pgfqpoint{\pgfdecoratedinputsegmentremainingdistance}
                      {\pgfdecorationsegmentamplitude}
        }
        {%  Control 1
        \pgfmathrand
        \pgfpointadd{\pgfqpoint{\pgfdecoratedinputsegmentremainingdistance}{0pt}}
                    {\pgfqpoint{-\pgfdecorationsegmentaspect
                     \pgfdecoratedinputsegmentremainingdistance}%
                               {\pgfmathresult\pgfdecorationsegmentamplitude}
                    }
        }
        {%TO 
        \pgfpointadd{\pgfpointdecoratedinputsegmentlast}{\pgfpoint{1pt}{1pt}}
        }
    }
    \state{final}{}
}
\begin{document}
\begin{tikzpicture}[decoration=penciline, decorate]
  \draw[decorate,style=help lines] (-2,-2) grid[step=1cm] (4,4);
  \draw[decorate,thick] (0,0) -- (0,3) -- (3,3);
  \draw[decorate,ultra thick,blue] (3,3) arc (0:-90:2cm);
        % supposed to be an arc
  \draw[decorate,thick,pattern=north east lines] (-0.4cm,-0.8cm)
    rectangle (1.2,-2);
  \node[decorate,draw,inner sep=0.5cm,fill=yellow,circle] (a) at (2,0) {};
        % That's not even an ellipse
  \node[decorate,draw,inner sep=0.3cm,fill=red] (b) at (2,-2) {};
  \draw[decorate] (b) to[in=-45,out=45] (a);
        % This was supposed to be an edge
  \node[decorate,draw,minimum height=2cm,minimum width=1cm] (c) at (-1.5,0) {};
  \draw[decorate,->,dashed] (-0.5cm,-0.5cm) -- (-0.5cm,3.5cm)  -| (c.north);
\end{tikzpicture}
\end{document}
$$
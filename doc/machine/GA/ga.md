# Genetic Algorithm

## Introduce

遗传算法（英语：genetic algorithm (GA) ）是计算数学中用于解决最优化的搜索算法，是进化算法的一种。进化算法最初是借鉴了进化生物学中的一些现象而发展起来的，这些现象包括遗传、突变、自然选择以及杂交等。

遗传算法通常实现方式为一种计算机模拟。对于一个最优化问题，一定数量的候选解（称为个体）可抽象表示为染色体，使种群向更好的解进化。传统上，解用二进制表示（即0和1的串），但也可以用其他表示方法。进化从完全随机个体的种群开始，之后一代一代发生。在每一代中评价整个种群的适应度，从当前种群中随机地选择多个个体（基于它们的适应度），通过自然选择和突变产生新的生命种群，该种群在算法的下一次迭代中成为当前种群。

### Problem Domain

全局最优化问题

## Methodology

### Concepts

- 基因(Gene)
- 个体(Individual)
- 种群(Population)
- 适应度(Fitness)
- 选择(Selection)
- 遗传/交叉(Crossover)
- 变异(Mutation)

### Flow

![ga.png](ga.png)

## Samples

### [N Queen Puzzle](https://en.wikipedia.org/wiki/Eight_queens_puzzle)

#### Classic Solution

回溯法(穷举)，时间复杂度**O(n!)**

![8-queens.gif](Eight-queens-animation.gif)

#### Encode


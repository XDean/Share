# Database Normalize

> Database normalization is the process of structuring a relational database in accordance with a series of so-called normal forms in order to reduce data redundancy and improve data integrity. It was first proposed by Edgar F. Codd as an integral part of his relational model.
> Normalization entails organizing the columns (attributes) and tables (relations) of a database to ensure that their dependencies are properly enforced by database integrity constraints. It is accomplished by applying some formal rules either by a process of synthesis (creating a new database design) or decomposition (improving an existing database design).

> 数据库规范化，又称正规化、标准化，是数据库设计的一系列原理和技术，以减少数据库中数据冗余，增进数据的一致性。
> 关系模型的发明者埃德加·科德最早提出这一概念，并于1970年代初定义了第一范式、第二范式和第三范式的概念，还与Raymond F. Boyce于1974年共同定义了第三范式的改进范式——BC范式。
> 除外还包括针对多值依赖的第四范式，连接依赖的第五范式、DK范式和第六范式。

# Concepts

### Superkey 超键

A superkey or super-key is defined in the relational model of database organization as a set of attributes of a relation variable for which it holds that in all relations assigned to that variable, there are no two distinct tuples (rows) that have the same values for the attributes in this set.[1] It can be defined as a set of attributes of a relation schema upon which all attributes of the schema are functionally dependent.

在关系模式中能够唯一标示多元组（即“行”）的属性集

### Candidate Key 候选键

A minimal superkey for that relation; that is, a set of attributes such that:

- the relation does not have two distinct tuples (i.e. rows or records in common database language) with the same values for these attributes (which means that the set of attributes is a superkey)
- there is no proper subset of these attributes for which (1) holds (which means that the set is minimal).

超键且不存在真子集亦为超键

### Functionally Dependent 函数依赖

Given a relation R, a set of attributes X in R is said to functionally determine another set of attributes Y, also in R, (written $X → Y$) if, and only if, each X value in R is associated with precisely one Y value in R; R is then said to satisfy the functional dependency $X → Y$. Equivalently, the projection $\Pi _{X,Y}R$ is a function, i.e. Y is a function of X. In simple words, if the values for the X attributes are known (say they are x), then the values for the Y attributes corresponding to x can be determined by looking them up in any tuple of R containing x. Customarily X is called the determinant set and Y the dependent set. A functional dependency FD: $X → Y$ is called trivial if Y is a subset of X.

在关系R中，X的值可以确定Y的值，称Y函数依赖X，记作 $X \rightarrow Y$

#### Fully Functionally Dependent 完全函数依赖

#### Partial Functionally Dependent 部分函数依赖

#### Delivery Functionally Dependent 传递函数依赖

### Multivalued Dependency 多值依赖

if we denote by $(x,y,z)$ the tuple having values for $\alpha, \beta, R-\alpha -\beta$ collectively equal to $x, y, z$, correspondingly, then whenever the tuples $(a,b,c)$ and $(a,d,e)$ exist in {\displaystyle r} r, the tuples {\displaystyle (a,b,e)} (a,b,e) and $(a,d,c)$ should also exist in $r$.

对于关系$R$的一组分解$(A,B,C)$，对于给定的$a \in A$，存在 $C^* \subseteq C$，对于任意 $b \in B$，值域 $C_(a,b)=C^*$，称C多值依赖A，记作 $X \twoheadrightarrow Y$

### Join Dependency 连接依赖

A table $T$ is subject to a join dependency if $T$ can always be recreated by joining multiple tables each having a subset of the attributes of $T$. If one of the tables in the join has all the attributes of the table $T$, the join dependency is called trivial.

对于关系$R$的一组分解$(A_1, A_2, ... A_n)$，可以通过连接重建关系$R$，称$R$有连接依赖 $*(A_1, A_2, ... A_n)$

# 1NF

## Definition

The domain of each attribute contains only atomic (indivisible) values, and the value of each attribute contains only a single value from that domain.

所有列的值域都是由原子值组成；所有字段的值都只能是单一值。

## Sample

| Name    | Type        | Phone                 |
|-------- |------------ |---------------------  |
| Dean    | CHD GUI     | 1801234567            |
| Ashkan  | CHD Server  | 9962653205,86010123   |
| Jason   | D4C GUI     | 1807654321            |
| David   | D4C Server  | 9962653208,86021230   |

**Defects**

- Find all GUI team members or find all CHD team members
- Dean get a new phone number or Ashkan need to remove a phone number

**With 1NF**

<details>
<summary>Click to show</summary>
<p>
<br>

| Name    | Product   | Type    |
|-------- |---------  |-------- |
| Dean    | CHD       | GUI     |
| Ashkan  | CHD       | Server  |
| Jason   | D4C       | GUI     |
| David   | D4C       | Server  |

<br>

| Name    | Phone       |
|-------- |------------ |
| Dean    | 1801234567  |
| Ashkan  | 9962653205  |
| Ashkan  | 86010123    |
| Jason   | 1807654321  |
| David   | 9962653208  |
| David   | 86021230    |

<br>
</p>
</details>

## Counterexample in CHD

`pfc_aberrations.delta_size_coefficients VARCHAR(512) NULL  Semicolon separated per fitting method, comma separated per Zernike order, space separated coefficients per fitting order`

`pfc_fem_markers.fem TEXT COMMENT '"," separated for each dose and " " separated for each focus'`

# 2NF

## Definition

Any non-prime attribute that is functionally dependent on any proper subset of any candidate key of the relation. A non-prime attribute of a relation is an attribute that is not a part of any candidate key of the relation.

任意非键字段都完全依赖每个候选键

## Sample

| Name      | Location  | Team | Manager   | Type        | ProjectName   |
|---------- |---------- |----- |---------- |-----------  |-------------  |
| WenZhe    | SZ        | PLT  | Liang     | SharkTank   | FindCode      |
| XiaoLong  | SJ        | PM   | George    | SharkTank   | STAR          |
| WenZhe    | SZ        | PLT  | Liang     | ABC         | Golang        |

**Defects**

- Some data duplicate many times
- WenZhe transfer to another team
- If XiaoLong exit the project, after we delete the row. Where is XiaoLong? Who is his manager?

**With 2NF**

<details>
<summary>Click to show</summary>
<p>
<br>

| Name      | Location  | Team | Manager   |
|---------- |---------- |----- |---------- |
| WenZhe    | SZ        | PLT  | Liang     |
| XiaoLong  | SJ        | PM   | George    |

<br>

| Name      | Type        | ProjectName   |
|---------- |-----------  |-------------  |
| WenZhe    | SharkTank   | FindCode      |
| XiaoLong  | SharkTank   | STAR          |
| WenZhe    | ABC         | Golang        |

<br>
</p>
</details>

## Counterexample in CHD

`pfc_fem_markers` has candidate key `{defect_id, target_id}`. But it also have field `detector_id` and `reticle`. While exist functionally dependence `defect_id -> detector_id` and `target_id -> reticle`

# 3NF

## Definition

No non-prime (non-key) attribute is transitively dependent of any key i.e. no non-prime attribute depends on other non-prime attributes. All the non-prime attributes must depend on the primary key only.

不存在非键字段对其他非键字段的依赖

## Sample

| Name      | Location  | Team | Manager   |
|---------- |---------- |----- |---------- |
| WenZhe    | SZ        | PLT  | Liang     |
| ZhengPu   | SZ        | PLT  | Liang     |
| XiaoLong  | SJ        | PM   | George    |

**Defects**

- WenZhe transfer to another team
- PLT transfer to another manager
- Team without member has no manager

**With 3NF**

<details>
<summary>Click to show</summary>
<p>
<br>

| Name      | Location  | Team |
|---------- |---------- |----- |
| WenZhe    | SZ        | PLT  |
| ZhengPu   | SZ        | PLT  |
| XiaoLong  | SJ        | PM   |

<br>

| Location  | Team | Manager   |
|---------- |----- |---------- |
| SZ        | PLT  | Liang     |
| SJ        | PM   | George    |

<br>
</p>
</details>

## Counterexample in CHD

`pfc_defects` had field `x` and `origin_x` while `x = origin_x + offset` i.e exist `origin_x -> x` functionally dependence

# BCNF

## Definition

For every dependency $X \rightarrow Y$, one of the following conditions hold true:
- $X \rightarrow Y$ is a trivial functional dependency (i.e., $Y$ is a subset of $X$)
- $X$ is a superkey for schema R

任意非平凡依赖$X \rightarrow Y$，$X$是超键

# 4NF

## Definition

For every one of its non-trivial multivalued dependencies $X \twoheadrightarrow  Y$, $X$ is a superkey—that is, $X$ is a super key.
 
任意非平凡多值依赖$X \twoheadrightarrow Y$，$X$是超键

## Sample

| Restaurant  |  Food   | Delivery Area   |
|------------ |:------: |--------------:  |
| KFC         | Burger  | W.Trimble       |
| KFC         | Burger  | E.Trimble       |
| KFC         | Rice    | W.Trimble       |
| KFC         | Rice    | E.Trimble       |
| BK          | Burger  | W.Trimble       |

*Suppose that for each area they have same food*

**Defects**

- KFC has new food: dumpling
- KFC has new Delivery Area: N.Trimble

**With 4NF**

<details>
<summary>Click to show</summary>
<p>
<br>

| Restaurant  |  Food   |
|------------ |:------: |
| KFC         | Burger  |
| KFC         | Rice    |
| BK          | Burger  |

<br>

| Restaurant  | Delivery Area   |
|------------ |--------------:  |
| KFC         | W.Trimble       |
| KFC         | E.Trimble       |
| BK          | W.Trimble       |

<br>
</p>
</details>

## Counterexample in CHD

`pfc_hotspots` has hotspots before and after decking. That means each origin defect has n copies.

# 5NF

## Definition

Every non-trivial join dependency in that table is implied by the candidate keys.

任意非平凡连接依赖 $\left\{A_1, A_2, A_3 ... A_n\right\}$, $A_i$是超键

## Sample

| Restaurant  |  Food   | Delivery Area   |
|------------ |:------: |--------------:  |
| KFC         | Burger  | China           |
| KFC         | Burger  | US              |
| KFC         | Rice    | China           |
| BK          | Burger  | US              |

**Defects**

- KFC has new food dumpling that only Chinese eat it
- KFC has new Delivery Area, North Korea, where need Rice but not Burger 

**With 4NF**

<details>
<summary>Click to show</summary>
<p>
<br>

| Restaurant  |  Food   |
|------------ |:------: |
| KFC         | Burger  |
| KFC         | Rice    |
| BK          | Burger  |

<br>

| Restaurant  | Delivery Area   |
|------------ |--------------:  |
| KFC         | China           |
| KFC         | US              |
| BK          | US              |

<br>

|  Food   | Delivery Area   |
|:------: |--------------:  |
| Burger  | China           |
| Burger  | US              |
| Rice    | China           |

<br>
</p>
</details>

# 6NF

## Definition

No nontrivial join dependencies at all.

不存在非平凡连接依赖

# Summary

数据库规范化旨在减少数据库中数据冗余，保证一致性，防止异常操作。通过各级范式可以快速发现数据库设计中的缺陷。

然而也有另一种声音存在，即反正则化。例如有些场景下会放弃伸缩性来获取性能。

但是这并不意味着正则化存在缺陷，反而只有在深入理解了正则化后才能设计出兼顾逻辑与性能的优秀数据库。

1. Un-Normalize 看山是山
2. Normalize 看山不是山
3. De-Normalize (Not-Only-Normalize) 看山还是山(看山不只是山)
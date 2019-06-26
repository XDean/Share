# Database Normalize

> Database normalization is the process of structuring a relational database in accordance with a series of so-called normal forms in order to reduce data redundancy and improve data integrity. It was first proposed by Edgar F. Codd as an integral part of his relational model.
> Normalization entails organizing the columns (attributes) and tables (relations) of a database to ensure that their dependencies are properly enforced by database integrity constraints. It is accomplished by applying some formal rules either by a process of synthesis (creating a new database design) or decomposition (improving an existing database design).

> 数据库规范化，又称正规化、标准化，是数据库设计的一系列原理和技术，以减少数据库中数据冗余，增进数据的一致性。
> 关系模型的发明者埃德加·科德最早提出这一概念，并于1970年代初定义了第一范式、第二范式和第三范式的概念，还与Raymond F. Boyce于1974年共同定义了第三范式的改进范式——BC范式。
> 除外还包括针对多值依赖的第四范式，连接依赖的第五范式、DK范式和第六范式。

# Concepts


# 1NF

## Definition

The domain of each attribute contains only atomic (indivisible) values, and the value of each attribute contains only a single value from that domain.

所有列的值域都是由原子值组成；所有字段的值都只能是单一值。

# 2NF

## Definition

Any non-prime attribute that is functionally dependent on any proper subset of any candidate key of the relation. A non-prime attribute of a relation is an attribute that is not a part of any candidate key of the relation.

所有非键字段都完全依赖每个候选键

# 3NF

## Definition

No non-prime (non-key) attribute is transitively dependent of any key i.e. no non-prime attribute depends on other non-prime attributes. All the non-prime attributes must depend on the primary key only.

不存在非键字段对其他非键字段的依赖

# BCNF

## Definition

For every one of its dependencies \\(X -> Y\\), one of the following conditions hold true:
- \\(X -> Y\\) is a trivial functional dependency (i.e., \\(Y\\) is a subset of \\(X\\))
- \\(X\\) is a superkey for schema R

任意非平凡依赖\\(X -> Y\\)，\\(X\\)是超键

# 4NF

任意非平凡多值依赖\\(X -> Y\\)，\\(X\\)是超键

# 5NF

# 6NF
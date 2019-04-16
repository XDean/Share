# Introduce

[Git Guide](https://git-scm.com/book/en/v2)

[Demo Project](https://git-brion-us.asml.com:8443/users/dxu/repos/git-demo/browse)

[Demo Project For Command](https://git-brion-us.asml.com:8443/users/dxu/repos/git-command-demo/browse)

[Git In Practice](https://git-brion-us.asml.com:8443/users/dxu/repos/git-in-practice/browse)

## Compare with Perforce

![P4Flow](images/P4Flow.png)

![GitFlow](images/GitFlow.png)

![GitTree](images/GitTree.png)

## Why Git?

- 70% developer
- 90% open source project

- distributed
- collaborate, branch model

## Basic Concept

### Repository

- a project
- a `.git` folder

### Commit

- a node in the commit tree
- record the 4 factors of VCS
- has 0 to n parent commit node
- has 0 to n child commit node

![commit-tree](images/commit-tree.png)


### Stage

![stage](https://git-scm.com/book/en/v2/images/areas.png)


### Branch and Merge

- a pointer to the commit tree
- usually is an independent develop process
- `HEAD`, a special pointer that direct to local branch you are currently on.

---
At the beginning, we have `master` branch

![branch-1](images/branch-1.png)

---
Create Branch `feature/GIT-1`

![branch-2](images/branch-2.png)

---
New change in `master` branch

![branch-3](images/branch-3.png)

---
New change in `feature/GIT-1` branch

![branch-4](images/branch-4.png)

---
Merge `feature/GIT-1` into `master`

![branch-5](images/branch-5.png)

### Remote

- A Git repository on server
- Default remote name is `origin` 

![remote](images/remote.png)

<!-- PAGE TABLE START -->

| Previous | Next |
| --- | --- |
| [Back to parent](.) | [Git Basic Commands](2-Basics.md) |

<!-- PAGE TABLE END -->

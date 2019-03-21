# Introduce

[Demo Project](https://git-brion-us.asml.com:8443/users/dxu/repos/git-demo/browse)

## What is Version Control System?

- Which changes were made?
- Who made the changes?
- When were the changes made?
- Why were changes needed?

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

### Branch and Merge

- a pointer to the commit tree
- usually is an independent develop process
- `HEAD`, a special pointer that direct to local branch you are currently on.
<p>
1. ![branch-1](images/branch-1.png)
2. ![branch-2](images/branch-2.png)
3. ![branch-3](images/branch-3.png)
4. ![branch-4](images/branch-4.png)
5. ![branch-5](images/branch-5.png)


### Stage

![stage](https://git-scm.com/book/en/v2/images/areas.png)


| Previous | Next |
| --- | --- |
|   | [Work Flow](2-workflow.md) |
# Introduce

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

### Branch

- a pointer to the commit tree
- usually is an independent develop process
- `HEAD`, a special pointer that direct to local branch you are currently on.

### Merge

![merge](images/merge.png)

- Fast-forward
- Merged
- Conflict
- Fail (rarely, such as not-related history)

### Stage

> ![stage](https://git-scm.com/book/en/v2/images/areas.png)


| Previous | Next |
| --- | --- |
|   | [Work Flow](2-workflow.md) |
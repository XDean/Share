# Git FAQ

## 如何revert一个commit

- 这个commit是普通的commit还是merge commit
    - 是普通的commit，`git revert <commit-id>`
    - 是merge commit， `git revert <commit-id> -m 1`

## 我必须要在网页上创建分支么

不。

Git是分布式的，本地和远程是完全平等的关系。

任何地方创建的分支或者其他操作的效果都是完全相同的。

你可以直接在本地创建分支，然后push到远程。

```
$ git checkout -b new-branch
$ git push -u origin new-branch
```

## 提交的内容有误，想要修改怎么办

- 想要修改的commit是刚刚提交的么？（即HEAD指向的commit）
    - 如果是，`git commit --amend`，详见下文
    - 如果不是，这些改动是否已经merge到master
        - 如果是，
        - 你真的真的需要清除这些历史么①？请自我确认三遍。
            - 如果是*3，使用`git rebase`，详见`rebase`
            - 如果不是，直接用新的commit修改错误的内容

①以下情形可以视为有必要改写历史

- 提交了大文件，如binary，dependencies
- 提交了敏感信息，如银行卡密码，身份证号码

### amend

你只需要正常修改你的文件然后在commit的时候指定为`--amend`。

```
$ vi <your-file>
$ git add <files>
$ git commit --amend
```

请注意，修订后的提交(无论是amend还是rebase)已经不是原来的提交，即他们和之前有不同的ID。

如果此前你已经push到了远程，则远程和本地分支不同步，需要强制push，`git push -f`

## 想要通过`git diff`来查看修改，但是发现没有diff或者是diff少了，怎么回事

- 先使用`git status`看一看你想要看的diff是否在stage上
    - 如果只想要看unstage上的改动，使用`git diff`
    - 如果只想要看stage的改动，使用`git diff --cached`
    - 如果想要both，使用`git diff HEAD`

- `git diff`命令默认拿工作树与`index`比较，所谓`index`就是`HEAD`+`stage`
- 加上`--cached`则变成拿`index`和HEAD比较
- 加上`HEAD`则指定拿工作树与HEAD比较

## `origin`是一个分支么，为什么可以作为`revision`使用(@vhu)

`origin`不是revision而是远程名(remote)。

之所以可以当作`revision`使用是因为当git无法resolve revision的时候，会尝试指向`HEAD`

所以`origin`事实上指向的时`origin/HEAD`。

但是`origin/HEAD`并不是默认生成的，你可以自己指定HEAD指向的位置

```
$ git remote set-head origin main
$ git branch --remote
  origin/HEAD -> origin/main
```

现在你就可以把`origin`当作`revision`使用

```
$ git log origin
$ git show origin
```

## Merge分支后status里有很多文件修改怎么办

说明你的merge有冲突。

这个时候`git status`里会有与你无关的文件出现在stage上。
这些改动是你要merge的分支里的所有改动。

然而stage上的这些改动你不需要你关心，他们已经被自动merge了。
你只需要解决有冲突的文件，然后正常提交即可。

```
$ git merge master
Auto-merging README.md
CONFLICT (content): Merge conflict in README.md
Automatic merge failed; fix conflicts and then commit the result.

$ git status
On branch test
You have unmerged paths.
(fix conflicts and run "git commit")
(use "git merge --abort" to abort the merge)

Unmerged paths:
(use "git add <file>..." to mark resolution)

both modified:   README.md

$ vi README.md
$ git add README.md
$ git commit
```

## 我能不能在master分支commit/push

Git 本身并不会强行禁止你在master上修改，但是Git Server可能会配置branch protect。

- 这个仓库是个人项目还是协作项目
    - 如果是个人项目（如gitea上个人仓库），你可以直接在master上commit/push
    - 如果是协作项目，请不要在master上提交，应当遵循开发流程(创建branch，push，创建PR)
        - 即使你强行push也会被server reject

## 远程分支被动过了导致push/pull失败怎么办 (@ddeng)

- 在开始解决之前，你需要考虑为什么你的远程分支会领先本地分支
    - 如果是因为本地修订(amend或者rebase)，则直接force push: `git push -f`
    - 如果是因为多人协作，建议不要共享分支，合理利用branch

解决的办法就是merge

```
$ git fetch
$ git merge orign/my-branch
$ git push
```

## 如何找到指定文件在区间内的修改 (@fyin)

你可以指定log参数来查询时间区间或者release区间，同时还可以指定文件路径

```
git log --since 2020-01-01 --until 2020-03-01 src
git log release/2020-06-23..release/2020-06-29 src
```

## 如何清理unstaged/untracked文件 (@kwang10)

- 清理unstaged files：`git checkout <path>...`
    - 在清理前请务必确认是否选中的unstaged的文件改动都不再需要，因为他们没有进index，一旦删除无法找回
- 清理untracked files: `git clean -f <path>...`
    - 在清理前请务必确认是否选中的untracked的文件都不再需要，因为他们没有进index，一旦删除无法找回
    - 强烈建议更新`.gitignore`以忽略某些常规的untracked files，而非清理它们

## 不小心`reset --hard`了（或者其他操作）导致文件丢失了，还能还原么

- 该修改是否曾经提交到Git
    - 如果是，你是否有其他引用指向该提交
        - 如果有，你可以通过该引用找到提交
        - 如果否，你可以通过`reflog`来查找
    - 如果否，没救了，Git管不了外部操作
    
## 如何找到某个文件的某一行谁动过 (@vhu)

`git blame --help`

```
$ git blame pom.xml
d59941e90 (Feng Yang - FENY 2019-10-01 15:40:24 -0700   1) <?xml version="1.0" encoding="UTF-8"?>
deb44cb58 (Alex Zhang       2019-11-17 12:03:08 -0800   2) <project xmlns="http://maven.apache.org/POM/4.0.0"
ec264cb2e (zhxie            2020-01-13 14:41:47 +0800   3)          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
ec264cb2e (zhxie            2020-01-13 14:41:47 +0800   4)          xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
ec264cb2e (zhxie            2020-01-13 14:41:47 +0800   5)     <modelVersion>4.0.0</modelVersion>
```

## 如何找到某个commit属于哪个PR(@Jen-Shiang Wang)

你可以用如下命令来查找该提交的所有merge操作，一般而言，第一条merge就是你要找的。

`git log <commit>..<main or release> --ancestry-path --merges --reverse`

```
$  git log SCL-496..main --ancestry-path --merges --reverse
8f69d54 Merge pull request #464 in BRIONOPSPM/pivt from SCL-496_show_resave_button_when_saving_failed to main
cfcee16 Merge pull request #465 in BRIONOPSPM/pivt from SCL-494-expire-time to main
ec506af Merge pull request #466 in BRIONOPSPM/pivt from SCL-496_show_resave_button_when_saving_failed to main
```
# Git概念

## Commit Tree

![](images/commit-tree.png)

## Ref and reflog 引用和引用日志

引用是指向commit tree的指针，一般有一下几种形式

- `HEAD`，即本地当前指针，一般指向当前分支
- `<branch>`，分支
- `<tag>`，标签

以上几项都可以有对应的remote版本，如常见的`origin/master`

reflog是指针的值的变化日志，使用`git reflog <ref>`命令可以查看，默认查看`HEAD`

```
$ git reflog
0fd9a8c7 (HEAD -> SCL-486, tag: release/2020-06-29, origin/SCL-486) HEAD@{0}: commit (amend): SCL-486 server mysql version doesn't support text default value
9151dce2 HEAD@{1}: commit: SCL-486 server mysql version doesn't support text default value
f46714fe (origin/main, origin/HEAD, main) HEAD@{2}: merge remotes/origin/main: Fast-forward
badbe282 HEAD@{3}: checkout: moving from main to SCL-486
f46714fe (origin/main, origin/HEAD, main) HEAD@{4}: pull --no-stat -v --progress origin main: Fast-forward
```

## Revision and Range 版本和版本区间

完整manual: `git help gitrevisions`

| Revision | 用法 | 示例 |
|---|---|--- |
| `<commit id>` | commit的SHA-1值，可以是完整的值，也可以是唯一前缀 | `dae86e1950b1277e545cee180551750029cfe735`, `dae86e` |
| `<ref>`  | 引用名称 | `HEAD`, `origin/master`|
| `<ref>@{<n>}` | 引用本地操作日志，可以使用`git reflog <rev>`查看 | `HEAD@{1}`|
| `<branch>@{upstream}` | Branch对应的远程分支，可以简写为`u` | `master@{u}` |
| `<rev>^<n>` | 版本的第N个父节点，若N=0则指向自己，N可省略默认为1 | `master^`, `HEAD^3` |
| `<rev>~<n>` | 版本的第N代父节点，仅跟随第一个父节点，即`~3=^^^` | `master~3` |

| Range | 用法 | 示例 |
|---|---|--- |
| `<rev>` | 包含改rev及所有祖先 | `HEAD` |
| `^<rev>` | 除外改版本及所有祖先 | `^HEAD~3` |
| `<rev1>..<rev2>` | 包含rev2及所有祖先，但排除rev1及所有祖先，可以理解为差集`rev2 \ rev1` | `HEAD~3..HEAD` |
| `<rev1>...<rev2>` | 包含仅被rev1和rev2及他们的所有祖先，但排除共同祖先，可以理解为差集之并`rev1 \ rev2 ∪ rev2 \ rev1` | `master..feature-branch` |

## Track branch 追踪分支

每一个本地分支可以设定一个Track Branch追踪分支，又叫upstream。

当进行远程操作时，track branch是默认的目标。如`git pull`, `git push`。

你可以用`git branch -vv`来查看追踪分支

```
$ git branch -vv
  bugfix/SCL-473-create-epic                   e9380a5 [origin/bugfix/SCL-473-create-epic] Merge remote-tracking branch 'remotes/origin/main' into bugfix/SCL-473-create-epic
  feature/SCL-450-change-log                   4550f04 SCL-450 entity api
* main                                         28fcd18 [origin/main] Merge pull request #457 in BRIONOPSPM/pivt from SCL-486 to main
  release-06-15                                e9380a5 [origin/release-06-15] Merge remote-tracking branch 'remotes/origin/main' into bugfix/SCL-473-create-epic
```

当进行默认远程操作，如`git push`，而没有配置追踪分支的时候，git会报错。
你可以使用`git branch -u <upstream>`来配置远程分支。
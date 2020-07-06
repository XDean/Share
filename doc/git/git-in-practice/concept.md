# Git概念和解读

## Commit Tree

![](images/commit-tree.png)

## Ref and reflog 引用和引用日志

引用是指向commit的指针，一般有一下几种形式

- `HEAD`，即本地当前指针，一般指向当前分支
- `<branch>`，分支
- `<tag>`，标签

以上几项都可以有对应的remote版本，如常见的`origin/master`

注意`origin/master`和`master`虽然是相对应的分支，但是他们是不同的ref，使用的时候要区分。

例如大部分时候你想要的是`git merge origin/master`而非`git merge master`。

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

## Tracking branch 追踪分支

每一个本地分支可以设定一个Tracking Branch追踪分支，又叫upstream。

当进行远程操作时，tracking branch是默认的目标。如`git pull`, `git push`。

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

## rebase 变基

rebase的作用是将一系列的提交像补丁一样应用到新的位置。
它和merge一样可以用来整合不同分支的修改。

分支原本的样子：
![base](images/base.png)

使用merge后：
![base](images/merge.png)

使用rebase后
![base](images/rebase.png)

### 我该不该用rebase

> Now that you’ve seen rebasing and merging in action, you may be wondering which one is better. Before we can answer this, let’s step back a bit and talk about what history means.
>
> 至此，你已在实战中学习了变基和合并的用法，你一定会想问，到底哪种方式更好。 在回答这个问题之前，让我们退后一步，想讨论一下提交历史到底意味着什么。
>
> One point of view on this is that your repository’s commit history is a record of what actually happened. It’s a historical document, valuable in its own right, and shouldn’t be tampered with. From this angle, changing the commit history is almost blasphemous; you’re lying about what actually transpired. So what if there was a messy series of merge commits? That’s how it happened, and the repository should preserve that for posterity.
>
> 有一种观点认为，仓库的提交历史即是 记录实际发生过什么。 它是针对历史的文档，本身就有价值，不能乱改。 从这个角度看来，改变提交历史是一种亵渎，你使用 谎言 掩盖了实际发生过的事情。 如果由合并产生的提交历史是一团糟怎么办？ 既然事实就是如此，那么这些痕迹就应该被保留下来，让后人能够查阅。
>
> The opposing point of view is that the commit history is the story of how your project was made. You wouldn’t publish the first draft of a book, and the manual for how to maintain your software deserves careful editing. This is the camp that uses tools like rebase and filter-branch to tell the story in the way that’s best for future readers.
>
> 另一种观点则正好相反，他们认为提交历史是 项目过程中发生的事。 没人会出版一本书的第一版草稿，软件维护手册也是需要反复修订才能方便使用。 持这一观点的人会使用 rebase 及 filter-branch 等工具来编写故事，怎么方便后来的读者就怎么写。
>
> 现在，让我们回到之前的问题上来，到底合并还是变基好？希望你能明白，这并没有一个简单的答案。 Git 是一个非常强大的工具，它允许你对提交历史做许多事情，但每个团队、每个项目对此的需求并不相同。 既然你已经分别学习了两者的用法，相信你能够根据实际情况作出明智的选择。
>
> Now, to the question of whether merging or rebasing is better: hopefully you’ll see that it’s not that simple. Git is a powerful tool, and allows you to do many things to and with your history, but every team and every project is different. Now that you know how both of these things work, it’s up to you to decide which one is best for your particular situation.
>
> In general the way to get the best of both worlds is to rebase local changes you’ve made but haven’t shared yet before you push them in order to clean up your story, but never rebase anything you’ve pushed somewhere.
>
> 总的原则是，只对尚未推送或分享给别人的本地修改执行变基操作清理历史， 从不对已推送至别处的提交执行变基操作，这样，你才能享受到两种方式带来的便利。

个人观点：除了reword，否则不rebase
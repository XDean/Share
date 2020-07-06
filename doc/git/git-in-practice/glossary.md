# Glossary 词汇表

`git help glossary`

| 英文 | 中文 | 释义 | 别名 |
|:---:|:---:|:---:|:---:|
| HEAD | 当前指针 | 指向当前分支的指针（除非处于`detached HEAD`）
| ahead | 领先 | 当版本A包含了版本B中不存在的提交，则称之为A领先于B<br>（注意与behind不冲突，A可以同时领先并落后于B）
| ancestor | 祖先节点 | 一个commit的所有parent以及parent的所有祖先
| behind | 落后 | 当版本A包含了版本B中不存在的提交，则称之为B落后于A
| branch | 分支 | 一个动态的指针，当前分支会随着新的commit连同HEAD一起移动
| checkout | 检出 | 将文件更新到目标状态，如果目标是一个revision则同时将HEAD指向它
| cherry-pick | 挑选 | 将一个已经存在的commit重新应用到当前分支
| clean | | 工作树没有未提交的修改内容，即WorkingTree=HEAD
| commit | 提交 | 作为名词，即git历史记录中的一个节点，记录了当时的文件状态和相关信息<br>作为动词，即向git仓库提交新的记录的动作，会创建新的commit并使HEAD指向它
| detached HEAD | 脱离模式 | HEAD指针直接指向commit/tag而非branch
| dirty | | 工作树上包含未提交的修改
| fast-forward | 快进 | 当要merge的分支完全领先于当前分支，则不需要进行merge操作而是直接把当前分支指向目标分支
| fetch | | 从远程仓库获取新的提交信息
| index | 索引 | 一组文件状态
| object name | 对象名 | Git仓库中对象的唯一标识符，采用SHA-1算法 | SHA-1 / hash / identifier
| object | 对象 | Git仓库中的对象（不可变），类型有：commit/tag/blob/tree
| origin | | 默认的远程仓库名
| parent | 父节点 | 一个commit的在commit tree上的parent，如果一个commit有多于1个parent，则他是一个merge commit
| pull | 拉取 | pull是fetch + merge的简写
| push | 推送 | 如果远程分支头是本地分支头的祖先，则推送`remote..local`，否则push失败
| reachable | 可达 | 如果A是B的祖先，则A相对于B可达
| rebase | 变基 | 重新应用一系列commit到一个新的位置
| ref | 引用 | a name that指向object或其他ref(后者又叫symbolic ref,符号引用) | 指针
| reflog | 引用日志 | 引用变化的本地日志
| remote repository | 远程仓库 | 在远程主机上用以追踪同一项目的的repository。用`fetch`和`push`与其通信
| resolve | 解决（冲突） |  手动修复合并操作中无法自动merge的部分
| tag | 标签 | 一类指针，通常指向一个tag object或者commit object，但是它不会跟随HEAD移动
| tracking branch | 追踪分支 | 用以追踪远程仓库变化的ref<br>它是本地分支进行远程操作的默认分支 | upstream branch
| working tree | 工作树 | 磁盘上的文件内容，一般包含`HEAD`和未提交的本地修改

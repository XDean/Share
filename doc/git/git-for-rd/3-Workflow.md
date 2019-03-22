## Create repository

### Brand New Project

- `git init`. Initialize current path to a git repository. It will create a hidden folder `.git` which save git's data.

### Exist Project from Remote

- `git clone https://git-brion-us.asml.com:8443/scm/~dxu/git-demo.git`. Clone a repository from remote url. It will create a folder named by the repository's name (here is `Share`) and download all data from remote repository. And the remote configuration is set up automatically (named `origin`).

`git clone url-to-remote/repo-name.git` is equals to following commands

- `mkdir repo-name`
- `cd repo-name`
- `git init`
- `git remote add origin url`
- `git pull`

## Work with a branch

### Create a branch

- `git branch -b issue-54`. Create a new branch named `issue-54` from `HEAD` and switch to the new branch

### Do your update on your working tree

### Commit the branch

- `git add README.md`. Add your changes onto stage.
- `git commit -m 'update README'`. Commit your stage with message.
- `git push origin issue-54`. Push the commit to remote.

*after commit*

![](https://git-scm.com/book/en/v2/images/small-team-4.png)

### Get up-to-date remote branches

- `git fetch`. Get data from remote.
- `git pull`. Get data from remote and merge automatically.

*after fetch*

![](https://git-scm.com/book/en/v2/images/small-team-5.png)

### Merge with remote branch

- `git checkout master`
- `git merge origin/master`
- `git merge issue-54`

![](https://git-scm.com/book/en/v2/images/small-team-6.png)

- `git push origin master`

![](https://git-scm.com/book/en/v2/images/small-team-7.png)




| Previous | Next |
| --- | --- |
| [Introduce](1-introduce.md) | [Bitbucket and EGit](3-egit-and-bitbucket.md) |
## git的学习和使用
### 简单的git操作
  - #### clone:

    clone是将远程仓库的文件克隆到本地仓库的操作
    具体的用法有: git clone 远程仓库的地址
    例如：git clone git@github/xxx/xxx.git(ssh链接)
  - #### add:
    add是将工作区的文件添加到暂存区的操作
    具体的用法是：git add 文件名
    例如：git add README.md
  - #### commit:

    commit是将暂存区的文件提交到本地仓库的操作
    具体的用法有: git commit -m "提交的信息"
    例如：git commit -m "提交了README.md文件"
  - #### push:

    push是将本地仓库的文件推送到远程仓库的操作
    git push origin master是指将本地仓库的文件推送到远程仓库的master分支上,可以将master替换成其他分支名。
    - 插入一些关于分支的知识：分支在多人进行同一个项目的工作中，可以将不同的工作分配给不同的人，这样就可以提高工作效率。分支的创建和合并都是在本地仓库进行的，无需联网。可以避免多人同时修改同一个文件造成的冲突。
  - #### pull：

    pull是将远程仓库的文件拉取到本地仓库的操作
    具体的用法有: git pull origin master
    当远程仓库的文件发生了改变，本地仓库的文件没有发生改变时，可以使用pull将远程仓库的文件拉取到本地仓库。
  - #### switch:

    switch是切换分支的操作
    具体的用法有: git switch 分支名
    例如：git switch master
  - #### restore:

    restore是撤销工作区的操作
    具体的用法有: git restore 文件名
    例如：git restore README.md
    - restore还可以撤销暂存区的操作
    - 使用时应当注意：restore是撤销工作区的操作，不是撤销本地仓库的操作，如果想要撤销本地仓库的操作，应当使用reset。
  > 工作区：在本地仓库的目录下，可以看到的文件夹和文件都是工作区。
  > 暂存区：在本地仓库的目录下，.git文件夹中的index文件就是暂存区。
  - #### reset:

    reset是撤销本地仓库的操作
    具体的用法有: git reset --hard 版本号
    例如：git reset --hard 1a2b3c4d5e6f7g8h9i0j
    - 版本号可以使用git log命令查看
### 如何回退版本
一共有三种回退模式：
  - #### git reset --mixed 版本号 ：
  此为默认方式，不带任何参数的git reset，即时这种方式，它回退到某个版本，只保留源码，回退commit和index信息

  - #### git reset --soft 版本号：
  回退到某个版本，只回退了commit的信息，不会恢复到index file一级。如果还要提交，直接commit即可本地工作目录内容以及暂存区内容全部回退至某一个版本

  - #### git reset --hard 版本号：
  彻底回退到某个版本，本地的源码也会变为该版本的内容
### 如何查看git日志
  - #### git log:
  git log是查看git日志的命令
  git log --oneline查看简化版的git日志的命令
  git log --oneline --graph是查看简化版的git日志的命令，并且可以查看分支的合并情况
  - #### git reflog:
  git reflog是查看git日志的命令，可以查看所有的git日志，包括回退的版本
### 如何修改commit信息
  - #### git commit --amend:
  git commit --amend是修改commit信息的命令
  git commit --amend --no-edi
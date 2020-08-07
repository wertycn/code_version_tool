# F10-CLI 使用指南
## 1. 安装
复制F10-CLI文件夹到你喜欢的位置
复制文件路径，添加环境变量
```
F10_CLI_HOME = ${YOUR F10-CLI FILE PATH}
PATH 添加 %F10_CLI_HOME%
```
重启电脑，然后在git项目路径范围内即可执行`F10-CLI`命令

## 2. 使用
### 生成版本变更描述文件
```
F10-CLI generate ${版本号} 
```
generate命令将对比当前分支与主分支（默认为master，可通过--branch=指定）差异，提取commit log 生成更新日志

使用示例：
    
    F10-CLI generate 版本(任务)编号

    指定对比分支 默认与master对比
        F10-CLI generate 版本(任务)编号 --branch master --remote origin

    强制重新生成 默认为false
        F10-CLI generate 版本(任务)编号 --overwrite

Usage:
  F10-CLI generate [flags]

Flags:
  -b, --branch string   对比分支 (default "master")
  -h, --help            help for generate
  -o, --overwrite       强制覆盖重复版本
  -r, --remote string   远程地址别名 (default "origin")

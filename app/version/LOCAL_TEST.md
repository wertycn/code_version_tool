---
title: LOCAL_TEST 版本变更信息 2020-08-09 23:37:26
categories:
- code_version_tool
date: 2020-08-09 23:37:26
updated: 2020-08-09 23:37:26
---
# LOCAL_TEST 版本变更信息 2020-08-09 23:37:26

## 一、更新日志

<!--简要描述本次修改影响了哪些功能，如新增XX功能，优化XX功能的XX问题等-->

<!--以下信息提取自 Commit Log -->
1. 新增本地交叉编译脚本 [c9d3698](https://github.com/wertycn/code_version_tool/commit/c9d3698)
2. add linux build file [cbe74cc](https://github.com/wertycn/code_version_tool/commit/cbe74cc)
3. add http service [99aed11](https://github.com/wertycn/code_version_tool/commit/99aed11)
4. 生成编译版本 [8dc21a8](https://github.com/wertycn/code_version_tool/commit/8dc21a8)
5. 增加check命令 [1ee403d](https://github.com/wertycn/code_version_tool/commit/1ee403d)
6. 优化变量替换细节 [82c3fb6](https://github.com/wertycn/code_version_tool/commit/82c3fb6)
7. 完成本地æ命令行调用测试 [6d0125c](https://github.com/wertycn/code_version_tool/commit/6d0125c)
8. 完成日志生ç版本生成功能æ [105a211](https://github.com/wertycn/code_version_tool/commit/105a211)
9. 新增：commit log 生成功能 [7b54645](https://github.com/wertycn/code_version_tool/commit/7b54645)
10. 新增：生成版本文件相关功能 [d3a3224](https://github.com/wertycn/code_version_tool/commit/d3a3224)
11. 完成git log 提信息提取封装 [254f0c0](https://github.com/wertycn/code_version_tool/commit/254f0c0)
12. a [2eb3b02](https://github.com/wertycn/code_version_tool/commit/2eb3b02)


<!--more-->
## 二、变动描述

### 1. 基本信息

* 项目名称：code_version_tool
* 项目地址：https://github.com/wertycn/code_version_tool.git
* 任务编号：LOCAL_TEST
* 开发成员：hajnxg@126.com; 
* 生成时间：2020-08-09 23:37:26

### 2. 代码变更

<!--本次修改影响了哪些代码，由程序自动统计生成-->
```
 F10-CLI                           | Bin 13000164 -> 0 bytes
 F10-CLI.exe                       | Bin 12748288 -> 11885568 bytes
 app/app.go                        |  63 ----------
 app/app_test.go                   |  75 ------------
 app/version.go                    | 244 +++++---------------------------------
 app/version/LOCAL_TEST.md         | 101 ----------------
 build/F10-CLI/F10-CLI             | Bin 12900671 -> 0 bytes
 build/F10-CLI/F10-CLI.exe         | Bin 12646912 -> 0 bytes
 build/F10-CLI/README.md           |  35 ------
 build/F10-CLI/template/version.md |  64 ----------
 ci-check.sh                       |   6 -
 cmd/check.go                      |  60 ----------
 cmd/generate.go                   |  25 +---
 cmd/http.go                       |  43 -------
 go.mod                            |   5 +-
 http/handler.go                   | 155 ------------------------
 http/http.go                      |  88 --------------
 local_build.sh                    |   7 --
 main.go                           |   9 +-
 service/version.go                |  48 --------
 static/a.sh                       |   2 -
 template/version.md               |  64 ----------
 upload/f10-cli/2020-08/F10-336.md |  86 --------------
 version/F10-336.md                |  82 -------------
 version/Test.md                   |   1 -
 version/shell_test.md             |  76 ------------
 version_check.sh                  |   2 -
 27 files changed, 44 insertions(+), 1297 deletions(-)

```

### 3. 数据库变更

<!--需要对数据库做那些修改，将SQL写在下方代码区，没有写则无-->
```SQL
# 服务器地址：
# 数据库名：
# 数据表名:
# 数据库变更SQL：

```

### 4. 其他变更

<!-- 配置文件，环境变量，容器平台配置，数据更新依赖等等其他变更请记录在这里-->


## 三、其他信息

### 1. 接口地址

<!--如果接口已上传到yapi,直接填写yapi地址即可-->
[Yapi-LOCAL_TEST]()

### 2. 任务命令

<!--请输入本次功能所需脚本的执行命令-->
```shell

```


---
title: {{ TASK_NO }} 版本变更信息 {{ DATE }}
categories:
- {{ PROJECT_GIT_NAME }}
date: {{ DATETIME }}
updated: {{ DATETIME }}
---
# {{ TASK_NO }} 版本变更信息 {{ DATE }}

## 一、更新日志

<!--简要描述本次修改影响了哪些功能，如新增XX功能，优化XX功能的XX问题等-->

<!--以下信息提取自 Commit Log -->
{{ CODE_CHANGE_COMMIT_INFO }}

<!--more-->
## 二、变动描述

### 1. 基本信息

* 项目名称：{{ PROJECT_GIT_NAME }}
* 项目地址：{{ PROJECT_GIT_URL }}
* 任务编号：{{ TASK_NO }}
* 开发成员：{{ DEVELOPMENT }}
* 生成时间：{{ DATETIME }}

### 2. 代码变更

<!--本次修改影响了哪些代码，由程序自动统计生成-->
```
{{ CODE_CHANGE_FILE_INFO }}
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
[Yapi-{{ TASK_NO }}]()

### 2. 任务命令

<!--请输入本次功能所需脚本的执行命令-->
```shell

```

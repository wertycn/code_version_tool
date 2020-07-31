# {{ TASK_NO }} {{ DATE }} 更新日志

## 一、任务信息
* 项目名称：{{ PROJECT_GIT_NAME }}
* 项目地址：{{ PROJECT_GIT_URL }}
* 任务编号：{{ TASK_NO }}
* 开发成员：{{ DEVELOPMENT }}
* 生成时间：{{ DATETIME }}

## 二、变更情况

### 1. 功能变更

<!--简要描述本次修改影响了哪些功能，如新增XX功能，优化XX功能的XX问题等-->



<!--more-->
### 2. 代码变更

<!--本次修改影响了哪些代码，由程序自动统计生成-->

* 变动统计

{{ CODE_CHANGE_FILE_COUNT }}

* 文件变动

```

{{ CODE_CHANGE_FILE_INFO }}

```

* COMMIT LOG

{{ CODE_CHANGE_COMMIT_INFO }}

统计方式：
```shell
    # 获取两个分支之间的提交
    git log this_branch...master_branch 
    # 获取master branch 的提交记录commit id 
    git log master_branch --format="%H"
    git log  --format="%H --- %ci --- %ce --- %s"
    # 分支间提交去除 master branch 提交就是本次合并后的提交
    
```



### 3. 数据库变更

<!--本次更新需要对数据库做那些修改，将SQL写在下方代码区，没有写则无-->

```SQL
# 服务器地址：
# 数据库名：
# 数据表名:
# 数据库变更SQL：


```



### 3. 其他变更

<!--配置文件，环境变量，容器平台配置，数据更新依赖等等其他变更请记录在这里-->



## 三、其他

### 1. 接口地址

<!--如果接口已上传到yapi,直接填写yapi地址即可-->

### 2. 任务命令

<!-- 如需要定时任务执行脚本，请输入本次功能所需脚本的执行命令 -->

```shell
# 

```






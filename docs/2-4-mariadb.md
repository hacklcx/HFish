### 数据库选择方案

| 部署情况   | 内网和小规模部署情况 | 外网及大规模部署 |
| ---------- | -------------------- | ---------------- |
| 适用数据库 | Sqlite/Mysql均可     | Mysql            |

`综合来说，Mysql数据库适合于所有部署情况，也是我们优先推荐的数据库，其他数据处理能力和并发兼容上都要优于SQLite`



### SQlite详述

HFish系统默认使用的sqlite数据库，具体见 usr/share/db/hfish.db（自带的已经初始化好的db）



### Sqlite更换为Mysql

HFish当前提供「数据库配置」功能，可快速更换数据库

![image-20211116210129137](http://img.threatbook.cn/hfish/image-20211116210129137.png)

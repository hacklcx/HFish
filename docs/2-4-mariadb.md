### 数据库选择方案

综合来说，MySQL数据库适合于所有部署情况，也是我们优先推荐的数据库，其他数据处理能力和并发兼容上都要优于SQLite。

| 部署情况   | 内网和小规模部署情况 | 外网及大规模部署 |
| ---------- | -------------------- | ---------------- |
| 适用数据库 | SQLite/MySQL均可     | MySQL            |



### SQLite详述

HFish系统默认使用的SQLite数据库，自带的已经初始化好的db具体路径为/usr/share/db/hfish.db



### SQLite更换为MySQL

HFish当前提供「数据库配置」功能，可快速更换数据库

![image-20211116210129137](http://img.threatbook.cn/hfish/image-20211116210129137.png)

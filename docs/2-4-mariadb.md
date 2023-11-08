#### 数据库选择

除非性能极度紧张或环境所限，否则HFish官方**强烈建议使用MySQL/MariaDB数据库！**

经过实战测评，MySQL/MariaDB数据库可以适应目前绝大多数场景，其数据处理和并发兼容能力都要优于SQLite。

> ##### 关于SQLite ##### 

出于开箱即用考虑，HFish系统默认使用的SQLite数据库，自带的已经初始化好的db具体路径为/usr/share/db/hfish.db

SQLite数据库仅适用于功能预览、小规模内网环境失陷感知等有限场景。

> ##### SQLite更换为MySQL/MariaDB数据库  ##### 

HFish提供两种更换数据库的机会：

1、在首次安装时，用户可以选择使用SQLite或MySQL/MariaDB数据库

2、如果已经选择了SQLite，以管理员身份登录后，在「数据库配置」页面，根据指南可快速更换数据库

![image-20211116210129137](https://hfish.net/images/image-20211116210129137.png)

`特别注意：HFish只支持SQLite向MySQL/MariaDB数据库迁移，不支持反向迁移`

- ### 数据库相关

HFish系统默认使用的sqlite数据库，具体见 db/hfish.db（自带的已经初始化好的db），相关的初始化脚本见 db/sql/sqlite/V<% version %>__sqlite.sql 

如果您想要重置 hfish.db， 可以通过下面命令生成新的 db 文件（请确保安装了sqlite3数据库）。 替换 db/hfish.db 即可。

```
sqlite3 hfish.db < db/sql/sqlite/V<% version %>__sqlite.sql
```



**sqlite数据库无需安装，使用方便，但在遭到大规模攻击，及当前版本升级时候会存在数据丢失的问题。**

因此，HFish同时**支持mysql**数据库，相关的初始化脚本见 db/sql/mysql/V<% version %>__mysql.sql。

如果您想要切换到mysql数据库，可以进行以下操作（请确认已经安装了mysql数据库，推荐5.7及以上版本）

> 1. 初始化数据库

linux环境可以在命令行执行下述命令，然后输入密码（root用户密码）。

```
mysql -u root -p < db/sql/mysql/V<% version %>__mysql.sql
```

windows环境可以使用远程连接工具（比如sqlyog等）导入db/sql/mysql/V<% version %>__mysql.sql 脚本。



> 2. 修改config.ini配置文件，数据库的连接方式，主要需要修改type和url，如下：

```
[database]
type = sqlite3
max_open = 50
max_idle = 50
url = ./db/hfish.db?cache=shared&mode=rwc
# type = mysql
# url = root:HFish312@tcp(:3306)/hfish?charset=utf8&parseTime=true&loc=Local
```




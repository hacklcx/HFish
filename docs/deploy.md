# 控制端部署

## Linux安装说明

### 一键安装脚本

如果您部署的环境为Linux，且可以访问互联网。我们为您准备了一键部署脚本进行安装和配置，请用root用户，运行下面的脚本。

```
bash <(curl -sS -L https://hfish.io/install)
```

![image-20210616163833456](http://img.threatbook.cn/hfish/20210616163834.png)

> 安装并运行单机版

这种安装方式，在安装完成后，自动在控制端上创建一个节点，并同时把控制端进程和节点端进程启动。

```wiki
#控制端所在目录
/opt/hfish/

#节点端所在目录
/opt/hfish/client
```

> 安装并运行集群版

这种安装方式，会在安装完成后，启动控制端进程，需要我们后续完成添加节点的操作。

!> 控制端部署完成后，请继续参考下面的【控制端配置】完成配置



### 手动安装

!> 如果上述的安装脚本您无法使用，您可以尝试用手动安装完成部署。

到官网 https://hfish.io   下载HFish最新版本安装包，按如下步骤进行安装 （以linux64位系统为例）：

> 第一步： 在当前目录创建一个路径解压安装包

```
mkdir hfish
```

> 第二步：将安装文件包解压到hfish目录下

```
tar zxvf hfish-*-linux-amd64.tar.gz -C hfish
```

> 第三步：请防火墙开启4433或者4434，确认返回success（如果有其他服务需要打开端口，使用相同命令打开。

```
firewall-cmd --add-port=4433/tcp --permanent
firewall-cmd --add-port=4434/tcp --permanent
firewall-cmd --reload
```

> 第四步：进入安装目录直接运行server，或者后台运行 nohup ./server &

```
cd hfish
nohup ./server &
```

> 第五步：登陆web界面

```
登陆链接：https:// [ip]:4433/web
账号：admin
密码：HFish2021
```

例：如果控制端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web

!> 控制端部署完成后，请继续参考下面的【控制端配置】完成配置



## Windows安装说明

> 第一步：下载HFish

​	访问我们官网的[下载页面](https://hfish.io/#/download)，下载最新版的服务端并解压。

> 第二步：运行文件目录下的server.exe

> 第三步：登陆web界面

```
登陆链接：https:// [ip]:4433/web
账号：admin
密码：HFish2021
```

例：如果控制端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web

!> 控制端部署完成后，请继续参考下面的【控制端配置】完成配置



## Docker安装说明

Docker也是我们推荐的蜜罐交付方式。而且因为容器环境本身就有一层权限隔离的原因，合理配置过的Docker运行环境，能获得更高的业务安全性。

> Docker镜像的下载

```shell
docker pull registry.cn-beijing.aliyuncs.com/threatbook/hfish-amd64
```

> 镜像的运行

```shell
docker run -d -p 4433:4433 -p 4434:4434 --name=hfish --restart=always registry.cn-beijing.aliyuncs.com/threatbook/hfish-amd64
```

!> 控制端部署完成后，请继续参考下面的【控制端配置】完成配置



## 数据库切换MySQL

HFish系统默认使用的sqlite数据库，具体见 db/hfish.db（自带的已经初始化好的db），相关的初始化脚本见 db/sql/sqlite/V2.4.0__sqlite.sql 

如果您想要重置 hfish.db， 可以通过下面命令生成新的 db 文件（请确保安装了sqlite3数据库）。 替换 db/hfish.db 即可。

```
sqlite3 hfish.db < db/sql/sqlite/V2.4.0__sqlite.sql
```



**sqlite数据库无需安装，使用方便，但在遭到大规模攻击，及当前版本升级时候会存在数据丢失的问题。**

因此，HFish同时**支持mysql**数据库，相关的初始化脚本见 db/sql/mysql/V2.4.0__mysql.sql。

如果您想要切换到mysql数据库，可以进行以下操作（请确认已经安装了mysql数据库，推荐5.7及以上版本）

> 1. 初始化数据库

linux环境可以在命令行执行下述命令，然后输入密码（root用户密码）。

```
mysql -u root -p < db/sql/mysql/V2.4.0__mysql.sql
```

windows环境可以使用远程连接工具（比如sqlyog等）导入db/sql/mysql/V2.4.0__mysql.sql 脚本。



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




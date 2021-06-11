# 控制端部署

## linux安装说明

到官网 https://hfish.io   下载HFish最新版本安装包 hfish-2.3-linux-amd64.tar.gz ，按如下步骤进行安装 （以linux64位系统为例）：

> 第一步： 在当前目录创建一个路径解压安装包

```
mkdir hfish
```

> 第二步：将安装文件包解压到hfish目录下

```
tar zxvf hfish-2.3.0-linux-amd64.tar.gz -C hfish
```

> 第三步：请防火墙开启4433或者4434，确认返回success（如果有其他服务需要打开端口，使用相同命令打开。

```
firewall-cmd --add-port=4433/tcp --permanent
firewall-cmd --add-port=4434/tcp --permanent
```

> 第四步：进入安装目录直接运行server，或者后台运行 nohup ./server &

```
cd hfish
./server
```

> 第五步：登陆web界面

```
登陆链接：https:// [ip]:4433/web
账号：admin
密码：HFish2021
```

例：如果您的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web



> 补充服务包

服务端自带的Linux-amd64的服务包，如果您有部署其它系统架构节点的需求，请额外下载下面的服务包，并解压到packages目录下。

```shell
#X86架构32位Linux
http://hfish.cn-bj.ufileos.com/services-linux-386.tar.gz

#ARM架构64位Linux
http://hfish.cn-bj.ufileos.com/services-linux-arm64.tar.gz

#X86架构32位Windows
http://hfish.cn-bj.ufileos.com/services-windows-386.tar.gz

#X86架构64位Windows
http://hfish.cn-bj.ufileos.com/services-windows-amd64.tar.gz
```

```shell
#下载解压示范,请自行替换链接、文件夹路径和包文件名
wget http://hfish.cn-bj.ufileos.com/services-windows-amd64.tar.gz
cd packages
tar -xzvf services-windows-amd64.tar.gz
```



**注意：安装server后不会自动开启服务，需要通过增加节点来运行服务，详情可以看「增加节点」模块。**

## Windows安装说明

> 下载HFish

​	访问我们官网的[下载页面](https://hfish.io/download.html)，下载最新版的服务端并解压。

> 运行文件目录下的server.exe



> 登陆web界面

```
登陆链接：https:// [ip]:4433/web
账号：admin
密码：HFish2021
```

例：如果您的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web



**注意：安装server后不会自动开启服务，需要通过增加节点来运行服务，详情可以看「增加节点」模块。**



## 蜜罐服务说明

> 高交互蜜罐和低交互蜜罐

接触蜜罐产品中我们经常会听到关于高交互蜜罐和低交互蜜罐，两种蜜罐根据与攻击者的交互程度不同而不同。两种类型的蜜罐在使用场景上有所区别。

- 高交互蜜罐

  高交互蜜罐可以跟攻击者进行更多的互动，从而收集更多的攻击者行为数据，能为溯源工作提供更多的参考信息。

  高交互蜜罐也有自己的局限性，高交互蜜罐通常意味着更高的资源占用，和相应更高的业务风险。

  

- 低交互蜜罐

  低交互蜜罐通常跟攻击者之间的交互行为要少一些，不过因为蜜罐产品具有“访问既是风险”的特性，只要能够实现访问就报警，就能解决失陷报警的使用场景。

  低交互蜜罐的局限性跟高交互蜜罐的局限性恰好相反，低交互蜜罐服务对节点资源的占用更少、业务安全性更高。但因为无法获得更多元的攻击信息，对于溯源工作的帮助不如高交互蜜罐高。



高交互蜜罐和低交互蜜罐并无好坏差异，根据自己的业务选择更适合自己的方案即可。



> 目前蜜罐的服务支持4大类，13种不同的蜜罐服务（蜜罐服务还在持续增加），截止本文档发布，OneFish的蜜罐都是低交互蜜罐。

1. 常见的Linux服务

   包括SSH蜜罐、FTP蜜罐、TFTP蜜罐、TELNET蜜罐、VNC蜜罐、HTTP蜜罐

2. 常见Web应用仿真

   包括WordPress仿真登陆蜜罐、通用OA系统仿真登录蜜罐

3. 常见的数据库服务

   MYSQL蜜罐、Redis蜜罐、Memcache蜜罐、Elasticsearch蜜罐

4. 用户自定义蜜罐

## Docker部署



> 作为新兴的虚拟化方式，Docker以优秀的性能和便捷的部署与维护特性，获得大家的认可。

Docker也是我们推荐的蜜罐交付方式。而且因为是虚拟化的原因，合理配置过的Docker网络环境，能获得更高的业务安全性。

> Docker镜像的下载

```shell
docker pull registry.cn-beijing.aliyuncs.com/threatbook/hfish-amd64
```

> 镜像的运行

```shell
docker run -d -p 4433:4433 -p 4434:4434 --name=hfish --restart=always registry.cn-beijing.aliyuncs.com/threatbook/hfish-amd64
```

## 数据库切换MySQL

HFish系统默认使用的sqlite数据库，具体见 db/hfish.db（自带的已经初始化好的db），相关的初始化脚本见 db/sql/sqlite/V2.3.0__sqlite.sql 

如果您想要重置 hfish.db， 可以通过下面命令生成新的 db 文件（请确保安装了sqlite3数据库）。 替换 db/hfish.db 即可。

```
sqlite3 hfish.db < db/sql/V2.4.0__sqlite.sql
```



**sqlite数据库无需安装，使用方便，但在遭到大规模攻击，及当前版本升级时候会存在数据丢失的问题。**

因此，HFis同时**支持mysql**数据库，相关的初始化脚本见 db/sql/mysql/V2.3.0__mysql.sql。

如果您想要切换到mysql数据库，可以进行以下操作（请确认已经安装了mysql数据库，推荐5.7及以上版本）

> 1> 初始化数据库

linux环境可以在命令行执行下述命令，然后输入密码（root用户密码）。

```
mysql -u root -p < db/sql/mysql/V2.3.0__mysql.sql
```

windows环境可以使用远程连接工具（比如sqlyog等）导入db/sql/V2.3.0__mysql.sql 脚本。



> 2> 修改config.ini配置文件，数据库的连接方式，主要需要修改type和url，如下：

```
[database]
type = mysql
url = root:HFish316@tcp(:3306)/hfish?charset=utf8&parseTime=true&loc=Local

# type = sqlite3
# url = ./db/hfish.db?cache=shared&mode=rwc
```

## 一键安装

如果您部署的环境为Linux，且可以访问互联网。我们为您准备了一键部署脚本进行安装和配置，请用root用户，运行下面的脚本。

```
bash <(curl -sS -L https://hfish.io/install)
```

<img src="https://hfish.cn-bj.ufileos.com/images/20210607174007.png" alt="image-20210607154624162" style="zoom: 33%;" />

## 部署运行

下载HFish的控制端压缩包包，二进制文件无依赖，也不需安装，下载后可以直接解压运行。

<!-- tabs:start -->

#### ** Windows **

下载解压得到 `server.exe` 文件，运行

> 用浏览器访问控制台web页面

```wiki
登陆链接：https:// [ip]:4433/web
账号：admin
密码：HFish2021

例：如果您的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web
```

!> 注意：安装server后不会自动开启服务，需要通过增加节点来运行服务，详情可以看「[增加节点](https://hfish.io/docs/#/manage/add)」模块。

#### ** Linux **

Linux系统的部署配置，详见「[Linux主机部署](「https://hfish.io/docs/#/deploy/linux」)」

<!-- tabs:end -->



## 
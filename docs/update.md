

在2.7.0版本中，我们开发了HFish程序作为整个系统的管理进程，帮助您检查HFish管理端的运行情况，整体的升级方案和底层架构做了较多的调整。也因此，我们无法在此版本支持您在页面上的顺滑升级。

不过对于2.5.0和2.6.0的用户，我们提供了升级脚本，支持您的数据顺滑迁移。在到老版本安装目录解压安装包并执行内置的install脚本，即可完成升级与数据迁移。

当然，如果您本身的数据不再需要，我们也推荐你重新安装。

### Linux用户升级

> 1.下载安装包至2.6.2目录中（如果您之前使用脚本安装，那么默认目录为opt/hfish)

[HFish-Linux-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-linux-amd64.tgz) （ Linux x86 架构 64 位系统）

> 2.进入2.6.2安装目录解压安装包

```
tar zxvf hfish-2.7.0-linux-amd64.tar.gz
```



> 3.运行install.sh

```
sudo ./install.sh
```



### Windows用户升级

> 1.下载安装包至2.6.2目录中

[HFish-Windows-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-windows-amd64.tgz) （Windows x86 架构 64 位系统）

> 2.进入2.6.2安装目录解压安装包



> 3.双击运行install.bat





### 升级后，数据迁移失败

./hfish-server --help看一下，当前的hfish-server这个软件我们提供了一个mode参数，一个是重置密码（resetpwd），一个是数据迁移（migrate）

<img src="http://img.threatbook.cn/hfish/image-20211117105331269.png" alt="image-20211117105331269" style="zoom:50%;" />

##### ./server -mode migrate时候，为数据迁移工具，使用填写方法为：

```
. /server -mode migrate -newlink -newtype(填写sqlite或者mysql) -2.7.0 -oldlink -oldtype(旧数据库类型，填写sqlite或者mysql) -oldversion(例如2.5.0/2.6.2等)
```

 按照这个顺序依次填写 新数据库链接、新数据库类型（sqlite/mysql）、新版本号（2.7.0），旧数据库链接、旧数据库类型（sqlite/mysql）、旧版本号，就可以了

##### 填写样例：

```shell
. /server -mode migrate -newlink root:1234567@tcp(127.0.0.1:3306)/hfish?charset=utf8mb4&parseTime=true&loc=Local -newtype mysql -newversion 2.7.0 -oldlink root:1234567@tcp(127.0.0.1:3306)/hfish?charset=utf8mb4&parseTime=true&loc=Local -oldtype mysql -oldversion 2.6.2
```


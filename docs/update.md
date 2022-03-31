### V2.7.0及其以上用户升级

V2.7.0及其以上用户全部支持顺滑升级。

联网情况下，检测到包，「点击右上角火箭」--->「确认升级」。即可完成升级。

非联网情况下，「点击右上角火箭」--->「上传安装包」--->「确认升级」。即可完成升级。

<img src="https://hfish.net/images/image-20220118114818278.png" alt="image-20220118114818278" style="zoom:50%;" />

### V2.5和V2.6版本升级

在2.9.1版本中，我们开发了HFish程序作为整个系统的管理进程，帮助您检查HFish管理端的运行情况，整体的升级方案和底层架构做了较多的调整。也因此，我们无法在此版本支持您在页面上的顺滑升级。

不过对于2.5.0和2.6.0的用户，我们提供了升级脚本，支持您的数据顺滑迁移。在到老版本安装目录解压安装包并执行内置的install.sh脚本，即可完成升级与数据迁移。

当然，如果您不再需要用原有老数据，我们更推荐你重新安装。

#### Linux用户升级

> 1.下载安装包至2.6.2目录中（如果您之前使用脚本安装，那么默认目录为opt/hfish)

[HFish-Linux-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-linux-amd64.tgz) （ Linux x86 架构 64 位系统）

> 2.进入2.6.2安装目录解压安装包

```
tar zxvf hfish-2.9.1-linux-amd64.tgz
```
> 3.运行install.sh

```
sudo ./install.sh
```



#### Windows用户升级

> 1.下载安装包至2.6.2目录中

[HFish-Windows-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-windows-amd64.tgz) （Windows x86 架构 64 位系统）

> 2.进入2.6.2安装目录解压安装包

> 3.双击运行install.bat



#### 数据迁移

执行./hfish-server --help会发现，当前版本提供了resetpwd和migrate两种模式，其中resetpwd模式用于管理员密码重置，migrate模式用于数据迁移，具体如下图：

<img src="https://hfish.net/images/image-20211117105331269.png" alt="image-20211117105331269" style="zoom:50%;" />

##### 当执行./server -mode migrate时，提供数据迁移功能，具体使用方法为：

```
. /hfish-server -mode migrate -newlink -newtype(填写sqlite或者mysql) -2.8.1 -oldlink -oldtype(旧数据库类型，填写sqlite或者mysql) -oldversion(例如2.5.0/2.6.2等)
```

 按照这个顺序依次填写 新数据库链接、新数据库类型（SQLite/MySQL）、新版本号（2.7.0），旧数据库链接、旧数据库类型（SQLite/MySQL）、旧版本号

##### 一个从老MySQL迁移数据到新MySQL的样例：

```
./hfish-server -mode migrate -newLink "root:1234567@tcp(127.0.0.1:3306)/hfish?charset=utf8mb4&parseTime=true&loc=Local" -newType mysql -newVersion 2.8.1 -oldLink "root:1234567@tcp(127.0.0.1:3306)/hfish?charset=utf8mb4&parseTime=true&loc=Local" -oldType mysql -oldVersion 2.6.2
```

##### 一个从老SQLite迁移数据到新SQLite的样例：

```
./hfish-server -mode migrate -newLink ""/Users/Shared/.hfish/database/hfish.db?cache=shared&mode=rwc" -newType sqlite -newVersion 2.8.1 -oldLink ""/Users/Shared/.hfish/database/hfish.db?cache=shared&mode=rwc" -oldType sqlite -oldVersion 2.6.2
```


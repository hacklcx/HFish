### 一键安装脚本

如果您部署的环境为Linux，且可以访问互联网。我们为您准备了一键部署脚本进行安装和配置，请用root用户，运行下面的脚本。

```
bash <(curl -sS -L https://hfish.io/install.sh)
```

[![image-20210616163833456](https://camo.githubusercontent.com/138f103b1cf034b7e493f298b453a43af20628a712f75c80a58d95a3a54b94ee/687474703a2f2f696d672e746872656174626f6f6b2e636e2f68666973682f32303231303631363136333833342e706e67)](https://camo.githubusercontent.com/138f103b1cf034b7e493f298b453a43af20628a712f75c80a58d95a3a54b94ee/687474703a2f2f696d672e746872656174626f6f6b2e636e2f68666973682f32303231303631363136333833342e706e67)

> 安装并运行单机版

这种安装方式，在安装完成后，自动在控制端上创建一个节点，并同时把控制端进程和节点端进程启动。

```
#控制端所在目录
/opt/hfish/

#节点端所在目录
/opt/hfish/client
```

> 安装并运行集群版

这种安装方式，会在安装完成后，启动控制端进程，需要我们后续完成添加节点的操作。

!> 控制端部署完成后，请继续参考下面的【控制端配置】完成配置



### 配置开机自启动

进入管理端（server端）安装目录，执行

```
sh <(curl -sSL https://hfish.io/autorun.sh)
```

即可配置开机自启动。

（该方法通用于节点端自启动）



### 手动安装

如果上述的安装脚本您无法使用，您可以尝试用手动安装完成部署。

到官网 [https://hfish.io](https://hfish.io/) 下载HFish最新版本安装包，按如下步骤进行安装 （以linux64位系统为例）：

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

控制端部署完成后，请继续参考下面的【控制端配置】完成配置


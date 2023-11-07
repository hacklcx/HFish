#### Linux联网环境，一键安装（强烈推荐）

> ##### 特别注意：CentOS 是HFish团队主力开发和测试系统，推荐选用CentOS系统部署管理端 #####

如果部署的环境为Linux，且可以访问互联网，强烈建议使用一键部署脚本进行安装和配置，在使用一键脚本前，请先配置防火墙。

> ##### 以root权限运行以下命令，确保配置防火墙开启TCP/4433、TCP/4434 #####

```
firewall-cmd --add-port=4433/tcp --permanent   #（用于web界面启动）
firewall-cmd --add-port=4434/tcp --permanent   #（用于节点与管理端通信）
firewall-cmd --reload
```

如之后蜜罐服务需要占用其他端口，可使用相同命令打开

> ##### 以root权限运行以下一键部署命令 #####

```
bash <(curl -sS -L https://hfish.net/webinstall.sh)
```

> ##### 完成安装后，通过以下网址、账号密码登录 ##### 

```
登陆链接：https://[ip]:4433/web/
账号：admin
密码：HFish2021
```

`特别注意：
`1、如果管理端的IP是192.168.1.1，则登陆链接为：https://192.168.1.1:4433/web/
`2、/web/这个URL路径不能少

安装完成登录后，在「节点管理」页面中可看到管理端服务器上的默认节点，如下图：

<img src="https://hfish.net/images/image-20210914113134975.png" alt="image-20210914113134975" style="zoom: 25%;" />



#### Linux无法联网，手动安装

如果您的环境无法联网，可以尝试手动安装。

> ##### 第一步：下载安装包：[HFish-Linux-amd64](https://hfish.cn-bj.ufileos.com/hfish-3.3.4-linux-amd64.tgz) （ Linux x86 架构 64 位系统）

按如下步骤进行安装 （以Linux 64位系统为例）：

> ##### 第二步： 在当前目录创建一个路径解压安装包

```
mkdir hfish
```

> ##### 第三步：将安装文件包解压到hfish目录下

```
tar zxvf hfish-3.3.4-linux-amd64.tgz -C hfish
```

> ##### 第四步：请防火墙开启4433、4434和7879，确认返回success（如果有其他服务需要打开端口，使用相同命令打开。

```
firewall-cmd --add-port=4433/tcp --permanent   （用于web界面启动）
firewall-cmd --add-port=4434/tcp --permanent   （用于节点与管理端通信）
firewall-cmd --reload
```

> ##### 第五步：进入安装目录直接运行install.sh

```
cd hfish
sudo ./install.sh
```

> ##### 第六步：登陆web界面

```
登陆链接：https://[ip]:4433/web/
账号：admin
密码：HFish2021
```

如果管理端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web/

在安装完成后，HFish会自动在管理端上创建一个节点。可在节点管理进行查看。

<img src="https://hfish.net/images/image-20210914113134975.png" alt="image-20210914113134975" style="zoom: 25%;" />





特别注意：

```
如果部署节点在外网，可能会出现tcp连接超过最大连接数限制（1024个），导致其他连接被拒绝的情况。在这种情况下，可以手动放开机器tcp最大连接数。

参考解决链接
https://www.cnblogs.com/lemon-flm/p/7975812.html
```


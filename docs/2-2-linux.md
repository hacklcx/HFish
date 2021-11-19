### 联网，使用一键安装脚本

`当前hfish分为两个进程,"hfish"进程为管理进程，负责监测、拉起和升级蜜罐主程序。"server"进程为蜜罐主程序进程，其执行蜜罐软件程序。因此，安装时候，请务必按照要求，执行hfish进程；如果直接执行server程序，可能会导致程序不稳定，升级失败等情况。`

`hfish-Linux的数据库文件及配置文件都当前存储在 /usr/share/hfish 下，在当前机器进行重装时候，会自动读取目录下的配置和数据。`

如果您部署的环境为Linux，且可以访问互联网。我们为您准备了一键部署脚本进行安装和配置。

在使用一键脚本前，请先配置防火墙

> 请防火墙开启4433、4434，确认返回success（如之后蜜罐服务需要占用其他端口，可使用相同命令打开。）

```
firewall-cmd --add-port=4433/tcp --permanent   #（用于web界面启动）
firewall-cmd --add-port=4434/tcp --permanent   #（用于节点与server端通信）
firewall-cmd --reload
```

> 使用root用户，运行下面的脚本。

```
bash <(curl -sS -L https://hfish.io/install.sh)
```

<img src="http://img.threatbook.cn/hfish/image-20210917162839603.png" alt="image-20210917162839603" style="zoom:50%;" />

> 完成安装

```
登陆链接：https://[ip]:4433/web/
账号：admin
密码：HFish2021
```

`例：如果管理端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web/`

在安装完成后，HFish会自动在管理端上创建一个节点。可进行登录后，在「节点管理」列表中进行查看。

<img src="http://img.threatbook.cn/hfish/image-20210914113134975.png" alt="image-20210914113134975" style="zoom: 25%;" />



### 无法联网，手动安装

如果上述的安装脚本您无法使用，您可以尝试用手动安装完成部署。

注意：`当前hfish分为两个进程,"hfish"进程为管理进程，负责监测、拉起和升级蜜罐主程序。"server"进程为蜜罐主程序进程，其执行蜜罐软件程序。因此，安装时候，请务必按照要求，执行hfish进程；如果直接执行server程序，可能会导致程序不稳定，升级失败等情况。`

`hfish-Linux的数据库文件及配置文件都当前存储在 /usr/share/hfish 下，在当前机器进行重装时候，会自动读取目录下的配置和数据。`

> ##### **第一步：下载安装包**：[HFish-Linux-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-linux-amd64.tgz) （ Linux x86 架构 64 位系统）

使用按如下步骤进行安装 （以linux64位系统为例）：

> 第二步： 在当前目录创建一个路径解压安装包

```
mkdir hfish
```

> 第三步：将安装文件包解压到hfish目录下

```
tar zxvf hfish-2.7.0-linux-amd64.tgz -C hfish
```

> 第四步：请防火墙开启4433、4434和7879，确认返回success（如果有其他服务需要打开端口，使用相同命令打开。

```
firewall-cmd --add-port=4433/tcp --permanent   （用于web界面启动）
firewall-cmd --add-port=4434/tcp --permanent   （用于节点与server端通信）
firewall-cmd --reload
```

> 第五步：进入安装目录直接运行install.sh

```
cd hfish
sudo ./install.sh
```

> 第六步：登陆web界面

```
登陆链接：https://[ip]:4433/web/
账号：admin
密码：HFish2021
```

`例：如果管理端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web/`

在安装完成后，HFish会自动在管理端上创建一个节点。可在节点管理进行查看。

<img src="http://img.threatbook.cn/hfish/image-20210914113134975.png" alt="image-20210914113134975" style="zoom: 25%;" />

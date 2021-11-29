### Docker版本简介

Docker是我们推荐的部署方式，当前在2.7.0版本，我们增加了以下特性

- 每隔1小时，会自动请求最新hfish镜像，进行自动升级。过程中数据不出现任何丢失。

如需关闭自动升级，可查看下方文档

- 进行数据持久化保存。

在usr/share/hfish目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。

`特别注意：当前Docker版本使用host模式启动，如果您不希望docker的server开放除4433和4434以外的端口，可以暂停本机节点。`



### Docker安装说明

#### 在docker中安装控制端：

**<u>请确认自己的环境已经安装并启动docker</u>**

> 步骤1:运行版本hfish（框内全部复制，粘贴，执行即可）

```shell
docker run -itd --name hfish \
-v /usr/share/hfish:/usr/share/hfish \
--network host \
threatbook/hfish-server:latest
```

<img src="http://img.threatbook.cn/hfish/4351638188574_.pic_hd.jpg" alt="4351638188574_.pic_hd" style="zoom:50%;" />



> 步骤2:配置后续自动升级（框内全部复制，粘贴，执行即可）

```shell
docker run -d    \
 --name watchtower \
 --restart unless-stopped \
  -v /var/run/docker.sock:/var/run/docker.sock  \
  --label=com.centurylinklabs.watchtower.enable=false \
  containrrr/watchtower  \
  --cleanup  \
  hfish \
  --interval 3600
```

![4381638189986_.pic_hd](http://img.threatbook.cn/hfish/4381638189986_.pic_hd.jpg)



> 步骤3:登陆HFish

登陆地址:https://ip:4433/web

初始用户名:admin

初始密码:HFish2021



#### 在docker中安装节点端：

> 先在控制端中正常添加节点：

<img src="http://img.threatbook.cn/hfish/202111071541552.png" alt="image-20211107152418598" style="zoom:50%;" />



> 记录下主机执行命令中的sh文件的url：

~~~wiki
如上图示例，url为：https://10.53.7.96:4434/tmp/jibrZM5VHMVN.sh
~~~

> 把该url拼接在如下的命令之后：

~~~dockerfile
docker run -d --net=host --name hfish-client --restart=always threatbook/hfishnode-amd64

如上图示例，拼接后的完整命令如下：
docker run -d --net=host --name hfish-client --restart=always threatbook/hfishnode-amd64 https://10.53.7.96:4434/tmp/jibrZM5VHMVN.sh
~~~

> 如节点运行环境为64位arm则拼接如下命令

~~~dockerfile
docker run -d --net=host --name hfish-client --restart=always threatbook/hfishnode-arm64

如上图示例，拼接后的完整命令如下：
docker run -d --net=host --name hfish-client --restart=always threatbook/hfishnode-arm64 https://10.53.7.96:4434/tmp/jibrZM5VHMVN.sh
~~~

> 执行上面的完整命令即可在docker容器中启动节点

~~~wiki
注意：本启动方式使用了host网络模式，容器内的开放的端口会同样在宿主机上开放，需要注意跟宿主机上端口是否冲突。
~~~





### 查看HFish日志

usr/share/hfish目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。

![image-20211012223326542](http://img.threatbook.cn/hfish/image-20211012223326542.png)



### 关闭自启动

不执行第二步，自启动即可



### 查看Docker日志

执行

```shell
docker logs hfish
```


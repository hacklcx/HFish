### Docker版本简介

Docker是我们推荐的部署方式，当前在2.6.2版本，我们增加了以下特性

- 使用docker-compose，默认拉起**HFish**以及**MySql镜像**

HFish镜像采用host模式启动，MySql镜像bridge模式启动。可在docker-compose.yml中修改MySql镜像的密码和映射端口。

- 进行数据持久化保存。

docker-compose.yml同目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。



### Docker安装说明

#### 在docker中安装控制端：

**<u>请确认自己的环境已经安装docker和docker compose</u>**

> 步骤1:将docker-compose.yml放置到需要启动HFish的服务器

下载:[docker_compose.yml](https://hfish.cn-bj.ufileos.com/docker-compose/2.6.2/docker-compose.yml)



> 步骤2:在dockercompose.yml中，按需修改MySql容器的对外映射端口和登陆密码

`注意，无论是修改端口还是密码，都需要将箭头所指的两个位置同步做修改`

<img src="http://img.threatbook.cn/hfish/image-20211013175549337.png" alt="image-20211013175549337" style="zoom:40%;" />



> 步骤3: 使用docker-compose up命令，启动docker

```
 docker-compose up -d
```



> 注意:在docker启动后，会有几分钟的mysql容器初始化，请耐心等待

<img src="http://img.threatbook.cn/hfish/image-20211012222554572.png" alt="image-20211012222554572" style="zoom: 20%;" />

<img src="http://img.threatbook.cn/hfish/image-20211013163538978.png" alt="image-20211013163538978" style="zoom:20%;" />



> 步骤4:登陆HFish

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

docker-compose.yml同目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。

![image-20211012223326542](http://img.threatbook.cn/hfish/image-20211012223326542.png)



### 查看Docker日志

进入docker-compose.yml的目录，执行

```shell
docker-compose logs
```


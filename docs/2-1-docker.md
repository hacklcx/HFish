### Docker版本简介

Docker是我们推荐的部署方式，当前在2.6.2版本，我们增加了以下特性

- 使用docker-compose，默认拉起**HFish**以及**MySql镜像**

HFish镜像采用host模式启动，MySql镜像bridge模式启动。可在docker-compose.yml中修改MySql镜像的密码和映射端口。

- 进行数据持久化保存。

docker-compose.yml同目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。



### Docker安装说明

**<u>请确认自己的环境已经安装docker和docker compose</u>**

> 步骤1:将docker-compose.yml放置到需要启动HFish的服务器

下载:[docker_compose.yml](http://hfish.cn-bj.ufileos.com/docker-compose/2.6.2/docker-compose.yml)



> 步骤2:在dockercompose.yml中，按需修改MySql的映射端口和密码

`注意，无论是修改端口还是密码，都需要将箭头所指的两个位置同步做修改`

<img src="http://img.threatbook.cn/hfish/image-20211012222209522.png" alt="image-20211012222209522" style="zoom:43%;" />



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



### 查看HFish日志

docker-compose.yml同目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。

![image-20211012223326542](http://img.threatbook.cn/hfish/image-20211012223326542.png)



### 查看Docker日志

进入docker-compose.yml的目录，执行

```shell
docker-compose logs
```


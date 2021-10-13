#### Docker版本简介

Docker是我们推荐的部署方式，当前在2.6.2版本，我们增加了以下特性

- 使用docker-compose，默认拉起**HFish**以及**MySql镜像**

HFish镜像采用host模式启动，MySql镜像bridge模式启动。可在docker-compose.yml中修改MySql镜像的密码和映射端口。

- 进行数据持久化保存。

docker-compose.yml同目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。



#### Docker使用说明

**<u>请确认自己的环境已经安装docker和docker compose</u>**

> 步骤1:将docker-compose.yml放置到需要启动HFish的服务器

下载:[docker_compose.yml](http://hfish.cn-bj.ufileos.com/docker-compose/2.6.2/docker-compose.yml)

复制:

```shell
version: "3.7"
services:
  web:
    image: threatbook/hfish-server:2.6.2
    network_mode: "host"
    container_name: hfish-server # 容器名
    restart: always
    volumes:
      - "./logs:/opt/hfish/logs"
      - "/root/.hfish:/root/.hfish"
    depends_on:
      - db
    command: sh /wait.sh hfish root 1234567 3306 /opt/hfish/server

  db:
    image: threatbook/hfish-mysql:2.6.2
#    build: ./mysql
    restart: always
    container_name: hfish-mysql-db # 容器名
    environment:
      - MYSQL_ROOT_PASSWORD=1234567
      - TZ=Asia/Shanghai
    ports:
      - 3306:3306
    volumes:
      - ./data:/var/lib/mysql
    command: --character-set-server=utf8mb4
      --collation-server=utf8mb4_general_ci
      --explicit_defaults_for_timestamp=true
      --lower_case_table_names=1
      --default-time-zone=+08:00
```

> 步骤2:在dockercompose.yml中，按需修改MySql的映射端口和密码

`注意，无论是修改端口还是密码，都需要将箭头所指的两个位置同步做修改`

![image-20211012222209522](http://img.threatbook.cn/hfish/image-20211012222209522.png)



> 步骤3: 使用docker-compose up命令，启动docker

```
docker-compose up
```



> 注意:在docker启动后，会有几分钟的mysql容器初始化，请耐心等待

<img src="http://img.threatbook.cn/hfish/image-20211012222554572.png" alt="image-20211012222554572" style="zoom:50%;" />

<img src="http://img.threatbook.cn/hfish/image-20211012222730930.png" alt="image-20211012222730930" style="zoom:50%;" />



> 步骤4:登陆HFish

登陆地址:https://ip:4433/web

初始用户名:admin

初始密码:HFish2021



> 数据持久化查询

docker-compose.yml同目录下，会生成一个data文件夹，用于存放当前的所有攻击数据。另外，会生成一个logs文件夹，存放当前的所有日志。

![image-20211012223326542](http://img.threatbook.cn/hfish/image-20211012223326542.png)
Docker是我们推荐的蜜罐交付方式。而且因为容器环境本身就有一层权限隔离的原因，合理配置过的Docker运行环境，能获得更高的业务安全性。

当前，我们只提供Linux amd64版本的docker镜像

> Linux amd64 Docker镜像的下载

```shell
docker pull dskyz/hfish:latest
```

> 镜像的运行

```shell
docker run -d -p 4433:4433 -p 4434:4434 --name=hfish --restart=always dskyz/hfish:latest
```

例：如果控制端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web

> 登陆web界面

```
登陆链接：https:// [ip]:4433/web
账号：admin
密码：HFish2021
```

例：如果控制端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web

控制端部署完成后，请继续参考下面的【控制端配置】完成配置


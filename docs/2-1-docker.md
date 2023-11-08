### Docker下载部署

Docker是我们推荐的部署方式之一，当前的版本拥有以下特性：

- 自动升级：每小时请求最新镜像进行升级，升级不会丢失数据。

- 数据持久化：在宿主机/usr/share/hfish目录下建立data目录用于存放攻击数据，建立logs目录用于存放日志。

`注意：当前Docker版本使用host模式启动，如果您不希望Docker的管理端开放除4433和4434以外的端口，可暂停本机节点。`



#### Docker默认安装说明

> ##### 确认主机中已安装并启动Docker #####

```
docker version
```

> ##### 运行版本HFish（框内全部复制，粘贴，执行即可） ##### 

```
docker run -itd --name hfish \
-v /usr/share/hfish:/usr/share/hfish \
--network host \
--privileged=true \
threatbook/hfish-server:latest
```

<img src="https://hfish.net/images/4351638188574_.pic_hd.jpg" alt="4351638188574_.pic_hd" style="zoom:50%;" />



> ##### 配置后续自动升级（框内全部复制，粘贴，执行即可） ##### 

```
docker run -d    \
 --name watchtower \
 --restart unless-stopped \
  -v /var/run/docker.sock:/var/run/docker.sock  \
  --label=com.centurylinklabs.watchtower.enable=false \
--privileged=true \
  containrrr/watchtower  \
  --cleanup  \
  hfish \
  --interval 3600
```

![4381638189986_.pic_hd](https://hfish.net/images/4381638189986_.pic_hd.jpg)



>  ##### 登陆HFish  ##### 

```
登陆地址：https://[server]:4433/web/
初始用户名：admin
初始密码：HFish2021
```


#### Docker升级失败

如果已经配置了Docker镜像代理，有可能会导致watchower无法生效，手动执行：

```
docker pull threatbook/hfish-server:3.3.4  
docker tag threatbook/hfish-server:3.3.4  threatbook/hfish-server:latest  
docker rm -f hfish  
docker run -itd --name hfish \
-v /usr/share/hfish:/usr/share/hfish \
--network host \
--privileged=true \
threatbook/hfish-server:latest
```


#### 未配置自动升级情况下，Docker单次手动升级说明

> ##### 配置watchover（框内全部复制，粘贴，执行即可）

```
docker run -d    \
 --name watchtower \
 --restart unless-stopped \
  -v /var/run/docker.sock:/var/run/docker.sock  \
  --label=com.centurylinklabs.watchtower.enable=false \
--privileged=true \
  containrrr/watchtower  \
  --cleanup  \
  hfish \
  --interval 10
```

> ##### 等待升级成功后，登录页面，确认升级完成

> ##### 取消watchover自动升级

```
docker stop watchtower
```

配置后，如果之后定期升级，就可以不断的通过 docker start watchtower 和 docker stop watchtower来完成对hfish镜像的手动升级



#### Docker修改持久化配置后重启说明

> ##### 在usr/share/hfish/config.toml下面修改配置

> ##### 重启docker容器

```
docker restart hfish
```

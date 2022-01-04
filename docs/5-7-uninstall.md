### 卸载管理端

#### 1. 卸载Linux管理端

1、删除计划任务进程

`注意：不同的linux版本结束方式不同，需要自己确认`


2、结束管理端进程

```
# 结束hfish和hfish-server的进程
root@HFish~# ps ax | grep ./hfish | grep -v grep
8435 ?        Sl    97:59 ./hfish
8436 ?        Sl    97:59 ./hfish-server


root@HFish:~# kill -9 8435
root@HFish:~# kill -9 8436
```

3、删除文件夹

```
# 使用install.sh安装的HFish会被部署到/opt/hfish目标，将整个删除即可
root@HFish~# rm -rf /opt/hfish
```

4、清理所有配置（如果后续还要安装，建议不要删除，否则下次使用需要完全重新配置）

```
# 使用install.sh安装的HFish会在/usr/share/hfish下建立全局变量
root@HFish~# rm -rf /usr/share/hfish
```

5、删除MySQL数据库配置（SQLite可忽略）

```
# 删除HFish数据库
root@HFish:~# mysql -h127.0.0.1 -uroot -p
Enter password:*******（默认密码详见config.ini配置文件）
mysql> DROP DATABASE hfish;

# 停止MySQL服务
root@HFish:~# systemctl stop mysqld
root@HFish:~# systemctl disable mysqld
```

6、可还原SSH和Firewall配置

```
# 清除SSH config内对于访问来源的限制
root@HFish~# vi /etc/ssh/sshd_config
注释掉以 AllowUsers root@ 开头的行

# 重启SSH服务
root@HFish~# systemctl restart sshd

# 清除Firewall服务的规则（请根据实际情况删除！）
root@HFish~# firewall-cmd --permanent --list-all | grep ports | head -n 1 | \
cut -d: -f2 | tr ' ' '\n' | xargs -I {} firewall-cmd --permanent --remove-port={}

# 重启Firewall服务
root@HFish~# systemctl restart firewalld
```

#### 2、卸载Windows管理端

1、删除计划任务进程HFish管理端

<img src="/images/image-20211206115017049.png" alt="image-20211206115017049" style="zoom: 33%;" />

<img src="/images/image-20211206115035865.png" alt="image-20211206115035865" style="zoom:33%;" />

2、结束hfish进程

在任务管理器，结束hfish和hfish-server的进程


3、删除管理端文件夹

### 节点端

#### 1、卸载Linux节点

> 卸载节点

1、删除计划任务进程

```
# 打开计划任务，删除带有hfish字样的行
root@HFish~# crontab -e
```

2、结束hfish进程

```
# 结束hfish和hfish-server的进程
root@HFish~# ps ax | grep ./hfish | grep -v grep
8435 ?        Sl    97:59 ./hfish
8436 ?        Sl    97:59 ./hfish-server


root@HFish:~# kill -9 8435
root@HFish:~# kill -9 8436
```

3、删除client文件夹

文件夹路径为按照自己的安装路径，client端没有全局配置，删除安装文件夹即可



#### 2、删除Windows节点

1、关闭计划任务 HFishClient

<img src="/images/image-20211206115017049.png" alt="image-20211206115017049" style="zoom: 33%;" />

2、结束client进程

在任务管理器，结束hfish和hfish-client的进程


3、删除client文件夹

文件夹路径为按照自己的安装路径，client端没有全局配置，删除安装文件夹即可


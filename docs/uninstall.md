# 卸载蜜罐的流程

> 卸载 server 端

1. 杀死 server 进程

```shell
# 结束 server 进程
root@HFish~$ ps ax | grep ./server | grep -v grep
 8435 ?        Sl    97:59 ./server

root@HFish:~$ sudo kill -9 8435
```

2. 删除 server 文件夹

```shell
# 使用 install.sh 安装的 HFish 会被部署到/opt/hfish目标，删除即可
root@HFish~$ sudo rm -rf /opt/hfish
```

3. 清理数据库（如果使用的是 SQLite 数据库请忽略）

```shell
# 删除 HFish 数据库
root@HFish:~$ mysql -h127.0.0.1 -uroot -p
Enter password:*******（默认密码详见config.ini配置文件）
mysql> DROP DATABASE hfish;

# 停止 MySQL 服务
root@HFish:~$ sudo systemctl stop mysqld
root@HFish:~$ sudo systemctl disable mysqld
```

4. 还原 SSH 和 Firewall 配置

```shell
# 清除 SSH config 内对于访问来源的限制
root@HFish~$ sudo vi /etc/ssh/sshd_config
注释掉以 AllowUsers root@ 开头的行

# 重启 SSH 服务
root@HFish~$ sudo systemctl restart sshd

# 清除 Firewall 服务的规则（请根据实际情况删除！）
root@HFish~$ sudo firewall-cmd --permanent --list-all | grep ports | head -n 1 | \
cut -d: -f2 | tr ' ' '\n' | xargs -I {} firewall-cmd --permanent --remove-port={}

# 重启 Firewall 服务
root@HFish~$ sudo systemctl restart firewalld
```



> 卸载节点端

1. 杀死 client 和 services 进程

```shell
# 结束 client 和 services 进程
root@HFish~$ sudo ps ax | grep -E 'services|./client' | grep -v grep
  10506 ?        Sl   134:20 ./client
1685739 ?        Sl     0:00 services/mysql/mysql 0.0.0.0:3306 true
1685743 ?        Sl     0:00 services/redis/redis 0.0.0.0:6379 true
1685747 ?        Sl     0:00 services/telnet/telnet 0.0.0.0:23 true
1685751 ?        Sl     0:00 services/ftp/ftp 0.0.0.0:21 true
1685755 ?        Sl     0:00 services/es/es 0.0.0.0:9200 true

root@HFish:~$ sudo kill -9 10506 1685739 1685743 1685747 1685751 1685755
```

2. 删除 client 文件夹

```shell
# 使用 install.sh 安装的 HFish 会被部署到/opt/hfish目标，删除即可
root@HFish~$ sudo rm -rf /opt/hfish
```

3. 还原 SSH 和 Firewall 配置

```shell
# 还原默认 SSH 端口
root@HFish~$ sudo vi /etc/ssh/sshd_config
- 把 Port 22122 注释掉或修改为默认的22

# 清除 SSH config 内对于访问来源的限制
root@HFish~$ sudo vi /etc/ssh/sshd_config
注释掉以 AllowUsers root@ 开头的行

# 重启 SSH 服务
root@HFish~$ sudo systemctl restart sshd

# 清除 Firewall 服务规则（请根据实际情况删除！）
root@HFish~$ sudo firewall-cmd --permanent --list-all | grep ports | head -n 1 | \
cut -d: -f2 | tr ' ' '\n' | xargs -I {} firewall-cmd --permanent --remove-port={}

# 重启 Firewall 服务
root@HFish~$ sudo systemctl restart firewalld
```


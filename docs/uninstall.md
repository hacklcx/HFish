# 卸载蜜罐的流程

> 卸载server端

1. 杀死server进程

```shell
# 结束serverhe进程
root@HFish~# ps ax | grep ./server | grep -v grep
 8435 ?        Sl    97:59 ./server

root@HFish:~# kill -9 8435
```

2. 删除server文件夹

```shell
# 使用install.sh安装的HFish会被部署到/opt/hfish目标，删除即可
root@HFish~# rm -rf /usr/share/
```

3. 清理数据库（如果使用的是SQLite数据库请忽略）

```shell
# 删除HFish数据库
root@HFish:~# mysql -h127.0.0.1 -uroot -p
Enter password:*******（默认密码详见config.ini配置文件）
mysql> DROP DATABASE hfish;

# 停止MySQL服务
root@HFish:~# systemctl stop mysqld
root@HFish:~# systemctl disable mysqld
```

4. 还原SSH和Firewall配置

```shell
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



> 卸载节点端

1. 杀死clien

```shell
# 结束 client 和 services 进程
root@HFish~# ps ax | grep -E 'services|./client' | grep -v grep
  10506 ?        Sl   134:20 ./client

root@HFish:~# kill -9 10506
```

2. 删除client文件夹

```shell
# 使用 install.sh 安装的 HFish 会被部署到/opt/hfish目标，删除即可
root@HFish~# rm -rf /opt/hfish
```

3. 还原SSH和Firewall配置

```shell
# 还原默认SSH端口
root@HFish~# vi /etc/ssh/sshd_config
- 把 Port 22122 注释掉或修改为默认的22

# 清除SSH config内对于访问来源的限制
root@HFish~# vi /etc/ssh/sshd_config
注释掉以 AllowUsers root@ 开头的行

# 重启SSH服务
root@HFish~# systemctl restart sshd

# 清除Firewall服务规则（请根据实际情况删除！）
root@HFish~# firewall-cmd --permanent --list-all | grep ports | head -n 1 | \
cut -d: -f2 | tr ' ' '\n' | xargs -I {} firewall-cmd --permanent --remove-port={}

# 重启Firewall服务
root@HFish~# systemctl restart firewalld
```


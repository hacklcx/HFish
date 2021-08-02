# 卸载蜜罐的流程

> 卸载server端

1. 退出server进程

```shell
# 结束./server 进程
root@HFish~$ ps aux | grep server
root       8435  0.2 10.2 997804 188176 ?       Sl   Mar23  23:54 ./server

root@HFish:~$ kill 8435
```

2. 删除server文件夹

```shell
# 默认情况下 OneFish 统一被部署到/opt/onefish目标，删除即可
root@HFish~$ rm -rf /opt/onefish
```

3. 清理数据库

```shell
#删除 OneFish 数据库
root@HFish:~$ mysql -h127.0.0.1 -uroot -p
Enter password:*******（默认为OneFish210!）
mysql> DROP DATABASE onefish;

# 停止 MySQL 服务
root@HFish:~$ systemctl stop mysqld
root@HFish:~$ systemctl disable mysqld
```

4. 还原SSH和Firewall配置

```shell
# 删除SSH config内对于访问来源的限制
root@HFish~$ vi /etc/ssh/sshd_config
删除 AllowUsers root@xxx 这行

# 重启 SSH 服务
root@HFish~$ systemctl restart sshd

# 清除Firewall服务的规则
root@HFish~$ firewall-cmd --permanent --list-all | grep ports | head -n 1 | \
cut -d: -f2 | tr ' ' '\n' | xargs -I {} firewall-cmd --permanent --remove-port={}

# 重启 Firewall 服务
root@HFish~$ systemctl restart firewalld
```



> 卸载节点端

1. 退出client进程

```shell
# 结束./client 进程
root@HFish~$ ps aux | grep client
root       1012  0.2 10.2 997804 188176 ?       Sl   Mar23  23:54 ./client

root@HFish:~$ kill -8 1012
```

2. 删除client文件夹

```shell
# 默认情况下 OneFish 统一被部署到/opt/onefish目标，删除即可
root@HFish~$ rm -rf /opt/onefish
```

3. 还原SSH和Firewall配置

```shell
# 还原默认 SSH 端口
root@HFish~$ vi /etc/ssh/sshd_config
- 把 Port 22122 注释掉或修改为默认的22

# 删除 SSH config 内对于访问来源的限制
root@HFish~$ vi /etc/ssh/sshd_config
删除 AllowUsers root@xxx 这行

# 重启 SSH 服务
root@HFish~$ systemctl restart sshd

# 清除Firewall服务的规则
root@HFish~$ firewall-cmd --permanent --list-all | grep ports | head -n 1 | \
cut -d: -f2 | tr ' ' '\n' | xargs -I {} firewall-cmd --permanent --remove-port={}

# 重启 Firewall 服务
root@HFish~$ systemctl restart firewalld
```


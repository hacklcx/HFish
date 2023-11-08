#### 错误排查

> ##### 管理端部署完成后，访问Web管理页面始终无法打开 #####

1. 确认浏览器访问地址是 https://[server]:4433/web/，注意不可缺少“/web/”这个路径
2. 确认管理端进程的运行情况和TCP/4433端口开放情况，如果不正常需要重启管理端进程

```
# 检查 hfish-server的进程是否运行正常
ps ax | grep ./hfish | grep -v grep
​
# 检查TCP/4433端口是否正常开放
ss -ntpl
```

3. 检查管理端主机是否开启了防火墙，导致目前无法访问，必要情况，考虑关闭防火墙

```
#centos7 检查防火墙状态
systemctl status firewalld

#centos7 检查防火墙开放端口
firewall-cmd --list-ports
```

4. Linux环境使用date命令确认系统时间的准确
5. 如果以上都没有问题，请将server和client日志提供给我们

```
节点端日志在安装目录的logs文件夹内，文件名为client.log
Linux管理端日志在/usr/share/hfish/log文件夹内，文件名为server.log
Windows管理端日志在C:\Users\Public\hfish\log文件夹内，文件名为server.log
```

> ##### 节点状态为红色离线 #####

1. 检查节点到管理端的网络连通情况，以下是几种常见情况

```
节点每60秒连接管理端的TCP/4434端口一次，180秒内连接不上即显示为离线。
刚完成部署或网络不稳定的时候会出现显示为离线。
通常情况，等待2~3分钟，如果节点恢复绿色在线，那蜜罐服务也会从绿色启用，变成绿色在线。
```

2. 检查节点上的进程运行情况，如果进程运行异常，杀死全部关联进程后重启进程，并查看错误日志

```
# 检查./client的进程是否运行正常
ps ax | grep -E 'services|./client' | grep -v grep        
```

3. 如果以上都没有问题，请将server和client日志提供给我们

```
节点端日志在安装目录的logs目录内，文件名为client.log
Linux管理端日志在/usr/share/hfish/log文件夹内，文件名为server.log
Linux管理端日志在C:\Users\Public\hfish\log文件夹内，文件名为server.log
​```


> ##### 节点在线，部分蜜罐服务在线，部分蜜罐服务离线 #####

通常情况，用户可以登录Web管理界面，将鼠标悬停在蜜罐后面的问号图标，查看离线原因，如下图：

![image-20220721111203773](http://img.threatbook.cn/hfish/image-20220721111203773.png)


> ##### 蜜罐离线原因 bind:address already in use #####

此蜜罐占用的端口已经被别的进程占用。

新手常见的错误之一是使用HFish模拟SSH服务，而又没有修改真正的SSH服务，导致蜜罐SSH服务和真正SSH服务争夺TCP/22端口，观测到的现象是刚启动节点时SSH蜜罐服务在线，刷新页面后显示离线。

使用ss -ntpl命令检查该蜜罐服务的端口是否被占用，如果被占用，建议修改该业务的默认端口。
​
Windows操作系统上，如果用户启用了tcp端口监听，大概率会发现TCP 135、139、445、3389端口冲突，
这是用于Windows默认占用了这些端口，不建议在Windows上监听TCP 135、139、445、3389端口。
​
Linux操作系统中，可以通过netstat -ant查看端口是否已被占用（需root权限）。

> ##### 我通过SSH登录运行了节点进程，如何防止SSH退出后节点进程终止？ ##### 
```
nohup .~/client >>nohup.out 2>&1 &
```
​
> ##### 如何在Linux上设定节点进程开机自启动 #####
```
echo 'nohup .~/client >>nohup.out 2>&1 &' >> /etc/rc.local
```
​
> ##### 如何在Linux上通过计划任务自动重启节点进程 #####
```
echo '* * * * * nohup .~/client >>nohup.out 2>&1 &' >> /var/spool/cron/crontabs/root
```
`当前hfish分为两个进程,"hfish"进程为管理进程，负责监测、拉起和升级蜜罐主程序。"server"进程为蜜罐主程序进程，其执行蜜罐软件程序。因此，安装时候，请务必按照要求，执行hfish进程；如果直接执行server程序，可能会导致程序不稳定，升级失败等情况。`

`hfish-windows的数据库文件当前存储在 C:\\Users\\Public\\hfish 下，在当前机器进行重装时候，会自动默认读取文件下配置和数据"。`



> **第一步：下载安装包**[HFish-Windows-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-windows-amd64.tgz) （Windows x86 架构 64 位系统）

> 第二步：防火墙上进出站双向打开TCP4433、4434端口放行（如果需要使用其他服务，也对应打开端口）

> 第二步：运行文件目录下的install.sh

 autorun.bat将帮助server.exe转为后台运行，启动后直接运行。(启动在一分钟内)

> 第三步：登陆web界面

```
登陆链接：https://[ip]:4433/web/
账号：admin
密码：HFish2021
```

例：如果控制端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web

`如果您启动页面失败，请查看windows防火墙是否放开4433与4434`

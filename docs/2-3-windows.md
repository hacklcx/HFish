#### Windows 3.1.4 环境安装流程

> 第一步：下载安装包[HFish-Windows-amd64](https://hfish.cn-bj.ufileos.com/hfish-3.1.4-windows-amd64.tgz) （Windows x86 架构 64 位系统），解压缩

> 第二步：防火墙上进出站双向打开TCP/4433、4434端口放行（如需使用其他服务，也需要打开端口）

> 第三步：进入HFish-Windows-amd64文件夹内，运行文件目录下的install.bat （脚本会在当前目录进行安装HFish）

> 第四步：登陆web界面

```
登陆链接：https://[ip]:4433/web/
账号：admin
密码：HFish2021
```

`如果管理端的地址是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web/`

注意事项：

`1、当前HFish分为两个进程,"hfish"进程为管理进程，负责监测、拉起和升级蜜罐主程序。"管理端"进程为蜜罐主程序进程，其执行蜜罐软件程序。因此，安装时候，请务必按照要求，执行hfish进程；如果直接执行管理端程序，可能会导致程序不稳定，升级失败等情况。`

`2、HFish的windows版数据库文件当前存储在 C:\Users\Public\hfish 目录，重装HFish后，HFish默认自动读取该目录内的配置和数据"。`

`3、如果访问页面失败，检查Windows防火墙是否放开TCP/4433和4434`


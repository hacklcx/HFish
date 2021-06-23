# 下载

### HFish v<% version %>

```wiki
2021年6月16日发布

1.新增扫描感知功能，可感知到全端口范围内的TCP、UDP、ICMP扫描，支持IPv4与IPv6。
2.新增服务在线下载及上传功能，新版发布后，进入服务列表，即可看到最新的服务。
3.新增多用户管理功能，支持管理员与普通用户的权限区分。
4.修复节点多网卡导致的报错问题，每个节点最高支持50个不同ip地址。
5.修复邮件服务器配置问题，填写发件邮箱进行邮件配置测试。
6.修复情报页面api报错问题。
```

!> 注意：如果当前使用 sqlite 数据库的话，升级时，hfish.db 文件将会被覆盖，导致之前的攻击记录丢失，请注意进行备份。如果要将之前的 db 文件导入当前版本时，请参考 mysql.sql 的语句修改 db 文件，执行导入。

## 下载安装

### 控制端安装包

- [HFish-Linux-amd64](http://hfish.cn-bj.ufileos.com/hfish-<% version %>-linux-amd64.tar.gz) 为 Linux x86 架构 64 位系统使用
- [HFish-Windows-amd64](http://hfish.cn-bj.ufileos.com/hfish-<% version %>-windows-amd64.tar.gz) 为 Windows x86 架构 64 位系统使用
- [HFish-Linux-arm64](http://hfish.cn-bj.ufileos.com/hfish-<% version %>-linux-arm64.tar.gz) 为 Linux Arm 架构 64 位系统使用，常见于 NAS、路由器、树莓派等……

## 文件结构

```wiki
HFish
│   server  #控制端文件
│   config.ini  #控制端配置文件
│   README.md  #安装使用说明
│   version  #版本信息
│   ssl.key  #SSL私钥
│   ssl.pem  #SSL证书
│   tools  #控制台工具，目前功能为重置控制台web登录密码
│
└───db
│   │   hfish.db  #sqlite数据
│   │   ipip.ipdb  #ip归属地信息
│   │
│   └───sql
│       └───mysql
│       │   V<% version %>__mysql.sql  #mysql数据库用户升级文件
│       │
│       └───sqlite
│       │   V<% version %>__sqlite.sql  #sqlite数据库用户升级文件
│
└───logs
│   │   server-年-月-日.log  #server日志文件
│
└───packages
│   │   install.sh  #节点部署时安装脚本
│   │   node_account.conf  #蜜饵源文件
│   │
│   └───linux-x86  #Linux 服务包
│       │   client
│       │   service-*.tar.gz
│
└───static  #web服务预览图
│   └───services
│       │   ……
│
└───web  #控制台web文件
│   │   ……

```

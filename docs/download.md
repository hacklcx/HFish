# 下载

### HFish v<% version %>

```wiki
2021年7月25日发布

特别注意：只支持2.4.0及其以上的版本支持顺滑升级，其他版本需要重新进行部署安装。

1.新增API配置功能，支持用户对攻击IP、攻击详细信息、攻击者所使用攻击账号密码导出。
2.新增TCP端口监听服务，支持最高对10个自定义端口的灵活监听。
3.新增蜜罐服务的支持，单个节点，最高可添加10种蜜罐服务。
4.新增windows开机自启动能力，防止意外关机导致的程序退出。
5.修复告警策略中，修改配置不生效的问题 。
6.修复数据清理时，扫描数据，攻击IP及账号资产未进行清理的问题。
7.修复部分使用交互问题。
```



## 下载安装

### 控制端安装包

- [HFish-Linux-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-linux-amd64.tar.gz) 为 Linux x86 架构 64 位系统使用
- [HFish-Windows-amd64](https://hfish.cn-bj.ufileos.com/hfish-<% version %>-windows-amd64.tar.gz) 为 Windows x86 架构 64 位系统使用
- [HFish-Linux-arm64](https://hfish.cn-bj.ufileos.com/hfish-2.5.0-linux-arm64.tar.gz) 为 Linux Arm 架构 64 位系统使用，常见于 NAS、路由器、树莓派等……

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
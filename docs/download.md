# 下载

### HFish v2.3.0

```wiki
2021年4月30日发布

1. 新增主机蜜饵功能，其可将仿真的账号配置文件，作为蜜饵分发到主机，通过高仿真的假配置文件，诱导攻击者攻击蜜罐，触发内网失陷警告，精准溯源失陷主机。

2. 新增web蜜罐上传功能，可根据指导文档，自主开发web蜜罐，在服务页面上传，即可运行使用。

3. 新增「账号资产」页面，内覆盖所有被攻击时使用的用户名密码。其中，页面支持对企业关键字进行监测，可以导入企业名字和员工姓名，实时查询是否有对应信息被攻击者尝试，一旦发现，说明该攻击风险较高。

4. 新增攻击者IP陈列，辅助用户进行高级溯源。

5. 新增本地升级功能，可点击升级按钮，选择本地上传升级包升级。

6. web蜜罐增加对HTTPS的支持，可在模版页面进行操作修改。

7. 新增web蜜罐访问记录，支持记录子目录与url扫描访问请求。

8. 新增web蜜罐预览功能，支持预览web蜜罐详情界面，帮助进行模版组合。
```



!> 注意：如果当前使用sqlite数据库的话，升级时，hfish.db文件将会被覆盖，导致之前的攻击记录丢失，请注意进行备份。如果要将之前的db文件导入当前版本时，请参考mysql.sql的语句修改db文件，执行导入。



## 下载安装

请根据自己要部署系统的架构选择「控制端安装包」，如果您需要用不同架构设备进行集群部署，您需要补充相应的「增添服务包」。

<!-- tabs:start -->

#### ** 控制端安装包 **

+ [HFish-Linux-amd64](http://hfish.cn-bj.ufileos.com/hfish-linux-amd64.tar.gz) 为 Linux x86架构64位系统使用
+ [HFish-Linux-386](http://hfish.cn-bj.ufileos.com/hfish-linux-386.tar.gz) 为 Linux x86架构32位系统使用
+ [HFish-Windows-amd64](http://hfish.cn-bj.ufileos.com/hfish-windows-amd64.tar.gz) 为 Windows x86架构64位系统使用
+ [HFish-Windows-386](http://hfish.cn-bj.ufileos.com/hfish-windows-386.tar.gz)* 为 Windows x86架构32位系统使用
+ [HFish-Linux-arm64](http://hfish.cn-bj.ufileos.com/hfish-windows-arm64.tar.gz) 为 Linux Arm架构64位系统使用，常见于NAS、路由器、树莓派等……



#### ** 增添服务包 **

请把增添的服务包下载到控制端文件目录下 `packages`目录下.

+ [Services-Linux-amd64](http://hfish.cn-bj.ufileos.com/services-linux-amd64.tar.gz) 为 Linux x86架构64位系统使用
+ [Services-Linux-386](http://hfish.cn-bj.ufileos.com/services-linux-386.tar.gz) 为 Linux x86架构32位系统使用
+ [Services-Windows-amd64](http://hfish.cn-bj.ufileos.com/services-windows-amd64.tar.gz) 为 Windows x86架构64位系统使用
+ [Services-Windows-386](http://hfish.cn-bj.ufileos.com/services-windows-386.tar.gz)为 Windows x86架构32位系统使用
+ [Services-Linux-arm64](http://hfish.cn-bj.ufileos.com/services-linux-arm64.tar.gz) 为 Linux Arm架构64位系统使用，常见于NAS、路由器、树莓派等……

<!-- tabs:end -->



## 文件结构

```wiki
HFish 
│   server  #控制端文件 
│   config.ini  #控制端配置文件
│   README.md  #安装使用说明
│   version  #版本信息
│   ssl.key  #SSL私钥
│   ssl.pem  #SSL证书
│
└───db
│   │   hfish.db  #sqlite数据
│   │   ipip.ipdb  #ip归属地信息
│   │
│   └───sql
│       └───mysql
│       │   V2.3.0__mysql.sql  #mysql数据库用户升级文件
│       │         
│       └───sqlite
│       │   V2.3.0__sqlite.sql  #sqlite数据库用户升级文件
│   
└───logs
│   │   server-年-月-日.log  #server日志文件
│   
└───packages
│   │   install.sh  #节点部署时安装脚本
│   │   node_account.conf  #蜜饵源文件
│   │
│   └───linux-amd64  #Linux 64位服务包
│       │   client
│       │   service-*.tar.gz
│   
└───static  #web服务预览图
│   └───services  
│       │   ……
│
└───web  #控制台web文件
│   │   ……
│ 
└───tools
    │   tools  #控制台工具，目前功能为重置控制台web登录密码
```


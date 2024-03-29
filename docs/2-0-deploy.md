#### 快速部署

HFish采用B/S架构，系统由管理端和节点端组成，管理端用来生成和管理节点端，并接收、分析和展示节点端回传的数据，节点端接受管理端的控制并负责构建蜜罐服务。


> ##### HFish支持架构列表 ##### 

|                           | Windows            | Linux X86          |
| ----------------- | ----------------- | ----------------- |
| 管理端（Server)  | 支持64位            | 支持64位            |
| 节点端（Client） | 支持64位和32位 | 支持64为和32位 |


> ##### HFish部署在内网所需配置 ##### 

部署在内网的蜜罐对性能要求较低，针对过往测试情况，我们给出两个配置。

|               | 管理端         | 节点端       |
| --------- | ------------ | ----------- |
| 建议配置 | 2核4g200G | 1核2g50G |
| 最低配置 | 1核2g100G | 1核1g50G |

`注意：日志磁盘占用情况受攻击数量影响很大，建议管理端配置200G以上硬盘空间。`


> ##### HFish部署在外网所需配置（必须更换MySQL数据库） ##### 

部署在公网的蜜罐因遭受更多攻击，因此会有更大性能需求。

|               | 管理端         | 节点端       |
| --------- | ------------- | ----------- |
| 建议配置 | 4核8g200G | 1核2g50G |
| 最低配置 | 2核4g100G | 1核1g50G |

`注意：日志磁盘占用情况受攻击数量影响较大，建议管理端配置200G以上硬盘空间。`



> ##### HFish部署权限要求 ##### 

1、如果使用官网推荐的install.sh脚本安装，必须具备root权限，安装目录会位于opt目录下；

2、如果下载安装包手动安装，在默认使用SQLite数据库情况下，管理端的部署和使用不需要root权限，但如果后续需要替换SQLite改为MySQL数据，则MySQL安装和配置需要root权限；

3、节点端安装和运行无需root权限，但是由于操作系统限制，非root权限运行的节点无法监听低于TCP/1024的端口；

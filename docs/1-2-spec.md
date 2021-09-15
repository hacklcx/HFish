### HFish架构

HFish采用B/S架构，系统由控制端和节点端组成，控制端用来生成和管理节点端，并接收、分析和展示节点端回传的数据，节点端接受控制端的控制并负责构建蜜罐服务。

在HFish中，**管理端**只用于**数据的分析和展示**，**节点端**进行**虚拟蜜罐**，最后由**蜜罐来承受攻击**。

<img src="http://img.threatbook.cn/hfish/image-20210902163914134.png" alt="image-20210902163914134" style="zoom:50%;" />



### HFish特点

HFish当前具备如下几个特点：

- 安全可靠：主打低中交互蜜罐，简单有效；

- 蜜罐丰富：支持SSH、FTP、TFTP、MySQL、Redis、Telnet、VNC、Memcache、Elasticsearch、Wordpress、OA系统等10多种蜜罐服务，支持用户制作自定义Web蜜罐；

- 开放透明：支持对接微步在线X社区API、五路syslog输出、支持邮件、钉钉、企业微信、飞书、自定义WebHook告警输出；

- 快捷管理：支持单个安装包批量部署，支持批量修改端口和服务；

- 跨平台：支持Linux x32/x64/ARM、Windows x32/x64平台




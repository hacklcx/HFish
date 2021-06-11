# 节点端部署

当您在【模板管理】页面添加完需要的蜜罐模板后，您就可以进行【增加节点】的操作了。

> 点击增加节点

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318190541416.png" alt="image-20210318190541416" style="zoom:50%;" />



> 对要增加的节点进行具体的配置。

当前，管理端启动会自带两个默认模版，分别对应【研发测试环境】和【通用web环境】。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318190641973.png" alt="image-20210318190641973" style="zoom:50%;" />



- 节点名称：可以自定义，用于您个人管理中识别设备的名称，长度不低于6个字符
- 部署位置：可以自定义，用于您个人管理中识别设备的位置
- 节点安装包：目前我们支持linux_x86、linux_arm、windows，三大平台下的32/64位的设备作为节点，您可以选择相应的版本

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318190656299.png" alt="image-20210318190656299" style="zoom:50%;" />



> 选择配置好之后，我们就会生成一个节点安装包

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318191631254.png" alt="image-20210318191631254" style="zoom:50%;" />



**值得一提的是，由于windows环境的特殊原因，并不能支持命令行部署，需要关闭掉页面后，在节点栏点击，下载安装包。**

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318191736929.png" alt="image-20210318191736929" style="zoom:50%;" />



节点配置：节点配置下拉框中，内容就是用户设置的所有模版。此外，**节点添加以后仍然可以修改模版**。

服务器地址：该服务器地址写您的s端地址，我们支持内网与云主机。请注意一定要注意您的s端地址书写。**假设c端添加上长期没有流量，我们建议您查看一下服务器地址是否正确。**



下面来跟大家一起看一下具体界面

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318192101619.png" alt="image-20210318192101619" style="zoom:50%;" />



> 搜索框（红色）

搜索节点ip，筛选节点状态，查看流量以及模版。

> 导出框（绿色）

这里可以导出所有的节点内容，生成csv格式的文件。

> 节点列表界面（蓝色）

这里显示具体的节点名称，流量。当我们点击的时候可以展开，查看详细信息。其中女服务状态的在线、离线和禁用，可以在模版中进行修改。

![image-20210318192240339](https://hfish.cn-bj.ufileos.com/images/image-20210318192240339.png)





**注意事项**

```shell
1. 建议节点是全新的主机环境，除了节点的服务外不要运行其他的任何程序或开放相应的端口。
2. 建议在安装前修改默认的ssh端口，因为ssh蜜罐服务的端口是22，不修改端口的情况下部署ssh蜜罐，会让该蜜罐无法正常启动。
3. 节点的联通性测试结束，正式使用前，记得配置相应的安全策略。SSH 端口，只能被安全区的设备访问。蜜罐服务端口应该对所有网段开放。非蜜罐服务端口应该结束相关进程或调整 iptables&firewalld 规则限制访问。
```

下面是环境管理的使用方式。



首先要解释的是，在OneFish中，为了蜜罐更好的可拓展性，我们将蜜罐的层次分为三类

> 蜜罐服务

蜜罐的最基础元素是服务，截止到2021.5.6日更新，目前提供 MYSQL、SSH、FTP、Elasticsearch、REDIS、通用OA系统、WordPress、HTTP、VNC、Telnet、MEMCACHE、TFTP、CUSTOM 共13种服务。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318194212394.png" alt="image-20210318194212394" style="zoom:50%;" />



服务管理分为两个模块：

- 服务搜索与管理模块

该模块可以搜索服务名称，筛选服务类型与筛选监听端口，辅助搜寻服务。

- 服务列表

该模块主要是展示服务名称、被模版引用数、默认监听端口和服务概述。

**其中，默认监听端口可以在模版详情中进行操作修改。**





> 蜜罐模版

在服务之上，我们支持对蜜罐服务的组合，并称之为模版。

在软件初始化阶段，我们会自带两种模版。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318194755834.png" alt="image-20210318194755834" style="zoom:50%;" />



另外，可以通过新建模版自由添加

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318195331781.png" alt="image-20210318195331781" style="zoom:50%;" />



<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318195941660.png" alt="image-20210318195941660" style="zoom:50%;" />










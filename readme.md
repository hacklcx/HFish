<p align="center">
  <a href="https://hfish.net/" target="_blank">
    <img width="200" src="images/logo.png">
  </a>
</p>


<h1 align="center">HFish </h1>
<p align="center">HFish是一款安全、简单可信赖的跨平台蜜罐软件，允许商业和个人用户免费使用。</p>

<p  align="center">
<a href="https://hfish.net/" target="_bank">官网</a>
<span>|</span>
<a href="https://hfish.net/#/?id=hfish%e8%ae%be%e8%ae%a1%e7%90%86%e5%bf%b5" target="_bank">使用手册</a>
<span>|</span>
<a href="https://github.com/hacklcx/HFish" target="_bank">Github</a>
<span>|</span>
<a href="https://gitee.com/lauix/HFish" target="_bank">Gitee</a>
<span>|</span>
<a href="https://hfish.net/#/download" target="_bank">下载部署</a>
</p>

## 二维码

## 特点

+ 安全可靠：主打低中交互蜜罐，简单有效；云端高交互蜜罐，方便安全。

+ 功能丰富：含有43高低交互蜜罐，支持基本网络服务、OA系统、CRM系统、NAS存储系统、Web服务器、运维平台、无线AP、交换机/路由器、邮件系统、IoT设备等40多种蜜罐服务，支持用户制作自定义Web蜜罐，支持用户进行流量牵引到云蜜网、可开关的扫描感知能力、支持可自定义的蜜饵配置；

+ 开放透明：支持对接微步在线X社区API、五路syslog输出、支持邮件、钉钉、企业威胁、飞书、自定义WebHook告警输出；
+ 快捷管理：支持单个安装包批量部署，支持批量修改端口和服务；

+ 跨平台：支持Linux x32/x64/ARM、Windows x32/x64平台、国产操作系统、龙芯、海光、飞腾、鲲鹏、腾云、兆芯硬件；



## 快速开始

[官方网站](https://hfish.net/)：更多使用蜜罐、使用场景和玩法详见官网

[详细文档](https://hfish.net/docs/#/)：更详细的功能说明、故障排错指南



## 架构

HFish由管理端和节点端组成，管理端用来生成和管理节点端，并接收、分析和展示节点端回传的数据，节点端接受管理端的控制并负责构建蜜罐服务。

> 蜜罐工作原理

![image-20210611130621311](images/20210616174908.png)





> 融合在企业网络中

![image-20210611130733084](images/20210616174930.png)

## 注意

+ Linux 安装无需root权限，但是会导致无法监听低于TCP/1024以下端口

+ 管理端使用 TCP/4433 和 TCP/4434 端口，节点端监听端口根据模拟的服务不同而不同

+ 节点端需要可访问管理端的 TCP/4434 端口，管理端不会主动访问节点端

+ 管理端默认用户名/密码：**`admin / HFish2021`**



## 部署管理端

先部署管理端，再通过管理端的Web页面配置节点端，安装包仅包含管理端和节点端，蜜罐服务包需要部署管理端后从 **`服务管理`** 页面联网下载或离线上传



可联网环境：如果用户的环境允许联网，建议使用以下快速部署步骤：

+ Linux 环境：
  + 在线安装：在shell中运行命令：**`sh | curl https://hfish.net/install.sh`**
  + 离线安装：请访问页面 https://hfish.net/#/2-2-linux?id=%e6%97%a0%e6%b3%95%e8%81%94%e7%bd%91%ef%bc%8c%e6%89%8b%e5%8a%a8%e5%ae%89%e8%a3%85
+ Windows x64 环境：
  + 安装：请访问下载页面 https://hfish.net/#/2-3-windows

离线部署：如果用户为隔离网络环境，请使用以下部署方式



## 配置管理端

+ 浏览器中输入 **`https://server_ip:4433/web/`**，登录管理端

  安装HFish管理端后，默认在管理端所在机器上建立节点感知攻击，该节点被命名为「内置节点」。

  该节点将默认开启部分服务，包括FTP、SSH、Telnet、Zabbix监控系统、Nginx蜜罐、MySQL蜜罐、Redis蜜罐、HTTP代理蜜罐、ElasticSearch蜜罐和通用TCP端口监听。

  ![image-20220127160600755](http://img.threatbook.cn/hfish/image-20220127160600755.png)

  `注意：该节点不能被删除，但可以暂停。`

+ 新增节点：

  - 进入【节点管理】页面，点击【增加节点】

  ![image-20220127160623423](http://img.threatbook.cn/hfish/image-20220127160623423.png)

  - 根据节点设备类型选择对应的安装包和回连地址

    <img src="/Users/maqian/Library/Application Support/typora-user-images/image-20220127160717724.png" alt="image-20220127160717724" style="zoom:50%;" />

  <img src="/Users/maqian/Library/Application Support/typora-user-images/image-20220127160641955.png" alt="image-20220127160641955" style="zoom:50%;" />

  - 在节点机器执行命令语句或安装包，即可成功部署节点。



## 效果图

+ 攻击详情：记录所有对蜜罐的访问请求，包括正常请求、攻击行为、暴力破解

![image2021-6-7_13-57-19](images/20210611114902.png)



+ 扫描详情：记录对所有节点主机的UDP和SYN扫描

![image2021-6-7_17-18-11](images/20210611114934.png)

+ 蜜饵管理

![image2021-6-7_19-6-21](images/20210611115053.png)

+ 节点信息

![image-20220127160421053](http://img.threatbook.cn/hfish/image-20220127160421053.png)

+ 模板管理

<img src="http://img.threatbook.cn/hfish/image-20220127155314024.png" alt="image-20220127155314024" style="zoom:67%;" />



+ 威胁情报对接

![image2021-6-7_19-0-11](images/20210611115158.png)

+ 告警配置

![image2021-6-7_19-4-10](images/20210611115224.png)

## 致谢

## wx群

如何大家有更多的建议希望能够更便捷的交流，可以添加我们的wx群。



![HFish官方群的qr](images/20210611115258.png)

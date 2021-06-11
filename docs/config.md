# 服务端配置

OneFish支持您自定义不同的告警策略。您可以为您不同类型的内容，进行不同方式的通知，以及通知给不同的人。

- 在您配置告警策略之前，您需要先完成上面的【通知配置】



> 添加一个新的策略

![image-20210318173355287](https://hfish.cn-bj.ufileos.com/images/2021-03-18-093357.png)

> 对策略进行配置

![image-20210209202737224](https://qiniu.cuiqingcai.com/2021-03-18-090743.png)

> 对接精准的云端的威胁情报后，可以对攻击行为进行更准的研判，帮助我们更科学的进行处置。

对接了威胁情报后，当OneFish捕获到了来自外网的攻击行为后，我们可以在攻击列表中了解攻击者的IP情报。OneFish会把您在云端查询到的情报在本地缓存3天，保持您攻击情报时效性的同时，节省您的查询次数。

<img src="/Users/water/git/hfish.io/safeset/images/image-20210209191721156.png" alt="image-20210209191721156" style="zoom:50%;" />





- 我们支持对接两种来自微步在线的威胁情报

> 对接微步在线云API（IP信誉接口）

关于该接口完整的说明，可以参考[微步在线云API文档](https://x.threatbook.cn/nodev4/vb4/API)

本接口在注册后可以获得每日50条云端情报的查询额度，给微步发送扩容邮件后，可以提升到每日200条的额度。详情访问[微步在线X社区](https://x.threatbook.cn/nodev4/vb4/article?threatInfoID=3101)。

<img src="/Users/water/git/hfish.io/safeset/images/image-20210209172440082.png" alt="image-20210209172440082" style="zoom:50%;" />



> 对接TIP的本地情报，您可以跟据页面的描述进行注册和使用。

使用该接口需要购买微步在线的TIP本地情报系统

<img src="images/image-20210209192626875.png" alt="image-20210209192626875" style="zoom:50%;" />



> 通知功能是蜜罐的核心功能之一

对于蜜罐捕获到的信息，跟据您不同的安全运营流程，您可能需要把该信息第一时间通知其它的安全设备，也可能需要把该信息通知给相关的安全运营人员。OneFish用三种方式满足您的需求。

- Syslog通知
- 邮件通知
- Webhook通知

<img src="/Users/water/git/hfish.io/safeset/images/image-20210209173435065.png" alt="image-20210209173435065" style="zoom:50%;" />

> 用 Syslog 联动其它安全设备

您可以自定义接受通知设备的地址、协议和端口，用来接受OneFish捕获的攻击信息和报警。OneFish最多支持3路syslog进行通知。



> 用邮件通知相关安全人员

您可以通过配置相关的邮件服务器信息，来接受OneFish的通知和报警。



> Webhook通知其它设备/人

很多的场景下我们都可以方便的使用webhook联动人或者设备。

- 对于当前企业办公中最为流行的3大即时通讯软件企业微信、钉钉、飞书的机器人，我们也做了适配，您在IM中建立一个机器人，把机器人的token复制到OneFish的webhook配置中，就可以第一时间在IM中获取蜜罐捕获的攻击告警了。
- 三家IM的官方文档如下，您可以对照进行参考

```wiki
- 企业微信官方文档

  https://work.weixin.qq.com/help?doc_id=13376#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8%E7%BE%A4%E6%9C%BA%E5%99%A8%E4%BA%BA

- 钉钉官方文档

  https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq

- 飞书官方文档
  https://www.feishu.cn/hc/zh-CN/articles/360040553973
```







> 初次登陆账号密码为

```wiki
admin  /  OneFish2021
```

<img src="/Users/water/git/hfish.io/safeset/images/image-20210209185036616.png" alt="image-20210209185036616" style="zoom:50%;" />



> 登陆后请及时修改您的密码

<img src="/Users/water/git/hfish.io/safeset/images/image-20210209172137133.png" alt="image-20210209172137133" style="zoom:50%;" />


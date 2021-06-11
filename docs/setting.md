# 服务端配置

### 告警策略



> 添加一个新的策略

![image-20210318173355287](https://hfish.cn-bj.ufileos.com/images/2021-03-18-093357.png)



> 对策略进行配置

![image-20210209202737224](https://qiniu.cuiqingcai.com/2021-03-18-090743.png)



> 通知当前分为威胁告警和系统通知两种类型

威胁告警是系统感知攻击时的告警；系统通知是系统自身运行状态的告警。



> 在设置通知方式前，您应该先完成了前边的通知配置

如果您完成了通知配置，那么这里三种不同的通知方式中就会出现您之前的配置，勾选即可。

> 对接精准的云端的威胁情报后，可以对攻击行为进行更准的研判，帮助我们更科学的进行处置。

对接了威胁情报后，当HFish捕获到了来自外网的攻击行为后，我们可以在攻击列表中了解攻击者的IP情报。HFish会把您在云端查询到的情报在本地缓存7天，保持您攻击情报时效性的同时，节省您的查询次数。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318220204897.png" alt="image-20210318220204897" style="zoom:50%;" />



- 我们支持对接两种来自微步在线的威胁情报

> 对接微步在线云API（IP信誉接口）

关于该接口完整的说明，可以参考[微步在线云API文档](https://x.threatbook.cn/nodev4/vb4/API)

本接口在注册后可以获得每日50条云端情报的查询额度，给微步发送扩容邮件后，可以提升到每日200条的额度。详情访问[微步在线X社区](https://x.threatbook.cn/nodev4/vb4/article?threatInfoID=3101)。

如果有企业化需求，可以邮件 honeypot@threatbook.cn



> 对接TIP的本地情报，您可以跟据页面的描述进行注册和使用。

使用该接口需要购买微步在线的TIP本地情报系统

<img src="/Users/maqian/Library/Application%2520Support/typora-user-images/image-20210318221558637.png" alt="image-20210318221558637" style="zoom:50%;" />

> 通知功能是蜜罐的核心功能之一

对于蜜罐捕获到的信息，跟据您不同的安全运营流程，您可能需要把该信息第一时间通知其它的安全设备，也可能需要把该信息通知给相关的安全运营人员。HFish用三种方式满足您的需求。

- Syslog通知
- 邮件通知
- Webhook通知

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318215610629.png" alt="image-20210318215610629" style="zoom:50%;" />

> 用 Syslog 联动其它安全设备

您可以自定义接受通知设备的地址、协议和端口，用来接受OneFish捕获的攻击信息和报警。HFish最多支持5路syslog进行通知。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318215642971.png" alt="image-20210318215642971" style="zoom:50%;" />

> 用邮件通知相关安全人员

您可以通过配置相关的邮件服务器信息，来接受OneFish的通知和报警。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210318215718987.png" alt="image-20210318215718987" style="zoom:50%;" />

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






### 通知配置



![image-20210806094214284](http://img.threatbook.cn/hfish/image-20210806094214284.png)



> 通知功能是蜜罐的核心功能之一

对于蜜罐捕获到的信息，跟据您不同的安全运营流程，您可能需要把该信息第一时间通知其它的安全设备，也可能需要把该信息通知给相关的安全运营人员。HFish用三种方式满足您的需求。

- Syslog通知
- 邮件通知
- Webhook通知



> 用 Syslog 联动其它安全设备

您可以自定义接受通知设备的地址、协议和端口，用来接受HFish捕获的攻击信息和报警。OneFish最多支持3路syslog进行通知。



> 用邮件通知相关安全人员

您可以通过配置相关的邮件服务器信息，来接受HFish的通知和报警。



> Webhook通知其它设备/人

很多的场景下我们都可以方便的使用webhook联动人或者设备。

- 对于当前企业办公中最为流行的3大即时通讯软件企业微信、钉钉、飞书的机器人，我们也做了适配，您在IM中建立一个机器人，把机器人的token复制到HFish的webhook配置中，就可以第一时间在IM中获取蜜罐捕获的攻击告警了。
- 三家IM的对接文档如下，您可以对照进行参考

```wiki
- 企业微信官方文档

  https://work.weixin.qq.com/help?doc_id=13376#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8%E7%BE%A4%E6%9C%BA%E5%99%A8%E4%BA%BA

- 钉钉对接文档

  https://hfish.io/#/6-2-1dingtalk

- 飞书官方文档
  https://www.feishu.cn/hc/zh-CN/articles/360040553973
```


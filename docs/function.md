# 功能说明

### 账号资产

> 用户名密码页面收集了所有被用来攻击的账号密码，可以对企业账号资产有效监控

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210506152344041.png" alt="image-20210506152344041" style="zoom:50%;" />

> 为辅助企业进行内部账号监控，设定高级监测策略，建议输入企业的邮箱、员工姓名、企业名称等信息进行监控，从而随时监控泄漏情况

1.点击界面右上角查看高级监测策略

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210401150526485.png" alt="image-20210401150526485" style="zoom: 50%;" />

2.按照规则要求，导入csv文件。

**注意！务必按照提示规则进行写入**

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210506153037454.png" alt="image-20210506153037454" style="zoom:33%;" />

3.页面可查看到所有匹配高级监测策略的数据，从而帮助运维人员精准排查泄漏账号，实现企业账号资产安全防护。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210506153041469.png" alt="image-20210506153041469" style="zoom:50%;" />

### 主机失陷检测

失陷蜜饵是部署在业务主机上的失陷检测蜜饵。在主机失陷情况下，通过部署虚假的账号、本地证书等失陷蜜饵，诱导攻击者转移攻击目标，并触发失陷告警。

其中，主机蜜饵是一种基于部署虚假的账号密码配置文件，诱导转移攻击者攻击目标的防御手段。

命令在主机运行后，会在本地生成一份虚假的“账号密码备份文件”。 当该主机被攻陷时，攻击者将被诱导，使用文件中的账号信息进行登录。借此，安全人员发现主机失陷情况。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210506162347469.png" alt="image-20210506162347469" style="zoom:50%;" />



![image-20210506162447618](https://hfish.cn-bj.ufileos.com/images/image-20210506162447618.png)



### 恶意IP

> 恶意IP页面将监控所有攻击IP的相关信息，包括微步情报及企业自定义情报。
>
> 另外，所有的溯源信息，最终都会呈现在恶意IP页面，并成为企业的私有情报库。



![image-20210506150145273](https://hfish.cn-bj.ufileos.com/images/image-20210506150145273.png)

### 自定义蜜罐传输协议

针对Web应用仿真、网络设备服务、安全设备服务以及IOT服务，可以根据自身业务场景和网络情况，选择其具体的传输协议（HTTP或者HTTPS），从而让蜜罐更符合当前网络结构，更好吸引攻击者视线。

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210506155628363.png" alt="image-20210506155628363" style="zoom:50%;" />



### WEB蜜罐自定义开发



为了方便企业的定制业务，管理段提供了上传自定义web服务的内容，可根据微步在线的开发规范和原则，自己对web界面进行开发，修改，并上传，使其成为真正的蜜罐服务。

我们为大家准备了一个样例，请先下载我们的web模板样例。

http://threatbook-user-img.cn-bj.ufileos.com/hfish/svc/web-demo.zip



> 1.web蜜罐文件所在目录

```shell
- index.html 
在节点client安装目录./services/service_id/root 下面

- 其它格式的文件
在节点client安装目录./services/service_id/root下的所有目录都可以自行定义、上传文件，用户可以在不同目录下面上传自己的样式文件和图片。
```

> 2.修改页面元素

根据index.html文件中的信息，替换和修改相关的文件。

> 3.制作全新的登陆页面

我们可以自己制作一个全新的登陆页面，通过替换表单元素实现“定制开发”

```shell
- 删除client安装目录./services/service_id/root下所有文件后，自行上传编辑完成的html页面和相关文件

- 修改主页文件名为index.html

- 按照下面图片的要求，修改表单元素。
```

![蜜罐web页面表哥元素](https://hfish.cn-bj.ufileos.com/images/20210406150240.png)



> 4.将修改完成的服务包进行上传，完成web服务添加

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210506162036933.png" alt="image-20210506162036933" style="zoom:50%;" />



<img src="https://hfish.cn-bj.ufileos.com/images/image-20210506162100883.png" alt="image-20210506162100883" style="zoom:50%;" />





最后，如果您希望微步为您进行规范统一开发，请邮件发送给honeypot@threatbook.cn。


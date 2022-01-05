### 高交互蜜罐

当前，HFish提供两种高交互蜜罐，包括**高交互SSH蜜罐**和**高交互Telnet**蜜罐。



#### 高交互蜜罐原理

HFish高交互蜜罐部署在云服务器，由HFish团队做统一运维，在云上，我们部署了多台服务器，通过nginx负载均衡，处理本地到云端的攻击流量转发。

<img src="/Users/maqian/Library/Application Support/typora-user-images/image-20220105214958606.png" alt="image-20220105214958606" style="zoom:33%;" />

在高交互蜜罐中，**攻击者与用户本地发生的流量交互都会被转发到云端蜜网，所有的威胁行为都在云端蜜网发生**。

<img src="/Users/maqian/Library/Application Support/typora-user-images/image-20220105220938586.png" alt="image-20220105220938586" style="zoom:40%;" />



#### 使用高交互蜜罐

当前，云端高交互蜜罐默认配置在服务管理中，使用需要

**节点可联通zoo.hfish.net**

**管理端可联通zoo.hfish.net**

确认网络联通后，直接在节点中添加高交互蜜罐服务即可。

<img src="/Users/maqian/Library/Application Support/typora-user-images/image-20220105221149928.png" alt="image-20220105221149928" style="zoom:42%;" />

<img src="http://img.threatbook.cn/hfish/image-20220105221346398.png" alt="image-20220105221346398" style="zoom: 28%;" />



#### 高交互蜜罐数据查阅

HFish管理端每五分钟向api.hfish.net发起高交互蜜罐请求。在这个过程内的攻击数据将被下拉。

数据中，我们可以看到

- 攻击者进行的登陆行为
- 此次登陆为成功还是失败，
- 登录成功后执行了那些命令

如果在聚合的攻击中存在样本，样本将展示在关联信息下载中。

<img src="http://img.threatbook.cn/hfish/image-20220105221536927.png" alt="image-20220105221536927" style="zoom:50%;" />

#### Custom蜜罐

一个经常容易误解的事实是，Custom蜜罐是用于**接受**其他蜜罐平台数据的接口，其目的是为了实现蜜罐告警数据的统一管理，便于用户通过HFish观测整个企业内部各种蜜罐告警数据。

#### 使用方法

> 1、调用apikey：进入该蜜罐的后台，可以看到蜜罐的config文件，文件内有一个apikey。

<div align="center"><img src="/images/image-20211027204924883.png" alt="" height="400px" /></div>

<div align="center"><img src="/images/image-20211027205006150.png" alt="" height="400px" /></div>


> 2、使用该apikey，将数据上传到HFish管理端

传输方式：post

传输地址：https://ip:8989/api/v1/report

body内使用form格式填写apikey和info，info为您希望打入hfish的其他蜜罐信息。

**参考示例：**

<div align="center"><img src="/images/image-20211027205827645.png" alt="" height="400px" /></div>

<div align="center"><img src="/images/image-20211027205939646.png" alt="" style="zoom:50%;" /></div>

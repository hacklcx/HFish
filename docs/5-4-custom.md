#### Custom蜜罐

Custom蜜罐是我们制作的，可用于接受其他蜜罐平台或者自建蜜罐的数据，实现蜜罐告警数据的统一管理。将HFish做成一个企业内蜜罐数据集合体的蜜罐

#### 使用方法

> 1. 调用apikey：进入该蜜罐的后台，可以看到蜜罐的config文件，文件内有一个apikey。

![image-20211027204924883](http://img.threatbook.cn/hfish/image-20211027204924883.png)

![image-20211027205006150](http://img.threatbook.cn/hfish/image-20211027205006150.png)



> 2.使用该apikey，将数据上传到HFish管理端

传输方式：post

传输地址：https://ip:8989/api/v1/report

body内使用form格式填写apikey和info，info为您希望打入hfish的其他蜜罐信息。

**参考示例：**

<img src="http://img.threatbook.cn/hfish/image-20211027205827645.png" alt="image-20211027205827645" style="zoom:50%;" />

<img src="http://img.threatbook.cn/hfish/image-20211027205939646.png" alt="image-20211027205939646" style="zoom:50%;" />

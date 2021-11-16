
### 什么是蜜饵

蜜饵泛指任意伪造的高价值文件（例如运维手册、邮件、配置文件等），用于引诱和转移攻击者视线，最终达到牵引攻击者离开真实的高价值资产并进入陷阱的目的。

### 使用场景

HFish的蜜饵在 **牵引** 攻击者的功能上增加了 **精确定位失陷** 能力，即每个蜜饵都是 **唯一的**，攻击者入侵用户主机后，如果盗取蜜饵文件中的数据并从任意主机发起攻击，防守者仍能知道失陷源头在哪里。

举个例子：

```wiki
攻击者侵入企业内部某台服务器，在其目录中找到一个payment_config.ini文件，文件中包含数据库主机IP地址和账号密码，
攻击者为隐藏自己真实入侵路径，通过第三台主机访问数据库主机……
```

在以上场景中，payment_config.ini为蜜饵，所谓的数据库主机是另外一台位于安全区域的蜜罐，而攻击者得到的所谓账号密码也是虚假且唯一的，防守者可以根据其得到攻击者真实的横向移动路径。

由于蜜饵只是静态文件，所以蜜饵适合部署在任何主机和场景中，例如作为附件通过邮件发送（检测邮件是否被盗）、在攻防演练期间上传到百度网盘或github上混淆攻击者视线、压缩改名成backup.zip放置在Web目录下守株待兔等待攻击者扫描器上钩……


### HFish的蜜饵

HFish的蜜饵模块由**蜜饵管理**、 **分发接口**和**蜜饵信息**三部分组成，

#### 1.蜜饵管理

HFish支持自定义蜜饵的名字，下发路径和蜜饵文件的具体内容。用户可以按照需求进行自定义蜜饵的设置。



<img src="http://img.threatbook.cn/hfish/2781635152098_.pic_hd.jpg" alt="2781635152098_.pic_hd" style="zoom:33%;" />

在蜜饵内容中，$username$、$password$和$honeypot$分别代表账号、密码和蜜罐变量，以上为必填变量，必须进行引用，才能让蜜饵功能生效。
三个变量，按照文件想呈现给攻击者的效果进行引用。
$username$变量如果未填写账号字典，则默认用root作为所有蜜饵的账号名。
$password$变量按照选取的位数，自动生成蜜饵的密码。
$honeypot$变量按照蜜饵下拉节点的部署服务，自动生成IP和端口。

<img src="/Users/maqian/Library/Containers/com.tencent.xinWeChat/Data/Library/Application Support/com.tencent.xinWeChat/2.0b4.0.9/d4fd4109804e4614bec96b26017afad7/Message/MessageTemp/9e20f478899dc29eb19741386f9343c8/Image/2801635164451_.pic.jpg" alt="2801635164451_.pic" style="zoom:33%;" />

点击预览，可以查看当前的蜜饵内容，在实际被下拉时的显示内容

<img src="http://img.threatbook.cn/hfish/2811635164463_.pic_hd.jpg" alt="2811635164463_.pic_hd" style="zoom: 33%;" />

点击确定，即可新增一条文件蜜饵。

<img src="http://img.threatbook.cn/hfish/2821635164487_.pic_hd.jpg" alt="2821635164487_.pic_hd" style="zoom: 33%;" />



#### 2.分发接口

为了符合企业内的网络情况，蜜饵 **分发接口** 实际位于节点端，启用或禁用开关位于控制端的节点管理页面任意一个节点的详情页面中，默认监听tcp/7878端口，

任何一台节点都可以作为节点分发服务器使用，如下图：

<img src="http://img.threatbook.cn/hfish/image-20211025205553278.png" alt="image-20211025205553278" style="zoom:50%;" />

启用后，用户可以从需要部署蜜饵的主机上访问如下地址，得到一个唯一的蜜饵文件：

<img src="/Users/maqian/Library/Application Support/typora-user-images/image-20211025205700560.png" alt="image-20211025205700560" style="zoom:33%;" />

复制该下发指令后，前往需进行布防的业务机器，执行即可。



#### 3.蜜饵信息

每条蜜饵指令在被执行的时候，都会在管理端生成一条记录。用于主机蜜饵布防情况的管理


当前，修改节点服务可分为两种，

### 直接修改服务

> 点击单个节点，可直接对节点上的服务进行添加和删除

<img src="http://img.threatbook.cn/hfish/image-20210914120052175.png" alt="image-20210914120052175" style="zoom:50%;" />





### 创建模版，应用到多节点

进行模版创建，即可以将一套配置批量应用到多个节点上，其操作步骤如下：

> 进入模版管理，进行模版创建

<img src="http://img.threatbook.cn/hfish/image-20210914115931102.png" alt="image-20210914115931102" style="zoom:25%;" />

> 展开蜜罐节点，选择上面创建的蜜罐模板

<img src="http://img.threatbook.cn/hfish/20210616173018.png" alt="image-20210616173015062" style="zoom: 33%;" />



> 刚变更模板后的蜜罐服务状态为【启用】

<img src="http://img.threatbook.cn/hfish/20210616173055.png" alt="image-20210616173053947" style="zoom: 33%;" />



> 节点正常完成模板加载后，服务状态应该为【在线】。如果是【离线】，说明蜜罐服务没有正常启动，请参考我们后面的【排错说明】，找到问题。

<img src="http://img.threatbook.cn/hfish/20210616173129.png" alt="image-20210616173128526" style="zoom: 33%;" />




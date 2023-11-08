#### 添加/删除节点

#### 内置节点

安装HFish管理端后，在管理端所在机器上会默认建立节点感知攻击，该节点被命名为【内置节点】。

该节点将默认开启部分服务，包括FTP、SSH、Telnet、Zabbix监控系统、Nginx蜜罐、MySQL蜜罐、Redis蜜罐、HTTP代理蜜罐、ElasticSearch蜜罐和通用TCP端口监听。

`注意：该节点不能被删除，但可以暂停。`

<img src="https://hfish.net/images/image-20210902210912371.png" alt="image-20210902210912371" style="zoom:50%;" />


#### 新增节点

> ##### 进入【节点管理】页面，点击【增加节点】

<img src="https://hfish.net/images/image-20210902172749029.png" alt="image-20210902172749029" style="zoom:33%;" />

> ##### 根据节点设备类型选择对应的安装包和回连地址

<img src="https://hfish.net/images/image-20210902172832815.png" alt="image-20210902172832815" style="zoom:33%;" />
<p>
<img src="https://hfish.net/images/image-20210902172916191.png" alt="image-20210902172916191" style="zoom:33%;" />

> ##### 在节点机器执行命令语句或安装包，即可成功部署节点。



#### 删除节点

> ##### 进入【节点管理】页面，在节点列表中找到要删除的节点，点击该节点右侧【删除】，HFish需要二次验证您的管理员身份，输入admin密码后，节点将被删除。

节点被删除后，节点端进程会自动退出，但程序会保留在原有路径，需要手动二次删除，管理端上已收集的关于该节点的所有攻击数据不会丢失，仍然能查看。

> 当前，我们会为所有用户，在Server机器上默认建立一个节点感知攻击。该节点命名为「内置节点」。

该节点将默认开启部分服务，包括FTP、SSH、Telnet、Zabbix监控系统、Nginx蜜罐、MySQL蜜罐、Redis蜜罐、HTTP代理蜜罐、ElasticSearch蜜罐和通用TCP端口监听。

<img src="http://img.threatbook.cn/hfish/image-20210902210912371.png" alt="image-20210902210912371" style="zoom:50%;" />



### 新增节点

> 进入节点管理页面，可进行 新增节点

<img src="http://img.threatbook.cn/hfish/image-20210902172749029.png" alt="image-20210902172749029" style="zoom:33%;" />

> 选择对应的安装包和回连地址

<img src="http://img.threatbook.cn/hfish/image-20210902172832815.png" alt="image-20210902172832815" style="zoom:33%;" />

<img src="http://img.threatbook.cn/hfish/image-20210902172916191.png" alt="image-20210902172916191" style="zoom:33%;" />

> 在节点机器执行命令语句或安装包，即可成功部署节点。

进入「节点管理」页面，对可对该节点进行管理。





### 删除节点

如果您进行节点删除，那么

- 节点端进程会自动退出，但程序会保留在原有路径，需要手动二次删除。
- 所有攻击数据不会丢失，仍然能在所有数据中进行查看

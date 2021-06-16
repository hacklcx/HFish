### WEB蜜罐自定义开发

为了方便企业的定制业务，管理段提供了上传自定义web服务的内容，可根据微步在线的开发规范和原则，自己对web界面进行开发，修改，并上传，使其成为真正的蜜罐服务。

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


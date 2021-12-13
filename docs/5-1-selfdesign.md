#### 通过样例了解功能实现方式

> web蜜罐样例

```wiki
# 下载自定义web蜜罐样例
http://img.threatbook.cn/hfish/svc/service-demo.zip
```

> 解压后获得index.html和portrait.js两个文件


> index.html文件中的代码功能

```wiki
<form>中的代码明确了页面上账密表单的提交方式，具体利用方式参考下文“制作全新的登陆页面”

<script>中的代码明确了调用jsonp的方式
```

> portrait.js 文件中的代码功能

```wiki
这个文件是jsonp溯源功能的利用代码，攻击者在已登录其他社交平台的情况下，蜜罐可以获得部分社交平台的账号信息。

本代码因为利用了浏览器的漏洞，有一定的时效性，随着攻击者更新自己的浏览器，利用代码可能失效。并有可能让攻击者在访问该页面时，触发杀毒软件的报警。

在利用代码失效后，您可以选择删除index.html中的利用代码。

同时请关注我们的官网（ https://hfish.io ）和x社区( https://x.threatbook.cn )，等待我们和社区用户更新漏洞利用代码，并替换本文件和index内的利用代码，恢复溯源能力。
```



### 制作全新的登陆页面

我们可以自己制作一个全新的登陆页面，通过替换表单元素实现“定制开发”

```shell
- 修改主页文件名为index.html

- 按照下面图片的要求，修改表单元素。
```

![蜜罐web页面表格元素](http://img.threatbook.cn/hfish/20210728213641.png)



### 打包并上传到蜜罐的管理后台

> 打包所有静态文件资源

把所有的静态文件文件打包名为“service-xxx.zip”文件。包括但不限于index.html 、portrait.js 和其它格式的静态文件、文件夹。

注意：文件命名为规范格式前缀 **必须** 为“service-” ； “xxx”可以自定义，但不能为“web”和“root”，且压缩包 **必须** 压缩为.zip格式文件。

![image-20210508222121613](http://img.threatbook.cn/hfish/20210728213740.png)



> 打开管理端后台

![image-20210508213915879](http://img.threatbook.cn/hfish/20210728213815.png)



> 配置新增服务页面

![image-20210508221316072](http://img.threatbook.cn/hfish/20210728213852.png)



> 自定义web蜜罐添加成功


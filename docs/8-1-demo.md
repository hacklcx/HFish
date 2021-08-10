

### SSH蜜罐

>  在终端里，尝试连接蜜罐的ssh端口，会显示“Permission denied, please try again.”

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113406672.png" alt="image-20210319113406672" style="zoom: 50%;" />

> 这时攻击列表会记录下所有测试过的用户名和密码

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113330040.png" alt="image-20210319113330040" style="zoom: 33%;" />



### FTP蜜罐

> 用FTP终端尝试连接FTP蜜罐端口，会在攻击列表中出现FTP蜜罐报警

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113309227.png" alt="image-20210319113309227" style="zoom: 33%;" />



### HTTP蜜罐

> HTTP蜜罐为http代理蜜罐，利用http代理工具连接蜜罐端口

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113242516.png" alt="image-20210319113242516" style="zoom: 33%;" />

> 攻击列表中的显示信息如下

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113211933.png" alt="image-20210319113211933" style="zoom: 33%;" />



### TELNET蜜罐

> 利用TELNET应用连接蜜罐端口

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113132880.png" alt="image-20210319113132880" style="zoom: 33%;" />

> 攻击列表中显示信息如下

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113101608.png" alt="image-20210319113101608" style="zoom: 33%;" />



### MYSQL蜜罐

> 用MYSQL工具连接蜜罐对应端口，可输入指令。

<img src="http://img.threatbook.cn/hfish/1521628589153_.pic_hd.jpg" alt="1521628589153_.pic_hd" style="zoom:50%;" />



> 在管理端，可以看到攻击者的攻击记录以及其本地的etc/group信息

<img src="http://img.threatbook.cn/hfish/1531628589485_.pic_hd.jpg" alt="1531628589485_.pic_hd" style="zoom: 33%;" />



### WEB蜜罐

WEB蜜罐用浏览器访问相应的端口，并尝试输入【用户名】和【密码】后

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319112136937.png" alt="image-20210319112136937" style="zoom:50%;" />

会提示用户名和密码错误

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319112522399.png" alt="image-20210319112522399" style="zoom:50%;" />



服务端后台会获取攻击者用于尝试的用户名和密码

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319112739513.png" alt="image-20210319112739513" style="zoom: 33%;" />



### VNC蜜罐

> 通过VNC viewer进行登陆常识，输入IP和端口

<img src="http://img.threatbook.cn/hfish/1591628590040_.pic_hd.jpg" alt="1591628590040_.pic_hd" style="zoom: 33%;" />

<img src="http://img.threatbook.cn/hfish/1611628590115_.pic_hd.jpg" alt="1611628590115_.pic_hd" style="zoom:33%;" />



> 在管理端，可以看到相关攻击记录

<img src="http://img.threatbook.cn/hfish/image-20210810195923459.png" alt="image-20210810195923459" style="zoom:33%;" />



### REDIS蜜罐

> 使用redis命令，远程登录redis

<img src="http://img.threatbook.cn/hfish/image-20210810200645587.png" alt="image-20210810200645587" style="zoom: 33%;" />

> 进入管理端，可以看到攻击详情

<img src="http://img.threatbook.cn/hfish/1641628594371_.pic_hd.jpg" alt="1641628594371_.pic_hd" style="zoom: 33%;">



### MEMCACHE蜜罐

> 通过Telnet（或其他方式）尝试连接MEMCACHE

<img src="http://img.threatbook.cn/hfish/image-20210810202312677.png" alt="image-20210810202312677" style="zoom: 33%;" />

> 进入管理端，可以查看攻击详情

<img src="http://img.threatbook.cn/hfish/1741628595751_.pic_hd.jpg" alt="1741628595751_.pic_hd" style="zoom: 33%;" />





### Elasticsearch蜜罐

> 通过IP和端口，可以登陆查看Elasticsearch蜜罐

<img src="http://img.threatbook.cn/hfish/1701628595182_.pic.jpg" alt="1701628595182_.pic" style="zoom:50%;" />、



> 进入管理端，可以看到攻击的请求详情

<img src="http://img.threatbook.cn/hfish/1721628595216_.pic_hd.jpg" alt="1721628595216_.pic_hd" style="zoom: 33%;" />






### SSH蜜罐

>  在终端里，尝试连接蜜罐的ssh端口，会显示“Permission denied, please try again.”

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113406672.png" alt="image-20210319113406672" style="zoom:50%;" />

> 这时攻击列表会记录下所有测试过的用户名和密码

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113330040.png" alt="image-20210319113330040" style="zoom:50%;" />



### FTP蜜罐

> 用FTP终端尝试连接FTP蜜罐端口，会在攻击列表中出现FTP蜜罐报警

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113309227.png" alt="image-20210319113309227" style="zoom:50%;" />



### HTTP蜜罐

> HTTP蜜罐为http代理蜜罐，利用http代理工具连接蜜罐端口

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113242516.png" alt="image-20210319113242516" style="zoom:50%;" />

> 攻击列表中的显示信息如下

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113211933.png" alt="image-20210319113211933" style="zoom:50%;" />

### TELNET蜜罐

> 利用TELNET应用连接蜜罐端口

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113132880.png" alt="image-20210319113132880" style="zoom:50%;" />

> 攻击列表中显示信息如下

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319113101608.png" alt="image-20210319113101608" style="zoom:50%;" />

### MYSQL蜜罐

> 用MYSQL工具连接蜜罐对应端口

![image-20210311174259758](https://hfish.cn-bj.ufileos.com/images/20210311174259.png)

### WEB蜜罐

WEB蜜罐用浏览器访问相应的端口，并尝试输入【用户名】和【密码】后

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319112136937.png" alt="image-20210319112136937" style="zoom:50%;" />

会提示用户名和密码错误

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319112522399.png" alt="image-20210319112522399" style="zoom:50%;" />



服务端后台会获取攻击者用于尝试的用户名和密码

<img src="https://hfish.cn-bj.ufileos.com/images/image-20210319112739513.png" alt="image-20210319112739513" style="zoom: 50%;" />


- ### 手动安装

  如果上述的安装脚本您无法使用，您可以尝试用手动安装完成部署。
  
  到官网 [https://hfish.io](https://hfish.io/) 下载HFish最新版本安装包，按如下步骤进行安装 （以linux64位系统为例）：
  
  > 第一步： 在当前目录创建一个路径解压安装包
  
  ```
  mkdir hfish
  ```
  
  > 第二步：将安装文件包解压到hfish目录下
  
  ```
  tar zxvf hfish-*-linux-amd64.tar.gz -C hfish
  ```
  
  > 第三步：请防火墙开启4433或者4434，确认返回success（如果有其他服务需要打开端口，使用相同命令打开。
  
  ```
  firewall-cmd --add-port=4433/tcp --permanent
  firewall-cmd --add-port=4434/tcp --permanent
  firewall-cmd --reload
  ```
  
  > 第四步：进入安装目录直接运行server，或者后台运行 nohup ./server &
  
  ```
  cd hfish
  nohup ./server &
  ```
  
  > 第五步：登陆web界面
  
  ```
  登陆链接：https:// [ip]:4433/web
  账号：admin
  密码：HFish2021
  ```
  
  例：如果控制端的ip是192.168.1.1，登陆链接为：https://192.168.1.1:4433/web
  
  控制端部署完成后，请继续参考下面的【控制端配置】完成配置


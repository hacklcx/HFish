![image-20210730154355445](http://img.threatbook.cn/hfish/20210730154357.png)

> 扫描感知可以感知到针对蜜罐节点的扫描行为

扫描感知通过对网卡抓包，可以感知到针对该节点全端口的扫描行为。支持TCP、UDP和ICMP扫描类型。



- 目前扫描感知列表内能够展示的信息如下：
  1. 扫描IP
  2. 威胁情报
  3. 被扫描节点
  4. 被扫描IP
  5. 扫描类型
  6. 被扫描端口
  7. 节点位置
  8. 扫描开始时间
  9. 扫描持续时间
  
  

> **注意！Windows节点的扫描感知依赖WinPcap，需要手动进行下载安装！**

WinPcap官方链接：https://www.winpcap.org/install/bin/WinPcap_4_1_3.exe

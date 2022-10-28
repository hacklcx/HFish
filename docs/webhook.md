### Syslog外发信息字段

#### 威胁告警外发字段：

**HFish-威胁告警**

```
client                //节点名称
client_ip             //节点IP
attack_type           //攻击类型，包含scan/attack/signon/hr_signon/compromise（扫描/攻击/登陆/高危登陆/失陷）
scan_type             //扫描类型（udp/tdp/icmp)
scan_port             //扫描端口 (为空时候为N/A）
type                  //攻击蜜罐 
class                 //蜜罐类型
account               //账号信息
src_ip                //攻击来源IP
labels                //威胁情报标签
dst_ip                //受害IP
geo                   //攻击来源ip的地理位置
time                  //攻击发生时间
threat_name: aaa,bbb,ccc //威胁行为名称
threat_level: high       //威胁行为等级
info                  //攻击详情(为空填写N/A）
```

#### 系统告警外发字段

**HFish-节点离线告警**

```
HFish-节点离线告警    //告警标题
Server_ip           //HFish系统IP
client              //离线节点名称
client_ip           //离线节点IP
```

**HFish-蜜罐离线告警**

```
title             //告警标题
Server_ip         //HFish系统IP
client            //节点名称
client_ip         //节点IP
class             //离线蜜罐名称（如果监测到多个蜜罐离线，多个蜜罐在一条记录中发送）
```



 

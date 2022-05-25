威胁检测

#### 检测功能简介

**适用情况：**

本地默认自带规则，通过设置自定义规则，可对关注的重点攻击行为做实时监控。在重大保障/攻防演练中，当有0day、N day等漏洞爆出，可自定义添加对应漏洞的yara规则，重点监控。另，威胁检测引擎会发现触发这些规则的告警，这些告警可帮助判断针对性攻击。

**检测数据：**

HFish的威胁检测引擎是基于攻击HFish的日志数据进行检测。日志原始数据可以查看攻击列表

这些数据包括有web蜜罐的url数据、UA等数据，高交互蜜罐内执行的命令等数据，普通蜜罐的记录日志。

原始数据参考可通过攻击列表下载查看：

<img src="http://img.threatbook.cn/hfish/image-20220525175235039.png" alt="image-20220525175235039" style="zoom:50%;" />

<img src="http://img.threatbook.cn/hfish/image-20220525175429375.png" alt="image-20220525175429375" style="zoom:50%;" />



#### 使用介绍

通过增加威胁检测引擎 ，支持规则检测，且支持上传自定义yara规则，系统自动化分析数据，判定攻击行为，丰富IP威胁判定，更清晰的感受攻击态势和攻击情况。

<img src="http://img.threatbook.cn/hfish/image-20220525180224940.png" alt="image-20220525180224940" style="zoom:50%;" />



#### 告警查看

1）点击列表中的【命令记录】下方的数字即可跳转到对应攻击

<img src="http://img.threatbook.cn/hfish/image-20220525221605474.png" alt="image-20220525221605474" style="zoom:50%;" />

2）也可在【威胁感知】-【攻击列表】进行搜索匹配到规则的告警，只有匹配过的告警记录，筛选条件【攻击行为类型】才有选项。

<img src="/Users/maqian/Library/Application Support/typora-user-images/image-20220525221529700.png" alt="image-20220525221529700" style="zoom:50%;" />



#### 新增自定义规则

1、点击右上角【新增检测规则】

<img src="http://img.threatbook.cn/hfish/image-20220525180831991.png" alt="image-20220525180831991" style="zoom:50%;" />

2、撰写新增检测规则。填写策略名称、描述、严重级别、yara规则。另最下方支持工具帮助测试写好的yara规则是否能对当前日志正确检出

<img src="http://img.threatbook.cn/hfish/image-20220525180854591-20220525180925867.png" alt="image-20220525180854591" style="zoom: 33%;" />




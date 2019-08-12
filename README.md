![dashboard.png](./images/dashboard.png)

# 介绍

> *本 Team 研发此平台，仅为企业安全测试使用，禁止其他人员使用非法用途！一切行为与本 Team 无关。*

**HFish** 是一款基于 Golang 开发的跨平台多功能主动诱导型蜜罐框架系统，为了企业安全防护测试做出了精心的打造


- 多功能 不仅仅支持 HTTP(S) 蜜罐，还支持 SSH、SFTP、Redis、Mysql、FTP、Telnet、暗网 等
- 扩展性 提供 API 接口，使用者可以随意扩展蜜罐模块 ( WEB、PC、APP )
- 便捷性 使用 Golang 开发，使用者可以在 Win + Mac + Linux 上快速部署一套蜜罐平台

# 地址

## Github

- Git: https://github.com/hacklcx/HFish
- Download: https://github.com/hacklcx/HFish/releases

## 码云(Gitee)

- Git: https://gitee.com/lauix/HFish
- Download: https://gitee.com/lauix/HFish/releases

# 快速部署

## 部署说明

- 下载当前系统二进制包
- cd 到程序根目录，修改 config.ini 配置文件
- 执行 ./HFish run 启动服务
- 浏览器输入 http://localhost:9001 打开

## 集群部署

- 复制 HFish、config.ini、web(不启动WEB蜜罐可以不复制) 目录文件到服务器上
- 修改 config.ini -> rpc -> status 为 2
- 修改 config.ini -> rpc -> addr  地址为 HFish 服务端地址

## 命令行帮助

![help.png](./images/help.png)

## 启动服务

![run.png](./images/run.png)

# 部分界面展示

## 登录页

![login.png](./images/login.png)

## 上钩页

![fish.png](./images/fish.png)

## 分布式集群

![colony.png](./images/colony.png)

## 邮件群发

![mail.png](./images/mail.png)

# 部分功能使用演示

## WEB 蜜罐

![web.png](./images/web.png)

## SSH 蜜罐

![ssh.png](./images/ssh.png)

## Redis 蜜罐

![redis.png](./images/redis.png)

## Mysql 蜜罐

![mysql.png](./images/mysql.png)

## FTP 蜜罐

![ftp.png](./images/ftp.png)

## Telnet 蜜罐

![telnet.png](./images/telnet.png)

# 注意事项

- 邮箱 SMTP 配置后需要开启方可使用
- API 接口 info 字段，&& 为换行符
- 启动 WEB 蜜罐，请先启动 API 模块
- WEB 插件 需在 WEB 目录下 编写
- WEB 插件 下面必须存在两个目录
- 集群 心跳为60秒,断开显示会延迟60秒
- 暗网蜜罐是支持的，但是目前Tor服务网上找不到，无法提供演示

# API 接口

## WEB 蜜罐

```
URL: http://localhost:9001/api/v1/post/report

POST：

    name    :   WEB管理后台蜜罐                     # 项目名
    info    :   admin&&12345                      # 上报信息，&& 为换行符号
    sec_key :   9cbf8a4dcb8e30682b927f352d6559a0  # API 安全密钥

特殊说明：

    URL api/v1/post/report 可在 config.ini 配置里修改
    sec_key 可在 config.ini 配置里修改，修改后 WEB 模板也需要同时修改
```

## 暗网 蜜罐

```
URL: http://localhost:9001/api/v1/post/deep_report

POST：

    name    :   暗网后台蜜罐                        # 项目名
    info    :   admin&&12345                      # 上报信息，&& 为换行符号
    sec_key :   9cbf8a4dcb8e30682b927f352d6559a0  # API 安全密钥

特殊说明：

    URL api/v1/post/deep_report 可在 config.ini 配置里修改
    sec_key 可在 config.ini 配置里修改，修改后 暗网 模板也需要同时修改
```

## 黑名单IP

```
URL(Get): http://localhost:9001/api/v1/get/ip

特殊说明：

    提供此接口为了配合防火墙使用，具体方案欢迎来讨论！
```

# TODO

- [x] 登录模块
- [x] 仪表盘模块
- [x] 上钩列表
- [x] 邮件群发
- [x] 命令行优化
- [x] 支持自定义 WEB 模板
- [x] 支持 Mysql 服务端获取连接客户端电脑任意文件
- [x] 支持 HTTP(S)、SSH、SFTP、Redis、Mysql、FTP、Telnet、暗网 蜜罐
- [x] 日记完善优化
- [x] 支持分布式架构
- [x] 支持分页
- [x] 支持 ip 地理信息
- [x] 提供黑名单IP接口
- [ ] 支持 SMTP、POP3、TFTP、Oracle、VPN 等
- [ ] WIFI 蜜罐支持
- [ ] 自动化蜜罐支持
- [ ] 蜜罐报告生成
- [ ] 邮件发送支持编辑器
- [ ] 支持邮件模板选择
- [ ] 蜜罐高交互完善
- [ ] 支持 Ngrok 一键映射
- [ ] 支持更多的图表统计
- [ ] Mysql 支持
- [ ] 规划更多的功能...

# 关于

- Team: HackLC
- URL: https://hack.lc

# 反馈群

加微信拉人，请备注 **HackLC**

![wechat.png](./images/wechat.jpg)

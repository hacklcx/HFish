#### 忘记密码

如果用户不慎忘记HFish Web管理端admin密码，可以按照如下步骤强行重置。

> ##### Linux #####

1、进入HFish管理端安装目录，执行./tools -mode resetpwd
2、kill hfish-server 等待几秒钟，进程会被自动拉起
3、使用默认账号密码：admin/HFish2021进行登录



> ##### Windows #####

1. 运行cmd，进入HFish管理端安装目录，进入当前版本目录，执行tools.exe -mode resetpwd  
2. 打开任务管理器，结束hfish-server进程，等待几秒钟，进程会被自动拉起
3. 使用默认账号密码：admin/HFish2021进行登录



> ##### Docker  #####

1. docker exec -it hfish /bin/sh
2. cd 3.0.1(版本号）
3. 执行./tools -mode resetpwd
4. exit  退出容器
5. docker restart hfish
6. 使用默认账号密码：admin/HFish2021进行登录

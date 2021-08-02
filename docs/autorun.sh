# /bash/sh

if [ -n $(find /bin /usr/bin -name "systemctl") ]; then
    if [ -f ./server ]; then
    DESC=HFish-server
    RUN="./server"
    elif [ -f ./client ]; then
    DESC=HFish-client
    RUN="./client"
    else
    echo "Error! No Exist Program，请在HFish的程序目录下运行，或参阅 https://hfish.io 官网文档手动配置！\n" && exit 1
	fi

    if [ $(ps -ef | grep ${RUN} | grep -v grep | wc -l) -gt 0 ]; then
    ps -ef | grep ${RUN} | grep -v grep | awk '{print $2}' | xargs kill
    fi

	rm -rf /etc/systemd/system/${DESC}.service
	
    echo "[Unit]" >> /etc/systemd/system/${DESC}.service
    echo "Description=${DESC}" >> /etc/systemd/system/${DESC}.service
    echo "After=network.target" >> /etc/systemd/system/${DESC}.service
    echo "Wants=mariadb.service syslog.target remote-fs.target \n" >> /etc/systemd/system/${DESC}.service
    echo "[Service]" >> /etc/systemd/system/${DESC}.service
    echo "Type=simple" >> /etc/systemd/system/${DESC}.service
    echo "ExecStart=/bin/bash -c 'cd $(pwd) && ${RUN}'" >> /etc/systemd/system/${DESC}.service
    echo "ExecReload=/usr/bin/kill -s HUP $MAINPID" >> /etc/systemd/system/${DESC}.service
    echo "ExecStop=/usr/bin/kill -s QUIT $MAINPID" >> /etc/systemd/system/${DESC}.service
    echo "Restart=on-failure" >> /etc/systemd/system/${DESC}.service
    echo "RestartSec=30 \n" >> /etc/systemd/system/${DESC}.service
    echo "[Install]" >> /etc/systemd/system/${DESC}.service
    echo "WantedBy=multi-user.target" >> /etc/systemd/system/${DESC}.service

    systemctl daemon-reload
    systemctl start ${DESC}

else
echo "未发现systemctl程序，服务脚本无法工作，请参阅 https://hfish.io 官网文档手动配置！\n" && exit 1
fi
#!/bin/bash

#初始化
initVar() {
	installType='yum -y install'
	removeType='yum -y remove'
	upgrade="yum -y update"
	echoType='echo -e'
	version='2.5.0'
}
initVar
export LANG=en_US.UTF-8

#字体颜色
echoContent() {
	case $1 in
	# 红色
	"red")
		# shellcheck disable=SC2154
		${echoType} "\033[31m${printN}$2 \033[0m"
		;;
		# 天蓝色
	"skyBlue")
		${echoType} "\033[1;36m${printN}$2 \033[0m"
		;;
		# 绿色
	"green")
		${echoType} "\033[32m${printN}$2 \033[0m"
		;;
		# 白色
	"white")
		${echoType} "\033[37m${printN}$2 \033[0m"
		;;
	"magenta")
		${echoType} "\033[31m${printN}$2 \033[0m"
		;;
		# 黄色
	"yellow")
		${echoType} "\033[33m${printN}$2 \033[0m"
		;;
	esac
}


#首页菜单
menu() {
	echoContent red "\n==============================================================\n"
	echoContent green "当前版本：v${version}"
	echoContent green "HFish官网 https://hfish.io "
	echoContent red "\n==============================================================\n"
	echoContent skyBlue "-------------------------安装部署-----------------------------\n"
	echoContent yellow "1.安装并运行HFish单机版"
	echoContent yellow "2.安装并运行HFish集群版控制端"
	echoContent yellow "3.退出安装"
	# echoContent yellow "4.用Docker运行HFish控制端"
	# echoContent skyBlue "\n-------------------------配置管理-----------------------------\n"
	# echoContent yellow "5.防火墙放通控制端端口（coming soon）"
	# echoContent yellow "6.将HFish添加为系统服务（coming soon）"
	# echoContent yellow "7.将控制端数据库替换为MariaDB（coming soon）"
    # echoContent skyBlue "\n-------------------------运维管理-----------------------------\n"
	# echoContent yellow "8.将错误日志反馈给开发者（coming soon）"
	# echoContent yellow "9.卸载HFish（coming soon）"
	echoContent red "\n=============================================================="

	read -r -p "请选择:" selectMenuType
	case ${selectMenuType} in
    1):
        standaloneInstall
        ;;
	2):
		serverInstall
		;;
	3)
		exitInstall
		;;
	*)
		echoContent red ' ---> 选择错误，重新选择'
		selectMenuType
		;;
	esac
}

standaloneInstall(){
  cd /opt
    if [ $(uname -s) = 'Linux' ] && [ $(uname -m) = 'x86_64' ] && [ $(getconf LONG_BIT) = '64' ]; then
    wget -N --no-check-certificate http://hfish.cn-bj.ufileos.com/hfish-standalone-${version}-linux-amd64.tar.gz
	elif [ $(uname -m) = 'aarch64' ] && [ $(getconf LONG_BIT) = '64' ]; then
    wget -N --no-check-certificate http://hfish.cn-bj.ufileos.com/hfish-standalone-${version}-linux-arm64.tar.gz
	else
    echoContent red "未检测到系统版本，请参阅 https://hfish.io 官网文档手动安装！\n" && exit 1
	fi
	
	tar -zxvf /opt/hfish-standalone*.tar.gz
	cd /opt/hfish && nohup ./server &
	sleep 2
    cd /opt/hfish/client && nohup ./client &
}

serverInstall() {
  cd /opt
	if [ $(uname -s) = 'Linux' ] && [ $(uname -m) = 'x86_64' ] && [ $(getconf LONG_BIT) = '64' ]; then
    wget -N --no-check-certificate http://hfish.cn-bj.ufileos.com/hfish-${version}-linux-amd64.tar.gz
	elif [ $(uname -m) = 'aarch64' ] && [ $(getconf LONG_BIT) = '64' ]; then
    wget -N --no-check-certificate http://hfish.cn-bj.ufileos.com/hfish-${version}-linux-arm64.tar.gz
	else
    echoContent red "未检测到系统版本，请参阅 https://hfish.io 官网文档手动安装！\n" && exit 1
	fi

	mkdir -p hfish
	tar -zxvf /opt/hfish*.tar.gz -C hfish
	cd hfish
	nohup ./server &
}

exitInstall() {
	exit 1
}

# selectServiceInstall() {
# 	if [ -d "/opt/hfish/packages" ]; then
#         cd /opt/hfish/packages
#         wget http://img.threatbook.cn/hfish/svc/services-2.4.0.tar.gz
#         tar zxvf services*.tar.gz
#         rm -f services-2.4.0.tar.gz
#     else 
#         echoContent red "未检测到安装目录，请参阅 https://hfish.io 官网文档手动安装！\n" && exit 1
#     fi 
# }
cd /opt
menu

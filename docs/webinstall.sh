#!/bin/bash

#初始化
initVar() {
	installType='yum -y install'
	removeType='yum -y remove'
	upgrade="yum -y update"
	echoType='echo -e'
	version='2.7.0'
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
    echoContent red " _   _   _____   _         _     " 
    echoContent red "| | | | |  ___| (_)  ___  | |__  "
    echoContent red "| |_| | | |_    | | / __| | '_ \ "
    echoContent red "|  _  | |  _|   | | \__ \ | | | |"
    echoContent red "|_| |_| |_|     |_| |___/ |_| |_|"
    echoContent green "当前版本：v${version}"
    echoContent green "https://hfish.io \n"
    echoContent red "----------"
    echoContent yellow "按 1 安装并运行HFish"
    echoContent yellow "按 2 退出脚本"
	# echoContent skyBlue "\n-------------------------配置管理-----------------------------\n"
	# echoContent yellow "5.防火墙放通控制端端口（coming soon）"
	# echoContent yellow "6.将HFish添加为系统服务（coming soon）"
    # echoContent skyBlue "\n-------------------------运维管理-----------------------------\n"
	# echoContent yellow "8.将错误日志反馈给开发者（coming soon）"
	# echoContent yellow "9.卸载HFish（coming soon）"
	echoContent red "----------"

	while [ 1 ]; do
		read -r -p "请选择:" selectMenuType

		case ${selectMenuType} in
		1):
			serverInstall
			;;
		2)
			exitInstall
			;;
		*)
			continue
			;;
		esac
		break
	done
}

serverInstall() {
  cd /opt
	if [ $(uname -s) = 'Linux' ] && [ $(uname -m) = 'x86_64' ] && [ $(getconf LONG_BIT) = '64' ]; then
    curl -k http://hfish.cn-bj.ufileos.com/hfish-${version}-linux-amd64.tgz -o hfish-${version}-linux-amd64.tgz
	elif [ $(uname -m) = 'aarch64' ] && [ $(getconf LONG_BIT) = '64' ]; then
    curl -k http://hfish.cn-bj.ufileos.com/hfish-${version}-linux-arm64.tgz hfish-${version}-linux-arm64.tgz
	else
    echoContent red "未检测到系统版本，请参阅 https://hfish.io 官网文档手动安装！\n" && exit 1
	fi

	mkdir -p hfish
	tar -zxvf /opt/hfish*.tgz -C hfish
	cd hfish
	sudo ./install.sh
}

exitInstall() {
	exit 1
}

cd /opt
menu

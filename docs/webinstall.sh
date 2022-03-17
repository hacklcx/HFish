#!/bin/bash

#init
initVar() {
	installType='yum -y install'
	removeType='yum -y remove'
	upgrade="yum -y update"
	echoType='echo -e'
	version='2.9.1'
}
initVar
export LANG=en_US.UTF-8

#the color of font
echoContent() {
	case $1 in
	"red")
		# shellcheck disable=SC2154
		${echoType} "\033[31m${printN}$2 \033[0m"
		;;
	"skyBlue")
		${echoType} "\033[1;36m${printN}$2 \033[0m"
		;;
	"green")
		${echoType} "\033[32m${printN}$2 \033[0m"
		;;
	"white")
		${echoType} "\033[37m${printN}$2 \033[0m"
		;;
	"magenta")
		${echoType} "\033[31m${printN}$2 \033[0m"
		;;
	"yellow")
		${echoType} "\033[33m${printN}$2 \033[0m"
		;;
	esac
}


#start
menu() {
	echoContent red " _   _   _____   _         _     " 
	echoContent red "| | | | |  ___| (_)  ___  | |__  "
	echoContent red "| |_| | | |_    | | / __| | '_ \ "
	echoContent red "|  _  | |  _|   | | \__ \ | | | |"
	echoContent red "|_| |_| |_|     |_| |___/ |_| |_|  v${version}"
	echoContent green "https://hfish.io\n\n"
	echoContent white "Press 1 : Install and run HFish"
	# echoContent yellow "Press 2 : Add mgmt ports to the firewall（coming soon）"
	# echoContent yellow "Press 3 : Install as service（coming soon）"
	# echoContent yellow "Press 4 : Post error logfile to HFish team（coming soon）"
	# echoContent yellow "Press 9 : Uninstall HFish（coming soon）"
	echoContent white "Press 0 : Exit"
	echoContent white "----------"

	while [ 1 ]; do
		read -r -p "Input: " selectMenuType

		case ${selectMenuType} in
		1):
			serverInstall
			;;
		0)
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
	else
		echoContent red "No OS version is detected. Please refer to https://hfish.io for manual installation\n" && exit 1
	fi

	mkdir -p hfish
	tar -zxvf /opt/hfish-${version}*.tgz -C hfish
	cd hfish
	sudo ./install.sh
}

exitInstall() {
	exit 1
}

cd /opt
menu

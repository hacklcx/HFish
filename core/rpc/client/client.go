package client

import (
	"HFish/utils/log"
	"HFish/utils/conf"
	"HFish/core/rpc/core"
	"strings"
	"fmt"
)

// 上报状态结构
type Status struct {
	AgentIp                                                                       string
	AgentName                                                                     string
	Web, Deep, Ssh, Redis, Mysql, Http, Telnet, Ftp, MemCahe, Plug, ES, TFtp, Vnc string
}

// 上报结果结构
type Result struct {
	AgentIp     string
	AgentName   string
	Type        string
	ProjectName string
	SourceIp    string
	Info        string
	Id          string // 数据库ID，更新用 0 为新插入数据
}

var rpcClient *rpc.Client
var ipAddr string

func RpcInit() {
	rpcAddr := conf.Get("rpc", "addr")
	c, conn, err := rpc.Dial("tcp", rpcAddr)

	if err != nil {
		rpcClient = nil
		log.Pr("RPC", "127.0.0.1", "连接 RPC Server 失败")
		ipAddr = ""
	} else {
		rpcClient = c
		ipAddr = strings.Split(conn.LocalAddr().String(), ":")[0]
		fmt.Println("连接RPC Server 成功")
	}
}

func reportStatus(rpcName string, ftpStatus string, telnetStatus string, httpStatus string, mysqlStatus string, redisStatus string, sshStatus string, webStatus string, darkStatus string, memCacheStatus string, plugStatus string, esStatus string, tftpStatus string, vncStatus string) {
	if (rpcClient != nil) {
		status := Status{
			ipAddr,
			rpcName,
			webStatus,
			darkStatus,
			sshStatus,
			redisStatus,
			mysqlStatus,
			httpStatus,
			telnetStatus,
			ftpStatus,
			memCacheStatus,
			plugStatus,
			esStatus,
			tftpStatus,
			vncStatus,
		}

		var reply string
		err := rpcClient.Call("HFishRPCService.ReportStatus", status, &reply)

		if err != nil {
			log.Pr("RPC", "127.0.0.1", "上报服务状态失败", err)
			RpcInit()
		} else {
			fmt.Println("上报服务状态成功")
		}
	} else {
		RpcInit()
	}
}

func ReportResult(typex string, projectName string, sourceIp string, info string, id string) string {
	var reply string

	if (rpcClient != nil) {
		// projectName 只有 WEB 才需要传项目名 其他协议空即可
		// id 0 为 新插入数据，非 0 都是更新数据
		// id 非 0 的时候 sourceIp 为空
		rpcName := conf.Get("rpc", "name")

		result := Result{
			ipAddr,
			rpcName,
			typex,
			projectName,
			sourceIp,
			info,
			id,
		}

		err := rpcClient.Call("HFishRPCService.ReportResult", result, &reply)

		if err != nil {
			log.Pr("RPC", "127.0.0.1", "上报上钩结果失败")
			RpcInit()
		} else {
			fmt.Println("上报上钩结果成功")
		}

		return reply
	} else {
		return "0"
	}
}

func Start(rpcName string, ftpStatus string, telnetStatus string, httpStatus string, mysqlStatus string, redisStatus string, sshStatus string, webStatus string, darkStatus string, memCacheStatus string, plugStatus string, esStatus string, tftpStatus string, vncStatus string) {
	reportStatus(rpcName, ftpStatus, telnetStatus, httpStatus, mysqlStatus, redisStatus, sshStatus, webStatus, darkStatus, memCacheStatus, plugStatus, esStatus, tftpStatus, vncStatus)
}

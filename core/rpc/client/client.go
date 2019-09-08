package client

import (
	"HFish/utils/log"
	"HFish/utils/conf"
	"HFish/core/rpc/core"
	"strings"
)

// 上报状态结构
type Status struct {
	AgentIp                                                        string
	AgentName                                                      string
	Web, Deep, Ssh, Redis, Mysql, Http, Telnet, Ftp, MemCahe, Plug string
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

func createClient() (*rpc.Client, string, bool) {
	rpcAddr := conf.Get("rpc", "addr")
	client, conn, err := rpc.Dial("tcp", rpcAddr)

	if err != nil {
		log.Pr("RPC", "127.0.0.1", "连接 RPC Server 失败")
		return client, "", false
	}

	ipArr := strings.Split(conn.LocalAddr().String(), ":")

	return client, ipArr[0], true
}

func reportStatus(rpcName string, ftpStatus string, telnetStatus string, httpStatus string, mysqlStatus string, redisStatus string, sshStatus string, webStatus string, darkStatus string, memCacheStatus string, plugStatus string) {
	client, addr, boolStatus := createClient()

	if boolStatus {
		defer client.Close()

		status := Status{
			addr,
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
		}

		var reply string
		err := client.Call("HFishRPCService.ReportStatus", status, &reply)

		if err != nil {
			log.Pr("RPC", "127.0.0.1", "上报服务状态失败", err)
		}
	}
}

func ReportResult(typex string, projectName string, sourceIp string, info string, id string) string {
	// projectName 只有 WEB 才需要传项目名 其他协议空即可
	// id 0 为 新插入数据，非 0 都是更新数据
	// id 非 0 的时候 sourceIp 为空
	client, addr, boolStatus := createClient()

	if boolStatus {
		defer client.Close()

		rpcName := conf.Get("rpc", "name")

		result := Result{
			addr,
			rpcName,
			typex,
			projectName,
			sourceIp,
			info,
			id,
		}

		var reply string
		err := client.Call("HFishRPCService.ReportResult", result, &reply)

		if err != nil {
			log.Pr("RPC", "127.0.0.1", "上报上钩结果失败")
		}

		return reply
	}
	return ""
}

func Start(rpcName string, ftpStatus string, telnetStatus string, httpStatus string, mysqlStatus string, redisStatus string, sshStatus string, webStatus string, darkStatus string, memCacheStatus string, plugStatus string) {
	reportStatus(rpcName, ftpStatus, telnetStatus, httpStatus, mysqlStatus, redisStatus, sshStatus, webStatus, darkStatus, memCacheStatus, plugStatus)
}

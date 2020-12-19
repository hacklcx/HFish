package client

import (
	"HFish/utils/log"
	"HFish/utils/conf"
	"HFish/core/rpc/core"
	"strings"
	"fmt"
)

// Report status structure
type Status struct {
	AgentIp                                                                               string
	AgentName                                                                             string
	Web, Deep, Ssh, Redis, Mysql, Http, Telnet, Ftp, MemCahe, Plug, ES, TFtp, Vnc, Custom string
}

// Report result structure
type Result struct {
	AgentIp     string
	AgentName   string
	Type        string
	ProjectName string
	SourceIp    string
	Info        string
	Id          string // Database ID, update with 0 for newly inserted data
}

var rpcClient *rpc.Client
var ipAddr string

func RpcInit() {
	rpcAddr := conf.Get("rpc", "addr")
	c, conn, err := rpc.Dial("tcp", rpcAddr)

	if err != nil {
		rpcClient = nil
		log.Pr("RPC", "127.0.0.1", "Failed to connect to RPC Server")
		ipAddr = ""
	} else {
		rpcClient = c
		ipAddr = strings.Split(conn.LocalAddr().String(), ":")[0]
		fmt.Println("Successfully connected to RPC Server")
	}
}

func reportStatus(rpcName string, ftpStatus string, telnetStatus string, httpStatus string, mysqlStatus string, redisStatus string, sshStatus string, webStatus string, darkStatus string, memCacheStatus string, plugStatus string, esStatus string, tftpStatus string, vncStatus string, customStatus string) {
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
			customStatus,
		}

		var reply string
		err := rpcClient.Call("HFishRPCService.ReportStatus", status, &reply)

		if err != nil {
			log.Pr("RPC", "127.0.0.1", "Failed to report service status", err)
			RpcInit()
		} else {
			fmt.Println("Report service status successfully")
		}
	} else {
		RpcInit()
	}
}

func ReportResult(typex string, projectName string, sourceIp string, info string, id string) string {
	var reply string

	if (rpcClient != nil) {
		// projectName Only WEB need to pass the project name, other protocols can be empty
		// id 0 is newly inserted data, non-zero is updated data
		// sourceIp is empty when id is not 0
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
			log.Pr("RPC", "127.0.0.1", "Failed to report the hook result")
			RpcInit()
		} else {
			fmt.Println("Successfully reported the hook result")
		}

		return reply
	} else {
		return "0"
	}
}

func Start(rpcName string, ftpStatus string, telnetStatus string, httpStatus string, mysqlStatus string, redisStatus string, sshStatus string, webStatus string, darkStatus string, memCacheStatus string, plugStatus string, esStatus string, tftpStatus string, vncStatus string, customStatus string) {
	reportStatus(rpcName, ftpStatus, telnetStatus, httpStatus, mysqlStatus, redisStatus, sshStatus, webStatus, darkStatus, memCacheStatus, plugStatus, esStatus, tftpStatus, vncStatus, customStatus)
}

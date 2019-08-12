package server

import (
	"net/rpc"
	"net"
	"HFish/utils/log"
	"HFish/core/report"
	"strconv"
)

// 上报状态结构
type Status struct {
	AgentIp                                         string
	AgentName                                       string
	Web, Deep, Ssh, Redis, Mysql, Http, Telnet, Ftp string
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

type HFishRPCService int

// 上报状态 RPC 方法
func (t *HFishRPCService) ReportStatus(s *Status, reply *string) error {

	// 上报 客户端 状态
	go report.ReportAgentStatus(
		s.AgentName,
		s.AgentIp,
		s.Web,
		s.Deep,
		s.Ssh,
		s.Redis,
		s.Mysql,
		s.Http,
		s.Telnet,
		s.Ftp,
	)

	return nil
}

// 上报结果 RPC 方法
func (t *HFishRPCService) ReportResult(r *Result, reply *string) error {
	var idx string

	switch r.Type {
	case "WEB":
		go report.ReportWeb(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "DEEP":
		go report.ReportDeepWeb(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "SSH":
		go report.ReportSSH(r.SourceIp, r.AgentName, r.Info)
	case "REDIS":
		if r.Id == "0" {
			id := report.ReportRedis(r.SourceIp, r.AgentName, r.Info)
			idx = strconv.FormatInt(id, 10)
		} else {
			go report.ReportUpdateRedis(r.Id, r.Info)
		}
	case "MYSQL":
		if r.Id == "0" {
			id := report.ReportMysql(r.SourceIp, r.AgentName, r.Info)
			idx = strconv.FormatInt(id, 10)
		} else {
			go report.ReportUpdateMysql(r.Id, r.Info)
		}
	case "TELNET":
		if r.Id == "0" {
			id := report.ReportTelnet(r.SourceIp, r.AgentName, r.Info)
			idx = strconv.FormatInt(id, 10)
		} else {
			go report.ReportUpdateTelnet(r.Id, r.Info)
		}
	case "FTP":
		go report.ReportFTP(r.SourceIp, r.AgentName, r.Info)
	}

	*reply = idx
	return nil
}

// 启动 RPC 服务端
func Start(addr string) {
	rpcService := new(HFishRPCService)
	rpc.Register(rpcService)

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Pr("RPC", "127.0.0.1", "RPC Server 启动失败", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Pr("RPC", "127.0.0.1", "RPC Server 监听地址失败", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

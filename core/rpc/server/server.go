package server

import (
	"HFish/core/rpc/core"
	"HFish/utils/log"
	"HFish/core/report"
	"strconv"
	"net"
	"fmt"
	"HFish/core/pool"
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

type HFishRPCService int

// Report status RPC method
func (t *HFishRPCService) ReportStatus(s *Status, reply *string) error {
	// Report client status
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
		s.MemCahe,
		s.Plug,
		s.ES,
		s.TFtp,
		s.Vnc,
		s.Custom,
	)

	return nil
}

// Report result RPC method
func (t *HFishRPCService) ReportResult(r *Result, reply *string) error {
	var idx string

	switch r.Type {
	case "PLUG":
		go report.ReportPlugWeb(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "WEB":
		go report.ReportWeb(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "DEEP":
		go report.ReportDeepWeb(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "HTTP":
		go report.ReportHttp(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "ES":
		go report.ReportEs(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "VNC":
		go report.ReportVnc(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "CUSTOM":
		go report.ReportCustom(r.ProjectName, r.AgentName, r.SourceIp, r.Info)
	case "FTP":
		go report.ReportFTP(r.SourceIp, r.AgentName, r.Info)
	case "TFTP":
		if r.Id == "0" {
			id := report.ReportTFtp(r.SourceIp, r.AgentName, r.Info)
			idx = strconv.FormatInt(id, 10)
		} else {
			go report.ReportUpdateTFtp(r.Id, r.Info)
		}
	case "SSH":
		if r.Id == "0" {
			id := report.ReportSSH(r.SourceIp, r.AgentName, r.Info)
			idx = strconv.FormatInt(id, 10)
		} else {
			go report.ReportUpdateSSH(r.Id, r.Info)
		}
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
	case "MEMCACHE":
		if r.Id == "0" {
			id := report.ReportMemCche(r.SourceIp, r.AgentName, r.Info)
			idx = strconv.FormatInt(id, 10)
		} else {
			go report.ReportUpdateMemCche(r.Id, r.Info)
		}
	}

	*reply = idx
	return nil
}

// Start the RPC server
func Start(addr string) {
	rpcService := new(HFishRPCService)
	rpc.Register(rpcService)

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Pr("RPC", "127.0.0.1", "RPC Server failed to activate", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Pr("RPC", "127.0.0.1", "RPC Server Failed to listen to the address", err)
	}

	wg, poolX := pool.New(500)
	defer poolX.Release()

	for {
		wg.Add(1)
		poolX.Submit(func() {
			conn, err := listener.Accept()

			if err != nil {
				log.Pr("RPC", "127.0.0.1", "The client failed to connect to the RPC Server", err)
			}

			fmt.Println(conn.RemoteAddr())

			rpc.ServeConn(conn)

			wg.Done()
		})
	}
}

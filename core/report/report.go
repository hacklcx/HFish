package report

import (
	"HFish/core/dbUtil"
	"time"
	"HFish/utils/ip"
	"strings"
	"HFish/utils/send"
	"strconv"
)

func alert(title string, agent string, ipx string, infox string) {
	sql := `select status,info from hfish_setting where type = "alertMail"`
	isAlertStatus := dbUtil.Query(sql)

	status := strconv.FormatInt(isAlertStatus[0]["status"].(int64), 10)

	// 判断是否启用通知
	if status == "1" {
		info := isAlertStatus[0]["info"]
		config := strings.Split(info.(string), "&&")

		text := `
		<div><b>Hi，上钩了！</b></div>
		<div><b><br /></b></div>
		<div><b>集群名称：</b>` + agent + `</div>
		<div><b>攻击IP：</b>` + ipx + `</div>
		<div><b>上钩内容：</b>` + infox + `</div>
		<div><br /></div>
		<div><span style="color: rgb(128, 128, 128); font-size: 10px;">(HFish 自动发送)</span></div>
		`

		send.SendMail(config[4:], "[HFish]提醒你，"+title+"有鱼上钩!", text, config)
	}
}

// 上报 集群 状态
func ReportAgentStatus(agentName string, agentIp string, webStatus string, deepStatus string, sshStatus string, redisStatus string, mysqlStatus string, httpStatus string, telnetStatus string, ftpStatus string) {
	sql := `
	INSERT INTO hfish_colony (
		agent_name,
		agent_ip,
		web_status,
		deep_status,
		ssh_status,
		redis_status,
		mysql_status,
		http_status,
		telnet_status,
		ftp_status,
		last_update_time
	)
	VALUES
		(?,?,?,?,?,?,?,?,?,?,?);
	`

	id := dbUtil.Insert(sql, agentName, agentIp, webStatus, deepStatus, sshStatus, redisStatus, mysqlStatus, httpStatus, telnetStatus, ftpStatus, time.Now().Format("2006-01-02 15:04:05"))

	// 如果 ID 等于0 代表 该数据以及存在
	if id == 0 {
		sql := `
		UPDATE hfish_colony
		SET agent_ip = ?, web_status = ?, deep_status = ?, ssh_status = ?, redis_status = ?, mysql_status = ?, http_status = ?, telnet_status = ?, ftp_status = ?, last_update_time =?
		WHERE
			agent_name =?;
		`

		dbUtil.Update(sql, agentIp, webStatus, deepStatus, sshStatus, redisStatus, mysqlStatus, httpStatus, telnetStatus, ftpStatus, time.Now().Format("2006-01-02 15:04:05"), agentName)
	}
}

// 上报 WEB
func ReportWeb(projectName string, agent string, ipx string, info string) {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	dbUtil.Insert(sql, "WEB", projectName, agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert("WEB", agent, ipx, info)
}

// 上报 暗网 WEB
func ReportDeepWeb(projectName string, agent string, ipx string, info string) {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	dbUtil.Insert(sql, "DEEP", projectName, agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert("DEEP", agent, ipx, info)
}

// 上报 SSH
func ReportSSH(ipx string, agent string, info string) {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	dbUtil.Insert(sql, "SSH", "SSH蜜罐", agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert("SSH", agent, ipx, info)
}

// 上报 Redis
func ReportRedis(ipx string, agent string, info string) int64 {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	go alert("REDIS", agent, ipx, info)
	return dbUtil.Insert(sql, "REDIS", "Redis蜜罐", agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 更新 Redis 操作
func ReportUpdateRedis(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
}

// 上报 Mysql
func ReportMysql(ipx string, agent string, info string) int64 {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	go alert("MYSQL", agent, ipx, info)
	return dbUtil.Insert(sql, "MYSQL", "Mysql蜜罐", agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 更新 Mysql 操作
func ReportUpdateMysql(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
}

// 上报 FTP
func ReportFTP(ipx string, agent string, info string) {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	dbUtil.Insert(sql, "FTP", "FTP蜜罐", agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert("FTP", agent, ipx, info)
}

// 上报 Telnet
func ReportTelnet(ipx string, agent string, info string) int64 {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	go alert("TELNET", agent, ipx, info)
	return dbUtil.Insert(sql, "TELNET", "Telnet蜜罐", agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 更新 Telnet 操作
func ReportUpdateTelnet(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
}

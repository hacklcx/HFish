package report

import (
	"HFish/core/dbUtil"
	"time"
	"HFish/utils/ip"
)

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
}

// 上报 暗网 WEB
func ReportDeepWeb(projectName string, agent string, ipx string, info string) {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	dbUtil.Insert(sql, "DEEP", projectName, agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 SSH
func ReportSSH(ipx string, agent string, info string) {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	dbUtil.Insert(sql, "SSH", "SSH蜜罐", agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 Redis
func ReportRedis(ipx string, agent string, info string) int64 {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
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
}

// 上报 Telnet
func ReportTelnet(ipx string, agent string, info string) int64 {
	ipInfo := ip.Get(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,ip_info,info,create_time) values(?,?,?,?,?,?,?);`
	return dbUtil.Insert(sql, "TELNET", "Telnet蜜罐", agent, ipx, ipInfo, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 更新 Telnet 操作
func ReportUpdateTelnet(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
}

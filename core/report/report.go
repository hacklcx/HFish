package report

import (
	"HFish/core/dbUtil"
	"time"
	"HFish/utils/ip"
	"strings"
	"HFish/utils/send"
	"strconv"
	"HFish/utils/try"
	"encoding/json"
	"bytes"
	"net/http"
	"HFish/utils/log"
)

type HFishInfo struct {
	id      string
	model   string
	project string
	typex   string
	agent   string
	ip      string
	country string
	region  string
	city    string
	info    string
	time    string
}

func alert(id string, model string, typex string, projectName string, agent string, ipx string, country string, region string, city string, infox string, time string) {
	// 判断邮件通知
	try.Try(func() {
		// 只有新加入才会发送邮件通知
		if (model == "new") {
			sql := `select status,info from hfish_setting where type = "alertMail"`
			isAlertStatus := dbUtil.Query(sql)

			status := strconv.FormatInt(isAlertStatus[0]["status"].(int64), 10)

			// 判断是否启用通知
			if status == "1" {
				info := isAlertStatus[0]["info"]
				config := strings.Split(info.(string), "&&")

				if (country == "本地地址") {
					region = ""
					city = ""
				} else if (country == "局域网") {
					region = ""
					city = ""
				}

				text := `
				<div><b>Hi，上钩了！</b></div>
				<div><b><br /></b></div>
				<div><b>集群名称：</b>` + agent + `</div>
				<div><b>攻击IP：</b>` + ipx + `</div>
				<div><b>地理信息：</b>` + country + ` ` + region + ` ` + city + `</div>
				<div><b>上钩内容：</b>` + infox + `</div>
				<div><br /></div>
				<div><span style="color: rgb(128, 128, 128); font-size: 10px;">(HFish 自动发送)</span></div>
				`

				send.SendMail(config[4:], "[HFish]提醒你，"+typex+"有鱼上钩!", text, config)
			}
		}
	}).Catch(func() {
	})

	// 判断 WebHook 通知
	try.Try(func() {
		sql := `select status,info from hfish_setting where type = "webHook"`
		isAlertStatus := dbUtil.Query(sql)

		status := strconv.FormatInt(isAlertStatus[0]["status"].(int64), 10)

		// 判断是否启用通知
		if status == "1" {
			info := isAlertStatus[0]["info"]

			fishInfo := HFishInfo{
				id,
				model,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
			}

			b, _ := json.Marshal(fishInfo)
			body := bytes.NewBuffer(b)

			resp, err := http.Post(info.(string), "application/json;charset=utf-8", body)

			if err != nil {
				log.Pr("HFish", "127.0.0.1", "WebHook 调用失败", err)
			} else {
				log.Pr("HFish", "127.0.0.1", "WebHook 调用成功")
			}

			defer resp.Body.Close()
		}
	}).Catch(func() {
	})
}

// 上报 集群 状态
func ReportAgentStatus(agentName string, agentIp string, webStatus string, deepStatus string, sshStatus string, redisStatus string, mysqlStatus string, httpStatus string, telnetStatus string, ftpStatus string, memCacheStatus string, plugStatus string) {
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
		mem_cache_status,
		plug_status,
		last_update_time
	)
	VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?);
	`

	id := dbUtil.Insert(sql, agentName, agentIp, webStatus, deepStatus, sshStatus, redisStatus, mysqlStatus, httpStatus, telnetStatus, ftpStatus, memCacheStatus, plugStatus, time.Now().Format("2006-01-02 15:04:05"))

	// 如果 ID 等于0 代表 该数据以及存在
	if id == 0 {
		sql := `
		UPDATE hfish_colony
		SET agent_ip = ?, web_status = ?, deep_status = ?, ssh_status = ?, redis_status = ?, mysql_status = ?, http_status = ?, telnet_status = ?, ftp_status = ?, mem_cache_status = ?, plug_status = ?, last_update_time = ?
		WHERE
			agent_name =?;
		`

		dbUtil.Update(sql, agentIp, webStatus, deepStatus, sshStatus, redisStatus, mysqlStatus, httpStatus, telnetStatus, ftpStatus, memCacheStatus, plugStatus, time.Now().Format("2006-01-02 15:04:05"), agentName)
	}
}

// 上报 WEB
func ReportWeb(projectName string, agent string, ipx string, info string) {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "WEB", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "WEB", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 暗网 WEB
func ReportDeepWeb(projectName string, agent string, ipx string, info string) {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "DEEP", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "DEEP", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 蜜罐插件
func ReportPlugWeb(projectName string, agent string, ipx string, info string) {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "PLUG", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "PLUG", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 SSH
func ReportSSH(ipx string, agent string, info string) int64 {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "SSH", "SSH蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "SSH", "SSH蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	return id
}

// 更新 SSH 操作
func ReportUpdateSSH(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
	go alert(id, "update", "SSH", "SSH蜜罐", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 Redis
func ReportRedis(ipx string, agent string, info string) int64 {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "REDIS", "Redis蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "REDIS", "Redis蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	return id
}

// 更新 Redis 操作
func ReportUpdateRedis(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
	go alert(id, "update", "REDIS", "Redis蜜罐", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 Mysql
func ReportMysql(ipx string, agent string, info string) int64 {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "MYSQL", "Mysql蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "MYSQL", "Mysql蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	return id
}

// 更新 Mysql 操作
func ReportUpdateMysql(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
	go alert(id, "update", "MYSQL", "Mysql蜜罐", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 FTP
func ReportFTP(ipx string, agent string, info string) {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "FTP", "FTP蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "FTP", "FTP蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 Telnet
func ReportTelnet(ipx string, agent string, info string) int64 {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "TELNET", "Telnet蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "TELNET", "Telnet蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	return id
}

// 更新 Telnet 操作
func ReportUpdateTelnet(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
	go alert(id, "update", "TELNET", "Telnet蜜罐", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 MemCache
func ReportMemCche(ipx string, agent string, info string) int64 {
	country, region, city := ip.GetIp(ipx)
	sql := `INSERT INTO hfish_info(type,project_name,agent,ip,country,region,city,info,create_time) values(?,?,?,?,?,?,?,?,?);`
	id := dbUtil.Insert(sql, "MEMCACHE", "MemCache蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	go alert(strconv.FormatInt(id, 10), "new", "MEMCACHE", "MemCache蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	return id
}

// 更新 MemCache 操作
func ReportUpdateMemCche(id string, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
	go alert(id, "update", "MEMCACHE", "MemCache蜜罐", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
}

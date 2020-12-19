package report

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"HFish/core/alert"
	"HFish/core/dbUtil"
	"HFish/core/pool"
	"HFish/utils/cache"
	"HFish/utils/conf"
	"HFish/utils/geo"
	"HFish/utils/ip"
	"HFish/utils/log"
	"HFish/utils/try"
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

// Notification module
func alertx(id string, model string, typex string, projectName string, agent string, ipx string, country string, region string, city string, infox string, timex string) {
	// Syslog notification
	alert.AlertSyslog(model, projectName, typex, agent, ipx, country, region, city, infox, timex)

	// E-mail notification
	alert.AlertMail(model, typex, agent, ipx, country, region, city, infox)

	// WebHook
	alert.AlertWebHook(id, model, typex, projectName, agent, ipx, country, region, city, infox, timex)

	// Big data display
	alert.AlertDataWs(model, typex, projectName, agent, ipx, country, region, city, timex)
}

// Report cluster status
func ReportAgentStatus(agentName string, agentIp string, webStatus string, deepStatus string, sshStatus string, redisStatus string, mysqlStatus string, httpStatus string, telnetStatus string, ftpStatus string, memCacheStatus string, plugStatus string, esStatus string, tftpStatus string, vncStatus string, customStatus string) {
	_, err := dbUtil.DB().Table("hfish_colony").Data(map[string]interface{}{
		"agent_name":       agentName,
		"agent_ip":         agentIp,
		"web_status":       webStatus,
		"deep_status":      deepStatus,
		"ssh_status":       sshStatus,
		"redis_status":     redisStatus,
		"mysql_status":     mysqlStatus,
		"http_status":      httpStatus,
		"telnet_status":    telnetStatus,
		"ftp_status":       ftpStatus,
		"mem_cache_status": memCacheStatus,
		"plug_status":      plugStatus,
		"es_status":        esStatus,
		"tftp_status":      tftpStatus,
		"vnc_status":       vncStatus,
		"custom_status":    customStatus,
		"last_update_time": time.Now().Format("2006-01-02 15:04:05"),
	}).InsertGetId()

	if err != nil {
		// If it is abnormal, it means that the unique index is triggered and go directly to the update operation
		_, err := dbUtil.DB().
			Table("hfish_colony").Data(map[string]interface{}{
			"agent_ip":         agentIp,
			"web_status":       webStatus,
			"deep_status":      deepStatus,
			"ssh_status":       sshStatus,
			"redis_status":     redisStatus,
			"mysql_status":     mysqlStatus,
			"http_status":      httpStatus,
			"telnet_status":    telnetStatus,
			"ftp_status":       ftpStatus,
			"mem_cache_status": memCacheStatus,
			"plug_status":      plugStatus,
			"es_status":        esStatus,
			"tftp_status":      tftpStatus,
			"vnc_status":       vncStatus,
			"custom_status":    customStatus,
			"last_update_time": time.Now().Format("2006-01-02 15:04:05"),
		}).Where("agent_name", agentName).Update()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "Failed to update cluster information", err)
		}
	}
}

// Determine whether it is a whitelisted IP
func isWhiteIp(ip string) bool {
	var isWhite = false

	try.Try(func() {
		status, _ := cache.Get("IpConfigStatus")

		// Determine whether to enable notification
		if status == "1" {
			info, _ := cache.Get("IpConfigInfo")
			ipArr := strings.Split(info.(string), "&&")

			for _, val := range ipArr {
				if (ip == val) {
					isWhite = true
				}
			}
		}

	}).Catch(func() {
	})

	return isWhite
}

// Universal insert
func insertInfo(typex string, projectName string, agent string, ipx string, country string, region string, city string, info string) int64 {
	timex := time.Now().Format("2006-01-02 15:04:05")
	text := fmt.Sprintf("project: %s, type: %s, agent: %s, ip: %s, geo: %s, info: %s, time: %s",
		projectName, typex, agent, ipx, geo.Format(country, region, city, "-"), info, timex)
	collectIntelligenceData(text)

	intelligence, err := fetchIntelligenceData(ipx)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "fetch intelligence data err", err)
		intelligence = err.Error()
	}

	id, err := dbUtil.DB().Table("hfish_info").Data(map[string]interface{}{
		"type":         typex,
		"project_name": projectName,
		"agent":        agent,
		"ip":           ipx,
		"country":      country,
		"region":       region,
		"city":         city,
		"intelligence": intelligence,
		"info":         info,
		"info_len":     len(info),
		"create_time":  timex,
	}).InsertGetId()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to insert hook information", err)
	}

	return id
}

// Insert account password
func insertAccountPasswd(typex string, account string, passwd string) {
	_, err := dbUtil.DB().Table("hfish_passwd").Data(map[string]interface{}{
		"type":        typex,
		"account":     account,
		"password":    passwd,
		"create_time": time.Now().Format("2006-01-02 15:04:05"),
	}).Insert()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to insert account password information", err)
	}
}

// Update
func updateInfoCore(id string, info string) {
	try.Try(func() {
		var sql string

		// Here for compatibility Mysql + Sqlite
		dbType := conf.Get("admin", "db_type")

		if dbType == "sqlite" {
			sql = `
				UPDATE hfish_info
				SET info = info||?
				WHERE
					id = ?;
				`
		} else if dbType == "mysql" {
			sql = `
				UPDATE hfish_info
				SET info = CONCAT(info, ?)
				WHERE
					id = ?;
				`
		}

		_, err := dbUtil.DB().Execute(sql, info, id)

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "Failed to update hook information", err)
		}
	}).Catch(func() {
	})
}

// General update
func updateInfo(id string, info string) {
	wgUpdate, poolUpdateX := pool.New(10)

	defer poolUpdateX.Release()

	wgUpdate.Add(1)
	poolUpdateX.Submit(func() {
		time.Sleep(time.Second * 2)
		go updateInfoCore(id, info)
		wgUpdate.Done()
	})
}

// Report WEB
func ReportWeb(projectName string, agent string, ipx string, info string) {
	// IP Not in the whitelist, report
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("WEB", projectName, agent, ipx, country, region, city, info)

		// Insert account password
		arr := strings.Split(info, "&&")
		insertAccountPasswd("WEB", arr[0], arr[1])

		go alertx(strconv.FormatInt(id, 10), "new", "WEB", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report Dark Web WEB
func ReportDeepWeb(projectName string, agent string, ipx string, info string) {
	// IP Not in the whitelist, report
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)

		// Insert account password
		arr := strings.Split(info, "&&")
		insertAccountPasswd("DEEP", arr[0], arr[1])

		id := insertInfo("DEEP", projectName, agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "DEEP", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report honeypot plugin
func ReportPlugWeb(projectName string, agent string, ipx string, info string) {
	// IP is not in the whitelist, report it
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("PLUG", projectName, agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "PLUG", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report SSH
func ReportSSH(ipx string, agent string, info string) int64 {
	defer func() {
		if err := recover(); err != nil {
			log.Pr("HFish", "127.0.0.1", "执行SSH插入失败", err)
		}
	}()

	// IP is not in the whitelist, report it
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("SSH", "SSH蜜罐", agent, ipx, country, region, city, info)

		// Insert account password
		arr := strings.Split(info, "&&")
		insertAccountPasswd("SSH", arr[0], arr[1])

		go alertx(strconv.FormatInt(id, 10), "new", "SSH", "SSH蜜罐", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
		return id
	}
	return 0
}

// Update SSH operation
func ReportUpdateSSH(id string, info string) {
	if (id != "0") {
		go updateInfo(id, info)
		go alertx(id, "update", "SSH", "SSH honeypot", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report Redis
func ReportRedis(ipx string, agent string, info string) int64 {
	// IP is not in the whitelist, report it
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("REDIS", "Redis honeypot", agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "REDIS", "Redis honeypot", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
		return id
	}
	return 0
}

// 更新 Redis 操作
func ReportUpdateRedis(id string, info string) {
	if (id != "0") {
		go updateInfo(id, info)
		go alertx(id, "update", "REDIS", "Redis honeypot", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report Mysql
func ReportMysql(ipx string, agent string, info string) int64 {
	// IP 不在白名单，进行上报
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("MYSQL", "Mysql honeypot", agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "MYSQL", "Mysql honeypot", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
		return id
	}
	return 0
}

// Update Mysql operation
func ReportUpdateMysql(id string, info string) {
	if (id != "0") {
		go updateInfo(id, info)
		go alertx(id, "update", "MYSQL", "Mysql honeypot", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report FTP
func ReportFTP(ipx string, agent string, info string) {
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("FTP", "FTP Honeypot", agent, ipx, country, region, city, info)

		// Insert account password
		arr := strings.Split(info, "&&")
		insertAccountPasswd("FTP", arr[0], arr[1])

		go alertx(strconv.FormatInt(id, 10), "new", "FTP", "FTP honeypot", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report Telnet
func ReportTelnet(ipx string, agent string, info string) int64 {
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("TELNET", "Telnet Honeypot", agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "TELNET", "Telnet honeypot", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
		return id
	}
	return 0
}

// Update Telnet operation
func ReportUpdateTelnet(id string, info string) {
	if (id != "0") {
		go updateInfo(id, info)
		go alertx(id, "update", "TELNET", "Telnet Honeypot", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report MemCache
func ReportMemCche(ipx string, agent string, info string) int64 {
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("MEMCACHE", "MemCache honeypot", agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "MEMCACHE", "MemCache honeypot", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
		return id
	}
	return 0
}

// Update MemCache operation
func ReportUpdateMemCche(id string, info string) {
	if (id != "0") {
		go updateInfo(id, info)
		go alertx(id, "update", "MEMCACHE", "MemCache honeypot", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report HTTP proxy
func ReportHttp(projectName string, agent string, ipx string, info string) {
	// IP Not in the whitelist, report
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("HTTP", projectName, agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "HTTP", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report ES
func ReportEs(projectName string, agent string, ipx string, info string) {
	// IP Not in the whitelist, report
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("ES", projectName, agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "ES", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report VNC
func ReportVnc(projectName string, agent string, ipx string, info string) {
	// IP Not in the whitelist, report
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("VNC", projectName, agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "VNC", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report TFTP
func ReportTFtp(ipx string, agent string, info string) int64 {
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("TFTP", "TFTP honeypot", agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "TFTP", "TFTP honeypot", agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
		return id
	}
	return 0
}

// Update TFTP operation
func ReportUpdateTFtp(id string, info string) {
	if (id != "0") {
		go updateInfo(id, info)
		go alertx(id, "update", "TFTP", "TFTP honeypot", "", "", "", "", "", info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

// Report custom honeypot
func ReportCustom(projectName string, agent string, ipx string, info string) {
	// IP Not in the whitelist, report
	if (isWhiteIp(ipx) == false) {
		country, region, city := ip.GetIp(ipx)
		id := insertInfo("CUSTOM", projectName, agent, ipx, country, region, city, info)
		go alertx(strconv.FormatInt(id, 10), "new", "CUSTOM", projectName, agent, ipx, country, region, city, info, time.Now().Format("2006-01-02 15:04:05"))
	}
}

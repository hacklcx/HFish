package report

import (
	"HFish/core/dbUtil"
	"time"
)

// 上报 WEB
func ReportWeb(projectName string, ip string, info string) {
	sql := `INSERT INTO hfish_info(type,project_name,ip,info,create_time) values(?,?,?,?,?);`
	dbUtil.Insert(sql, "WEB", projectName, ip, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 SSH
func ReportSSH(ip string, info string) {
	sql := `INSERT INTO hfish_info(type,project_name,ip,info,create_time) values(?,?,?,?,?);`
	dbUtil.Insert(sql, "SSH", "SSH钓鱼", ip, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 上报 Redis
func ReportRedis(ip string, info string) int64 {
	sql := `INSERT INTO hfish_info(type,project_name,ip,info,create_time) values(?,?,?,?,?);`
	return dbUtil.Insert(sql, "REDIS", "Redis钓鱼", ip, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 更新 Redis 操作
func ReportUpdateRedis(id int64, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
}

// 上报 Mysql
func ReportMysql(ip string, info string) int64 {
	sql := `INSERT INTO hfish_info(type,project_name,ip,info,create_time) values(?,?,?,?,?);`
	return dbUtil.Insert(sql, "MYSQL", "Mysql钓鱼", ip, info, time.Now().Format("2006-01-02 15:04:05"))
}

// 更新 Redis 操作
func ReportUpdateMysql(id int64, info string) {
	sql := `UPDATE hfish_info SET info = info||? WHERE id = ?;`
	dbUtil.Update(sql, info, id)
}

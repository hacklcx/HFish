package dashboard

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"HFish/utils/conf"
	"strconv"
	"HFish/error"
	"HFish/utils/log"
	"fmt"
)

func Html(c *gin.Context) {
	// 查询上钩数量
	webSum, err := dbUtil.DB().Table("hfish_info").Where("type", "=", "WEB").Count()

	fmt.Println(err)

	sshSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "SSH").Count()
	redisSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "REDIS").Count()
	mysqlSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "MYSQL").Count()
	deepSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "DEEP").Count()
	telnetSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "TELNET").Count()
	ftpSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "FTP").Count()
	memCacheSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "MEMCACHE").Count()

	// 读取服务运行状态
	mysqlStatus := conf.Get("mysql", "status")
	redisStatus := conf.Get("redis", "status")
	sshStatus := conf.Get("ssh", "status")
	webStatus := conf.Get("web", "status")
	apiStatus := conf.Get("api", "status")
	deepStatus := conf.Get("deep", "status")
	telnetStatus := conf.Get("telnet", "status")
	ftpStatus := conf.Get("ftp", "status")
	memCacheStatus := conf.Get("mem_cache", "status")

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"webSum":         webSum,
		"sshSum":         sshSum,
		"redisSum":       redisSum,
		"mysqlSum":       mysqlSum,
		"deepSum":        deepSum,
		"telnetSum":      telnetSum,
		"ftpSum":         ftpSum,
		"memCacheSum":    memCacheSum,
		"webStatus":      webStatus,
		"sshStatus":      sshStatus,
		"redisStatus":    redisStatus,
		"mysqlStatus":    mysqlStatus,
		"apiStatus":      apiStatus,
		"deepStatus":     deepStatus,
		"telnetStatus":   telnetStatus,
		"ftpStatus":      ftpStatus,
		"memCacheStatus": memCacheStatus,
	})
}

// 仪表盘折线图 统计
func GetFishData(c *gin.Context) {
	//// 统计 web
	//sqlWeb := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="WEB"
	//GROUP BY
	//hour;
	//`
	//
	//resultWeb := dbUtil.Query(sqlWeb)
	//
	webMap := make(map[string]int64)
	//for k := range resultWeb {
	//	webMap[resultWeb[k]["hour"].(string)] = resultWeb[k]["sum"].(int64)
	//}
	//
	//// 统计 ssh
	//sqlSsh := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="SSH"
	//GROUP BY
	//hour;
	//`
	//
	//resultSSH := dbUtil.Query(sqlSsh)
	//
	sshMap := make(map[string]int64)
	//for k := range resultSSH {
	//	sshMap[resultSSH[k]["hour"].(string)] = resultSSH[k]["sum"].(int64)
	//}
	//
	//// 统计 redis
	//sqlRedis := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="REDIS"
	//GROUP BY
	//hour;
	//`
	//
	//resultRedis := dbUtil.Query(sqlRedis)
	//
	redisMap := make(map[string]int64)
	//for k := range resultRedis {
	//	redisMap[resultRedis[k]["hour"].(string)] = resultRedis[k]["sum"].(int64)
	//}
	//
	//// 统计 mysql
	//sqlMysql := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="MYSQL"
	//GROUP BY
	//hour;
	//`
	//
	//resultMysql := dbUtil.Query(sqlMysql)
	//
	mysqlMap := make(map[string]int64)
	//for k := range resultMysql {
	//	mysqlMap[resultMysql[k]["hour"].(string)] = resultMysql[k]["sum"].(int64)
	//}
	//
	//// 统计 deep
	//sqlDeep := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="DEEP"
	//GROUP BY
	//hour;
	//`
	//
	//resultDeep := dbUtil.Query(sqlDeep)
	//
	deepMap := make(map[string]int64)
	//for k := range resultDeep {
	//	deepMap[resultDeep[k]["hour"].(string)] = resultDeep[k]["sum"].(int64)
	//}
	//
	//// 统计 ftp
	//sqlFtp := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="FTP"
	//GROUP BY
	//hour;
	//`
	//
	//resultFtp := dbUtil.Query(sqlFtp)
	//
	ftpMap := make(map[string]int64)
	//for k := range resultFtp {
	//	ftpMap[resultFtp[k]["hour"].(string)] = resultFtp[k]["sum"].(int64)
	//}
	//
	//// 统计 Telnet
	//sqlTelnet := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="TELNET"
	//GROUP BY
	//hour;
	//`
	//
	//resultTelnet := dbUtil.Query(sqlTelnet)
	//
	telnetMap := make(map[string]int64)
	//for k := range resultTelnet {
	//	telnetMap[resultTelnet[k]["hour"].(string)] = resultTelnet[k]["sum"].(int64)
	//}
	//
	//// 统计 MemCache
	//sqlMemCache := `
	//SELECT
	//	strftime("%H", create_time) AS hour,
	//	sum(1) AS sum
	//FROM
	//	hfish_info
	//WHERE
	//	strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
	//AND type="MEMCACHE"
	//GROUP BY
	//hour;
	//`
	//
	//resultMemCache := dbUtil.Query(sqlMemCache)
	//
	memCacheMap := make(map[string]int64)
	//for k := range resultMemCache {
	//	memCacheMap[resultMemCache[k]["hour"].(string)] = resultMemCache[k]["sum"].(int64)
	//}

	// 拼接 json
	data := map[string]map[string]int64{
		"web":      webMap,
		"ssh":      sshMap,
		"redis":    redisMap,
		"mysql":    mysqlMap,
		"deep":     deepMap,
		"ftp":      ftpMap,
		"telnet":   telnetMap,
		"memCache": memCacheMap,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

// 仪表盘攻击饼图统计
func GetFishPieData(c *gin.Context) {
	// 统计攻击地区
	resultRegion, errRegion := dbUtil.DB().Table("hfish_info").Fields("country", "count(1) AS sum").Where("country", "!=", "").GroupBy("country").OrderBy("sum desc").Limit(10).Get()

	if errRegion != nil {
		log.Pr("HFish", "127.0.0.1", "统计攻击地区失败", errRegion)
	}

	var regionList []map[string]string

	for k := range resultRegion {
		regionMap := make(map[string]string)
		regionMap["name"] = resultRegion[k]["country"].(string)
		regionMap["value"] = strconv.FormatInt(resultRegion[k]["sum"].(int64), 10)
		regionList = append(regionList, regionMap)
	}

	// 统计攻击IP
	resultIP, errIp := dbUtil.DB().Table("hfish_info").Fields("ip", "count(1) AS sum").Where("ip", "!=", "").GroupBy("ip").OrderBy("sum desc").Limit(10).Get()

	if errIp != nil {
		log.Pr("HFish", "127.0.0.1", "统计攻击IP失败", errIp)
	}

	var ipList []map[string]string

	for k := range resultIP {
		ipMap := make(map[string]string)
		ipMap["name"] = resultIP[k]["ip"].(string)
		ipMap["value"] = strconv.FormatInt(resultIP[k]["sum"].(int64), 10)
		ipList = append(ipList, ipMap)
	}

	data := map[string]interface{}{
		"regionList": regionList,
		"ipList":     ipList,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

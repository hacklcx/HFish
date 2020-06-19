package dashboard

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"HFish/utils/conf"
	"strconv"
	"HFish/error"
	"HFish/utils/log"
	"HFish/utils/cache"
)

func Html(c *gin.Context) {
	// 查询上钩数量
	webSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "WEB").Count()
	sshSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "SSH").Count()
	redisSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "REDIS").Count()
	mysqlSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "MYSQL").Count()
	deepSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "DEEP").Count()
	telnetSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "TELNET").Count()
	ftpSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "FTP").Count()
	memCacheSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "MEMCACHE").Count()
	httpSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "HTTP").Count()
	tftpSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "TFTP").Count()
	esSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "ES").Count()
	vncSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "VNC").Count()
	customSum, _ := dbUtil.DB().Table("hfish_info").Where("type", "=", "CUSTOM").Count()

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
	httpStatus := conf.Get("http", "status")
	tftpStatus := conf.Get("tftp", "status")
	esStatus := conf.Get("elasticsearch", "status")
	vncStatus := conf.Get("vnc", "status")

	// 判断自定义蜜罐是否启动
	customStatus := "0"

	customNames := conf.GetCustomName()
	if len(customNames) > 0 {
		customStatus = "1"
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"webSum":         webSum,
		"sshSum":         sshSum,
		"redisSum":       redisSum,
		"mysqlSum":       mysqlSum,
		"deepSum":        deepSum,
		"telnetSum":      telnetSum,
		"ftpSum":         ftpSum,
		"memCacheSum":    memCacheSum,
		"httpSum":        httpSum,
		"tftpSum":        tftpSum,
		"esSum":          esSum,
		"vncSum":         vncSum,
		"customSum":      customSum,
		"webStatus":      webStatus,
		"sshStatus":      sshStatus,
		"redisStatus":    redisStatus,
		"mysqlStatus":    mysqlStatus,
		"apiStatus":      apiStatus,
		"deepStatus":     deepStatus,
		"telnetStatus":   telnetStatus,
		"ftpStatus":      ftpStatus,
		"memCacheStatus": memCacheStatus,
		"httpStatus":     httpStatus,
		"tftpStatus":     tftpStatus,
		"esStatus":       esStatus,
		"vncStatus":      vncStatus,
		"customStatus":   customStatus,
	})
}

func getData(sql string) map[string]interface{} {
	var result []map[string]interface{}
	err := dbUtil.DB().Table(&result).Query(sql)

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询SQL失败", err)
	}

	resultMap := make(map[string]interface{})

	for k := range result {
		resultMap[result[k]["hour"].(string)] = result[k]["sum"]
	}

	return resultMap
}

// 仪表盘折线图 统计
func GetFishData(c *gin.Context) {
	var sqlWeb string
	var sqlSsh string
	var sqlRedis string
	var sqlMysql string
	var sqlDeep string
	var sqlFtp string
	var sqlTelnet string
	var sqlMemCache string
	var sqlHttp string
	var sqlTftp string
	var sqlVnc string
	var sqlEs string
	var sqlCustom string

	// 此处为了兼容 Mysql + Sqlite
	dbType := conf.Get("admin", "db_type")

	if dbType == "sqlite" {
		// 统计 web
		sqlWeb = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="WEB"
		GROUP BY
		hour;
		`

		// 统计 ssh
		sqlSsh = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="SSH"
		GROUP BY
		hour;
		`

		// 统计 redis
		sqlRedis = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="REDIS"
		GROUP BY
		hour;
		`

		// 统计 mysql
		sqlMysql = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="MYSQL"
		GROUP BY
		hour;
		`

		// 统计 deep
		sqlDeep = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="DEEP"
		GROUP BY
		hour;
		`

		// 统计 ftp
		sqlFtp = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="FTP"
		GROUP BY
		hour;
		`

		// 统计 Telnet
		sqlTelnet = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="TELNET"
		GROUP BY
		hour;
		`

		// 统计 MemCache
		sqlMemCache = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="MEMCACHE"
		GROUP BY
		hour;
		`

		// 统计 HTTP
		sqlHttp = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="HTTP"
		GROUP BY
		hour;
		`

		// 统计 TFTP
		sqlTftp = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="TFTP"
		GROUP BY
		hour;
		`

		// 统计 VNC
		sqlVnc = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="VNC"
		GROUP BY
		hour;
		`

		// 统计 ES
		sqlEs = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="ES"
		GROUP BY
		hour;
		`

		// 统计 CUSTOM
		sqlCustom = `
		SELECT
			strftime("%H", create_time) AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			strftime('%s', datetime('now')) - strftime('%s', create_time) < (24 * 3600)
		AND type="CUSTOM"
		GROUP BY
		hour;
		`

	} else if dbType == "mysql" {
		// 统计 web
		sqlWeb = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "WEB"
		GROUP BY
			hour;
		`

		// 统计 ssh
		sqlSsh = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "SSH"
		GROUP BY
			hour;
		`

		// 统计 redis
		sqlRedis = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "REDIS"
		GROUP BY
			hour;
		`

		// 统计 mysql
		sqlMysql = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "MYSQL"
		GROUP BY
			hour;
		`

		// 统计 deep
		sqlDeep = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "DEEP"
		GROUP BY
			hour;
		`

		// 统计 ftp
		sqlFtp = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "FTP"
		GROUP BY
			hour;
		`

		// 统计 Telnet
		sqlTelnet = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "TELNET"
		GROUP BY
			hour;
		`

		// 统计 MemCache
		sqlMemCache = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "MEMCACHE"
		GROUP BY
			hour;
		`

		// 统计 HTTP
		sqlHttp = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "HTTP"
		GROUP BY
			hour;
		`

		// 统计 TFTP
		sqlTftp = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "TFTP"
		GROUP BY
			hour;
		`

		// 统计 VNC
		sqlVnc = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "VNC"
		GROUP BY
			hour;
		`

		// 统计 ES
		sqlEs = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "ES"
		GROUP BY
			hour;
		`

		// 统计 CUSTOM
		sqlCustom = `
		SELECT
			DATE_FORMAT(create_time,"%H") AS hour,
			sum(1) AS sum
		FROM
			hfish_info
		WHERE
			create_time >= (NOW() - INTERVAL 24 HOUR)
		AND type = "CUSTOM"
		GROUP BY
			hour;
		`
	}

	var data interface{}

	val, is := cache.Get("DashboardZxDq")

	if is {
		data = val
	} else {
		webMap := getData(sqlWeb)
		sshMap := getData(sqlSsh)
		redisMap := getData(sqlRedis)
		mysqlMap := getData(sqlMysql)
		deepMap := getData(sqlDeep)
		ftpMap := getData(sqlFtp)
		telnetMap := getData(sqlTelnet)
		memCacheMap := getData(sqlMemCache)
		httpMap := getData(sqlHttp)
		tftpMap := getData(sqlTftp)
		esMap := getData(sqlEs)
		vncMap := getData(sqlVnc)
		customMap := getData(sqlCustom)

		// 拼接 json
		data = map[string]interface{}{
			"web":       webMap,
			"ssh":       sshMap,
			"redis":     redisMap,
			"mysql":     mysqlMap,
			"deep":      deepMap,
			"ftp":       ftpMap,
			"telnet":    telnetMap,
			"memCache":  memCacheMap,
			"httpMap":   httpMap,
			"tftpMap":   tftpMap,
			"vncMap":    vncMap,
			"esMap":     esMap,
			"customMap": customMap,
		}

		cache.Set("DashboardZxDq", data)
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
	var data interface{}

	val, is := cache.Get("DashboardBarDq")

	if is {
		data = val
	} else {
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

		data = map[string]interface{}{
			"regionList": regionList,
			"ipList":     ipList,
		}

		cache.Set("DashboardBarDq", data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

package dashboard

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"strconv"
	"HFish/utils/conf"
	"HFish/error"
)

func Html(c *gin.Context) {
	// 查询上钩数量
	sqlWeb := `select count(1) as sum from hfish_info where type="WEB";`
	sqlSsh := `select count(1) as sum from hfish_info where type="SSH";`
	sqlRedis := `select count(1) as sum from hfish_info where type="REDIS";`
	sqlMysql := `select count(1) as sum from hfish_info where type="MYSQL";`

	resultWeb := dbUtil.Query(sqlWeb)
	resultSsh := dbUtil.Query(sqlSsh)
	resultRedis := dbUtil.Query(sqlRedis)
	resultMysql := dbUtil.Query(sqlMysql)

	webSum := strconv.FormatInt(resultWeb[0]["sum"].(int64), 10)
	sshSum := strconv.FormatInt(resultSsh[0]["sum"].(int64), 10)
	redisSum := strconv.FormatInt(resultRedis[0]["sum"].(int64), 10)
	mysqlSum := strconv.FormatInt(resultMysql[0]["sum"].(int64), 10)

	// 读取服务运行状态
	mysqlStatus := conf.Get("mysql", "status")
	redisStatus := conf.Get("redis", "status")
	sshStatus := conf.Get("ssh", "status")
	webStatus := conf.Get("web", "status")
	apiStatus := conf.Get("api", "status")

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"webSum":      webSum,
		"sshSum":      sshSum,
		"redisSum":    redisSum,
		"mysqlSum":    mysqlSum,
		"webStatus":   webStatus,
		"sshStatus":   sshStatus,
		"redisStatus": redisStatus,
		"mysqlStatus": mysqlStatus,
		"apiStatus":   apiStatus,
	})
}

// 仪表盘折线图 统计
func GetFishData(c *gin.Context) {
	// 统计 web
	sqlWeb := `
	SELECT
		strftime("%H", create_time) AS hour,
		sum(1) AS sum
	FROM
		hfish_info
	WHERE
		strftime("%Y-%d", create_time) = strftime("%Y-%d", 'now')
	AND type="WEB"
	GROUP BY
	hour;
   `

	resultWeb := dbUtil.Query(sqlWeb)

	webMap := make(map[string]int64)
	for k := range resultWeb {
		webMap[resultWeb[k]["hour"].(string)] = resultWeb[k]["sum"].(int64)
	}

	// 统计 ssh
	sqlSsh := `
	SELECT
		strftime("%H", create_time) AS hour,
		sum(1) AS sum
	FROM
		hfish_info
	WHERE
		strftime("%Y-%d", create_time) = strftime("%Y-%d", 'now')
	AND type="SSH"
	GROUP BY
	hour;
   `

	resultSSH := dbUtil.Query(sqlSsh)

	sshMap := make(map[string]int64)
	for k := range resultSSH {
		sshMap[resultSSH[k]["hour"].(string)] = resultSSH[k]["sum"].(int64)
	}

	// 统计 redis
	sqlRedis := `
	SELECT
		strftime("%H", create_time) AS hour,
		sum(1) AS sum
	FROM
		hfish_info
	WHERE
		strftime("%Y-%d", create_time) = strftime("%Y-%d", 'now')
	AND type="REDIS"
	GROUP BY
	hour;
   `

	resultRedis := dbUtil.Query(sqlRedis)

	redisMap := make(map[string]int64)
	for k := range resultRedis {
		redisMap[resultRedis[k]["hour"].(string)] = resultRedis[k]["sum"].(int64)
	}

	// 统计 mysql
	sqlMysql := `
	SELECT
		strftime("%H", create_time) AS hour,
		sum(1) AS sum
	FROM
		hfish_info
	WHERE
		strftime("%Y-%d", create_time) = strftime("%Y-%d", 'now')
	AND type="MYSQL"
	GROUP BY
	hour;
   `

	resultMysql := dbUtil.Query(sqlMysql)

	mysqlMap := make(map[string]int64)
	for k := range resultMysql {
		mysqlMap[resultMysql[k]["hour"].(string)] = resultMysql[k]["sum"].(int64)
	}

	// 拼接 json
	s := map[string]map[string]int64{
		"web":   webMap,
		"ssh":   sshMap,
		"redis": redisMap,
		"mysql": mysqlMap,
	}

	c.JSON(http.StatusOK, error.ErrSuccessEdit(s))
}

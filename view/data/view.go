package data

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"HFish/core/dbUtil"
	"strconv"
	"HFish/utils/log"
	"HFish/error"
	"HFish/utils/conf"
)

func Html(c *gin.Context) {
	attackCity := conf.Get("admin", "attack_city")
	c.HTML(http.StatusOK, "data.html", gin.H{
		"dataAttack": attackCity,
	})
}

// 统计中国攻击地区
func GetChina(c *gin.Context) {
	resultRegion, errRegion := dbUtil.DB().Table("hfish_info").Fields("region", "count(1) AS sum").Where("country", "=", "中国").GroupBy("region").OrderBy("sum desc").Limit(8).Get()

	if errRegion != nil {
		log.Pr("HFish", "127.0.0.1", "统计攻击地区失败", errRegion)
	}

	var regionList []map[string]string

	for k := range resultRegion {
		regionMap := make(map[string]string)
		regionMap["name"] = resultRegion[k]["region"].(string)
		regionMap["value"] = strconv.FormatInt(resultRegion[k]["sum"].(int64), 10)
		regionList = append(regionList, regionMap)
	}

	data := map[string]interface{}{
		"regionList": regionList,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

// 统计国家攻击地区
func GetCountry(c *gin.Context) {
	resultRegion, errRegion := dbUtil.DB().Table("hfish_info").Fields("country", "count(1) AS sum").Where("country", "!=", "").GroupBy("country").OrderBy("sum desc").Limit(8).Get()

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

	data := map[string]interface{}{
		"regionList": regionList,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

// 统计攻击IP地区
func GetIp(c *gin.Context) {
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
		"ipList": ipList,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

// 统计攻击类型
func GetType(c *gin.Context) {
	// 统计攻击IP
	resultType, errType := dbUtil.DB().Table("hfish_info").Fields("type", "count(1) AS sum").Where("type", "!=", "").GroupBy("type").OrderBy("sum desc").Limit(10).Get()

	if errType != nil {
		log.Pr("HFish", "127.0.0.1", "统计攻击IP失败", errType)
	}

	var typeList []map[string]string

	for k := range resultType {
		typeMap := make(map[string]string)
		typeMap["name"] = resultType[k]["type"].(string)
		typeMap["value"] = strconv.FormatInt(resultType[k]["sum"].(int64), 10)
		typeList = append(typeList, typeMap)
	}

	data := map[string]interface{}{
		"typeList": typeList,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

// 获取最新数据流
func GetNewInfo(c *gin.Context) {
	db := dbUtil.DB().Table("hfish_info").OrderBy("id desc").Limit(20)

	result, err := db.Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取最新数据流失败", err)
	}

	data := map[string]interface{}{
		"result": result,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": data,
	})
}

// 获取统计账号
func GetAccountInfo(c *gin.Context) {
	var result []map[string]interface{}
	err := dbUtil.DB().Table(&result).Query("select account,count(0) as sum from hfish_passwd GROUP BY account ORDER BY sum desc;")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询SQL失败", err)
	}

	var resultMap []map[string]string

	for k := range result {
		rMap := make(map[string]string)
		rMap["name"] = result[k]["account"].(string)
		rMap["value"] = strconv.FormatInt(result[k]["sum"].(int64), 10)
		resultMap = append(resultMap, rMap)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": resultMap,
	})
}

// 获取统计密码
func GetPasswdInfo(c *gin.Context) {
	var result []map[string]interface{}
	err := dbUtil.DB().Table(&result).Query("select password,count(0) as sum from hfish_passwd GROUP BY password ORDER BY sum desc;")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询SQL失败", err)
	}

	var resultMap []map[string]string

	for k := range result {
		rMap := make(map[string]string)
		rMap["name"] = result[k]["password"].(string)
		rMap["value"] = strconv.FormatInt(result[k]["sum"].(int64), 10)
		resultMap = append(resultMap, rMap)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": resultMap,
	})
}

// 获取全球攻击数量
func GetWordInfo(c *gin.Context) {
	var result []map[string]interface{}
	err := dbUtil.DB().Table(&result).Query("select region,count(1) as sum from hfish_info GROUP BY region;")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询SQL失败", err)
	}

	var resultMap []map[string]string

	for k := range result {
		rMap := make(map[string]string)
		rMap["name"] = result[k]["region"].(string)
		rMap["value"] = strconv.FormatInt(result[k]["sum"].(int64), 10)
		resultMap = append(resultMap, rMap)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": resultMap,
	})
}

// 往下是 Web Socket 代码

// 存储全部客户端连接
var connClient = make(map[*websocket.Conn]bool)

// 去除跨域限制
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端连接
func Ws(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		// 创建 WebSocket 失败
		return
	}

	connClient[ws] = true

	defer ws.Close()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			// 客户端断开
			connClient[ws] = false
			break
		}
	}
}

// 发送消息
func Send(data map[string]interface{}) {
	for k, v := range connClient {
		if v {
			err := k.WriteJSON(data)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}

// 生成数据 JSON
func MakeDataJson(typex string, data map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"type": typex,
		"data": data,
	}

	return result
}

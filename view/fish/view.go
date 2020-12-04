package fish

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/log"
	"HFish/utils/page"
)

var header = []string{"编号", "蜜罐类型", "蜜罐名称", "集群名称", "攻击IP", "地理信息", "威胁情报", "详情信息", "详情长度", "攻击时间"}

var judgmentMap = map[string]string{
	"C2": "远控",
	"Botnet": "僵尸网络",
	"Hijacked": "劫持",
	"Phishing": "钓鱼",
	"Malware": "恶意软件",
	"Exploit": "漏洞利用",
	"Scanner": "扫描",
	"Zombie": "傀儡机",
	"Spam": "垃圾邮件",
	"Suspicious": "可疑",
	"Compromised": "失陷主机",
	"Whitelist": "白名单",
	"Brute Force": "暴力破解",
	"Proxy": "代理",
	"Info": "基础信息",
	"MiningPool": "矿池",
	"CoinMiner": "私有矿池",
	"Sinkhole C2": "安全机构接管C2",
	"SSH Brute Force": "SSH 暴力破解",
	"FTP Brute Force": "FTP 暴力破解",
	"SMTP Brute Force": "SMTP 暴力破解",
	"Http Brute Force": "HTTP AUTH 暴力破解",
	"Web Login Brute Force": "撞库",
	"HTTP Proxy": "HTTP Proxy",
	"HTTP Proxy In": "HTTP 代理入口",
	"HTTP Proxy Out": "HTTP 代理出口",
	"Socks Proxy": "Socks 代理",
	"Socks Proxy In": "Socks 代理入口",
	"Socks Proxy Out": "Socks 代理出口",
	"VPN": "VPN 代理",
	"VPN In": "VPN 代理入口",
	"VPN Out": "VPN 代理出口",
	"Tor": "Tor 代理",
	"Tor Proxy In": "Tor入口",
	"Tor Proxy Out": "Tor出口",
	"Bogon": "保留地址",
	"FullBogon": "未启用IP",
	"Gateway": "网关",
	"IDC": "IDC 服务器",
	"Dynamic IP": "动态IP",
	"Edu": "教育",
	"DDNS": "动态域名",
	"Mobile": "移动基站",
	"Search Engine Crawler": "搜索引擎爬虫",
	"CDN": "CDN 服务器",
	"Advertisement": "广告",
	"DNS": "DNS 服务器",
	"BTtracker": "BT 服务器",
	"Backbone": "骨干网",
}

// 蜜罐 页面
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "fish.html", gin.H{})
}

func isCondition(condition string) bool {
	if condition == ">" || condition == "<" ||
		condition == ">=" ||	condition == "<=" ||
		condition == "=" {
		return true
	}
	return false
}

func convertJudgment(judgmentArray []interface{}) string {
	var judgments string
	for _, j := range judgmentArray {
		js := j.(string)
		judgment, ok := judgmentMap[js]
		if !ok {
			judgment = js
		}
		if len(judgments) == 0 {
			judgments = judgment
		} else {
			judgments += " " + judgment
		}
	}
	return judgments
}

// 获取上钩列表
func GetFishList(c *gin.Context) {
	pageNo, _ := c.GetQuery("page_no")
	pageSize, _ := c.GetQuery("page_size")
	typex, _ := c.GetQuery("type")
	colony, _ := c.GetQuery("colony")
	soText, _ := c.GetQuery("so_text")
	condition, _ := c.GetQuery("condition")
	length, _ := c.GetQuery("length")
	startTime, _ := c.GetQuery("start_time")
	endTime, _ := c.GetQuery("end_time")

	// 拼接 SQL
	db := dbUtil.DB().Table("hfish_info").Fields("id", "type", "project_name", "agent", "ip", "country", "region", "city", "intelligence", "create_time", "info", "info_len")
	dbCount := dbUtil.DB().Table("hfish_info")

	if typex != "all" {
		db.Where("type", "=", typex)
		dbCount.Where("type", "=", typex)
	}

	if colony != "all" {
		db.Where("agent", "=", colony)
		dbCount.Where("agent", "=", colony)
	}

	if isCondition(condition) && length != "" {
		db.Where("info_len", condition, length)
		dbCount.Where("info_len", condition, length)
	}

	if startTime != "" && endTime != "" {
		stInt, err1 := strconv.ParseInt(startTime, 10, 64)
		etInt, err2 := strconv.ParseInt(endTime, 10, 64)
		if err1 == nil && err2 == nil {
			st := time.Unix(stInt, 0).Format("2006-01-02 15:04:05")
			et := time.Unix(etInt, 0).Format("2006-01-02 15:04:05")
			db.Where("create_time", ">=", st).Where("create_time", "<=", et)
			dbCount.Where("create_time", ">=", st).Where("create_time", "<=", et)
		} else {
			log.Pr("HFish", "127.0.0.1", "parseInt startTime err", err1)
			log.Pr("HFish", "127.0.0.1", "parseInt endTime err", err2)
		}
	}

	if soText != "" {
		db.Where("project_name", "like", "%"+soText+"%").OrWhere("ip", "like", "%"+soText+"%")
		dbCount.Where("project_name", "like", "%"+soText+"%").OrWhere("ip", "like", "%"+soText+"%")
	}
	// 统计查询数量
	totalCount, errCount := dbCount.Count()

	if errCount != nil {
		log.Pr("HFish", "127.0.0.1", "统计分页总数失败", errCount)
	}

	// 查询列表
	pageNoInt, _ := strconv.Atoi(pageNo)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	pageStart := page.Start(pageNoInt, pageSizeInt)

	result, err := db.OrderBy("id desc").Limit(pageSizeInt).Offset(pageStart).Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询上钩信息列表失败", err)
	}

	totalCountString := strconv.FormatInt(totalCount, 10)
	totalCountInt, _ := strconv.Atoi(totalCountString)

	pageCount := page.TotalPage(totalCountInt, pageSizeInt)

	for i, info := range result {
		intelligence, ok := info["intelligence"].(string)
		if !ok || len(intelligence) == 0 {
			continue
		}
		var intelligenceData json.RawMessage
		err := json.Unmarshal([]byte(intelligence), &intelligenceData)
		if err != nil {
			log.Pr("HFish", "127.0.0.1", "json unmarshal err", err)
			continue
		}
		result[i]["intelligence"] = intelligenceData
	}

	data := map[string]interface{}{
		"result":     result,
		"pageCount":  pageCount,
		"totalCount": totalCount,
		"page":       pageNo,
	}

	c.JSON(http.StatusOK, error.ErrSuccessWithData(data))
}

// 批量导出蜜罐数据
func ExportFishList(c *gin.Context) {
	typex, _ := c.GetQuery("type")
	colony, _ := c.GetQuery("colony")
	soText, _ := c.GetQuery("so_text")
	condition, _ := c.GetQuery("condition")
	length, _ := c.GetQuery("length")
	startTime, _ := c.GetQuery("start_time")
	endTime, _ := c.GetQuery("end_time")

	// 拼接 SQL
	db := dbUtil.DB().Table("hfish_info").Fields("id", "type", "project_name", "agent", "ip", "country", "region", "city", "intelligence", "create_time", "info", "info_len")

	if typex != "all" {
		db.Where("type", "=", typex)
	}

	if colony != "all" {
		db.Where("agent", "=", colony)
	}

	if isCondition(condition) && length != "" {
		db.Where("info_len", condition, length)
	}

	if startTime != "" && endTime != "" {
		stInt, err1 := strconv.ParseInt(startTime, 10, 64)
		etInt, err2 := strconv.ParseInt(endTime, 10, 64)
		if err1 == nil && err2 == nil {
			st := time.Unix(stInt, 0).Format("2006-01-02 15:04:05")
			et := time.Unix(etInt, 0).Format("2006-01-02 15:04:05")
			db.Where("create_time", ">=", st).Where("create_time", "<=", et)
		} else {
			log.Pr("HFish", "127.0.0.1", "parseInt startTime err", err1)
			log.Pr("HFish", "127.0.0.1", "parseInt endTime err", err2)
		}
	}

	if soText != "" {
		db.Where("project_name", "like", "%"+soText+"%").OrWhere("ip", "like", "%"+soText+"%")
	}
	result, err := db.OrderBy("id desc").Get()
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询上钩信息列表失败", err)
		c.JSON(http.StatusOK, error.ErrExportData)
		return
	}

	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	wr.Write(header)
	for i, info := range result {
		id := fmt.Sprintf("%d", i+1)
		ip := info["ip"].(string)

		var infoLen string
		iLen, ok := info["info_len"].(int64)
		if ok {
			infoLen = fmt.Sprintf("%d", iLen)
		}

		var intelligenceResult string
		intelligenceData := make(map[string]map[string]interface{})
		intelligence, _ := info["intelligence"].(string)
		if len(intelligence) > 20 { // 没有威胁情报的就不用反序列化了
			if err := json.Unmarshal([]byte(intelligence), &intelligenceData); err != nil {
				log.Pr("HFish", "127.0.0.1", "intelligence json unmarshal err", err)
				intelligenceResult = intelligence
			} else {
				judgments, ok := intelligenceData[ip]["judgments"].([]interface{})
				if ok {
					intelligenceResult = convertJudgment(judgments)
				} else {
					intelligenceResult = intelligence
				}
			}
		} else {
			intelligenceResult = intelligence
		}

		var geo string
		country, ok := info["country"].(string)
		if ok && country != "" {
			geo = country
		}
		region, ok := info["region"].(string)
		if ok && region != "" && region != country {
			geo += "-" + region
		}
		city, ok := info["city"].(string)
		if ok && city != "" && city != country && city != region {
			geo += "-" + city
		}

		data := []string{id, info["type"].(string), info["project_name"].(string),
			info["agent"].(string), info["ip"].(string), geo, intelligenceResult,
			info["info"].(string), infoLen, info["create_time"].(string)}
		wr.Write(data)
	}
	wr.Flush()
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=hfish.csv")
	c.String(http.StatusOK, b.String())
}

// 删除蜜罐
func PostFishDel(c *gin.Context) {
	id := c.PostForm("id")

	idx := strings.Split(id, ",")
	inId := make([]interface{}, 20)

	for _, x := range idx {
		inId = append(inId, x)
	}

	_, err := dbUtil.DB().Table("hfish_info").WhereIn("id", inId).Delete()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "删除蜜罐失败", err)
	}

	c.JSON(http.StatusOK, error.ErrSuccess)
}

// 获取蜜罐信息
func GetFishInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")

	result, err := dbUtil.DB().Table("hfish_info").Fields("info").Where("id", "=", id).First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取蜜罐信息失败", err)
	}

	c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
}

// 获取蜜罐分类信息,集群信息
func GetFishTypeInfo(c *gin.Context) {
	resultType, errType := dbUtil.DB().Table("hfish_info").Fields("type").GroupBy("type").Get()

	if errType != nil {
		log.Pr("HFish", "127.0.0.1", "获取蜜罐分类失败", errType)
	}

	resultAgent, errAgent := dbUtil.DB().Table("hfish_info").Fields("agent").GroupBy("agent").Get()

	if errAgent != nil {
		log.Pr("HFish", "127.0.0.1", "获取集群分类失败", errAgent)
	}

	data := map[string]interface{}{
		"resultInfoType":   resultType,
		"resultColonyName": resultAgent,
	}

	c.JSON(http.StatusOK, error.ErrSuccessWithData(data))
}

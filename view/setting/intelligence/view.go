package intelligence

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"HFish/core/dbUtil"
	merror "HFish/error"
	"HFish/utils/cache"
	"HFish/utils/log"
)

// 渲染威胁情报页面
func Html(c *gin.Context) {
	// 获取配置列表
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("type", "info", "status").
		Where("type", "apikey").
		OrWhere("type", "whiteIp").
		Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取配置列表失败", err)
		c.HTML(http.StatusOK, "setting.html", gin.H{
			"cloud_status": 0,
			"cloud_apikey": "",
			"local_status": 0,
			"local_iplist": "",
		})
		return
	}

	dataMap := make(map[string]map[string]interface{})
	for _, config := range result {
		cType, ok := config["type"].(string)
		if !ok {
			continue
		}
		dataMap[cType] = make(map[string]interface{})
		dataMap[cType]["status"] = config["status"]
		dataMap[cType]["info"] = config["info"]
	}

	c.HTML(http.StatusOK, "setting.html", gin.H{
		"cloud_status": dataMap["apikey"]["status"],
		"cloud_apikey": dataMap["apikey"]["info"],
		"local_status": dataMap["whiteIp"]["status"],
		"local_iplist": dataMap["whiteIp"]["info"],
	})
}

// 获取威胁情报配置信息
func GetIntelligence(c *gin.Context) {
	// 获取配置列表
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("type", "info", "status").
		Where("type", "apikey").
		OrWhere("type", "whiteIp").
		Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取配置列表失败", err)
		c.JSON(http.StatusOK, gin.H{ 
			"cloud_status": 0,
			"cloud_apikey": "",
			"local_status": 0,
			"local_iplist": "",
		})
		return
	}

	dataMap := make(map[string]map[string]interface{})
	for _, config := range result {
		cType, ok := config["type"].(string)
		if !ok {
			continue
		}
		dataMap[cType] = make(map[string]interface{})
		dataMap[cType]["status"] = config["status"]
		dataMap[cType]["info"] = config["info"]
	}

	c.JSON(http.StatusOK, merror.ErrSuccessWithData(gin.H{ 
		"cloud_status": dataMap["apikey"]["status"],
		"cloud_apikey": dataMap["apikey"]["info"],
		"local_status": dataMap["whiteIp"]["status"],
		"local_iplist": dataMap["whiteIp"]["info"],
	}))
}

// 更新威胁情报配置信息
func UpdateIntelligence(c *gin.Context) {
	cloudStatus := c.PostForm("cloud_status")
	cloudApikey := c.PostForm("cloud_apikey")
	localStatus := c.PostForm("local_status")
	localIpList := c.PostForm("local_iplist")

	if (cloudStatus != "0" && cloudStatus != "1") || (cloudStatus == "1" && len(cloudApikey) != 64) {
		log.Pr("HFish", "127.0.0.1", "请求数据非法", cloudApikey)
		c.JSON(http.StatusOK, merror.ErrInputData)
		return
	}

	localIpLists := strings.Split(localIpList, "&&")
	if (localStatus != "0" && localStatus != "1") || (localStatus == "1" && len(localIpLists) == 0) {
		log.Pr("HFish", "127.0.0.1", "请求数据非法", localIpList)
		c.JSON(http.StatusOK, merror.ErrInputData)
		return
	}

	nowTime := time.Now().Format("2006-01-02 15:04")
	// 更新
	_, err := dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": cloudStatus, "info": cloudApikey, "update_time": nowTime}).
		Where("type", "apikey").
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "更新威胁情报云端配置信息失败", err)
		c.JSON(http.StatusOK, merror.ErrUpdateData)
		return
	}

	_, err = dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": localStatus, "info": localIpList, "update_time": nowTime}).
		Where("type", "whiteIp").
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "更新威胁情报IP白名单配置信息失败", err)
		c.JSON(http.StatusOK, merror.ErrUpdateData)
		return
	}

	cache.Setx("ApikeyStatus", cloudStatus)
	cache.Setx("ApikeyInfo", cloudApikey)
	cache.Setx("IpConfigStatus", localStatus)
	cache.Setx("IpConfigInfo", localIpList)

	c.JSON(http.StatusOK, merror.ErrSuccess)
}

func httpGet(apikey string) error {
	url := fmt.Sprintf("https://api.threatbook.cn/v3/scene/ip_reputation?apikey=%s&resource=159.203.93.255", apikey)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result struct {
		ResponseCode int    `json:"response_code"`
		VerboseMsg   string `json:"verbose_msg"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.ResponseCode != 0 {
		return fmt.Errorf("%s:%d", result.VerboseMsg, result.ResponseCode)
	}
	return nil
}

// 测试获取云端威胁情报apikey是否有效
func TestIntelligence(c *gin.Context) {
	apikey := c.PostForm("apikey")

	err := httpGet(apikey)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "测试获取云端威胁情报失败", err)
		response := gin.H{
			"code": merror.ErrTestIntelligence["code"],
			"msg": merror.ErrTestIntelligence["msg"],
			"data": err.Error(),
		}
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, merror.ErrSuccess)
}

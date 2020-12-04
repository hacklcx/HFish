package alert

import (
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/cache"
	"HFish/utils/log"
	"HFish/utils/send"
)

// 渲染告警通知页面
func Html(c *gin.Context) {
	// 获取配置列表
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("type", "info", "status").
		Where("type", "alertMail").
		OrWhere("type", "webHook").
		OrWhere("type", "syslog").
		Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取告警通知配置信息失败", err)
		c.HTML(http.StatusOK, "setting.html", gin.H{
			"syslog_status": 0,
			"syslog_info": "",
			"email_status": 0,
			"email_info": "",
			"webhook_status": 0,
			"webhook_info": "",
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
		"syslog_status": dataMap["syslog"]["status"],
		"syslog_info": dataMap["syslog"]["info"],
		"email_status": dataMap["alertMail"]["status"],
		"email_info": dataMap["alertMail"]["info"],
		"webhook_status": dataMap["webHook"]["status"],
		"webhook_info": dataMap["webHook"]["info"],
	})
}

// 获取告警通知配置信息
func GetAlertData(c *gin.Context) {
	// 获取配置列表
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("type", "info", "status").
		Where("type", "alertMail").
		OrWhere("type", "webHook").
		OrWhere("type", "syslog").
		Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取告警通知配置信息失败", err)
		c.JSON(http.StatusOK, gin.H{ 
			"syslog_status": 0,
			"syslog_info": "",
			"email_status": 0,
			"email_info": "",
			"webhook_status": 0,
			"webhook_info": "",
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

	c.JSON(http.StatusOK, error.ErrSuccessWithData(gin.H{ 
		"syslog_status": dataMap["syslog"]["status"],
		"syslog_info": dataMap["syslog"]["info"],
		"email_status": dataMap["alertMail"]["status"],
		"email_info": dataMap["alertMail"]["info"],
		"webhook_status": dataMap["webHook"]["status"],
		"webhook_info": dataMap["webHook"]["info"],
	}))
}

// 更新告警通知配置信息
func UpdateAlertData(c *gin.Context) {
	syslogStatus := c.PostForm("syslog_status")
	syslogInfo := c.PostForm("syslog_info")
	emailStatus := c.PostForm("email_status")
	emailInfo := c.PostForm("email_info")
	webhookStatus := c.PostForm("webhook_status")
	webhookInfo := c.PostForm("webhook_info")

	syslogInfos := strings.Split(syslogInfo, "&&")
	if (syslogStatus != "0" && syslogStatus != "1") || (syslogStatus == "1" && len(syslogInfos) == 0) || (syslogStatus == "1" && len(syslogInfos) > 3) {
		log.Pr("HFish", "127.0.0.1", "请求数据非法", syslogInfo)
		c.JSON(http.StatusOK, error.ErrInputData)
		return
	}

	emailInfos := strings.Split(emailInfo, "&&")
	if (emailStatus != "0" && emailStatus != "1") || (emailStatus == "1" && len(emailInfos) < 6) {
		log.Pr("HFish", "127.0.0.1", "请求数据非法", emailInfo)
		c.JSON(http.StatusOK, error.ErrInputData)
		return
	}

	if (webhookStatus != "0" && webhookStatus != "1") || (webhookStatus == "1" && len(webhookStatus) == 0) {
		log.Pr("HFish", "127.0.0.1", "请求数据非法", webhookInfo)
		c.JSON(http.StatusOK, error.ErrInputData)
		return
	}

	nowTime := time.Now().Format("2006-01-02 15:04")
	// 更新syslog通知
	_, err := dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": syslogStatus, "info": syslogInfo, "update_time": nowTime}).
		Where("type", "syslog").
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "更新syslog告警通知配置信息失败", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}

	// 更新email通知
	_, err = dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": emailStatus, "info": emailInfo, "update_time": nowTime}).
		Where("type", "alertMail").
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "更新邮件告警通知配置信息失败", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}

	// 更新webhook通知
	_, err = dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": webhookStatus, "info": webhookInfo, "update_time": nowTime}).
		Where("type", "webHook").
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "更新webhook告警通知配置信息失败", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}

	// 更新syslog告警缓存
	cache.Setx("SyslogConfigStatus", syslogStatus)
	cache.Setx("SyslogConfigInfo", syslogInfo)

	// 更新邮件告警缓存
	cache.Setx("MailConfigStatus", emailStatus)
	cache.Setx("MailConfigInfo", emailInfo)

	// 更新webhook告警缓存
	cache.Setx("HookConfigStatus", webhookStatus)
	cache.Setx("HookConfigInfo", webhookInfo)

	c.JSON(http.StatusOK, error.ErrSuccess)
}

// 测试syslog服务器地址是否正常连通
func TestSyslog(c *gin.Context) {
	addr := c.PostForm("addr")
	protocol := c.PostForm("protocol")
	port := c.PostForm("port")

	err := send.TestSyslog(protocol, addr, port)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "测试syslog发送失败, 错误信息：", err)
		c.JSON(http.StatusOK, error.ErrTestSyslog)
		return
	}

	c.JSON(http.StatusOK, error.ErrSuccess)
}

// 测试邮件服务器地址是否正常连通
func TestEmail(c *gin.Context) {
	addr := c.PostForm("addr")
	protocol := c.PostForm("protocol")
	port := c.PostForm("port")
	account := c.PostForm("account")
	password := c.PostForm("password")
	emails := c.PostForm("emails")

	receivers := strings.Split(emails, ",")

	err := send.TestMail(addr, protocol, port, account, password, receivers)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "测试邮件发送失败, 错误信息：", err)
		c.JSON(http.StatusOK, error.ErrTestEmail)
		return
	}

	c.JSON(http.StatusOK, error.ErrSuccess)
}

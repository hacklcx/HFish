package setting

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"HFish/core/dbUtil"
	"strings"
	"time"
	"HFish/utils/log"
	"HFish/error"
	"HFish/utils/cache"
)

func Html(c *gin.Context) {
	// Get configuration list
	result, err := dbUtil.DB().Table("hfish_setting").Fields("id", "type", "info", "setting_name", "setting_dis", "update_time", "status").Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to get configuration list", err)
	}

	c.HTML(http.StatusOK, "setting.html", gin.H{
		"dataList": result,
	})
}

// Check whether the configuration information
func checkInfo(id string) bool {
	result, err := dbUtil.DB().Table("hfish_setting").Fields("id", "type", "info").Where("id", "=", id).First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Check whether the configuration information failed", err)
	}

	info := result["info"].(string)
	typeStr := result["type"].(string)
	infoArr := strings.Split(info, "&&")
	num := len(infoArr)

	if num == 4 && typeStr == "mail" {
		return true
	}
	if num == 2 && typeStr == "login" {
		return true
	}
	if num >= 4 && typeStr == "alertMail" {
		return true
	}
	if num >= 1 && typeStr == "whiteIp" {
		return true
	}
	if num >= 1 && typeStr == "webHook" {
		return true
	}
	if num >= 1 && typeStr == "passwdTM" {
		return true
	}
	return false
}

// Concatenated string
func joinInfo(args ...string) string {
	and := "&&"
	info := ""
	for _, value := range args {
		if value == "" {
			return ""
		}
		info += value + and
	}
	info = info[:len(info)-2]
	return info
}

// Update configuration information
func updateInfoBase(info string, id string) {
	_, err := dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"info": info, "update_time": time.Now().Format("2006-01-02 15:04")}).
		Where("id", id).
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to update configuration information", err)
	}
}

// Update mass mailing configuration
func UpdateEmailInfo(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")

	// Concatenated string
	info := joinInfo(host, port, email, pass)

	// Update
	updateInfoBase(info, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// Update warning email configuration
func UpdateAlertMail(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	receive := c.PostForm("receive")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")

	// Concatenated string
	receiveArr := strings.Split(receive, ",")
	receiveInfo := joinInfo(receiveArr...)
	info := joinInfo(host, port, email, pass, receiveInfo)

	// Update
	cache.Setx("MailConfigInfo", info)
	updateInfoBase(info, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// Update ip whitelist
func UpdateWhiteIp(c *gin.Context) {
	id := c.PostForm("id")
	whiteIpList := c.PostForm("whiteIpList")

	// Concatenated string
	Arr := strings.Split(whiteIpList, ",")
	info := joinInfo(Arr...)

	// Update
	cache.Setx("IpConfigInfo", info)
	updateInfoBase(info, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// Update webHook
func UpdateWebHook(c *gin.Context) {
	id := c.PostForm("id")
	webHookUrl := c.PostForm("webHookUrl")

	// Update
	cache.Setx("HookConfigInfo", webHookUrl)
	updateInfoBase(webHookUrl, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// Update password encryption symbol
func UpdatePasswdTM(c *gin.Context) {
	id := c.PostForm("id")
	text := c.PostForm("text")

	// Update
	cache.Setx("PasswdConfigInfo", text)
	updateInfoBase(text, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// Update setting status
func UpdateStatusSetting(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")

	if !checkInfo(id) && status == "1" {
		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrFailConfigCode,
			"msg":  error.ErrFailConfigMsg,
		})

		return
	}

	_, err := dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": status, "update_time": time.Now().Format("2006-01-02 15:04")}).
		Where("id", id).
		Update()

	if id == "2" {
		cache.Setx("MailConfigStatus", status)
	} else if id == "3" {
		cache.Setx("HookConfigStatus", status)
	} else if id == "4" {
		cache.Setx("IpConfigStatus", status)
	} else if id == "4" {
		cache.Setx("PasswdConfigStatus", status)
	}

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to update setting status", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// Get setting details according to id
func GetSettingInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")

	result, err := dbUtil.DB().Table("hfish_setting").Fields("id", "type", "info", "status").Where("id", "=", id).First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to get setting details", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": result,
	})
}

// Clear data
func ClearData(c *gin.Context) {
	tyep := c.PostForm("type")

	if tyep == "1" {
		_, err := dbUtil.DB().Table("hfish_info").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "Failed to clear the hook data", err)
		}
	} else if tyep == "2" {
		_, err := dbUtil.DB().Table("hfish_colony").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "Failed to clear cluster data", err)
		}
	} else if tyep == "3" {
		_, err := dbUtil.DB().Table("hfish_passwd").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "Failed to clear password data", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

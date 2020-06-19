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
	// 获取配置列表
	result, err := dbUtil.DB().Table("hfish_setting").Fields("id", "type", "info", "setting_name", "setting_dis", "update_time", "status").Get()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取配置列表失败", err)
	}

	c.HTML(http.StatusOK, "setting.html", gin.H{
		"dataList": result,
	})
}

// 检查是否配置信息
func checkInfo(id string) bool {
	result, err := dbUtil.DB().Table("hfish_setting").Fields("id", "type", "info").Where("id", "=", id).First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "检查是否配置信息失败", err)
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

// 拼接字符串
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

// 更新配置信息
func updateInfoBase(info string, id string) {
	_, err := dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"info": info, "update_time": time.Now().Format("2006-01-02 15:04")}).
		Where("id", id).
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "更新配置信息失败", err)
	}
}

// 更新邮件群发配置
func UpdateEmailInfo(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")

	// 拼接字符串
	info := joinInfo(host, port, email, pass)

	// 更新
	updateInfoBase(info, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// 更新警告邮件配置
func UpdateAlertMail(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	receive := c.PostForm("receive")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")

	// 拼接字符串
	receiveArr := strings.Split(receive, ",")
	receiveInfo := joinInfo(receiveArr...)
	info := joinInfo(host, port, email, pass, receiveInfo)

	// 更新
	cache.Setx("MailConfigInfo", info)
	updateInfoBase(info, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// 更新ip白名单
func UpdateWhiteIp(c *gin.Context) {
	id := c.PostForm("id")
	whiteIpList := c.PostForm("whiteIpList")

	// 拼接字符串
	Arr := strings.Split(whiteIpList, ",")
	info := joinInfo(Arr...)

	// 更新
	cache.Setx("IpConfigInfo", info)
	updateInfoBase(info, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// 更新 webHook
func UpdateWebHook(c *gin.Context) {
	id := c.PostForm("id")
	webHookUrl := c.PostForm("webHookUrl")

	// 更新
	cache.Setx("HookConfigInfo", webHookUrl)
	updateInfoBase(webHookUrl, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// 更新 密码加密符号
func UpdatePasswdTM(c *gin.Context) {
	id := c.PostForm("id")
	text := c.PostForm("text")

	// 更新
	cache.Setx("PasswdConfigInfo", text)
	updateInfoBase(text, id)

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// 更新设置状态
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
		log.Pr("HFish", "127.0.0.1", "更新设置状态失败", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

// 根据id获取设置详情
func GetSettingInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")

	result, err := dbUtil.DB().Table("hfish_setting").Fields("id", "type", "info", "status").Where("id", "=", id).First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取设置详情失败", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": result,
	})
}

// 清空数据
func ClearData(c *gin.Context) {
	tyep := c.PostForm("type")

	if tyep == "1" {
		_, err := dbUtil.DB().Table("hfish_info").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "清空上钩数据失败", err)
		}
	} else if tyep == "2" {
		_, err := dbUtil.DB().Table("hfish_colony").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "清空集群数据失败", err)
		}
	} else if tyep == "3" {
		_, err := dbUtil.DB().Table("hfish_passwd").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "清空密码数据失败", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
	})
}

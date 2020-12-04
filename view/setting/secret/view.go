package secret

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/cache"
	"HFish/utils/log"
)

// 渲染数据合规页面
func Html(c *gin.Context) {
	// 获取配置列表
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("info", "status").
		Where("type", "passwdTM").
		First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取配置列表失败", err)
		c.HTML(http.StatusOK, "setting.html", gin.H{
			"passwd_status": 0,
			"passwd_text": "",
		})
		return
	}

	c.HTML(http.StatusOK, "setting.html", gin.H{
		"passwd_status": result["status"],
		"passwd_text": result["info"],
	})
}

// 获取数据合规配置信息
func GetSecretData(c *gin.Context) {
	// 获取配置列表
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("info", "status").
		Where("type", "passwdTM").
		First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "获取配置列表失败", err)
		c.JSON(http.StatusOK, gin.H{
			"passwd_status": 0,
			"passwd_text": "",
		})
		return
	}

	c.JSON(http.StatusOK, error.ErrSuccessWithData(gin.H{ 
		"passwd_status": result["status"],
		"passwd_text": result["info"],
	}))
}

// 更新数据合规配置信息
func UpdateSecretData(c *gin.Context) {
	status := c.PostForm("passwd_status")
	text := c.PostForm("passwd_text")

	if (status != "0" && status != "1") || (status == "1" && len(text) == 0) {
		log.Pr("HFish", "127.0.0.1", "请求数据非法", text)
		c.JSON(http.StatusOK, error.ErrInputData)
		return
	}

	// 更新
	nowTime := time.Now().Format("2006-01-02 15:04")
	_, err := dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": status, "info": text, "update_time": nowTime}).
		Where("type", "passwdTM").
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "更新数据合规配置信息失败", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}

	cache.Setx("PasswdConfigStatus", status)
	cache.Setx("PasswdConfigInfo", text)

	c.JSON(http.StatusOK, error.ErrSuccess)
}

// 清空数据
func ClearData(c *gin.Context) {
	tyep := c.PostForm("type")

	if tyep == "1" {
		_, err := dbUtil.DB().Table("hfish_info").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "清空上钩数据失败", err)
			c.JSON(http.StatusOK, error.ErrDeleteData)
			return
		}
	} else if tyep == "2" {
		_, err := dbUtil.DB().Table("hfish_colony").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "清空集群数据失败", err)
			c.JSON(http.StatusOK, error.ErrDeleteData)
			return
		}
	} else if tyep == "3" {
		_, err := dbUtil.DB().Table("hfish_passwd").Force().Delete()

		if err != nil {
			log.Pr("HFish", "127.0.0.1", "清空密码数据失败", err)
			c.JSON(http.StatusOK, error.ErrDeleteData)
			return
		}
	}

	c.JSON(http.StatusOK, error.ErrSuccess)
}

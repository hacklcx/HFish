package view

import (
	"HFish/view/api"
	"HFish/view/dashboard"
	"HFish/view/fish"
	"HFish/view/mail"
	"HFish/view/colony"
	"HFish/view/setting"
	"github.com/gin-gonic/gin"
	"HFish/view/login"
	"HFish/utils/conf"
	"net/http"
	"HFish/utils/is"
)

// 解决跨域问题
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func LoadUrl(r *gin.Engine) {
	// 判断是否为 RPC 客户端
	if is.Rpc() {
		/* RPC 客户端 */

		// API 接口
		// WEB 上报蜜罐信息
		apiStatus := conf.Get("api", "status")

		// 判断 API 是否启用
		if apiStatus == "1" {
			r.Use(cors())

			webUrl := conf.Get("api", "web_url")
			deepUrl := conf.Get("api", "deep_url")

			r.POST(webUrl, api.ReportWeb)
			r.POST(deepUrl, api.ReportDeepWeb)
		}
	} else {
		/* RPC 服务端 */
		// 登录
		r.GET("/login", login.Html)
		r.POST("/login", login.Login)
		r.GET("/logout", login.Logout)

		// 仪表盘
		r.GET("/", login.Jump, dashboard.Html)
		r.GET("/dashboard", login.Jump, dashboard.Html)
		r.GET("/get/dashboard/data", login.Jump, dashboard.GetFishData)

		// 蜜罐列表
		r.GET("/fish", login.Jump, fish.Html)
		r.GET("/get/fish/list", login.Jump, fish.GetFishList)
		r.GET("/get/fish/info", login.Jump, fish.GetFishInfo)
		r.GET("/get/fish/typeList", login.Jump, fish.GetFishTypeInfo)
		r.POST("/post/fish/del", login.Jump, fish.PostFishDel)

		// 分布式集群
		r.GET("/colony", login.Jump, colony.Html)
		r.GET("/get/colony/list", login.Jump, colony.GetColony)

		// 邮件群发
		r.GET("/mail", login.Jump, mail.Html)
		r.POST("/post/mail/sendEmail", login.Jump, mail.SendEmailToUsers)

		// 设置
		r.GET("/setting", login.Jump, setting.Html)
		r.GET("/get/setting/info", login.Jump, setting.GetSettingInfo)
		r.POST("/post/setting/update", login.Jump, setting.UpdateEmailInfo)
		r.POST("/post/setting/updateAlertMail", login.Jump, setting.UpdateAlertMail)
		r.POST("/post/setting/checkSetting", login.Jump, setting.UpdateStatusSetting)

		// API 接口
		// WEB 上报蜜罐信息
		apiStatus := conf.Get("api", "status")

		// 判断 API 是否启用
		if apiStatus == "1" {
			r.Use(cors())

			webUrl := conf.Get("api", "web_url")
			deepUrl := conf.Get("api", "deep_url")

			r.POST(webUrl, api.ReportWeb)
			r.POST(deepUrl, api.ReportDeepWeb)

			r.GET("/api/v1/get/ip", api.GetIpList)
		}
	}
}

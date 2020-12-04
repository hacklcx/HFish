package view

import (
	"github.com/gin-gonic/gin"

	"HFish/utils/cors"
	"HFish/view/api"
	"HFish/view/colony"
	"HFish/view/data"
	"HFish/view/dashboard"
	"HFish/view/fish"
	"HFish/view/login"
	"HFish/view/mail"
	"HFish/view/setting/alert"
	"HFish/view/setting/intelligence"
	"HFish/view/setting/massset"
	"HFish/view/setting/secret"
)

func LoadUrl(r *gin.Engine) {
	/* RPC 服务端 */
	// 登录
	r.GET("/login", login.Html)
	r.POST("/login", login.Login)
	r.GET("/logout", login.Logout)

	// 仪表盘
	r.GET("/", login.Jump, dashboard.Html)
	r.GET("/dashboard", login.Jump, dashboard.Html)
	r.GET("/get/dashboard/data", login.Jump, dashboard.GetFishData)
	r.GET("/get/dashboard/pie_data", login.Jump, dashboard.GetFishPieData)

	// 蜜罐列表
	r.GET("/fish", login.Jump, fish.Html)
	r.GET("/get/fish/list", login.Jump, fish.GetFishList)
	r.GET("/get/fish/export", login.Jump, fish.ExportFishList)
	r.GET("/get/fish/info", login.Jump, fish.GetFishInfo)
	r.GET("/get/fish/typeList", login.Jump, fish.GetFishTypeInfo)
	r.POST("/post/fish/del", login.Jump, fish.PostFishDel)

	// 大数据仪表盘
	r.GET("/data", login.Jump, data.Html)
	r.GET("/data/get/china", login.Jump, data.GetChina)
	r.GET("/data/get/country", login.Jump, data.GetCountry)
	r.GET("/data/get/ip", login.Jump, data.GetIp)
	r.GET("/data/get/type", login.Jump, data.GetType)
	r.GET("/data/get/info", login.Jump, data.GetNewInfo)
	r.GET("/data/get/account", login.Jump, data.GetAccountInfo)
	r.GET("/data/get/password", login.Jump, data.GetPasswdInfo)
	r.GET("/data/get/word", login.Jump, data.GetWordInfo)
	r.GET("/data/ws", data.Ws)

	// 分布式集群
	r.GET("/colony", login.Jump, colony.Html)
	r.GET("/get/colony/list", login.Jump, colony.GetColony)
	r.POST("/post/colony/del", login.Jump, colony.PostColonyDel)

	// 邮件群发
	r.GET("/mail", login.Jump, mail.Html)
	r.POST("/post/mail/sendEmail", login.Jump, mail.SendEmailToUsers)

	// 系统设置->告警通知
	r.GET("/setting", login.Jump, alert.Html)
	r.GET("/get/setting/alert", login.Jump, alert.GetAlertData)
	r.POST("/post/setting/alert", login.Jump, alert.UpdateAlertData)
	r.POST("/post/setting/syslog/test", login.Jump, alert.TestSyslog)
	r.POST("/post/setting/email/test", login.Jump, alert.TestEmail)

	// 系统设置->群发设置
	r.GET("/get/setting/massset", login.Jump, massset.GetMassSet)
	r.POST("/post/setting/massset", login.Jump, massset.UpdateMassSet)

	// 系统设置->威胁情报
	r.GET("/get/setting/intelligence", login.Jump, intelligence.GetIntelligence)
	r.POST("/post/setting/intelligence", login.Jump, intelligence.UpdateIntelligence)
	r.POST("/post/setting/intelligence/test", login.Jump, intelligence.TestIntelligence)

	// 系统设置->数据合规
	r.GET("/get/setting/secret", login.Jump, secret.GetSecretData)
	r.POST("/post/setting/secret", login.Jump, secret.UpdateSecretData)
	r.POST("/post/setting/cleardata", login.Jump, secret.ClearData)

	// API 接口
	// 解决跨域问题
	r.Use(cors.Cors())
	r.GET("/api/v1/get/ip", api.GetIpList)
	r.GET("/api/v1/get/fish_info", api.GetFishInfo)
	r.GET("/api/v1/get/passwd_list", api.GetAccountPasswdInfo)
}

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
	"HFish/utils/cors"
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
	r.GET("/get/fish/info", login.Jump, fish.GetFishInfo)
	r.GET("/get/fish/typeList", login.Jump, fish.GetFishTypeInfo)
	r.POST("/post/fish/del", login.Jump, fish.PostFishDel)

	// 分布式集群
	r.GET("/colony", login.Jump, colony.Html)
	r.GET("/get/colony/list", login.Jump, colony.GetColony)
	r.POST("/post/colony/del", login.Jump, colony.PostColonyDel)

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
	// 解决跨域问题
	r.Use(cors.Cors())
	r.GET("/api/v1/get/ip", api.GetIpList)
	r.GET("/api/v1/get/fish_info", api.GetFishInfo)
}

package api

import (
	"github.com/gin-gonic/gin"
	"HFish/core/report"
	"net/http"
	"HFish/error"
	"HFish/utils/conf"
	"HFish/core/dbUtil"
	"HFish/core/rpc/client"
	"HFish/utils/is"
)

// 上报WEB蜜罐
func ReportWeb(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	secKey := c.PostForm("sec_key")
	ip := c.ClientIP()

	apiSecKey := conf.Get("api", "sec_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey())
	} else {

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("WEB", name, ip, info, "0")
		} else {
			go report.ReportWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, error.ErrSuccessNull())
	}
}

// 上报暗网蜜罐
func ReportDeepWeb(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	secKey := c.PostForm("sec_key")
	ip := c.ClientIP()

	apiSecKey := conf.Get("api", "sec_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey())
	} else {

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("DEEP", name, ip, info, "0")
		} else {
			go report.ReportDeepWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, error.ErrSuccessNull())
	}
}

// 蜜罐插件API
func ReportPlugWeb(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	secKey := c.PostForm("sec_key")
	ip := c.PostForm("ip")

	apiSecKey := conf.Get("api", "sec_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey())
	} else {
		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("PLUG", name, ip, info, "0")
		} else {
			go report.ReportPlugWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, error.ErrSuccessNull())
	}
}

// 获取黑名单 黑客IP 列表
func GetIpList(c *gin.Context) {
	sql := `select ip from hfish_info GROUP BY ip;`
	result := dbUtil.Query(sql)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

// 获取钓鱼列表 API
func GetFishInfo(c *gin.Context) {
	sql := `select * from hfish_info ORDER BY id desc`
	result := dbUtil.Query(sql)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

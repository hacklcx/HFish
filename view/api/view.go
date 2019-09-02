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
	"HFish/utils/log"
)

// 上报WEB蜜罐
func ReportWeb(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	secKey := c.PostForm("sec_key")
	ip := c.ClientIP()

	if (ip == "::1") {
		ip = "127.0.0.1"
	}

	apiSecKey := conf.Get("api", "sec_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrFailApiKeyCode,
			"msg":  error.ErrFailApiKeyMsg,
		})
	} else {

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("WEB", name, ip, info, "0")
		} else {
			go report.ReportWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrSuccessCode,
			"msg":  error.ErrSuccessMsg,
		})
	}
}

// 上报暗网蜜罐
func ReportDeepWeb(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	secKey := c.PostForm("sec_key")
	ip := c.ClientIP()

	if (ip == "::1") {
		ip = "127.0.0.1"
	}

	apiSecKey := conf.Get("api", "sec_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrFailApiKeyCode,
			"msg":  error.ErrFailApiKeyMsg,
		})
	} else {

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("DEEP", name, ip, info, "0")
		} else {
			go report.ReportDeepWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrSuccessCode,
			"msg":  error.ErrSuccessMsg,
		})
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
		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrFailApiKeyCode,
			"msg":  error.ErrFailApiKeyMsg,
		})
	} else {
		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("PLUG", name, ip, info, "0")
		} else {
			go report.ReportPlugWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrSuccessCode,
			"msg":  error.ErrSuccessMsg,
		})
	}
}

// 获取黑名单 黑客IP 列表
func GetIpList(c *gin.Context) {
	result, err := dbUtil.DB().Table("hfish_info").Fields("ip").GroupBy("ip").Get()

	if err != nil {
		log.Pr("API", "127.0.0.1", "查询黑名单IP列表失败", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": result,
	})
}

// 获取钓鱼列表 API
func GetFishInfo(c *gin.Context) {
	result, err := dbUtil.DB().Table("hfish_info").OrderBy("id desc").Get()

	if err != nil {
		log.Pr("API", "127.0.0.1", "获取钓鱼列表失败", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrSuccessCode,
		"msg":  error.ErrSuccessMsg,
		"data": result,
	})
}

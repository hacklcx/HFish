package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"HFish/core/report"
	"HFish/core/dbUtil"
	"HFish/core/rpc/client"
	"HFish/error"
	"HFish/utils/conf"
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

	apiSecKey := conf.Get("api", "report_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("WEB", name, ip, info, "0")
		} else {
			go report.ReportWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, error.ErrSuccess)
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

	apiSecKey := conf.Get("api", "report_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("DEEP", name, ip, info, "0")
		} else {
			go report.ReportDeepWeb(name, "本机", ip, info)
		}

		c.JSON(http.StatusOK, error.ErrSuccess)
	}
}

type PlugInfo struct {
	Name   string                 `json:"name"`
	Ip     string                 `json:"ip"`
	SecKey string                 `json:"sec_key"`
	Info   map[string]interface{} `json:"info"`
}

// 蜜罐插件API
func ReportPlugWeb(c *gin.Context) {
	var info PlugInfo
	err := c.BindJSON(&info)

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "插件上报信息错误", err)

		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrFailPlug["code"],
			"msg":  error.ErrFailPlug["msg"],
			"data": err,
		})
		return
	}

	args := ""

	if len(info.Info) != 0 {
		for k, v := range info.Info["args"].(map[string]interface{}) {
			if args == "" {
				args += k + "=" + v.(string)
			} else {
				args += "&" + k + "=" + v.(string)
			}
		}
	}

	data := "Host:" + info.Info["host"].(string) + "&&Url:" + info.Info["uri"].(string) + "&&Method:" + info.Info["method"].(string) + "&&Args:" + args + "&&UserAgent:" + info.Info["http_user_agent"].(string) + "&&RemoteAddr:" + info.Info["remote_addr"].(string) + "&&TimeLocal:" + info.Info["time_local"].(string)

	apiSecKey := conf.Get("api", "report_key")

	if info.SecKey != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("PLUG", info.Name, info.Ip, data, "0")
		} else {
			go report.ReportPlugWeb(info.Name, "本机", info.Ip, data)
		}

		c.JSON(http.StatusOK, error.ErrSuccess)
	}
}

// 获取黑名单 黑客IP 列表
func GetIpList(c *gin.Context) {
	key, _ := c.GetQuery("key")

	apiSecKey := conf.Get("api", "query_key")

	if key != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {
		result, err := dbUtil.DB().Table("hfish_info").Fields("ip").GroupBy("ip").Get()

		if err != nil {
			log.Pr("API", "127.0.0.1", "查询黑名单IP列表失败", err)
		}

		c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
	}
}

// 获取钓鱼列表 API
func GetFishInfo(c *gin.Context) {
	key, _ := c.GetQuery("key")

	apiSecKey := conf.Get("api", "query_key")

	if key != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {
		result, err := dbUtil.DB().Table("hfish_info").OrderBy("id desc").Get()

		if err != nil {
			log.Pr("API", "127.0.0.1", "获取钓鱼列表失败", err)
		}

		c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
	}
}

// 获取账号密码列表 API
func GetAccountPasswdInfo(c *gin.Context) {
	key, _ := c.GetQuery("key")

	apiSecKey := conf.Get("api", "query_key")

	if key != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {
		result, err := dbUtil.DB().Table("hfish_passwd").OrderBy("id desc").Get()

		if err != nil {
			log.Pr("API", "127.0.0.1", "获取账号密码列表失败", err)
		}

		c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
	}
}

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

// Report WEB honeypot
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

		// Determine whether it is an RPC client
		if is.Rpc() {
			go client.ReportResult("WEB", name, ip, info, "0")
		} else {
			go report.ReportWeb(name, "Native", ip, info)
		}

		c.JSON(http.StatusOK, error.ErrSuccess)
	}
}

// Report dark web honeypot
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

		// Determine whether it is an RPC client
		if is.Rpc() {
			go client.ReportResult("DEEP", name, ip, info, "0")
		} else {
			go report.ReportDeepWeb(name, "Native", ip, info)
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

// Honeypot plugin API
func ReportPlugWeb(c *gin.Context) {
	var info PlugInfo
	err := c.BindJSON(&info)

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Plug-in report error", err)

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

		// Determine whether it is an RPC client
		if is.Rpc() {
			go client.ReportResult("PLUG", info.Name, info.Ip, data, "0")
		} else {
			go report.ReportPlugWeb(info.Name, "Native", info.Ip, data)
		}

		c.JSON(http.StatusOK, error.ErrSuccess)
	}
}

// Get blacklist hacker IP list
func GetIpList(c *gin.Context) {
	key, _ := c.GetQuery("key")

	apiSecKey := conf.Get("api", "query_key")

	if key != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {
		result, err := dbUtil.DB().Table("hfish_info").Fields("ip").GroupBy("ip").Get()

		if err != nil {
			log.Pr("API", "127.0.0.1", "Failed to query the blacklist IP list", err)
		}

		c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
	}
}

// Get phishing list API
func GetFishInfo(c *gin.Context) {
	key, _ := c.GetQuery("key")

	apiSecKey := conf.Get("api", "query_key")

	if key != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {
		result, err := dbUtil.DB().Table("hfish_info").OrderBy("id desc").Get()

		if err != nil {
			log.Pr("API", "127.0.0.1", "Failed to get phishing list", err)
		}

		c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
	}
}

// Get account password list API
func GetAccountPasswdInfo(c *gin.Context) {
	key, _ := c.GetQuery("key")

	apiSecKey := conf.Get("api", "query_key")

	if key != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey)
		return
	} else {
		result, err := dbUtil.DB().Table("hfish_passwd").OrderBy("id desc").Get()

		if err != nil {
			log.Pr("API", "127.0.0.1", "Failed to get account password list", err)
		}

		c.JSON(http.StatusOK, error.ErrSuccessWithData(result))
	}
}

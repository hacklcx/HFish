package api

import (
	"github.com/gin-gonic/gin"
	"HFish/core/report"
	"net/http"
	"HFish/error"
	"HFish/utils/conf"
	"HFish/core/dbUtil"
)

func ReportWeb(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	secKey := c.PostForm("sec_key")
	ip := c.ClientIP()

	apiSecKey := conf.Get("api", "sec_key")

	if secKey != apiSecKey {
		c.JSON(http.StatusOK, error.ErrFailApiKey())
	} else {
		go report.ReportWeb(name, ip, info)
		c.JSON(http.StatusOK, error.ErrSuccessNull())
	}
}

// 获取记录黑客IP
func GetIpList(c *gin.Context) {
	sql := `select ip from hfish_info GROUP BY ip;`
	result := dbUtil.Query(sql)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

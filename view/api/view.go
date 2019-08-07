package api

import (
	"github.com/gin-gonic/gin"
	"HFish/core/report"
	"net/http"
	"HFish/error"
	"HFish/utils/conf"
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
		report.ReportWeb(name, ip, info)
		c.JSON(http.StatusOK, error.ErrSuccessNull())
	}
}

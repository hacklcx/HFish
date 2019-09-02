package mail

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"HFish/core/dbUtil"
	"HFish/utils/send"
	"HFish/error"
	"strconv"
	"HFish/utils/log"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "mail.html", gin.H{})
}

func SendEmailToUsers(c *gin.Context) {
	emails := c.PostForm("emails")
	title := c.PostForm("title")
	content := c.PostForm("content")

	eArr := strings.Split(emails, ",")

	result, err := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "mail").First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "查询邮件配置信息失败", err)
	}

	info := result["info"]
	config := strings.Split(info.(string), "&&")

	status := strconv.FormatInt(result["status"].(int64), 10)

	if status == "1" {
		send.SendMail(eArr, title, content, config)

		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrSuccessCode,
			"msg":  error.ErrSuccessMsg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": error.ErrFailMailCode,
			"msg":  error.ErrFailMailMsg,
		})
	}
}

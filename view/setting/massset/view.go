package massset

import (
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/log"
)

// Render the group setting page
func Html(c *gin.Context) {
// Get configuration list
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("info", "status").
		Where("type", "mail").
		First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to get the configuration information of group sending", err)
		c.HTML(http.StatusOK, "setting.html", gin.H{
			"email_status": 0,
			"email_info": "",
		})
		return
	}

	c.HTML(http.StatusOK, "setting.html", gin.H{
		"email_status": result["status"],
		"email_info": result["info"],
	})
}

// Get the configuration information of the group sending setting
func GetMassSet(c *gin.Context) {
	// Get configuration list
	result, err := dbUtil.DB().Table("hfish_setting").
		Fields("info", "status").
		Where("type", "mail").
		First()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to get the configuration information of group sending", err)
		c.JSON(http.StatusOK, gin.H{ 
			"mail_status": 0,
			"mail_info": "",
		})
		return
	}

	c.JSON(http.StatusOK, error.ErrSuccessWithData(gin.H{ 
		"mail_status": result["status"],
		"mail_info": result["info"],
	}))
}

// Update the configuration information of group sending settings
func UpdateMassSet(c *gin.Context) {
	status := c.PostForm("mail_status")
	info := c.PostForm("mail_info")

	infos := strings.Split(info, "&&")
	if (status != "0" && status != "1") || (status == "1" && len(infos) != 5){
		log.Pr("HFish", "127.0.0.1", "Illegal request data", info)
		c.JSON(http.StatusOK, error.ErrInputData)
		return
	}

	// Update
	nowTime := time.Now().Format("2006-01-02 15:04")
	_, err := dbUtil.DB().
		Table("hfish_setting").
		Data(map[string]interface{}{"status": status, "info": info, "update_time": nowTime}).
		Where("type", "mail").
		Update()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "Failed to update configuration information of group sending", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}

	c.JSON(http.StatusOK, error.ErrSuccess)
}

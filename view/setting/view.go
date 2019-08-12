package setting

import (
	"HFish/core/dbUtil"
	"HFish/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Html(c *gin.Context) {
	data := getSetting() //订阅通知等
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"dataList": data,
	})
}

/*获取配置*/
func getSetting() []map[string]interface{} {
	sql := "select id,type,info,setting_name,setting_dis,update_time,status from hfish_setting where setting_type!=-1"
	result := dbUtil.Query(sql)
	return result
}

/*检查是否配置信息*/
func checkInfo(id string) bool {
	sql := "select id,info,type from hfish_setting where id = ?"
	result := dbUtil.Query(sql, id)
	info := result[0]["info"].(string)
	typeStr := result[0]["type"].(string)
	infoArr := strings.Split(info, "&&")
	num := len(infoArr)

	if num == 4 && typeStr == "mail" {
		return true
	}
	if num == 2 && typeStr == "login" {
		return true
	}
	if num == 2 && typeStr == "alertOver" {
		return true
	}
	if num == 1 && typeStr == "pushBullet" {
		return true
	}
	if num == 1 && typeStr == "fangTang" {
		return true
	}
	if num >= 4 && typeStr == "alertMail" {
		return true
	}
	return false
}
func joinInfo(args ...string) string {
	and := "&&"
	info := ""
	for _, value := range args {
		if value == "" {
			return ""
		}
		info += value + and
	}
	info = info[:len(info)-2]
	return info
}

/*更新邮件通知*/
func UpdateEmailInfo(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")
	//subType := c.PostForm("type")
	info := joinInfo(host, port, email, pass)
	sql := `
		UPDATE  hfish_setting 
		set	info = ?,
			update_time = ?
		where id = ?;`
	dbUtil.Update(sql, info, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}
/*更新警告邮件通知*/
func UpdateAlertMail(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	receive:=c.PostForm("receive")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")
	//subType := c.PostForm("type")
	receiveArr:=strings.Split(receive,",")
	receiveInfo:=joinInfo(receiveArr...)
	info := joinInfo(host, port, email, pass,receiveInfo)
	sql := `
		UPDATE  hfish_setting 
		set	info = ?,
			update_time = ?
		where id = ?;`
	dbUtil.Update(sql, info, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

/*更新设置状态*/
func UpdateStatusSetting(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")

	if !checkInfo(id) && status == "1" {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "请配置后在启用", "data": nil})
		return
	}
	sql := `update hfish_setting
		set status = ?,
			update_time=?
		where id = ?`
	dbUtil.Update(sql, status, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

/*根据id获取设置详情*/
func GetSettingInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")
	sql := `select id,type,info,status from hfish_setting where id = ?`
	result := dbUtil.Query(sql, id)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

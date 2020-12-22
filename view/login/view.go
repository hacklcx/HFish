package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"HFish/core/dbUtil"
	"HFish/error"
	"HFish/utils/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const k30MinutesSec = 1800

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Jump(c *gin.Context) {
	session := sessions.Default(c)
	loginName, _ := session.Get("login_name").(string)
	expireTime, _ := session.Get("expire_time").(int64)

	nowTime := time.Now().Unix()
	if loginName == "" || nowTime > expireTime {
		session.Clear()
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	} else {
		if nowTime - expireTime < 60 {
			// 即将过期则重置过期时间
			session.Set("expire_time", nowTime + k30MinutesSec)
		}
		c.Next()
	}
}

func Login(c *gin.Context) {
	loginName := c.PostForm("loginName")
	loginPwd := c.PostForm("loginPwd")

	if len(loginName) < 3 || len(loginPwd) < 8 || len(loginPwd) > 20 {
		log.Pr("HFish", "127.0.0.1", "illegal username or password")
		c.JSON(http.StatusOK, error.ErrFailLogin)
		return
	}

	resultAdmin, err := dbUtil.DB().Table("hfish_admin").Where("username", loginName).First()
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "get username info error", err)
		c.JSON(http.StatusOK, error.ErrFailLogin)
		return
	}

	if resultAdmin["password"] != loginPwd {
		c.JSON(http.StatusOK, error.ErrFailLogin)
		return
	}

	nowTime := time.Now()
	_, err = dbUtil.DB().
		Table("hfish_admin").
		Data(map[string]interface{}{"last_login_time": nowTime.Format("2006-01-02 15:04:05")}).
		Where("username", loginName).
		Update()
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "update admin login info error", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}

	session := sessions.Default(c)
	session.Set("login_name", loginName)
	session.Set("login_pwd", loginPwd)
	session.Set("expire_time", nowTime.Unix() + k30MinutesSec)
	session.Save()

	c.JSON(http.StatusOK, error.ErrSuccess)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	c.Redirect(http.StatusFound, "/login")
}

func ChangePwd(c *gin.Context) {
	nowPwd := c.PostForm("now_pwd")
	newPwd := c.PostForm("new_pwd")
	dupPwd := c.PostForm("dup_pwd")

	if len(newPwd) < 8 || len(newPwd) > 20 {
		log.Pr("HFish", "127.0.0.1", "illegal new password")
		c.JSON(http.StatusOK, error.ErrInputPwd)
		return
	}

	if newPwd != dupPwd {
		log.Pr("HFish", "127.0.0.1", "new password and reset pwd not same")
		c.JSON(http.StatusOK, error.ErrSamePwd)
		return
	}

	session := sessions.Default(c)
	loginName, _ := session.Get("login_name").(string)
	loginPwd, _ := session.Get("login_pwd").(string)
	if nowPwd != loginPwd {
		log.Pr("HFish", "127.0.0.1", "input admin password error", nowPwd)
		c.JSON(http.StatusOK, error.ErrCheckPwd)
		return
	}

	nowTime := time.Now()
	_, err := dbUtil.DB().
		Table("hfish_admin").
		Data(map[string]interface{}{"password": newPwd, "update_time": nowTime.Format("2006-01-02 15:04:05")}).
		Where("username", loginName).
		Update()
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "update admin login info error", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}
	session.Clear()
	c.JSON(http.StatusOK, error.ErrSuccess)
}

func ResetPwd(c *gin.Context) {
	session := sessions.Default(c)
	loginName, _ := session.Get("login_name").(string)

	if loginName != "admin" {
		log.Pr("HFish", "127.0.0.1", "reset password need admin account")
		c.JSON(http.StatusOK, error.ErrSystem)
		return
	}

	nowTime := time.Now()
	_, err := dbUtil.DB().
		Table("hfish_admin").
		Data(map[string]interface{}{"password": "=HFish@2020=", "update_time": nowTime.Format("2006-01-02 15:04:05")}).
		Where("username", "admin").
		Update()
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "update admin password err", err)
		c.JSON(http.StatusOK, error.ErrUpdateData)
		return
	}
	session.Clear()
	c.JSON(http.StatusOK, error.ErrSuccess)
}

func existFile(fileName string) bool {
	stat, err := os.Stat(fileName)
	if err == nil {
		return !stat.IsDir()
	}
	return false
}

func CheckUpdate(c *gin.Context) {
	content, err := ioutil.ReadFile("version")
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "read version err:", err)
		c.JSON(http.StatusOK, error.ErrCheckNone)
		return
	}

	data := make(map[string]string)
	if err := json.Unmarshal(content, &data); err != nil {
		log.Pr("HFish", "127.0.0.1", "unmarshal version err", err)
		c.JSON(http.StatusOK, error.ErrCheckFail)
		return
	}

	if data["version"] <= "0.6.5" {
		log.Pr("HFish", "127.0.0.1", "no version update", data["version"])
		c.JSON(http.StatusOK, error.ErrCheckNone)
		return
	}
	c.JSON(http.StatusOK, error.ErrSuccessWithData(data))
}

func Upgrade(c *gin.Context) {
	version := c.PostForm("version")
	upgradePackage := fmt.Sprintf("HFish-%s.tar.gz", version)
	if !existFile(upgradePackage) {
		log.Pr("HFish", "127.0.0.1", "upgrade package not exist", version)
		c.JSON(http.StatusOK, error.ErrNoPackage)
		return
	}
	upgradeFlagFile := fmt.Sprintf(".hfish_%s_upgrade", version)
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	err := ioutil.WriteFile(upgradeFlagFile, []byte(nowTime), 0666)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "write upgrade version file err", err)
		c.JSON(http.StatusOK, error.ErrSystem)
		return
	}
	ch := make(chan bool)
	go func() {
		log.Pr("HFish", "127.0.0.1", "upgrade version", version)
		ticker := time.NewTicker(time.Second)
		for i := 0; i < 60; i++ {
			<-ticker.C
			if !existFile(upgradeFlagFile) {
				break
			}
		}
		ch <- true
	}()
	<-ch
	log.Pr("HFish", "127.0.0.1", "upgrade success")
	c.JSON(http.StatusOK, error.ErrSuccess)
}

func Session(c *gin.Context) {
	session := sessions.Default(c)
	loginName, _ := session.Get("login_name").(string)
	loginPwd, _ := session.Get("login_pwd").(string)
	expireTime, _ := session.Get("expire_time").(int64)
	c.JSON(http.StatusOK, gin.H{"login_name": loginName, "expire_time": expireTime, "login_pwd": loginPwd})
}

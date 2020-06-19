package login

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"HFish/error"
	"HFish/utils/conf"
	"time"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Jump(c *gin.Context) {
	account := conf.Get("admin", "account")

	session := sessions.Default(c)
	loginCookie := session.Get("is_login")

	if account != loginCookie {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	} else {
		c.Next()
	}
}

func Login(c *gin.Context) {
	loginName := c.PostForm("loginName")
	loginPwd := c.PostForm("loginPwd")

	account := conf.Get("admin", "account")
	password := conf.Get("admin", "password")

	if loginName == account {
		if loginPwd == password {
			session := sessions.Default(c)
			session.Set("is_login", loginName)
			session.Set("time", time.Now().Format("2006-01-02 15:04:05"))
			session.Save()

			c.JSON(http.StatusOK, gin.H{
				"code": error.ErrSuccessCode,
				"msg":  error.ErrSuccessMsg,
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": error.ErrFailLoginCode,
		"msg":  error.ErrFailLoginMsg,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	c.Redirect(http.StatusFound, "/login")
}

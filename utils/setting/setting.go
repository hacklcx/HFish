package setting

import (
	"HFish/core/exec"
	"HFish/utils/color"
	"HFish/view"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"net/http"
	"time"
	"HFish/utils/conf"
	"HFish/core/protocol/ssh"
	"HFish/core/protocol/redis"
	"HFish/core/protocol/mysql"
	"HFish/core/protocol/httpx"
)

func RunWeb(template string, static string, url string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("web/" + template + "/*")

	// 引入静态资源
	r.Static("/static", "./web/"+static)

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return r
}

func RunDark(template string, static string, url string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("web/" + template + "/*")

	// 引入静态资源
	r.Static("/static", "./web/"+static)

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return r
}

func RunAdmin() http.Handler {
	gin.DisableConsoleColor()

	f, _ := os.Create("./logs/hfish.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 引入gin
	r := gin.Default()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[HFish] %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("admin/*")

	// 引入静态资源
	r.Static("/static", "./static")

	// 加载路由
	view.LoadUrl(r)

	return r
}

func Run() {
	// 启动 HTTP 正向代理
	httpStatus := conf.Get("http", "status")

	// 判断 HTTP 正向代理 是否开启
	if httpStatus == "1" {
		httpAddr := conf.Get("http", "addr")
		go httpx.Start(httpAddr)
	}

	//=========================//

	// 启动 Mysql 钓鱼
	mysqlStatus := conf.Get("mysql", "status")

	// 判断 Mysql 钓鱼 是否开启
	if mysqlStatus == "1" {
		mysqlAddr := conf.Get("mysql", "addr")

		// 利用 Mysql 服务端 任意文件读取漏洞
		mysqlFiles := conf.Get("mysql", "files")

		go mysql.Start(mysqlAddr, mysqlFiles)
	}

	//=========================//

	// 启动 Redis 钓鱼
	redisStatus := conf.Get("redis", "status")

	// 判断 Redis 钓鱼 是否开启
	if redisStatus == "1" {
		redisAddr := conf.Get("redis", "addr")
		go redis.Start(redisAddr)
	}

	//=========================//

	// 启动 SSH 钓鱼
	sshStatus := conf.Get("ssh", "status")

	// 判断 SSG 钓鱼 是否开启
	if sshStatus == "1" {
		sshAddr := conf.Get("ssh", "addr")
		go ssh.Start(sshAddr)
	}

	//=========================//

	// 启动 Web 钓鱼
	webStatus := conf.Get("web", "status")

	// 判断 Web 钓鱼 是否开启
	if webStatus == "1" {
		webAddr := conf.Get("web", "addr")
		webTemplate := conf.Get("web", "template")
		webStatic := conf.Get("web", "static")
		webUrl := conf.Get("web", "url")

		serverWeb := &http.Server{
			Addr:         webAddr,
			Handler:      RunWeb(webTemplate, webStatic, webUrl),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		go serverWeb.ListenAndServe()
	}

	//=========================//

	// 启动 暗网 钓鱼
	darkStatus := conf.Get("dark_net", "status")

	// 判断 暗网 Web 钓鱼 是否开启
	if darkStatus == "1" {
		darkAddr := conf.Get("dark_net", "addr")
		darkTemplate := conf.Get("dark_net", "template")
		darkStatic := conf.Get("dark_net", "static")
		darkUrl := conf.Get("dark_net", "url")

		serverDark := &http.Server{
			Addr:         darkAddr,
			Handler:      RunDark(darkTemplate, darkStatic, darkUrl),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		go serverDark.ListenAndServe()
	}

	//=========================//

	// 启动 admin 管理后台
	adminAddr := conf.Get("admin", "addr")

	serverAdmin := &http.Server{
		Addr:         adminAddr,
		Handler:      RunAdmin(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverAdmin.ListenAndServe()
}

func Init() {
	fmt.Println("test")
}

func Help() {
	exec.Execute("clear")
	logo := ` o
  \_/\o
 ( Oo)                    \|/
 (_=-)  .===O- ~~~b~i~u~~ -O-
 /   \_/U'                /|\
 ||  |_/
 \\  |	     ~ By: HackLC Team
 {K ||       __ _______     __
  | PP      / // / __(_)__ / /
  | ||     / _  / _// (_-</ _ \
  (__\\   /_//_/_/ /_/___/_//_/ v0.1
`
	fmt.Println(color.Yellow(logo))
	fmt.Println(color.White(" A Safe and Active Attack Honeypot Fishing Framework System for Enterprises."))
	fmt.Println("")
	fmt.Println(color.Yellow(" + [ ABOUT ]----------------------------------------------------------- +"))
	fmt.Println("")
	fmt.Println(color.Green("   - Github:"), color.White("https://github.com/hacklcx/HFish"), color.Green(" - Team:"), color.White("https://hack.lc"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + [ ARGUMENTS ]------------------------------------------------------- +"))
	fmt.Println("")
	fmt.Println(color.Cyan("   run,--run"), color.White("	       Start up service"))
	//fmt.Println(color.Cyan("   init,--init"), color.White("		   Initialization, Wipe data"))
	fmt.Println(color.Cyan("   version,--version"), color.White("  HFish Version"))
	fmt.Println(color.Cyan("   help,--help"), color.White("	       Help"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + -------------------------------------------------------------------- +"))
	fmt.Println("")
}

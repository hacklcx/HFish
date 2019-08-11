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
	"HFish/core/protocol/ftp"
	"HFish/core/protocol/telnet"
	"HFish/core/rpc/server"
	"HFish/core/rpc/client"
)

func RunWeb(template string, index string, static string, url string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("web/" + template + "/*")

	// 引入静态资源
	r.Static("/static", "./web/"+static)

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, index, gin.H{})
	})

	return r
}

func RunDeep(template string, index string, static string, url string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("web/" + template + "/*")

	// 引入静态资源
	r.Static("/static", "./web/"+static)

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, index, gin.H{})
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
	// 启动 FTP 蜜罐
	ftpStatus := conf.Get("ftp", "status")

	// 判断 FTP 蜜罐 是否开启
	if ftpStatus == "1" {
		ftpAddr := conf.Get("ftp", "addr")
		go ftp.Start(ftpAddr)
	}

	//=========================//

	// 启动 Telnet 蜜罐
	telnetStatus := conf.Get("telnet", "status")

	// 判断 Telnet 蜜罐 是否开启
	if telnetStatus == "1" {
		telnetAddr := conf.Get("telnet", "addr")
		go telnet.Start(telnetAddr)
	}

	//=========================//

	//// 启动 HTTP 正向代理
	//httpStatus := conf.Get("http", "status")
	//
	//// 判断 HTTP 正向代理 是否开启
	//if httpStatus == "1" {
	//	httpAddr := conf.Get("http", "addr")
	//	go httpx.Start(httpAddr)
	//}

	//=========================//

	// 启动 Mysql 蜜罐
	mysqlStatus := conf.Get("mysql", "status")

	// 判断 Mysql 蜜罐 是否开启
	if mysqlStatus == "1" {
		mysqlAddr := conf.Get("mysql", "addr")

		// 利用 Mysql 服务端 任意文件读取漏洞
		mysqlFiles := conf.Get("mysql", "files")

		go mysql.Start(mysqlAddr, mysqlFiles)
	}

	//=========================//

	// 启动 Redis 蜜罐
	redisStatus := conf.Get("redis", "status")

	// 判断 Redis 蜜罐 是否开启
	if redisStatus == "1" {
		redisAddr := conf.Get("redis", "addr")
		go redis.Start(redisAddr)
	}

	//=========================//

	// 启动 SSH 蜜罐
	sshStatus := conf.Get("ssh", "status")

	// 判断 SSG 蜜罐 是否开启
	if sshStatus == "1" {
		sshAddr := conf.Get("ssh", "addr")
		go ssh.Start(sshAddr)
	}

	//=========================//

	// 启动 Web 蜜罐
	webStatus := conf.Get("web", "status")

	// 判断 Web 蜜罐 是否开启
	if webStatus == "1" {
		webAddr := conf.Get("web", "addr")
		webTemplate := conf.Get("web", "template")
		webStatic := conf.Get("web", "static")
		webUrl := conf.Get("web", "url")
		webIndex := conf.Get("web", "index")

		serverWeb := &http.Server{
			Addr:         webAddr,
			Handler:      RunWeb(webTemplate, webIndex, webStatic, webUrl),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		go serverWeb.ListenAndServe()
	}

	//=========================//

	// 启动 暗网 蜜罐
	deepStatus := conf.Get("deep", "status")

	// 判断 暗网 Web 蜜罐 是否开启
	if deepStatus == "1" {
		deepAddr := conf.Get("deep", "addr")
		deepTemplate := conf.Get("deep", "template")
		deepStatic := conf.Get("deep", "static")
		deepkUrl := conf.Get("deep", "url")
		deepIndex := conf.Get("deep", "index")

		serverDark := &http.Server{
			Addr:         deepAddr,
			Handler:      RunDeep(deepTemplate, deepIndex, deepStatic, deepkUrl),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		go serverDark.ListenAndServe()
	}

	//=========================//

	// 启动 RPC
	rpcStatus := conf.Get("rpc", "status")

	// 判断 RPC 是否开启 1 RPC 服务端 2 RPC 客户端
	if rpcStatus == "1" {
		// 服务端监听地址
		rpcAddr := conf.Get("rpc", "addr")
		go server.Start(rpcAddr)
	} else if rpcStatus == "2" {
		// 客户端连接服务端
		// 阻止进程，不启动 admin

		rpcName := conf.Get("rpc", "name")

		for {
			// 这样写 提高IO读写性能
			go client.Start(rpcName, ftpStatus, telnetStatus, "0", mysqlStatus, redisStatus, sshStatus, webStatus, deepStatus)

			time.Sleep(time.Duration(1) * time.Minute)
		}
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
  (__\\   /_//_/_/ /_/___/_//_/ v0.2
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

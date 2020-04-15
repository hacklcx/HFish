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
	"HFish/utils/cache"
	"HFish/utils/conf"
	"HFish/core/protocol/ssh"
	"HFish/core/protocol/redis"
	"HFish/core/protocol/mysql"
	"HFish/core/protocol/ftp"
	"HFish/core/protocol/telnet"
	"HFish/core/protocol/custom"
	"HFish/core/rpc/server"
	"HFish/core/rpc/client"
	"HFish/view/api"
	"HFish/utils/cors"
	"HFish/core/protocol/memcache"
	"HFish/core/protocol/tftp"
	"HFish/core/protocol/httpx"
	"HFish/core/protocol/elasticsearch"
	"HFish/core/protocol/vnc"
	"HFish/core/dbUtil"
	"strconv"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"syscall"
	"HFish/utils/ping"
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

	// API 启用状态
	apiStatus := conf.Get("api", "status")

	// 判断 API 是否启用
	if apiStatus == "1" {
		// 启动 WEB蜜罐 API
		r.Use(cors.Cors())
		webUrl := conf.Get("api", "web_url")
		r.POST(webUrl, api.ReportWeb)
	}

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

	// API 启用状态
	apiStatus := conf.Get("api", "status")

	// 判断 API 是否启用
	if apiStatus == "1" {
		// 启动 暗网蜜罐 API
		r.Use(cors.Cors())
		deepUrl := conf.Get("api", "deep_url")
		r.POST(deepUrl, api.ReportDeepWeb)
	}

	return r
}

func RunPlug() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// API 启用状态
	apiStatus := conf.Get("api", "status")

	// 判断 API 是否启用
	if apiStatus == "1" {
		// 启动 蜜罐插件 API
		r.Use(cors.Cors())
		plugUrl := conf.Get("api", "plug_url")
		r.POST(plugUrl, api.ReportPlugWeb)
	}

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

	store := cookie.NewStore([]byte("HFish"))
	r.Use(sessions.Sessions("HFish", store))

	r.Use(gin.Recovery())

	// 引入html资源
	r.LoadHTMLGlob("admin/*")

	// 引入静态资源
	r.Static("/static", "./static")

	// 加载路由
	view.LoadUrl(r)

	return r
}

// 初始化缓存
func initCahe() {
	resultMail, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "alertMail").First()
	resultHook, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "webHook").First()
	resultIp, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "whiteIp").First()
	resultPasswd, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "passwdTM").First()

	cache.Setx("MailConfigStatus", strconv.FormatInt(resultMail["status"].(int64), 10))
	cache.Setx("MailConfigInfo", resultMail["info"])

	cache.Setx("HookConfigStatus", strconv.FormatInt(resultHook["status"].(int64), 10))
	cache.Setx("HookConfigInfo", resultHook["info"])

	cache.Setx("IpConfigStatus", strconv.FormatInt(resultIp["status"].(int64), 10))
	cache.Setx("IpConfigInfo", resultIp["info"])

	cache.Setx("PasswdConfigStatus", strconv.FormatInt(resultPasswd["status"].(int64), 10))
	cache.Setx("PasswdConfigInfo", resultPasswd["info"])
}

func Run() {
	ping.Ping()

	// 启动 自定义 蜜罐
	custom.StartCustom()

	// 启动 vnc  蜜罐
	vncStatus := conf.Get("vnc", "status")

	// 判断 vnc 蜜罐 是否开启
	if vncStatus == "1" {
		vncAddr := conf.Get("vnc", "addr")
		go vnc.Start(vncAddr)

	}

	//=========================//

	// 启动 elasticsearch 蜜罐
	esStatus := conf.Get("elasticsearch", "status")

	// 判断 elasticsearch 蜜罐 是否开启
	if esStatus == "1" {
		esAddr := conf.Get("elasticsearch", "addr")
		go elasticsearch.Start(esAddr)
	}

	//=========================//

	// 启动 TFTP 蜜罐
	tftpStatus := conf.Get("tftp", "status")

	// 判断 TFTP 蜜罐 是否开启
	if tftpStatus == "1" {
		tftpAddr := conf.Get("tftp", "addr")
		go tftp.Start(tftpAddr)
	}

	//=========================//

	// 启动 MemCache 蜜罐
	memCacheStatus := conf.Get("mem_cache", "status")

	// 判断 MemCache 蜜罐 是否开启
	if memCacheStatus == "1" {
		memCacheAddr := conf.Get("mem_cache", "addr")
		go memcache.Start(memCacheAddr, "4")
	}

	//=========================//

	// 启动 FTP 蜜罐
	ftpStatus := conf.Get("ftp", "status")

	// 判断 FTP 蜜罐 是否开启
	if ftpStatus != "0" {
		ftpAddr := conf.Get("ftp", "addr")
		go ftp.Start(ftpAddr)
	}

	//=========================//

	// 启动 Telnet 蜜罐
	telnetStatus := conf.Get("telnet", "status")

	// 判断 Telnet 蜜罐 是否开启
	if telnetStatus != "0" {
		telnetAddr := conf.Get("telnet", "addr")
		go telnet.Start(telnetAddr)
	}

	//=========================//

	// 启动 HTTP 正向代理
	httpStatus := conf.Get("http", "status")

	// 判断 HTTP 正向代理 是否开启
	if httpStatus == "1" {
		httpAddr := conf.Get("http", "addr")
		go httpx.Start(httpAddr)
	}

	//=========================//

	// 启动 Mysql 蜜罐
	mysqlStatus := conf.Get("mysql", "status")

	// 判断 Mysql 蜜罐 是否开启
	if mysqlStatus != "0" {
		mysqlAddr := conf.Get("mysql", "addr")

		// 利用 Mysql 服务端 任意文件读取漏洞
		mysqlFiles := conf.Get("mysql", "files")

		go mysql.Start(mysqlAddr, mysqlFiles)
	}

	//=========================//

	// 启动 Redis 蜜罐
	redisStatus := conf.Get("redis", "status")

	// 判断 Redis 蜜罐 是否开启
	if redisStatus != "0" {
		redisAddr := conf.Get("redis", "addr")
		go redis.Start(redisAddr)
	}

	//=========================//

	// 启动 SSH 蜜罐
	sshStatus := conf.Get("ssh", "status")

	// 判断 SSG 蜜罐 是否开启
	if sshStatus != "0" {
		sshAddr := conf.Get("ssh", "addr")
		go ssh.Start(sshAddr)
	}

	//=========================//

	// 启动 Web 蜜罐
	webStatus := conf.Get("web", "status")

	// 判断 Web 蜜罐 是否开启
	if webStatus != "0" {
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
	if deepStatus != "0" {
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

	// 启动 蜜罐插件
	plugStatus := conf.Get("plug", "status")

	// 判断 蜜罐插件 是否开启
	if plugStatus != "0" {
		plugAddr := conf.Get("plug", "addr")

		serverPlug := &http.Server{
			Addr:         plugAddr,
			Handler:      RunPlug(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		go serverPlug.ListenAndServe()
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

		client.RpcInit()

		for {
			// 判断自定义蜜罐是否启动
			customStatus := "0"

			customNames := conf.GetCustomName()
			if len(customNames) > 0 {
				customStatus = "1"
			}

			// 这样写 提高IO读写性能
			go client.Start(rpcName, ftpStatus, telnetStatus, httpStatus, mysqlStatus, redisStatus, sshStatus, webStatus, deepStatus, memCacheStatus, plugStatus, esStatus, tftpStatus, vncStatus, customStatus)

			time.Sleep(time.Duration(1) * time.Minute)
		}
	}

	//=========================//
	// 初始化缓存
	initCahe()

	// 启动 admin 管理后台
	adminAddr := conf.Get("admin", "addr")

	serverAdmin := &http.Server{
		Addr:         adminAddr,
		Handler:      RunAdmin(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("pid is %d", syscall.Getpid())

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
  (__\\   /_//_/_/ /_/___/_//_/ v0.6.1
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

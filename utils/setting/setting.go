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

	// Import html resources
	r.LoadHTMLGlob("web/" + template + "/*")

	// Introduce static resources
	r.Static("/static", "./web/"+static)

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, index, gin.H{})
	})

	// API Enabled state
	apiStatus := conf.Get("api", "status")

	// Determine whether the API is enabled
	if apiStatus == "1" {
		// Start WEB honeypot API
		r.Use(cors.Cors())
		webUrl := conf.Get("api", "web_url")
		r.POST(webUrl, api.ReportWeb)
	}

	return r
}

func RunDeep(template string, index string, static string, url string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// Import html resources
	r.LoadHTMLGlob("web/" + template + "/*")

	// Introduce static resources
	r.Static("/static", "./web/"+static)

	r.GET(url, func(c *gin.Context) {
		c.HTML(http.StatusOK, index, gin.H{})
	})

	// API Enabled state
	apiStatus := conf.Get("api", "status")

	// Determine whether the API is enabled
	if apiStatus == "1" {
		// Start Darknet Honeypot API
		r.Use(cors.Cors())
		deepUrl := conf.Get("api", "deep_url")
		r.POST(deepUrl, api.ReportDeepWeb)
	}

	return r
}

func RunPlug() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	// API Enabled state
	apiStatus := conf.Get("api", "status")

	// Determine whether the API is enabled
	if apiStatus == "1" {
		// Start the honeypot plugin API
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

	// Introduce gin
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

	// Import html resources
	r.LoadHTMLGlob("admin/*")

	// Introduce static resources
	r.Static("/static", "./static")

	// Load route
	view.LoadUrl(r)

	return r
}

// Initialize the cache
func initCahe() {
	resultMail, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "alertMail").First()
	resultHook, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "webHook").First()
	resultIp, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "whiteIp").First()
	resultPasswd, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "passwdTM").First()
	resultApikey, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "apikey").First()
	resultSyslog, _ := dbUtil.DB().Table("hfish_setting").Fields("status", "info").Where("type", "=", "syslog").First()

	cache.Setx("MailConfigStatus", strconv.FormatInt(resultMail["status"].(int64), 10))
	cache.Setx("MailConfigInfo", resultMail["info"])

	cache.Setx("HookConfigStatus", strconv.FormatInt(resultHook["status"].(int64), 10))
	cache.Setx("HookConfigInfo", resultHook["info"])

	cache.Setx("IpConfigStatus", strconv.FormatInt(resultIp["status"].(int64), 10))
	cache.Setx("IpConfigInfo", resultIp["info"])

	cache.Setx("PasswdConfigStatus", strconv.FormatInt(resultPasswd["status"].(int64), 10))
	cache.Setx("PasswdConfigInfo", resultPasswd["info"])

	cache.Setx("ApikeyStatus", strconv.FormatInt(resultApikey["status"].(int64), 10))
	cache.Setx("ApikeyInfo", resultApikey["info"])

	cache.Setx("SyslogConfigStatus", strconv.FormatInt(resultSyslog["status"].(int64), 10))
	cache.Setx("SyslogConfigInfo", resultSyslog["info"])
}

func Run() {
	ping.Ping()

	// Start custom honeypot
	custom.StartCustom()

	// Start vnc honeypot
	vncStatus := conf.Get("vnc", "status")

	// Determine whether the vnc honeypot is open
	if vncStatus == "1" {
		vncAddr := conf.Get("vnc", "addr")
		go vnc.Start(vncAddr)

	}

	//=========================//

	// Active elasticsearch honeypot
	esStatus := conf.Get("elasticsearch", "status")

	// Whether to enable elasticsearch honeypot
	if esStatus == "1" {
		esAddr := conf.Get("elasticsearch", "addr")
		go elasticsearch.Start(esAddr)
	}

	//=========================//

	// Start TFTP honeypot
	tftpStatus := conf.Get("tftp", "status")

	// Wheather to enable Start TFTP honeypot
	if tftpStatus == "1" {
		tftpAddr := conf.Get("tftp", "addr")
		go tftp.Start(tftpAddr)
	}

	//=========================//

	// Start MemCache honeypot
	memCacheStatus := conf.Get("mem_cache", "status")

	// Weather to enale Start MemCache honeypot
	if memCacheStatus == "1" {
		memCacheAddr := conf.Get("mem_cache", "addr")
		go memcache.Start(memCacheAddr, "4")
	}

	//=========================//

	// Start FTP honeypot
	ftpStatus := conf.Get("ftp", "status")

	// Weather to enable Start FTP honeypot
	if ftpStatus != "0" {
		ftpAddr := conf.Get("ftp", "addr")
		go ftp.Start(ftpAddr)
	}

	//=========================//

	// Start Telnet honeypot
	telnetStatus := conf.Get("telnet", "status")

	// Weather to enable Start Telnet honeypot
	if telnetStatus != "0" {
		telnetAddr := conf.Get("telnet", "addr")
		go telnet.Start(telnetAddr)
	}

	//=========================//

	// Start HTTP forward proxy
	httpStatus := conf.Get("http", "status")

	// Weather to enable Start HTTP forward proxy
	if httpStatus == "1" {
		httpAddr := conf.Get("http", "addr")
		go httpx.Start(httpAddr)
	}

	//=========================//

	// Start Mysql honeypot
	mysqlStatus := conf.Get("mysql", "status")

	// Weather to enable Start Mysql honeypot
	if mysqlStatus != "0" {
		mysqlAddr := conf.Get("mysql", "addr")

		// Exploiting arbitrary file reading vulnerability on Mysql server
		mysqlFiles := conf.Get("mysql", "files")

		go mysql.Start(mysqlAddr, mysqlFiles)
	}

	//=========================//

	// Start Redis honeypot
	redisStatus := conf.Get("redis", "status")

	// Weather to enable Start Redis honeypot
	if redisStatus != "0" {
		redisAddr := conf.Get("redis", "addr")
		go redis.Start(redisAddr)
	}

	//=========================//

	// Start SSH honeypot
	sshStatus := conf.Get("ssh", "status")

	// Determine whether the SSG honeypot is open
	if sshStatus != "0" {
		sshAddr := conf.Get("ssh", "addr")
		go ssh.Start(sshAddr)
	}

	//=========================//

	// Start the web honeypot
	webStatus := conf.Get("web", "status")

	// Determine whether the web honeypot is open
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

	// Start dark web honeypot
	deepStatus := conf.Get("deep", "status")

	// Determine whether the dark web honeypot is open
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

	// Start the honeypot plugin
	plugStatus := conf.Get("plug", "status")

	// Determine whether the honeypot plugin is enabled
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

	// start up RPC
	rpcStatus := conf.Get("rpc", "status")

	// Determine whether RPC is enabled 1 RPC server 2 RPC client
	if rpcStatus == "1" {
		// Server listening address
		rpcAddr := conf.Get("rpc", "addr")
		go server.Start(rpcAddr)
	} else if rpcStatus == "2" {
		// Client connects to server
		// Block the process, do not start admin

		rpcName := conf.Get("rpc", "name")

		client.RpcInit()

		for {
			// Determine whether the custom honeypot is activated
			customStatus := "0"

			customNames := conf.GetCustomName()
			if len(customNames) > 0 {
				customStatus = "1"
			}

			// Write like this to improve IO read and write performance
			go client.Start(rpcName, ftpStatus, telnetStatus, httpStatus, mysqlStatus, redisStatus, sshStatus, webStatus, deepStatus, memCacheStatus, plugStatus, esStatus, tftpStatus, vncStatus, customStatus)

			time.Sleep(time.Duration(1) * time.Minute)
		}
	}

	//=========================//
	// Initialize the cache
	initCahe()

	// Start admin management background
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
  (__\\   /_//_/_/ /_/___/_//_/ v0.6.4
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

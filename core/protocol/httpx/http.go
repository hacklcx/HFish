package httpx

import (
	"net/http"
	"github.com/elazarl/goproxy"
	"net/url"
	"fmt"
)

/*http 正向代理*/

func Start(addr string, proxyUrl string) {

	gp := goproxy.NewProxyHttpServer()
	pu, err := url.Parse(proxyUrl)
	if err == nil {
		gp.Tr.Proxy = http.ProxyURL(&url.URL{
			Scheme: pu.Scheme,
			Host: pu.Host,
		})
	}

	gp.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	gp.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		// Report Send
		fmt.Println(req.RemoteAddr)
		return req, nil
	})
	http.ListenAndServe(addr, gp)
}

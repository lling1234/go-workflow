package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"act/api/internal/config"
	"act/api/internal/handler"
	"act/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/act-api.yaml", "the config file")

func swaggerUrl(w http.ResponseWriter, r *http.Request) {
	//这里定义代理目标地址，这里定义到本机9091端口
	u, _ := url.Parse("http://ling11.top:1122/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	// server := rest.MustNewServer(c.RestConf,rest.WithCors("http://ling11.top"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	server.AddRoute(rest.Route{Method: "GET", Path: "/swagger", Handler: swaggerUrl})
	server.AddRoute(rest.Route{Method: "GET", Path: "/swagger/swagger-ui.css", Handler: swaggerUrl})
	server.AddRoute(rest.Route{Method: "GET", Path: "/swagger/swagger-ui-bundle.js", Handler: swaggerUrl})
	server.AddRoute(rest.Route{Method: "GET", Path: "/swagger/swagger-ui-standalone-preset.js", Handler: swaggerUrl})
	server.AddRoute(rest.Route{Method: "GET", Path: "/swagger/favicon-32x32.png", Handler: swaggerUrl})
	server.AddRoute(rest.Route{Method: "GET", Path: "/swagger/favicon-16x16.png", Handler: swaggerUrl})
	server.AddRoute(rest.Route{Method: "GET", Path: "/swagger/swagger.json", Handler: swaggerUrl})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

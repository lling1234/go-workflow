package main

import (
	"flag"
	"fmt"
	"net/http"

	"go-wflow/app/wflow/api/internal/config"
	"go-wflow/app/wflow/api/internal/handler"
	"go-wflow/app/wflow/api/internal/svc"

	"github.com/qkbyte/go-zero/core/conf"
	"github.com/qkbyte/go-zero/rest"
)

var configFile = flag.String("f", "etc/wflow-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// swagger
	routes := make([]rest.Route, 0)
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/"))).ServeHTTP})
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/index.html", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/"))).ServeHTTP})
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/swagger.json", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/"))).ServeHTTP})
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/swagger-ui.css", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/css/"))).ServeHTTP})
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/favicon-32x32.png", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/image/"))).ServeHTTP})
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/favicon-16x16.png", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/image/"))).ServeHTTP})
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/swagger-ui-bundle.js", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/js/"))).ServeHTTP})
	routes = append(routes, rest.Route{Method: http.MethodGet, Path: "/swagger/swagger-ui-standalone-preset.js", Handler: http.StripPrefix("/swagger/", http.FileServer(http.Dir("../../../doc/swagger/js/"))).ServeHTTP})

	server.AddRoutes(routes)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

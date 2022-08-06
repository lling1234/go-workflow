package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	router "github.com/qkbyte/go-workflow/router"

	config "github.com/qkbyte/go-workflow/config"

	model "github.com/qkbyte/go-workflow/engine/model"
	"github.com/qkbyte/go-workflow/engine/service"
)

func main() {
	mux := router.Mux
	// 启动数据库连接
	model.Setup()
	// 启动redis连接
	model.SetRedis()
	// 启动定时任务
	service.CronJobs()
	// 配置111
	var conf = *config.Config
	// 启动服务
	readTimeout, err := strconv.Atoi(conf.ReadTimeout)
	if err != nil {
		panic(err)
	}
	writeTimeout, err := strconv.Atoi(conf.WriteTimeout)
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        mux,
		ReadTimeout:    time.Duration(readTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(writeTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("the application start up at port%s", server.Addr)
	if conf.TLSOpen == "true" {
		err = server.ListenAndServeTLS(conf.TLSCrt, conf.TLSKey)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/pkg/conf"
	"log"
	"net/http"
	"time"
)

type httpConf struct {
	runMode      string
	httpPort     int
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// init a http server
func New() (srv *http.Server) {

	e := gin.New()
	// get http configuration
	cfg := getConf()
	// set logger
	e.Use(gin.Logger())
	// set http run mode
	gin.SetMode(cfg.runMode)
	// initialize url request router
	InitRoutes(e)
	// set http server
	srv = &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.httpPort),
		Handler:        e,
		ReadTimeout:    cfg.readTimeout,
		WriteTimeout:   cfg.writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return

}

// get section "server" from configuration file "app.ini"
func getConf() (httpConf) {

	var cfg httpConf

	sec, err := conf.IniFile.GetSection("server")

	if err != nil {
		log.Fatalln("Fail to get section 'mode': ", err)
	}

	cfg.runMode  = sec.Key("RUN_MODE").String()
	cfg.httpPort = sec.Key("HTTP_PORT").MustInt(8080)
	cfg.readTimeout  = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	cfg.writeTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	return cfg
}




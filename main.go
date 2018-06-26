package main

import (
	"flag"
	log "github.com/cihub/seelog"
	"go_service/helpers"
	"go_service/routers"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	//读取配置
	configFilePath := flag.String("C", "conf/app.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")
	flag.Parse()

	logger, err := log.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		panic(err)
	}

	log.ReplaceLogger(logger)
	defer log.Flush()

	if err := helpers.LoadConfiguration(*configFilePath); err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	//运行模式
	if helpers.GetConfiguration().RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(helpers.GetConfiguration().AppHost)
}

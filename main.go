package main

import (
	"flag"
	"github.com/cihub/seelog"
	"go_service/helpers"
	"go_service/routers"
	"go_service/system"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	//读取配置
	configFilePath := flag.String("C", "conf/app.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")
	flag.Parse()

	logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		panic(err)
	}

	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	if err := system.LoadConfiguration(*configFilePath); err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	//运行模式
	if system.GetConfiguration().RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	helpers.InitRedis()

	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(system.GetConfiguration().AppHost)
}

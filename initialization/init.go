package initialization

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
)

var (
	ConfCenterService = &ConfCenter{}
	log               = &confCenterLog{}
)

type (
	ConfCenter struct {
		Addr string
	}
	confCenterLog struct {
		path  string
		level string
	}
)

func Init() error {
	initConfig()
	initLog()
	clearLog()
	return nil
}

func initConfig() {
	ConfCenterService.Addr = beego.AppConfig.String("confcenter_addr")
	if len(ConfCenterService.Addr) == 0 {
		logs.Error("read ConfCenter addr is null")
		panic(errors.New("read ConfCenter addr is null"))
		return
	}

	log.path = beego.AppConfig.String("log_path")
	if len(log.path) == 0 {
		logs.Error("read config log_path err ")
		panic(errors.New("read config log_path err "))
		return

	}

	log.level = beego.AppConfig.String("log_level")
	if len(log.level) == 0 {
		logs.Error("read config log_level err")
		panic(errors.New("read config log_level err"))
		return
	}
}

func initLog() {
	err := initLogger(log.path, log.level)
	if err != nil {
		panic(err)
		return
	}
}

func clearLog() {
	os.Truncate(log.path, 0)
}

func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}

func initLogger(logPath string, logLevel string) (err error) {

	config := make(map[string]interface{})
	config["filename"] = logPath
	config["level"] = convertLogLevel(logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

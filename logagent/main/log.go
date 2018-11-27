package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func converlogLevel(level string) int  {
	//var loglevel int
	switch (level) {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "error":
		return logs.LevelError
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug

}

func initlogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] =appConfig.logPath
	config["level"] = converlogLevel(appConfig.logLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("marshal failed,err: %v", err)
		return
	}
	logs.SetLogger(logs.AdapterConsole,string(configStr))

	return
}

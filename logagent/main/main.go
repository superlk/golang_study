package main

import (
	"awesomeProject/src/logagent/kafka"
	"awesomeProject/src/logagent/tailf"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "/Users/superlk/go/src/awesomeProject/src/logagent/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("laod conf failed,err:%v\n", err)
		panic("load conf error")
		return
	}
	err = initlogger()
	if err != nil {
		fmt.Sprintf("laod log failed,err:%v\n", err)
		panic("load log error")
		return
	}

	logs.Debug("load conf success config %v", appConfig)

	err = tailf.InitTail(appConfig.collectConf,appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed ,err:%v", err)
		return
	}
	logs.Debug("init  success")
	err = serverRun()
	if err != nil {
		logs.Error("server failed run  err:%v", err)
		return
	}
	logs.Info("program exited")
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed ,err:%v", err)
		return
	}
}

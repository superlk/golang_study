package main

import (
	"awesomeProject/src/logagent/tailf"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafa(msg)
		if err != nil {
			logs.Error("send to kafka failed .error", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return

}

func sendToKafa(msg *tailf.TextMsg)(err error)  {
	//err = kafka.SendToKafka(msg.Msg,msg.Topic)
	fmt.Printf("read msg:%s,topic:%s\n",msg.Msg,msg.Topic)
	return
}
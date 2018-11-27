package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}
type TailObjmgr struct {
	tailObjs   []*TailObj
	msgChan chan *TextMsg
}

var tailObjMgr *TailObjmgr

func GetOneLine()(msg *TextMsg)  {
	msg = <- tailObjMgr.msgChan
	return
}

func InitTail(conf []CollectConf, chanSize int) (err error) {
	if len(conf) == 0 {
		err = fmt.Errorf("invalid config for collect ,conf:%v", conf)
		return
	}
	tailObjMgr = &TailObjmgr{
		msgChan: make(chan *TextMsg, chanSize),
		tailObjs: make([]*TailObj,1),
	}
	for _, v := range conf {
		obj := &TailObj{
			conf: v,
		}
		tails, Terr := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			MustExist: false,
			Poll:      true,
		})
		if Terr != nil {
			fmt.Println("tail file err:", Terr)
			return Terr
		}
		obj.tail = tails
		fmt.Printf("obj===%+v\n",obj)
		fmt.Println(tailObjMgr,obj)
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)
		go readFromTail(obj)
	}

	return
}

func readFromTail(tailobj *TailObj) {
	for true {
		line, ok := <-tailobj.tail.Lines
		if !ok {
			logs.Info("0000")
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textmsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailobj.conf.Topic,
		}
		tailObjMgr.msgChan <- textmsg

	}

}

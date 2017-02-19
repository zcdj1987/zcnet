package main

import (
	"github.com/Sirupsen/logrus"
	// "github.com/paulstuart/ping"
	"os"
	"strconv"
	"time"
)

const (
	CONST_SYS_LOCAL_PORT = ":18666"
	CONST_SYS_KING_PORT  = ":16888"
	CONST_SYS_LOCAL_NAME = "GATE"
)

var (
	log         = logrus.New()
	addressKing string
	model       int
	masterList  []string
)

func main() {

	// 初始化启动传参数，必须传入king服务器ip和启动模式:0=正常模式；1=debug模式
	arg_num := len(os.Args)

	if arg_num < 3 {
		log.Println("传入参数不足,服务终止")
	}

	addressKing = os.Args[1]
	// if !ping.Ping(addressKing, 1) {
	// 	log.Println("king服务器未上线,服务终止")
	// }

	model, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Errorln(err)
		log.Println("模式参数错误，服务终止")
	}

	// 获取king服务器相关配置
	rpc_king_init()

	log.Println(addressKing, model)
	time.Sleep(time.Hour)
}

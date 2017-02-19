package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/paulstuart/ping"
	"os"
	"strconv"
	"time"
)

const (
	CONST_SYS_PORT = ":18666"
	CONST_SYS_NAME = "GATE"
)

var (
	address_king string
	model        int
)

func main() {

	// 启动传参数，必须传入king服务器ip和启动模式:0=正常模式；1=debug模式
	arg_num := len(os.Args)

	if arg_num < 4 {
		log.Println("传入参数不足,服务终止")
		os.Exit(1)
	}

	address_king = os.Args[1]
	if !ping.Ping(address_king, 1) {
		log.Println("king服务器未上线,服务终止")
		os.Exit(1)
	}

	model, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Errorln(err)
		log.Println("模式参数错误，服务终止")
		os.Exit(1)
	}

	// 获取系统
	log.Println(address_king, model)
	time.Sleep(time.Hour)
}

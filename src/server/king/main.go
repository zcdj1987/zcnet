package main

import (
	"github.com/Sirupsen/logrus"
	rpc "github.com/hprose/hprose-golang/rpc/fasthttp"
	"github.com/valyala/fasthttp"
	"time"
)

const (
	CONST_SYS_KING_PORT  = ":16888"
	CONST_SYS_LOCAL_NAME = "KING"
)

var (
	log = logrus.New()
)

var (
	service *rpc.FastHTTPService
)

func Get_master_server() (st []string, err error) {
	st = []string{"192.168.0.1"}
	return
}

func main() {
	service := rpc.NewFastHTTPService()
	service.AddFunction("Get_master_server", Get_master_server)
	fasthttp.ListenAndServe(CONST_SYS_KING_PORT, service.ServeFastHTTP)
	time.Sleep(time.Hour)
}

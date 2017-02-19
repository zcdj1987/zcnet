package king

import (
	log "github.com/Sirupsen/logrus"
	rpc "github.com/hprose/hprose-golang/rpc/fasthttp"
	"github.com/valyala/fasthttp"
)

var (
	service *rpc.FastHTTPService
)

func main() {
	service := rpc.NewFastHTTPService()
	fasthttp.ListenAndServe(":8080", service.ServeFastHTTP)
	select {}
}

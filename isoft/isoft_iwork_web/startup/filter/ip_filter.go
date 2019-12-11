package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

var whiteIps string

func init() {
	whiteIps = beego.AppConfig.String("iwork.check.white.ips")
}

func IPFilterFunc(ctx *context.Context) {
	if checkWhite(ctx.Request.URL.String()) {
		return
	}
	if !checkIp(ctx.Input.IP()) {
		ctx.ResponseWriter.WriteHeader(401)
	}
}

func checkWhite(url string) bool {
	return strings.HasPrefix(url, "/api/iwork/httpservice/")
}

func checkIp(ip string) bool {
	return strings.Contains(whiteIps, ip)
}

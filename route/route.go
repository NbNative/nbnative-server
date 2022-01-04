package route

import (
	"github.com/NbNative/nbnative-server/controller"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controller.HelloController{})
	esconfig, err := config.GetSection("es")
	if err != nil {
		logs.Error("es config is empty", err)
	}
	jaegerController, err := controller.NewJaegerController(esconfig)
	if err != nil {
		logs.Error("create jaeger controller is fail", err)
	}
	web.Router("/jaeger", jaegerController)
}

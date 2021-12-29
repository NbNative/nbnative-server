package route

import (
	"github.com/NbNative/nbnative-server/controller"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controller.HelloController{})
}

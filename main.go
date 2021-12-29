package main

import (
	_ "github.com/NbNative/nbnative-server/route"
	_ "github.com/NbNative/nbnative-server/util"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}

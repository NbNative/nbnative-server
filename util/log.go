package util

import "github.com/beego/beego/v2/core/logs"

func init() {
	logs.EnableFuncCallDepth(true)
	err := logs.SetLogger("file", `{"filename":"logs/server.log"}`)
	if err != nil {
		logs.Error("init file log export is panic", err)
	}
}

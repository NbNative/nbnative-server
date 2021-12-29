package controller

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	web.Controller
}

func (c *HelloController) Get() {
	c.SetData(map[string]interface{}{"msg": "hello word"})
	err := c.ServeJSON()
	if err != nil {
		logs.Error(err)
	}
}

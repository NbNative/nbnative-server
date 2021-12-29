package controller

import (
	"github.com/NbNative/nbnative-server/process"
	"github.com/beego/beego/v2/server/web"
	apiV1 "k8s.io/api/core/v1"
)

type CollectEventController struct {
	web.Controller
	EventProcess process.EventProcess
}

func NewCollectEventController() *CollectEventController {
	return &CollectEventController{
		EventProcess: process.NewEsProcessor(),
	}
}

func (c *CollectEventController) Exporter() {
	k8sEvent := apiV1.Event{}
	err := c.ParseForm(&k8sEvent)
	if err != nil {
		return
	}
	_, err = c.EventProcess.Process(&k8sEvent)
	if err != nil {
		return
	}
}

package controller

import (
	"github.com/NbNative/nbnative-server/drive/es"
	"github.com/NbNative/nbnative-server/storage"
	"github.com/beego/beego/v2/server/web"
	"strings"
)

type JaegerController struct {
	web.Controller
	storage storage.Storage
}

type Options struct {
}

func NewJaegerController(c map[string]string) (*JaegerController, error) {
	esClientStorage, err := storage.NewEsClientStorage(es.ClientConfig{
		Addresses: strings.Split(c["address"], ","),
		Username:  c["user-name"],
		Password:  c["password"],
	})
	if err != nil {
		return nil, err
	}
	return &JaegerController{
		storage: esClientStorage,
	}, nil
}

func (j *JaegerController) Traces() {
	j.storage.GetService()
}

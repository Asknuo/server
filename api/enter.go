package api

import "sever/service"

type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService

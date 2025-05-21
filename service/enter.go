package service

type ServiceGroup struct {
	BaseService
	EsService
	JwtService
	GaodeService
}

var ServiceGroupApp = new(ServiceGroup)

package application

import (
	"your/path/project/domain_mydomain/controller/restapi"
	"your/path/project/domain_mydomain/gateway/prod2"
	"your/path/project/domain_mydomain/usecase/getallperson"
	"your/path/project/domain_mydomain/usecase/runpersoncreate"
	"your/path/project/shared/driver"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"
	"your/path/project/shared/infrastructure/server"
	"your/path/project/shared/infrastructure/util"
)

type myperson struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c myperson) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewMyperson() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("myperson", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandlerDefault(log, appData, cfg)

		datasource := prod2.NewGateway(log, appData, cfg)

		return &myperson{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                   log,
				Config:                cfg,
				Router:                httpHandler.Router,
				RunPersonCreateInport: runpersoncreate.NewUsecase(datasource),
				GetAllPersonInport:    getallperson.NewUsecase(datasource),
			},
		}

	}
}

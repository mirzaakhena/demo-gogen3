package restapi

import (
	"github.com/gin-gonic/gin"

	"your/path/project/domain_mydomain/usecase/getallperson"
	"your/path/project/domain_mydomain/usecase/runpersoncreate"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"
)

type Controller struct {
	Router                gin.IRouter
	Config                *config.Config
	Log                   logger.Logger
	RunPersonCreateInport runpersoncreate.Inport
	GetAllPersonInport    getallperson.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/runpersoncreate", r.authorized(), r.runPersonCreateHandler(r.RunPersonCreateInport))
	r.Router.GET("/getallperson", r.authorized(), r.getAllPersonHandler(r.GetAllPersonInport))
}

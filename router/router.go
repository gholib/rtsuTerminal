package router

import (
	"github.com/gholib/rtsuTerminal/controllers"
	"github.com/julienschmidt/httprouter"
)

//Router
var (
	Router = httprouter.New()
)

//InitRouters : start listening routes
func InitRouters() {
	Router.GET("/v0/api/ping", controllers.Ping)
	Router.POST("/v0/api/user", controllers.CreateUser)

	// Terminals
	Router.POST("/v0/api/terminal", controllers.CreateTerminal)
	Router.GET("/v0/api/terminals", controllers.GetTermimals)
	// Service
	Router.POST("/v0/api/service", controllers.CreateService)
	Router.GET("/v0/api/service", controllers.GetServices)
}

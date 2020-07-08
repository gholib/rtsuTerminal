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
}

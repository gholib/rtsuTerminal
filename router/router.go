package router

import (
	"github.com/julienschmidt/httprouter"
)

var (
	Router = httprouter.New()
)

func InitRouters() {
	// Router.GET("/v0/api/ping", controllers.Ping)
}

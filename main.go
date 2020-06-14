package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gholib/rtsuTerminal/config"
	"github.com/julienschmidt/httprouter"
)

var (
	router = httprouter.New()
)

func main() {
	confPath := flag.String("config", "config/config.json", "path of the config file")

	conf, err := config.Parce(*confPath)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	// router.POST("/test", controllers.AuthorizationHandler)

	log.Fatal(http.ListenAndServe(conf.Server.Addr, router))
}

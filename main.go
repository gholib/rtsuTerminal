package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:         conf.Service.Addr,
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 40 * time.Second,
		Handler:      router,
	}
	ctx, cancelFun := context.WithCancel(context.Background())
	go func() { // gracefull shutdown
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		s := <-sigint
		log.Println("server received signal", s)
		defer cancelFun()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("server: couldn't shutdown because of " + err.Error())
		}
	}()

	initRouters()

	log.Fatal("Listening error:", srv.ListenAndServe())
}

func initRouters() {
	router.GET("/v0/api/ping", controllers.Ping)
}

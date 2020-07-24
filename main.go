package main

import (
	"context"
	"flag"
	"github.com/gholib/rtsuTerminal/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gholib/rtsuTerminal/config"
	"github.com/gholib/rtsuTerminal/migrations"
	"github.com/gholib/rtsuTerminal/router"
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
		Addr:         conf.Server.Addr,
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 40 * time.Second,
		Handler:      router.Router,
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

	//database
	db :=	models.Connect()

	defer db.Close()

	router.InitRouters()
	migrate()

	log.Fatal("Listening error:", srv.ListenAndServe())
}

func migrate() {
	migrations.Automigrate()
}

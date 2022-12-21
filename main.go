package main

import (
	"context"
	"gin-api-demo/pkg/settings"

	//_ "gin-api-demo/pkg/settings"
	"gin-api-demo/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	router := routers.NewRouter()

	srv := &http.Server{
		Addr:           ":" + settings.GetConfig().HTTPPort,
		Handler:        router,
		ReadTimeout:    settings.GetConfig().ReadTimeout,
		WriteTimeout:   settings.GetConfig().WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting...")
}

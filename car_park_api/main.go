package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "uc_task/car_park_api/config"
	"uc_task/car_park_api/cronjob"
	"uc_task/car_park_api/handler"
	"uc_task/car_park_api/metric"
	"uc_task/car_park_api/repo"

	"log"
	"uc_task/car_park_api/ds"

	"github.com/gin-gonic/gin"
)

func main() {
	// for metrics
	metric.NewMetric()

	ds, err := ds.NewDataSource()
	if err != nil {
		log.Fatal(err)
	}

	repository := repo.NewRepository(ds)

	cronPool := cronjob.NewCronPool(repository)
	cronPool.StartCronPool()

	router := gin.Default()
	h := handler.NewHandler(
		&handler.HConfig{
			R:    router,
			DS:   ds,
			Repo: repository,
		})

	h.Register()

	port := os.Getenv("APP_PORT")
	addr := fmt.Sprintf(":%s", port)

	server := http.Server{
		Addr:           addr,
		Handler:        h.R,
		ReadTimeout:    time.Duration(time.Minute * 3),
		WriteTimeout:   time.Duration(time.Minute * 3),
		MaxHeaderBytes: 10 << 20, //10MB
	}

	go func() {
		log.Printf("server started listening on port %v\n", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Println("server failed to initialized  on port ", addr)
			log.Fatalf("error on listening :%v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	// shutdown close
	if err := server.Shutdown(context.Background()); err != nil {
		log.Println("Failed to shutdown server: ", err.Error())
	}

}

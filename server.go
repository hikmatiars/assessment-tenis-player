package main

import (
	"assessment-tennis-player/http/api"
	"assessment-tennis-player/http/response"
	"assessment-tennis-player/router"
	"assessment-tennis-player/usecase"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	data []*response.DataContainer
	FlagReady int
)

func main() {
	var httpAddr = flag.String("http", ":"+"8080", "HTTP Listen address")

	//init context
	ctx := context.Background()
	uc := usecase.NewUseCase(ctx, data, FlagReady)

	//make error channel
	errs := make(chan error)

	defer func() {
		log.Print("ended")
	}()

	confHandler := &api.Handler{Usecase: uc}
	routeHttp := router.NewHttpServer(ctx, confHandler)

	go func() {
		fmt.Println("listening on port ", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, routeHttp)
	}()

	//Print if conditional error
	log.Printf("error %v", <-errs)
}

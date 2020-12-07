package main

import (
	"context"
	"core-project/http/api"
	"core-project/router"
	"core-project/usecase"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var httpAddr = flag.String("http", ":"+"8080", "HTTP Listen address")
	//init context
	ctx := context.Background()
	uc := usecase.NewUseCase(ctx)

	//make error channel
	errs := make(chan error)

	confHandler := &api.Handler{Usecase: uc}
	routeHttp := router.NewHttpServer(ctx, confHandler)

	go func() {
		fmt.Println("listening on port ", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, routeHttp)
	}()

	//Print if conditional error
	log.Printf("error %v", <-errs)
}

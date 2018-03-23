package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/salar-hafezi/go-taskmanager/common"
	"github.com/salar-hafezi/go-taskmanager/routers"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

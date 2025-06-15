package main

import (
	"log"
	"eman/passport/sync/src/service"
	"fmt"
	"net/http"
)

func main() {

	server := service.Bootstrap()


	fmt.Printf("Starting HTTP server on port %d \n", server.Configuration.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.Configuration.Port), server)

	if err != nil {
		log.Fatal("Error starting HTTP server:", err.Error())
	}
}



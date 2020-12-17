package main

import (
	"github.com/z13z/Kiosks/master-server/services/console"
	"log"
	"net/http"
)

//todo zaza move to configuration
const serviceAddress = ":8080"

func main() {
	startServices()
}

func startServices() {
	http.Handle("/kiosk", console.KiosksServiceHandler{})
	http.Handle("/image", console.ImageServiceHandler{})
	log.Fatal(http.ListenAndServe(serviceAddress, nil))
}

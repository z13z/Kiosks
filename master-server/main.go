package main

import (
	"github.com/z13z/Kiosks/master-server/services/console"
	"log"
	"net/http"
)

//todo zaza move to configuration
const kiosksServiceAddress = ":8080"

func main() {
	startServices()
}

func startServices() {
	http.Handle("/kiosk", console.KiosksServiceHandler{})
	log.Fatal(http.ListenAndServe(kiosksServiceAddress, nil))
}

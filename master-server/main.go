package main

import (
	"github.com/z13z/Kiosks/master-server/images"
	"github.com/z13z/Kiosks/master-server/services/console"
	"log"
	"net/http"
	"os"
)

//todo zaza move to configuration
const serviceAddress = ":8080"

func main() {
	log.SetOutput(os.Stdout)
	go images.BuildImagesJob()
	startServices()
}

func startServices() {
	http.Handle("/kiosk", console.KiosksServiceHandler{})
	http.Handle("/image", console.ImageServiceHandler{})
	http.Handle("/users", console.UserServiceHandler{})
	http.Handle("/defaultScript", console.DefaultImageScriptServiceHandler{})
	http.Handle("/login", console.LoginServiceHandler{})
	http.Handle("/imageDownload", console.ImageDownloadServiceHandler{})
	http.Handle("/kiosksConnector", console.KioskConnectorServiceHandler{})
	log.Print("Server started")
	log.Fatal(http.ListenAndServe(serviceAddress, nil))
}

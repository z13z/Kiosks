package console

import (
	"fmt"
	"github.com/z13z/Kiosks/master-server/services/controller"
	"log"
	"net/http"
	"strconv"
)

type KioskConnectorServiceHandler struct{}

func (KioskConnectorServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !CheckPermissionToCall(r, &w, permissionKiosksKey) {
		return
	}
	switch r.Method {
	case "GET":
		getKioskScreenshot(&w, r)
	case "POST":
		sendCommandToKiosk(&w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		errorMessage := "{\"message\": \"not found\"}"
		mustWrite := []byte(errorMessage)
		WriteBytesInResponse(&w, &mustWrite)
	}
}

func sendCommandToKiosk(w *http.ResponseWriter, r *http.Request) {
	kioskId, ok := getIntFromQuery(r.URL.Query()["id"])
	commandParam := r.URL.Query()["command"]
	if !ok || commandParam == nil || len(commandParam) == 0 {
		log.Print("Send Command To Kiosk was called without kiosk id or without command")
		writeBadRequestError(w, "Kiosk id and command must present in request")
		return
	}
	command := commandParam[0]
	kioskEntity := kiosksBean.GetKiosk(kioskId)
	if kioskEntity == nil {
		log.Printf("Send Command To Kiosk was called wrong kiosk id: %d", kioskId)
		writeBadRequestError(w, "Kiosk id must be in request")
		return
	}
	mustWrite, err := controller.SendCommandToKiosk(kioskEntity, command)
	if err != nil {
		log.Printf("Kiosk with id (%d) isn't available", kioskId)
		writeAnyErrorResponse(w, err, http.StatusServiceUnavailable, fmt.Sprintf("Kiosk with id (%d) isn't available", kioskId))
		return
	}
	(*w).Header().Set("Content-Type", "text/plain")
	(*w).Header().Set("Content-Length", strconv.Itoa(len(mustWrite)))
	(*w).WriteHeader(http.StatusOK)
	WriteBytesInResponse(w, &mustWrite)
}

func getKioskScreenshot(w *http.ResponseWriter, r *http.Request) {
	kioskId, ok := getIntFromQuery(r.URL.Query()["id"])
	if !ok {
		log.Print("Get Kiosk Screenshot was called without kiosk id")
		writeBadRequestError(w, "Kiosk id must be in request")
		return
	}
	kioskEntity := kiosksBean.GetKiosk(kioskId)
	if kioskEntity == nil {
		log.Printf("Get Kiosk Screenshot was called wrong kiosk id: %d", kioskId)
		writeBadRequestError(w, "Kiosk id must be in request")
		return
	}
	mustWrite, err := controller.GetKioskScreenshot(kioskEntity)
	if err != nil {
		log.Printf("Kiosk with id (%d) isn't available", kioskId)
		writeAnyErrorResponse(w, err, http.StatusServiceUnavailable, fmt.Sprintf("Kiosk with id (%d) isn't available", kioskId))
		return
	}
	(*w).Header().Set("Content-Type", controller.KioskImageScreenshotContentType)
	(*w).Header().Set("Content-Length", strconv.Itoa(len(mustWrite)))
	(*w).WriteHeader(http.StatusOK)
	WriteBytesInResponse(w, &mustWrite)
}

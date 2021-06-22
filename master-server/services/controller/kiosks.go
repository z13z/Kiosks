package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/z13z/Kiosks/master-server/crypto"
	"github.com/z13z/Kiosks/master-server/db/kiosks"
	"log"
	"net/http"
	"strconv"
)

type KiosksConnectorServiceHandler struct{}

const authenticationHeaderKey = "Authentication"
const fromIpHeaderKey = "X-From-Ip"

var kiosksBean = kiosks.NewBean()

func (KiosksConnectorServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "PUT":
		addKiosk(&w, r)
	case "POST":
		updateLastOnlineTime(&w, r)
	}
}

func updateLastOnlineTime(w *http.ResponseWriter, r *http.Request) {
	authenticationToken := r.Header.Get(authenticationHeaderKey)
	fromIp := r.Header.Get(fromIpHeaderKey)
	claims, ok := crypto.CheckJwtForUser(authenticationToken)
	if !ok {
		log.Print("Update last online Time called without authorization header")
		writeUnauthorizedError(w)
		return
	}
	if fromIp == "" {
		log.Printf("Missing %s header", fromIpHeaderKey)
		writeBadRequestErrorResponse(w, fmt.Errorf("missing %s header", fromIpHeaderKey))
		return
	}

	ok = kiosksBean.UpdateLastUpdateTimeForKiosk(claims.Id, fromIp)
	if !ok {
		log.Printf("Problem updating last online time for kiosk, id (%s): ", claims.Id)
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
	(*w).WriteHeader(http.StatusAccepted)
}

func addKiosk(w *http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	readBytes := buffer.Bytes()
	if err != nil {
		log.Print("Error while creating new kiosk: ", err)
		(*w).WriteHeader(http.StatusBadRequest)
		return
	}
	request := KioskCreateRequest{}
	err = json.Unmarshal(readBytes, &request)

	kiosk, err := kiosksBean.AddKiosk(request.KioskImageId, request.KioskAddress, crypto.Encrypt(request.ServicePassword))
	kioskStr, _ := json.Marshal(kiosk)
	log.Printf("Created kiosk (%s)", kioskStr)

	if err != nil {
		log.Print("Problem while creating new kiosk: ", err)
		(*w).WriteHeader(http.StatusBadRequest)
		return
	}
	jwtToken, err := crypto.GetJwtForKiosk(kiosk.GetId())
	if err != nil {
		log.Print("Problem while generating password for new kiosk, id "+strconv.FormatInt(kiosk.Id, 10)+": ", err)
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
	response := KioskCreateResponse{Id: kiosk.Id, Password: jwtToken}
	responseStr, err := json.Marshal(response)
	if err != nil {
		log.Print("Problem while marshaling response for new kiosk, id "+strconv.FormatInt(kiosk.Id, 10)+": ", err)
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}

	(*w).WriteHeader(http.StatusCreated)
	writeBytesInResponse(w, &responseStr)
}

func writeBadRequestErrorResponse(w *http.ResponseWriter, err error) {
	writeAnyErrorResponse(w, err, http.StatusBadRequest, "{\"message\": \"bad request\"}")
}

func writeUnauthorizedError(w *http.ResponseWriter) {
	writeAnyErrorResponse(w, nil, http.StatusUnauthorized, "{\"message\": \"unauthorized\"}")
}

func writeAnyErrorResponse(w *http.ResponseWriter, err error, errorCode int, errorMessage string) {
	(*w).WriteHeader(errorCode)
	mustWrite := []byte(errorMessage)
	writeBytesInResponse(w, &mustWrite)
	if err != nil {
		log.Print(err)
	}
}

func writeBytesInResponse(w *http.ResponseWriter, mustWrite *[]byte) {
	mustWriteLen := len(*mustWrite)
	writtenBytes, err := (*w).Write(*mustWrite)
	if err != nil {
		log.Fatal(err)
	}
	for writtenBytes < mustWriteLen {
		curWrittenBytes, err := (*w).Write((*mustWrite)[writtenBytes:])
		writtenBytes += curWrittenBytes
		if err != nil {
			log.Fatal(err)
		}
	}
}

type KioskCreateRequest struct {
	KioskImageId    int64  `json:"kioskImageId"`
	KioskAddress    string `json:"kioskAddress"`
	ServicePassword string `json:"servicePassword"`
}

type KioskCreateResponse struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}

package controller

import (
	"bytes"
	"encoding/json"
	"github.com/z13z/Kiosks/master-server/crypto"
	"github.com/z13z/Kiosks/master-server/db/kiosks"
	"log"
	"net/http"
	"strconv"
)

type KiosksConnectorServiceHandler struct{}

var kiosksBean = kiosks.NewBean()

func (KiosksConnectorServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "PUT":
		addKiosk(&w, r)
	}
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

	writeBytesInResponse(w, &responseStr)
	(*w).WriteHeader(http.StatusAccepted)
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

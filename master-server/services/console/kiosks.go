package console

import (
	"encoding/json"
	"github.com/z13z/Kiosks/master-server/db/kiosks"
	"net/http"
)

var kiosksBean = kiosks.NewBean()

func ServeKiosksRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getKiosksList(w)
	case "POST":
		//todo implement
		//editKiosk(w, r)
	case "PUT":
		//todo implement
		//addKiosk(w, r)
	case "DELETE":
		//todo implement
		//deleteKiosk()
	default:
		w.WriteHeader(http.StatusNotFound)
		errorMessage := "{\"message\": \"not found\"}"
		mustWrite := []byte(errorMessage)
		writeBytesInResponse(w, &mustWrite)
	}
}

func getKiosksList(w http.ResponseWriter) {
	mustWrite, err := json.Marshal(*kiosksBean.GetKiosks())
	if err != nil {
		panic(err.Error())
	}
	writeBytesInResponse(w, &mustWrite)

}

func writeBytesInResponse(w http.ResponseWriter, mustWrite *[]byte) {
	mustWriteLen := len(*mustWrite)
	writtenBytes, err := w.Write(*mustWrite)
	if err != nil {
		panic(err)
	}
	for writtenBytes < mustWriteLen {
		curWrittenBytes, err := w.Write((*mustWrite)[writtenBytes:])
		writtenBytes += curWrittenBytes
		if err != nil {
			panic(err)
		}
	}
}

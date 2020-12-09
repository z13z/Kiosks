package console

import (
	"encoding/json"
	"github.com/z13z/Kiosks/master-server/db/kiosks"
	"math"
	"net/http"
	"strconv"
)

type KiosksServiceHandler struct{}

var kiosksBean = kiosks.NewBean()

func (KiosksServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getKiosksList(&w, r)
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
		writeBytesInResponse(&w, &mustWrite)
	}
}

func getKiosksList(w *http.ResponseWriter, r *http.Request) {
	var found bool
	kioskId, specificKiosk := getIntFromQuery(r.URL.Query()["id"])
	offset, found := getIntFromQuery(r.URL.Query()["offset"])
	if !found {
		offset = 0
	}
	limit, found := getIntFromQuery(r.URL.Query()["limit"])
	if !found {
		limit = math.MaxInt32
	}
	var mustWrite []byte
	var err error
	if specificKiosk {
		mustWrite, err = json.Marshal(*kiosksBean.GetKiosk(kioskId))
	} else {
		mustWrite, err = json.Marshal(*kiosksBean.GetKiosks(offset, limit))
	}
	if err != nil {
		panic(err.Error())
	}
	writeBytesInResponse(w, &mustWrite)
}

func getIntFromQuery(param []string) (int, bool) {
	if param != nil {
		intParam, err := strconv.Atoi(param[0])
		if err != nil {
			return 0, false
		}
		return intParam, true
	}
	return 0, false
}

func writeBytesInResponse(w *http.ResponseWriter, mustWrite *[]byte) {
	mustWriteLen := len(*mustWrite)
	writtenBytes, err := (*w).Write(*mustWrite)
	if err != nil {
		panic(err)
	}
	for writtenBytes < mustWriteLen {
		curWrittenBytes, err := (*w).Write((*mustWrite)[writtenBytes:])
		writtenBytes += curWrittenBytes
		if err != nil {
			panic(err)
		}
	}
}

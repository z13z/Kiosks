package console

import (
	"encoding/json"
	"github.com/z13z/Kiosks/master-server/db/kiosks"
	"log"
	"math"
	"net/http"
)

type KiosksServiceHandler struct{}

var kiosksBean = kiosks.NewBean()

func (KiosksServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !CheckPermissionToCall(r, &w, permissionKiosksKey) {
		return
	}
	switch r.Method {
	case "GET":
		getKiosksList(&w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		errorMessage := "{\"message\": \"not found\"}"
		mustWrite := []byte(errorMessage)
		WriteBytesInResponse(&w, &mustWrite)
	}
}

func getKiosksList(w *http.ResponseWriter, r *http.Request) {
	var found bool
	kioskId, _ := getIntFromQuery(r.URL.Query()["id"])
	kioskNameParam := r.URL.Query()["name"]
	var kioskName string
	if kioskNameParam != nil && len(kioskNameParam) > 0 {
		kioskName = kioskNameParam[0]
	} else {
		kioskName = ""
	}
	oneKiosk := getBooleanFromQuery(r.URL.Query()["oneKiosk"])
	offset, found := getIntFromQuery(r.URL.Query()["offset"])
	if !found || offset < 0 {
		offset = 0
	}
	limit, found := getIntFromQuery(r.URL.Query()["limit"])
	if !found {
		limit = math.MaxInt32
	}
	var mustWrite []byte
	var err error
	if oneKiosk {
		mustWrite, err = json.Marshal(*kiosksBean.GetKiosk(kioskId))
	} else {
		response := kiosksListResponse{Rows: *kiosksBean.GetKiosks(kioskId, kioskName, offset, limit),
			RowsCount: kiosksBean.GetKiosksCount(kioskId, kioskName)}
		mustWrite, err = json.Marshal(response)
	}
	if err != nil {
		log.Fatal(err)
	}
	(*w).WriteHeader(http.StatusOK)
	WriteBytesInResponse(w, &mustWrite)
}

type kiosksListResponse struct {
	Rows      []kiosks.KioskEntity `json:"rows"`
	RowsCount int                  `json:"rowsCount"`
}

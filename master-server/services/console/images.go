package console

import (
	"encoding/json"
	"fmt"
	"github.com/z13z/Kiosks/master-server/db/images"
	"log"
	"math"
	"net/http"
)

type ImageServiceHandler struct{}

type DefaultImageScriptServiceHandler struct{}

var imagesBean = images.NewBean()

func (DefaultImageScriptServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		response, err := prepare_kioskBytes()
		if err != nil {
			writeServerErrorResponse(&w, err)
		} else {
			writeBytesInResponse(&w, &response)
		}
	} else {
		writeWrongHttpMethodResponse(&w, r.Method)
	}
}

func (ImageServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getImagesList(&w, r)
	case "POST":
		//todo implement
		//editImage(w, r)
	case "PUT":
		//todo implement
		//addImage(w, r)
	case "DELETE":
		//todo implement
		//deleteImage()
	default:
		writeWrongHttpMethodResponse(&w, r.Method)
	}
}

func writeWrongHttpMethodResponse(w *http.ResponseWriter, method string) {
	(*w).WriteHeader(http.StatusNotFound)
	errorMessage := fmt.Sprintf("{\"message\": \"unsupported http method [%s]\"}", method)
	mustWrite := []byte(errorMessage)
	writeBytesInResponse(w, &mustWrite)
}

func writeServerErrorResponse(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusInternalServerError)
	errorMessage := "{\"message\": \"internal server error\"}"
	mustWrite := []byte(errorMessage)
	writeBytesInResponse(w, &mustWrite)
	log.Fatal(err)
}

func getImagesList(w *http.ResponseWriter, r *http.Request) {
	var found bool
	imageId, _ := getIntFromQuery(r.URL.Query()["id"])
	imageNameParam := r.URL.Query()["name"]
	var imageName string
	if imageNameParam != nil && len(imageNameParam) > 0 {
		imageName = imageNameParam[0]
	} else {
		imageName = ""
	}
	oneImage := getBooleanFromQuery(r.URL.Query()["oneImage"])
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
	if oneImage {
		mustWrite, err = json.Marshal(*imagesBean.GetImage(imageId))
	} else {
		mustWrite, err = json.Marshal(*imagesBean.GetImages(imageId, imageName, offset, limit))
	}
	if err != nil {
		panic(err.Error())
	}
	writeBytesInResponse(w, &mustWrite)
}

package console

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/z13z/Kiosks/master-server/db/images"
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
	if !CheckPermissionToCall(r, &w, permissionImagesKey) {
		return
	}
	switch r.Method {
	case "GET":
		getImagesList(&w, r)
	case "POST":
		//todo implement
		//editImage(&w, r)
	case "PUT":
		addImage(&w, r)
	case "DELETE":
		//todo implement
		//deleteImage()
	default:
		writeWrongHttpMethodResponse(&w, r.Method)
	}
}

func addOrEditImage(w *http.ResponseWriter, r *http.Request, methodToCall func(entity *images.ImageEntity) error) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	readBytes := buffer.Bytes()
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}
	entity := images.ImageEntity{}
	err = json.Unmarshal(readBytes, &entity)
	if err != nil {
		writeBadRequestError(w, fmt.Sprintf("Can't unmarshal Image from json [%s]", string(readBytes)))
	}
	err = methodToCall(&entity)
	if err != nil {
		writeBadRequestError(w, string(readBytes))
	} else {
		(*w).WriteHeader(http.StatusAccepted)
	}
}

//todo zaza uncomment
//func editImage(w *http.ResponseWriter, r *http.Request) {
//	addOrEditImage(w, r, imagesBean.EditImage)
//}

func addImage(w *http.ResponseWriter, r *http.Request) {
	addOrEditImage(w, r, imagesBean.AddImage)
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
		response := imagesListResponse{Rows: *imagesBean.GetImages(imageId, imageName, offset, limit),
			RowsCount: imagesBean.GetImagesCount(imageId, imageName)}
		mustWrite, err = json.Marshal(response)
	}
	if err != nil {
		panic(err.Error())
	}
	writeBytesInResponse(w, &mustWrite)
}

type imagesListResponse struct {
	Rows      []images.ImageEntity `json:"rows"`
	RowsCount int                  `json:"rowsCount"`
}

package console

import (
	"fmt"
	"github.com/z13z/Kiosks/master-server/images"
	"net/http"
)

type ImageDownloadServiceHandler struct{}

func (ImageDownloadServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !CheckPermissionToCall(r, &w, permissionImagesKey) {
		return
	}
	if r.Method == "GET" {
		filePath := getImageFilePath(&w, r)
		if filePath == "" {
			return
		}
		http.ServeFile(w, r, filePath)
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

func getImageFilePath(w *http.ResponseWriter, r *http.Request) string {
	imageNameParam := r.URL.Query()["name"]
	if imageNameParam == nil || len(imageNameParam) == 0 {
		writeBadRequestError(w, "requested file doesn't exist")
		return ""
	}
	return fmt.Sprintf("%s/%s/%s/%s/%s", images.KiosksImagesDirectory, imageNameParam[0],
		images.KiosksImagesScriptsDirectoryName, images.OutputFileDir, images.OutputFileName)
}

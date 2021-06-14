package console

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/z13z/Kiosks/master-server/crypto"
	"log"
	"net/http"
)

type LoginServiceHandler struct{}

func loginUser(w *http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	readBytes := buffer.Bytes()
	if err != nil {
		writeBadRequestErrorResponse(w, err)
		return
	}
	requestWrapper := LoginRequest{}
	err = json.Unmarshal(readBytes, &requestWrapper)
	if err != nil {
		writeBadRequestErrorResponse(w, err)
		log.Printf("Error for login request %+v", requestWrapper)
		return
	}
	userFromDb := usersBean.GetUserByUsername(requestWrapper.Username)
	if userFromDb == nil {
		writeBadRequestErrorResponse(w, err)
		return
	}
	if userFromDb.Password != usersBean.GetPassword(requestWrapper.Password) {
		writeBadRequestErrorResponse(w, err)
		return
	}
	jwtToken, err := crypto.GetJwtForUser(*userFromDb)
	if err != nil {
		writeBadRequestErrorResponse(w, err)
		log.Print("Error while creating jwt token")
		return
	}
	writeJsonInResponse(w, fmt.Sprintf("{\"jwt\": \"%s\"}", jwtToken))
}

func (LoginServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		loginUser(&w, r)
	default:
		writeWrongHttpMethodResponse(&w, r.Method)
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

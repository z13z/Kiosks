package console

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/z13z/Kiosks/master-server/db/users"
	"math"
	"net/http"
)

type UserServiceHandler struct{}

var usersBean = users.NewBean()

func (UserServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !CheckPermissionToCall(r, &w, permissionUsersKey) {
		return
	}
	switch r.Method {
	case "GET":
		getUsersList(&w, r)
	case "POST":
		editUser(&w, r)
	case "PUT":
		//todo implement
		//addUser(w, r)
	case "DELETE":
		//todo implement
		//deleteUser()
	default:
		writeWrongHttpMethodResponse(&w, r.Method)
	}
}

func editUser(w *http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	readBytes := buffer.Bytes()
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}
	entity := users.UserEntity{}
	err = json.Unmarshal(readBytes, &entity)
	if err != nil {
		writeBadRequestError(w, fmt.Sprintf("Can't unmarshal User from json [%s]", string(readBytes)))
	}
	err = usersBean.EditUser(&entity)
	if err != nil {
		writeBadRequestError(w, string(readBytes))
	} else {
		(*w).WriteHeader(http.StatusAccepted)
	}
}

func getUsersList(w *http.ResponseWriter, r *http.Request) {
	var found bool
	userId, _ := getIntFromQuery(r.URL.Query()["id"])
	userNameParam := r.URL.Query()["name"]
	var userName string
	if userNameParam != nil && len(userNameParam) > 0 {
		userName = userNameParam[0]
	} else {
		userName = ""
	}
	oneUser := getBooleanFromQuery(r.URL.Query()["oneUser"])
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
	if oneUser {
		mustWrite, err = json.Marshal(*usersBean.GetUser(userId))
	} else {
		response := usersListResponse{Rows: *usersBean.GetUsers(userId, userName, offset, limit),
			RowsCount: usersBean.GetUsersCount(userId, userName)}
		mustWrite, err = json.Marshal(response)
	}
	if err != nil {
		panic(err.Error())
	}
	writeBytesInResponse(w, &mustWrite)
}

type usersListResponse struct {
	Rows      []users.UserEntity `json:"rows"`
	RowsCount int                `json:"rowsCount"`
}

package console

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/z13z/Kiosks/master-server/db/users"
	"math"
	"net/http"
	"strings"
	"time"
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
		addUser(&w, r)
	case "DELETE":
		deleteUser(&w, r)
	default:
		writeWrongHttpMethodResponse(&w, r.Method)
	}
}

func deleteUser(w *http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	readBytes := buffer.Bytes()
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}
	requestUser := DeleteRequest{}
	err = json.Unmarshal(readBytes, &requestUser)
	if err != nil {
		writeBadRequestError(w, fmt.Sprintf("Can't unmarshal User from json [%s]", string(readBytes)))
	}
	err = usersBean.DeleteUser(requestUser.RowId)
	if err != nil {
		writeBadRequestError(w, string(readBytes))
	} else {
		(*w).WriteHeader(http.StatusAccepted)
	}
}

func editUser(w *http.ResponseWriter, r *http.Request) {
	addOrEditUser(w, r, usersBean.EditUser)
}

func addUser(w *http.ResponseWriter, r *http.Request) {
	addOrEditUser(w, r, usersBean.AddUser)
}

func addOrEditUser(w *http.ResponseWriter, r *http.Request, action func(entity *users.UserEntity) error) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	readBytes := buffer.Bytes()
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}
	requestUser := RequestUserDAO{}
	err = json.Unmarshal(readBytes, &requestUser)
	if err != nil {
		writeBadRequestError(w, fmt.Sprintf("Can't unmarshal User from json [%s]", string(readBytes)))
	}
	entityUser, err := requestUser.getEntityObject()
	if err != nil {
		writeBadRequestError(w, string(readBytes))
		return
	}
	//edit or add user
	err = action(entityUser)
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

type DeleteRequest struct {
	RowId int64 `json:"rowId"`
}

type RequestUserDAO struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Password    string   `json:"password"`
}

func (req *RequestUserDAO) getEntityObject() (*users.UserEntity, error) {
	entity := users.UserEntity{}
	entity.Name = req.Name
	entity.Password = usersBean.GetPassword(req.Password)
	entity.UpdateTime = time.Now()
	if req.Permissions != nil {
		arr := pq.StringArray{}
		err := arr.Scan("{" + strings.Join(req.Permissions, ",") + "}")
		if err != nil {
			return nil, err
		}
		entity.Permissions = arr
	}
	return &entity, nil
}

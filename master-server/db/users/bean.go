package users

import (
	"crypto/sha256"
	"fmt"
	"github.com/z13z/Kiosks/master-server/db/common"
	"strconv"
)

type Bean struct {
	connector *common.DBConnector
}

func NewBean() *Bean {
	newBean := Bean{}
	newBean.connector = common.NewDBConnector()
	return &newBean
}

func (bean *Bean) GetUsers(id int, name string, offset, limit int) *[]UserEntity {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if name != "" {
		queryParams["name"] = name
	}
	resultFromDb := bean.connector.GetObjectsFromDb(&UserEntity{}, &queryParams, offset, limit)
	var resultUsers []UserEntity
	for _, userFromDb := range *resultFromDb {
		resultUsers = append(resultUsers, *userFromDb.(*UserEntity))
	}
	return &resultUsers
}

func (bean *Bean) GetUsersCount(id int, name string) int {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if name != "" {
		queryParams["name"] = name
	}
	return bean.connector.GetObjectsCountFromDb(&UserEntity{}, &queryParams)
}

func (bean *Bean) GetUser(id int) *UserEntity {
	resultFromDb := bean.GetUsers(id, "", 0, 1)
	if len(*resultFromDb) == 1 {
		return &((*resultFromDb)[0])
	}
	return nil
}

func (bean *Bean) EditUser(entity *UserEntity) error {
	entity.Password = bean.GetPassword(entity.Password)
	updated := bean.connector.UpdateObjectInDb(entity)
	if updated != 1 {
		return fmt.Errorf("user with id [%d] doesn't exist in database", entity.Id)
	}
	return nil
}

func (bean *Bean) GetUserByUsername(username string) *UserEntity {
	if username == "" {
		return nil
	}
	resultFromDb := bean.GetUsers(0, username, 0, 1)
	if len(*resultFromDb) == 1 {
		return &((*resultFromDb)[0])
	}
	return nil
}

func (bean *Bean) GetPassword(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

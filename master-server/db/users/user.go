package users

import (
	"github.com/lib/pq"
	"github.com/z13z/Kiosks/master-server/db/common"
	"time"
)

type UserEntity struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	UpdateTime  time.Time      `json:"updateTime"`
	Permissions pq.StringArray `sql:"type:KioskUserPermission ARRAY" json:"permissions"`
	Password    string         `json:"-"`
	LoggedIn    time.Time      `sql:"-" json:"loggedInTime"`
}

func (user *UserEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	updateTime := fields["update_time"]
	permissions := fields["permissions"]
	password := fields["password"]
	user.Id = *idValue.(*int64)
	user.Name = *nameValue.(*string)
	user.UpdateTime = *updateTime.(*time.Time)
	user.Permissions = *(permissions.(*pq.StringArray))
	user.Password = *password.(*string)
}

func (user *UserEntity) GetTableName() string {
	return "KioskUser"
}

func (user *UserEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name", "update_time", "permissions", "password"}
}

func (user *UserEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&user.Id, &user.Name, &user.UpdateTime, &user.Permissions, &user.Password})
}

func (user *UserEntity) NewEntity() common.IEntity {
	return new(UserEntity)
}

func (user *UserEntity) GetEditableFieldNames() *[]string {
	return &[]string{"name", "update_time", "permissions", "password"}
}

func (user *UserEntity) GetEditableFieldValueHolders() *[]interface{} {
	return &([]interface{}{&user.Name, &user.UpdateTime, &user.Permissions, &user.Password})
}

func (user *UserEntity) GetId() int64 {
	return user.Id
}

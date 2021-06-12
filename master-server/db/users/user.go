package users

import (
	"github.com/lib/pq"
	"github.com/z13z/Kiosks/master-server/db/common"
	"time"
)

type UserEntity struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	CreateTime  time.Time      `json:"createTime"`
	Permissions pq.StringArray `sql:"type:KioskUserPermission ARRAY" json:"permissions"`
}

func (user *UserEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	createTime := fields["create_time"]
	permissions := fields["permissions"]
	user.Id = *idValue.(*int64)
	user.Name = *nameValue.(*string)
	user.CreateTime = *createTime.(*time.Time)
	user.Permissions = *(permissions.(*pq.StringArray))
}

func (user *UserEntity) GetTableName() string {
	return "KioskUser"
}

func (user *UserEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name", "create_time", "permissions"}
}

func (user *UserEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&user.Id, &user.Name, &user.CreateTime, &user.Permissions})
}

func (user *UserEntity) NewEntity() common.IEntity {
	return new(UserEntity)
}

func (user *UserEntity) GetEditableFieldNames() *[]string {
	return &[]string{"name", "create_time", "permissions"}
}

func (user *UserEntity) GetEditableFieldValueHolders() *[]interface{} {
	return &([]interface{}{&user.Id, &user.Name, &user.CreateTime, &user.Permissions})
}

func (user *UserEntity) GetId() int64 {
	return user.Id
}

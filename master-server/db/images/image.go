package images

import (
	"github.com/z13z/Kiosks/master-server/db/common"
	"time"
)

type ImageEntity struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
	Script     string    `json:"script"`
	State      string    `sql:"type:KioskImageState" json:"state"`
}

func (image *ImageEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	createTime := fields["create_time"]
	state := fields["state"]
	script := fields["script"]
	image.Id = *idValue.(*int64)
	image.Name = *nameValue.(*string)
	image.CreateTime = *createTime.(*time.Time)
	image.State = *state.(*string)
	image.Script = *script.(*string)
}

func (image *ImageEntity) GetTableName() string {
	return "KioskImage"
}

func (image *ImageEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name", "create_time", "state", "script"}
}

func (image *ImageEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&image.Id, &image.Name, &image.CreateTime, &image.State, &image.Script})
}

func (image *ImageEntity) NewEntity() common.IEntity {
	return new(ImageEntity)
}

func (image *ImageEntity) GetEditableFieldNames() *[]string {
	return &[]string{"name", "create_time", "state", "script"}
}

func (image *ImageEntity) GetEditableFieldValueHolders() *[]interface{} {
	return &([]interface{}{&image.Name, &image.CreateTime, &image.State, &image.Script})
}

func (image *ImageEntity) GetId() int64 {
	return image.Id
}

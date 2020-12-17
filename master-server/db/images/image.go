package images

import (
	"github.com/z13z/Kiosks/master-server/db/common"
	"time"
)

type ImageEntity struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
}

func (image *ImageEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	createTime := fields["create_time"]
	image.Id = *idValue.(*int64)
	image.Name = *nameValue.(*string)
	image.CreateTime = *createTime.(*time.Time)
}

func (image *ImageEntity) GetTableName() string {
	return "KioskImage"
}

func (image *ImageEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name", "create_time"}
}

func (image *ImageEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&image.Id, &image.Name, &image.CreateTime})
}

func (image *ImageEntity) NewEntity() common.IEntity {
	return new(ImageEntity)
}

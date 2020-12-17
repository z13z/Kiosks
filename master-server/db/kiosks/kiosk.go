package kiosks

import (
	"github.com/z13z/Kiosks/master-server/db/common"
	"time"
)

type KioskEntity struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	KioskImageId int64     `json:"kioskImageId"`
	CreateTime   time.Time `json:"createTime"`
}

func (kiosk *KioskEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	kioskImageIdValue := fields["kiosk_image_id"]
	createTime := fields["create_time"]
	kiosk.Id = *idValue.(*int64)
	kiosk.Name = *nameValue.(*string)
	kiosk.KioskImageId = *kioskImageIdValue.(*int64)
	kiosk.CreateTime = *createTime.(*time.Time)
}

func (*KioskEntity) GetTableName() string {
	return "Kiosk"
}

func (*KioskEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name", "kiosk_image_id", "create_time"}
}

func (kiosk *KioskEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&kiosk.Id, &kiosk.Name, &kiosk.KioskImageId, &kiosk.CreateTime})
}

func (*KioskEntity) NewEntity() common.IEntity {
	return new(KioskEntity)
}

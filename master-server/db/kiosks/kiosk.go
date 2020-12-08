package kiosks

import (
	"github.com/z13z/Kiosks/master-server/db/common"
)

type KioskEntity struct {
	id           int64
	name         string
	kioskImageId int64
}

func (kiosk *KioskEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	nameValue := fields["name"]
	kioskImageIdValue := fields["kioskImageId"]
	kiosk.id = *idValue.(*int64)
	kiosk.name = *nameValue.(*string)
	kiosk.kioskImageId = *kioskImageIdValue.(*int64)
}

func (kiosk *KioskEntity) GetTableName() string {
	return "Kiosk"
}

func (kiosk *KioskEntity) GetFieldNames() *[]string {
	return &[]string{"id", "name", "kioskImageId"}
}

func (kiosk *KioskEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&kiosk.id, &kiosk.name, &kiosk.kioskImageId})
}

func (kiosk *KioskEntity) NewEntity() common.IEntity {
	return new(KioskEntity)
}

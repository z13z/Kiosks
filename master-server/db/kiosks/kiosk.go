package kiosks

import (
	"github.com/z13z/Kiosks/master-server/db/common"
	"time"
)

type KioskEntity struct {
	Id              int64     `json:"id"`
	Address         string    `json:"address"`
	KioskImageId    int64     `json:"kioskImageId"`
	LastOnline      time.Time `json:"lastOnline"`
	ServicePassword []byte    `json:"-"`
}

func (kiosk *KioskEntity) SetEntityFields(fields map[string]interface{}) {
	idValue := fields["id"]
	address := fields["address"]
	kioskImageIdValue := fields["kiosk_image_id"]
	lastOnline := fields["last_online"]
	servicePassword := fields["service_password"]
	kiosk.Id = *idValue.(*int64)
	kiosk.Address = *address.(*string)
	kiosk.KioskImageId = *kioskImageIdValue.(*int64)
	kiosk.LastOnline = *lastOnline.(*time.Time)
	kiosk.ServicePassword = *servicePassword.(*[]byte)
}

func (kiosk *KioskEntity) GetTableName() string {
	return "Kiosk"
}

func (kiosk *KioskEntity) GetFieldNames() *[]string {
	return &[]string{"id", "address", "kiosk_image_id", "last_online", "service_password"}
}

func (kiosk *KioskEntity) GetFieldValueHolders() *[]interface{} {
	return &([]interface{}{&kiosk.Id, &kiosk.Address, &kiosk.KioskImageId, &kiosk.LastOnline, &kiosk.ServicePassword})
}

func (kiosk *KioskEntity) NewEntity() common.IEntity {
	return new(KioskEntity)
}

func (kiosk *KioskEntity) GetEditableFieldValueHolders() *[]interface{} {
	return &([]interface{}{&kiosk.Address, &kiosk.KioskImageId, &kiosk.LastOnline, &kiosk.ServicePassword})
}

func (kiosk *KioskEntity) GetEditableFieldNames() *[]string {
	return &[]string{"address", "kiosk_image_id", "last_online", "service_password"}
}

func (kiosk *KioskEntity) GetId() int64 {
	return kiosk.Id
}

package kiosks

import (
	"fmt"
	"github.com/z13z/Kiosks/master-server/db/common"
	"strconv"
	"time"
)

type Bean struct {
	connector *common.DBConnector
}

func NewBean() *Bean {
	newBean := Bean{}
	newBean.connector = common.NewDBConnector()
	return &newBean
}

func (bean *Bean) GetKiosks(id int, kioskAddress string, offset, limit int) *[]KioskEntity {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if kioskAddress != "" {
		queryParams["address"] = kioskAddress
	}
	resultFromDb := bean.connector.GetObjectsFromDb(&KioskEntity{}, &queryParams, offset, limit)
	var resultKiosks []KioskEntity
	for _, kioskFromDb := range *resultFromDb {
		resultKiosks = append(resultKiosks, *kioskFromDb.(*KioskEntity))
	}
	return &resultKiosks
}

func (bean *Bean) GetKiosksCount(id int, kioskAddress string) int {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if kioskAddress != "" {
		queryParams["address"] = kioskAddress
	}
	return bean.connector.GetObjectsCountFromDb(&KioskEntity{}, &queryParams)
}

func (bean *Bean) GetKiosk(id int) *KioskEntity {
	resultFromDb := bean.GetKiosks(id, "", 0, 1)
	if len(*resultFromDb) == 1 {
		return &((*resultFromDb)[0])
	}
	return nil
}

func (bean *Bean) AddKiosk(kioskImageId int64, kioskAddress string, password []byte) (*KioskEntity, error) {
	kiosk := KioskEntity{
		KioskImageId:    kioskImageId,
		Address:         kioskAddress,
		LastOnline:      time.Now(),
		ServicePassword: password,
	}
	idToReturn, ok := bean.connector.InsertObjectInDb(&kiosk)
	if !ok {
		return nil, fmt.Errorf("kiosk wasn't saved in db")
	}
	kiosk.Id = idToReturn
	return &kiosk, nil
}

func (bean *Bean) UpdateLastUpdateTimeForKiosk(idStr string) bool {
	id, _ := strconv.Atoi(idStr)
	kiosk := bean.GetKiosk(id)
	(*kiosk).LastOnline = time.Now()
	return bean.connector.UpdateObjectInDb(kiosk) == 1
}

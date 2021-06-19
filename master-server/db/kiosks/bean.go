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

func (bean *Bean) GetKiosks(id int, name string, offset, limit int) *[]KioskEntity {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if name != "" {
		queryParams["name"] = name
	}
	resultFromDb := bean.connector.GetObjectsFromDb(&KioskEntity{}, &queryParams, offset, limit)
	var resultKiosks []KioskEntity
	for _, kioskFromDb := range *resultFromDb {
		resultKiosks = append(resultKiosks, *kioskFromDb.(*KioskEntity))
	}
	return &resultKiosks
}

func (bean *Bean) GetKiosksCount(id int, name string) int {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if name != "" {
		queryParams["name"] = name
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

func (bean *Bean) AddKiosk(kioskImageId int64, kioskAddress string) (*KioskEntity, error) {
	kiosk := KioskEntity{
		KioskImageId: kioskImageId,
		Address:      kioskAddress,
		LastOnline:   time.Now(),
	}
	idToReturn, ok := bean.connector.InsertObjectInDb(&kiosk)
	if !ok {
		return nil, fmt.Errorf("kiosk wasn't saved in db")
	}
	kiosk.Id = idToReturn
	return &kiosk, nil
}

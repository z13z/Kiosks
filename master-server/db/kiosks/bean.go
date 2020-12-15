package kiosks

import (
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

func (bean *Bean) GetKiosk(id int) *KioskEntity {
	resultFromDb := bean.GetKiosks(id, "", 0, 1)
	if len(*resultFromDb) == 1 {
		return &((*resultFromDb)[0])
	}
	return nil
}

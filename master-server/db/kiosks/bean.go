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

func (bean *Bean) GetKiosks(offset, limit int) *[]KioskEntity {
	resultFromDb := bean.connector.GetObjectsFromDb(&KioskEntity{}, nil, offset, limit)
	var resultKiosks []KioskEntity
	for _, kioskFromDb := range *resultFromDb {
		resultKiosks = append(resultKiosks, kioskFromDb.(KioskEntity))
	}
	return &resultKiosks
}

func (bean *Bean) GetKiosk(id int) *KioskEntity {
	whereClause := "WHERE id = " + strconv.Itoa(id)
	resultFromDb := bean.connector.GetObjectsFromDb(&KioskEntity{}, &whereClause, 0, 1)
	if len(*resultFromDb) == 1 {
		result := ((*resultFromDb)[0]).(KioskEntity)
		return &result
	}
	return nil
}

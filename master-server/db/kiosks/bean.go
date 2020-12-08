package kiosks

import "github.com/z13z/Kiosks/master-server/db/common"

type Bean struct {
	connector *common.DBConnector
}

func NewBean() *Bean {
	newBean := Bean{}
	newBean.connector = common.NewDBConnector()
	return &newBean
}

func (bean *Bean) GetKiosks() *[]KioskEntity {
	resultFromDb := bean.connector.GetObjectsFromDb(&KioskEntity{})
	var resultKiosks []KioskEntity
	for _, kioskFromDb := range *resultFromDb {
		resultKiosks = append(resultKiosks, kioskFromDb.(KioskEntity))
	}
	return &resultKiosks
}

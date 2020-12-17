package images

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

func (bean *Bean) GetImages(id int, name string, offset, limit int) *[]ImageEntity {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if name != "" {
		queryParams["name"] = name
	}
	resultFromDb := bean.connector.GetObjectsFromDb(&ImageEntity{}, &queryParams, offset, limit)
	var resultImages []ImageEntity
	for _, imageFromDb := range *resultFromDb {
		resultImages = append(resultImages, *imageFromDb.(*ImageEntity))
	}
	return &resultImages
}

func (bean *Bean) GetImage(id int) *ImageEntity {
	resultFromDb := bean.GetImages(id, "", 0, 1)
	if len(*resultFromDb) == 1 {
		return &((*resultFromDb)[0])
	}
	return nil
}

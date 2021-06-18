package images

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

func (bean *Bean) GetImagesCount(id int, name string) int {
	queryParams := make(map[string]string)
	if id != 0 {
		queryParams["id"] = strconv.Itoa(id)
	}
	if name != "" {
		queryParams["name"] = name
	}
	return bean.connector.GetObjectsCountFromDb(&ImageEntity{}, &queryParams)
}

func (bean *Bean) GetImage(id int) *ImageEntity {
	resultFromDb := bean.GetImages(id, "", 0, 1)
	if len(*resultFromDb) == 1 {
		return &((*resultFromDb)[0])
	}
	return nil
}

func (bean *Bean) EditImage(entity *ImageEntity) error {
	updated := bean.connector.UpdateObjectInDb(entity)
	if updated != 1 {
		return fmt.Errorf("image with id [%d] doesn't exist in database", entity.Id)
	}
	return nil
}

func (bean *Bean) AddImage(entity *ImageEntity) error {
	entity.State = "created"
	entity.CreateTime = time.Now()
	updated := bean.connector.InsertObjectInDb(entity)
	if !updated {
		return fmt.Errorf("image with name [%s] exists exist in database", entity.Name)
	}
	return nil
}

func (bean *Bean) DeleteImage(id int64) error {
	if !bean.connector.DeleteObjectInDb(&ImageEntity{Id: id}) {
		return fmt.Errorf("image with id [%d] doesn't in database", id)
	}
	return nil
}

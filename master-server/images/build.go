package images

import (
	"github.com/go-co-op/gocron"
	"github.com/z13z/Kiosks/master-server/db/images"
	"log"
	"time"
)

const mustBuildImageState = "waiting"
const batchSize = 2
const jobIntervalSecs = 1

func BuildImagesJob() {
	imagesBean := images.NewBean()
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(jobIntervalSecs).Minutes().Do(buildImagesJobRun, imagesBean)
	if err != nil {
		log.Fatal(err)
	}
	scheduler.StartAsync()
}

func buildImagesJobRun(imagesBean *images.Bean) {
	imagesToBuild := imagesBean.GetImages(0, "", mustBuildImageState, 0, batchSize)
	for _, image := range *imagesToBuild {
		go buildImage(image, imagesBean)
	}
}

func buildImage(image images.ImageEntity, bean *images.Bean) {
	image.State = "building"
	err := bean.EditImage(&image)
	if err != nil {
		log.Print("error occurred while setting image state to building", err)
		return
	}

}

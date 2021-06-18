package images

import (
	"github.com/go-co-op/gocron"
	"github.com/z13z/Kiosks/master-server/db/images"
	"log"
	"os"
	"os/exec"
	"time"
)

//todo zaza
//const mustBuildImageState = "waiting"
const mustBuildImageState = ""
const imageDoneState = "done"
const imageBuildingState = "building"
const imageFailedState = "failed"
const batchSize = 2
const jobIntervalSecs = 60
const kiosksImagesScriptsDirectoryName = "kiosk-image"
const kiosksImagesDirectory = "kiosk-image-result"
const buildOutputFileName = "output.txt"
const buildErrorFileName = "error.txt"

func BuildImagesJob() {
	err := downloadUbuntuImage()
	if err != nil {
		log.Fatal(err)
	}
	imagesBean := images.NewBean()
	scheduler := gocron.NewScheduler(time.UTC)
	_, err = scheduler.Every(jobIntervalSecs).Minutes().Do(buildImagesJobRun, imagesBean)
	if err != nil {
		log.Fatal(err)
	}
	scheduler.StartAsync()
}

func downloadUbuntuImage() error {
	err := RestoreAsset(".", "kiosk-image/download_ubuntu_image")
	if err != nil {
		return err
	}
	e := exec.Command("./kiosk-image/download_ubuntu_image")
	_, err = os.Create(buildOutputFileName)
	_, err = os.Create(buildErrorFileName)
	closeFile := func(errorFile *os.File) {
		_ = errorFile.Close()
	}
	outputFile, _ := os.Open(buildOutputFileName)
	defer closeFile(outputFile)
	errorFile, _ := os.Open(buildErrorFileName)
	defer closeFile(errorFile)
	e.Stdout = outputFile
	e.Stderr = errorFile
	err = e.Run()
	return err
}

func exportKioskBuildScriptsInDir(directory string) error {
	err := RestoreAsset(directory, "kiosk-image/Makefile")
	err = RestoreAsset(directory, "kiosk-image/create_custom_image")
	err = RestoreAsset(directory, "kiosk-image/chroot_commands")
	err = RestoreAsset(directory, "kiosk-image/prepare_kiosk")
	return err
}

func buildImagesJobRun(imagesBean *images.Bean) {
	imagesToBuild := imagesBean.GetImages(0, "", mustBuildImageState, 0, batchSize)
	for _, image := range *imagesToBuild {
		buildImage(&image, imagesBean)
	}
}

func buildImage(image *images.ImageEntity, bean *images.Bean) {
	if !setImageState(image, bean, imageBuildingState) {
		return
	}
	err := runMakeIsoScript(image)
	if err != nil {
		if !setImageState(image, bean, imageFailedState) {
			return
		}
	}
	setImageState(image, bean, imageDoneState)
}

func setImageState(image *images.ImageEntity, bean *images.Bean, state string) bool {
	image.State = state
	err := bean.EditImage(image)
	if err != nil {
		log.Print("error occurred while setting image state to "+state, err)
		return false
	}
	return true
}

func runMakeIsoScript(image *images.ImageEntity) error {
	err := createDirectoryAndLogFilesForKioskImage(image)
	if err != nil {
		return err
	}
	e := exec.Command("make")
	e.Dir = kiosksImagesDirectory + "/" + image.Name + "/" + kiosksImagesScriptsDirectoryName
	imageDir := kiosksImagesDirectory + "/" + image.Name
	e.Stdout, _ = os.Open(imageDir + "/" + buildOutputFileName)
	e.Stderr, _ = os.Open(imageDir + "/" + buildErrorFileName)
	err = e.Run()
	return err
}

// 1) mkdir for build directory and cd there.
// 2) mkdir for kiosk image directory (remove before if exists).
// 3) copy build scripts to kiosk image directory.
func createDirectoryAndLogFilesForKioskImage(image *images.ImageEntity) error {
	imageDir := kiosksImagesDirectory + "/" + image.Name
	err := os.Mkdir(kiosksImagesDirectory, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	err = os.RemoveAll(imageDir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	err = os.Mkdir(imageDir, 0755)
	if err != nil {
		return err
	}
	_, err = os.Create(imageDir + "/" + buildOutputFileName)
	if err != nil {
		return err
	}
	_, err = os.Create(imageDir + "/" + buildErrorFileName)
	if err != nil {
		return err
	}
	return exportKioskBuildScriptsInDir(imageDir)
}

package images

import (
	"bytes"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/z13z/Kiosks/master-server/db/images"
	"log"
	"os"
	"os/exec"
	"time"
)

const OutputFileName = "ubuntu-20.04.2-server-custom.iso"
const OutputFileDir = "build"
const KiosksImagesDirectory = "kiosk-image-result"
const KiosksImagesScriptsDirectoryName = "kiosk-image"
const mustBuildImageState = "waiting"
const imageDoneState = "done"
const imageBuildingState = "building"
const imageFailedState = "failed"
const batchSize = 2
const jobIntervalMins = 10 * batchSize
const buildOutputFileName = "output.txt"
const buildErrorFileName = "error.txt"

func BuildImagesJob() {
	err := downloadUbuntuImage()
	if err != nil {
		log.Fatal(err)
	}
	imagesBean := images.NewBean()
	scheduler := gocron.NewScheduler(time.UTC)
	_, err = scheduler.Every(jobIntervalMins).Minutes().Do(buildImagesJobRun, imagesBean)
	if err != nil {
		log.Fatal(err)
	}
	scheduler.StartAsync()
}

func downloadUbuntuImage() error {
	err := RestoreAsset(".", "kiosk-image/download_ubuntu_image")
	_ = os.Chmod("kiosk-image", 0777)
	_ = os.Chmod("kiosk-image/download_ubuntu_image", 0777)

	if err != nil {
		return err
	}
	// if command is not in pass call fails :)
	e := exec.Command("./kiosk-image/download_ubuntu_image")
	var outBuf, errBuf bytes.Buffer
	e.Stdout = &outBuf
	e.Stderr = &errBuf
	err = e.Run()
	errB := writeBufferToFile(&outBuf, buildOutputFileName)
	if errB != nil {
		log.Print("Error while writing buffer in: "+buildOutputFileName, errB)
	}
	errB = writeBufferToFile(&errBuf, buildErrorFileName)
	if errB != nil {
		log.Print("Error while writing buffer in: "+buildErrorFileName, errB)
	}
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
	log.Printf("Starging build of image (%s)", image.Name)
	defer log.Printf("Finished build of image (%s)", image.Name)
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
	defer deleteImageBuildFiles(image.Name)
	if err != nil {
		return err
	}
	e := exec.Command("make")
	e.Dir = KiosksImagesDirectory + "/" + image.Name + "/" + KiosksImagesScriptsDirectoryName
	imageDir := KiosksImagesDirectory + "/" + image.Name
	var outBuf, errBuf bytes.Buffer
	e.Stdout = &outBuf
	e.Stderr = &errBuf
	err = e.Run()
	errB := writeBufferToFile(&outBuf, imageDir+"/"+buildOutputFileName)
	if errB != nil {
		log.Print("Error while writing buffer in: "+imageDir+"/"+buildOutputFileName, errB)
	}
	errB = writeBufferToFile(&errBuf, imageDir+"/"+buildErrorFileName)
	if errB != nil {
		log.Print("Error while writing buffer in: "+imageDir+"/"+buildErrorFileName, errB)
	}
	return err
}

func deleteImageBuildFiles(imageName string) {
	imageDir := KiosksImagesDirectory + "/" + imageName
	err := os.Remove(imageDir + "/kiosk-image/Makefile")
	if err != nil {
		log.Print(fmt.Sprintf("Error while deleting image build script (%s)", imageDir+"/kiosk-image/Makefile"), err)
	}

	err = os.Remove(imageDir + "/kiosk-image/create_custom_image")
	if err != nil {
		log.Print(fmt.Sprintf("Error while deleting image build script (%s)", imageDir+"/kiosk-image/create_custom_image"), err)
	}

	err = os.Remove(imageDir + "/kiosk-image/chroot_commands")
	if err != nil {
		log.Print(fmt.Sprintf("Error while deleting image build script (%s)", imageDir+"/kiosk-image/create_custom_image"), err)
	}

	err = os.Remove(imageDir + "/kiosk-image/prepare_kiosk")
	if err != nil {
		log.Print(fmt.Sprintf("Error while deleting image build script (%s)", imageDir+"/kiosk-image/prepare_kiosk"), err)
	}
}

func writeBufferToFile(w *bytes.Buffer, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	_, err = f.Write(w.Bytes())
	return err
}

// 1) mkdir for build directory and cd there.
// 2) mkdir for kiosk image directory (remove before if exists).
// 3) copy build scripts to kiosk image directory.
func createDirectoryAndLogFilesForKioskImage(image *images.ImageEntity) error {
	imageDir := KiosksImagesDirectory + "/" + image.Name
	err := os.Mkdir(KiosksImagesDirectory, 0755)
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
	return exportKioskBuildScriptsInDir(imageDir)
}

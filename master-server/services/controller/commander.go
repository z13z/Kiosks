package controller

import (
	"fmt"
	"github.com/z13z/Kiosks/master-server/crypto"
	"github.com/z13z/Kiosks/master-server/db/kiosks"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//goland:noinspection HttpUrlsUsage
const kioskScreenshotServiceAddressTemplate = "http://%s/screenshot"

//goland:noinspection HttpUrlsUsage
const kioskExecuteServiceAddressTemplate = "http://%s/execute"
const KioskImageScreenshotContentType = "image/png"
const authenticationHeader = "Authentication"
const callTimeoutValue = 4 * time.Second

func sendRequestToKiosk(entity *kiosks.KioskEntity, method, address string, body io.Reader) ([]byte, error) {
	client := &http.Client{Timeout: callTimeoutValue}
	req, err := http.NewRequest(method, address, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authenticationHeader, crypto.Decrypt(entity.ServicePassword))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func GetKioskScreenshot(entity *kiosks.KioskEntity) ([]byte, error) {
	return sendRequestToKiosk(entity, "GET", fmt.Sprintf(kioskScreenshotServiceAddressTemplate, entity.Address), nil)
}

func SendCommandToKiosk(entity *kiosks.KioskEntity, command string) ([]byte, error) {
	reader := strings.Reader{}
	reader.Reset(command)
	return sendRequestToKiosk(entity, "POST", fmt.Sprintf(kioskExecuteServiceAddressTemplate, entity.Address), &reader)
}

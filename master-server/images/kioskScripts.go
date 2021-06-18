// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// kiosk-image/Makefile
// kiosk-image/chroot_commands
// kiosk-image/create_custom_image
// kiosk-image/download_ubuntu_image
// kiosk-image/prepare_kiosk
package images

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _kioskImageMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xb1\x6a\xc3\x30\x10\x86\x67\xdd\x53\xdc\xd0\xa9\x20\x39\x84\xd0\xc1\x10\xa8\x03\xa6\xcd\x10\xa7\x34\x75\xa1\x93\x50\x24\x25\x11\xb6\x2c\x23\x59\x59\xda\xbc\x7b\xa9\xec\xd1\xa1\x70\xd3\x1d\x1f\xdf\x7f\x3f\x6c\x8a\x43\xc9\xeb\x4d\x5d\x7d\xd4\x7c\xbb\x2b\x5e\x4a\x5e\x15\xbb\x72\x1d\x8f\xb1\x1b\x22\x5d\x2e\xd8\x62\xc5\x96\xb4\x35\x57\x4d\x83\xf6\x57\xed\xa9\xb0\xea\x69\xc5\x4c\x70\x33\xec\x67\xf9\x7e\xd8\xee\xab\xf5\xc4\x01\x1c\xa3\x69\x55\x8e\xc6\x8a\xb3\x0e\xd9\xc3\xf7\xbc\xee\x86\xf2\xe2\x9d\x1b\xb8\x74\xd6\x8a\x4e\x05\x94\x5e\x8b\x41\x73\x19\xc3\xe0\x2c\x4f\x38\xf6\x5e\xf7\xc2\x6b\xde\x18\x17\x1a\x20\x84\x90\x67\x2d\x2f\x0e\x93\xc3\x74\x67\x4c\x87\xd1\x85\x27\xef\x2c\xde\xf5\x25\x3a\xc1\x2c\x9b\x53\xdd\x0f\xfa\x83\x21\x2a\x87\x21\x02\xc8\x56\x8b\x2e\x07\x92\x16\xde\x22\xf5\xa7\x31\x0b\xc0\x7f\xff\xe6\x40\x6c\xa3\x8c\x47\xda\x4f\xdd\x00\x91\x3d\x32\x96\x8d\xf3\xf8\xd7\xef\x6c\x8a\xa9\xe1\x1b\x00\x7b\x7b\xdd\x57\x5f\xf9\xa4\xfc\x0d\x00\x00\xff\xff\x63\x45\xef\x8f\xc9\x01\x00\x00")

func kioskImageMakefileBytes() ([]byte, error) {
	return bindataRead(
		_kioskImageMakefile,
		"kiosk-image/Makefile",
	)
}

func kioskImageMakefile() (*asset, error) {
	bytes, err := kioskImageMakefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kiosk-image/Makefile", size: 457, mode: os.FileMode(436), modTime: time.Unix(1624042352, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kioskImageChroot_commands = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xc1\x8a\x1b\x31\x10\x44\xef\xfa\x8a\x0e\x7b\x4b\x90\x3b\x3f\xb0\x81\xb0\x04\x72\x70\xc8\x27\x18\x8d\xd4\x63\x0b\x8f\xba\x45\xab\x35\xb1\x61\xc9\xb7\x87\x99\x78\x76\x63\x43\x6e\xd5\xf5\xa4\x43\x55\x3d\x7d\xc0\x21\x33\x0e\xa1\x9d\x5c\x91\xce\x06\xde\xa0\xaa\x44\x60\x61\x02\x5c\xe4\x3b\x68\xd7\x36\xb6\x1b\x69\xd7\xf6\x0e\x12\xcd\xd5\x36\x92\x68\xc6\x6a\xcd\xd1\xa5\x8a\x1a\x7c\xff\xf9\xe3\xdb\x33\xaa\x88\x6d\xce\xfe\xe5\xf0\x75\xbf\x7f\x7e\x71\x69\xe8\xcd\xf7\x9e\xd3\x91\x18\xbe\x00\xce\x41\x71\xca\x03\x2e\x3e\x96\x10\x4f\x99\xc9\xe7\xe4\x52\x3d\x1f\x7d\xca\x33\xa9\x81\xf7\x93\xc4\x30\x81\xf7\x4a\x1c\x0a\x81\xf7\x21\x25\xc0\xb6\xe4\xc8\x9c\x2d\xda\xe4\x26\x06\xdf\x60\x8d\x66\xda\xe9\x81\xba\x27\xed\xcc\x99\x8f\x10\xa5\x94\xc0\xa9\xc1\x28\x0a\xe7\x2c\xed\xec\x73\x09\x47\x82\xa8\x14\x2c\x0b\xc3\xa8\x52\x80\x4a\xb5\x2b\xf4\xa1\xb3\x75\x17\x4f\x45\x12\x7c\xba\x40\x55\xaa\x41\xe9\xb0\xfe\x73\x3b\xc0\x5f\xa2\xe7\x56\x43\x24\xbc\x47\x2e\x54\x83\x38\x51\x60\xa7\x05\xbc\x8e\x80\x56\x2a\x7e\x84\xdf\xb8\x5b\x9a\x3f\x9c\x72\x33\xd1\xeb\x42\x91\x2c\xa2\x52\x93\x69\xde\x45\xe1\x71\xf5\xfe\xd7\xcb\xc2\xee\x92\xdd\x17\xf5\x56\x90\x52\x91\xf9\xb1\x85\xfe\x77\xbd\x75\x62\x78\x7d\x85\xdb\xed\xa7\xf1\x36\xfb\xf6\x60\x59\x7a\xd3\xdb\xb6\xff\xdc\x6f\x5a\x3b\x3b\xba\x64\x83\xcf\x7f\x02\x00\x00\xff\xff\xa6\x79\x67\xf1\x56\x02\x00\x00")

func kioskImageChroot_commandsBytes() ([]byte, error) {
	return bindataRead(
		_kioskImageChroot_commands,
		"kiosk-image/chroot_commands",
	)
}

func kioskImageChroot_commands() (*asset, error) {
	bytes, err := kioskImageChroot_commandsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kiosk-image/chroot_commands", size: 598, mode: os.FileMode(509), modTime: time.Unix(1624020774, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kioskImageCreate_custom_image = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\x6f\x4f\x23\xb7\x13\x7e\x9f\x4f\x31\x3f\x12\xc1\xaf\xaa\xbc\x1b\x4e\x5c\x5b\x55\xa2\x12\x05\x54\x5d\xdb\x6b\x4f\x42\xba\xbe\x41\x42\x8e\x3d\x9b\xb5\x62\x7b\x16\x8f\x1d\x92\x03\xbe\x7b\x65\x27\x61\x17\x48\xaf\xd7\xf6\x4d\x14\xcf\x3e\xf3\x78\xe6\x99\x3f\x1e\xff\xaf\x9e\x19\x5f\x73\x3b\x1a\x5f\xa9\x60\xba\x08\x86\xc1\x49\x8d\xd0\x04\x72\xa0\x49\x25\x87\x3e\xca\x68\xc8\x43\x1b\x63\xc7\xdf\xd7\x75\x8b\xb6\xab\xd2\x2c\xf9\x98\x2a\x45\xae\x56\xe4\x5c\xf2\x26\xae\xeb\x5f\xcd\x12\xcf\x2f\xce\x13\x47\x72\xe6\x53\xf1\x1a\xa1\x6a\x09\x0e\x54\x40\x19\x11\x54\xf9\x04\xc6\xc9\x39\x1e\xc0\xe1\x21\x5c\x8f\x7e\x3c\xbb\xba\xbc\x79\xf7\xfe\xec\xa7\xcb\x9b\xdf\xce\xde\x5f\x9e\x4e\x8e\xe1\xf0\xf0\x7a\x34\xb0\xfc\x62\x88\x17\xc5\xc8\x49\x13\xc8\x1c\xa5\xe7\x28\xad\x05\xb1\x06\xbe\x4d\x92\xdb\x86\x45\x24\xb2\x0c\x73\xf4\x86\xa9\x5c\x50\x5c\xdc\x42\x9b\x00\xb3\x64\xac\x2e\x67\xa5\x07\x87\xcd\x47\xd1\x81\xf3\xf1\xb9\x01\x57\x31\x48\x15\x85\xda\x00\x59\x75\x50\x55\x75\xa1\xe5\x7a\x72\xff\x22\xe8\x47\xa8\xfa\xf8\x1c\x25\x1f\x41\x10\x58\xa2\x0e\xf6\x60\x77\x97\x15\x74\xe0\xb5\x57\x20\x04\xae\x94\x4d\x1a\x4f\x6b\x25\xb9\xc3\x50\x37\xc6\x22\xaf\x39\xa2\xab\x76\x19\x82\x90\xd9\xb7\x7e\x15\x5c\xe6\x49\xfe\x09\x96\x31\x9f\x61\xe9\x03\x5d\xf6\xe2\x05\xa2\x08\xa8\xcd\x20\x32\xd5\x41\x8d\x51\xd5\x01\x99\xec\xb2\x52\xe4\x9b\x82\x28\xc6\x3d\xe9\xce\x8c\xd7\x50\x87\xe4\xeb\x0d\x2c\x24\xff\x0a\x25\x36\x20\x8d\xcb\x2d\x48\xe3\x72\x00\x2a\xea\x17\xfb\x1d\x85\x05\x77\x52\xe1\x50\xfe\x2e\x60\x27\x03\xde\x2c\x4a\x43\xfc\x35\x4e\xb5\x39\x9d\x9b\xdc\x97\xd2\x6b\xde\x8b\x2c\x19\xb6\x8e\x34\x7c\x7d\x37\x10\x74\x8f\x70\x4e\x7a\xd3\x20\x0f\x95\x69\x7b\xb9\x7a\xde\x57\xd7\x66\x7c\xe9\xfe\x0f\x81\x74\x52\xc6\xcf\x21\xb6\x08\xe7\x17\xd0\xb7\xe7\x68\x5c\x18\x39\x95\xd3\xf8\x15\xbd\xee\x16\x73\x71\x9b\x30\xac\x41\xfc\x01\x42\x70\x4b\x77\x0d\x05\x27\xe3\xe9\xd1\xe4\xfe\x83\x54\x0b\x39\xc7\x47\x98\xdc\x7f\xc4\xc0\x86\xfc\xe3\xb5\x3f\x82\x1f\xfe\x41\x42\x63\x5c\xbd\x28\xfa\x97\xf9\x7e\x11\x4a\x68\xe4\x45\xa4\xae\xe7\x67\xd4\x20\x0c\x1c\xd5\x69\x66\x6e\x53\xde\x19\xfa\xe8\x3f\x73\x6d\x9d\xfe\x15\x53\x29\xd0\x39\xb9\x2e\x20\x67\x01\x87\xdd\xf8\x34\x32\xa5\x14\x9f\xe5\xee\x67\x74\x06\xc7\xd3\x93\xef\xde\x7e\xfb\x4d\x61\xea\x82\xf1\xb1\x81\xc9\xff\x75\x02\xc1\xab\x3c\x01\x96\xd4\x42\xb0\xf9\x84\xa7\xc7\x1b\xe2\x07\x50\x29\x82\x68\x8e\xbf\xfa\xbb\xca\x65\xaf\x3e\x6a\xb2\x1a\x9c\x7e\xcb\xc9\x55\x71\x15\x41\x7a\x0d\x4a\x5a\x95\x6c\xde\xb4\x1e\xef\xf2\x47\xe0\xe4\x78\xb7\xf8\xf6\xed\x8d\xe0\x86\x1c\xd9\xdc\xe4\x11\x15\x71\xdd\x21\x34\x20\x4a\x02\x53\x78\x80\x82\x5e\xc9\x30\x67\x10\xd3\xad\x0f\x3c\xc0\x3c\x60\x07\x62\x09\x86\xc9\x1a\x9f\x56\xf5\x8c\x28\x56\x4a\xc6\x9d\x4b\x44\x7c\x79\x43\x89\x3e\x8f\xc2\xbb\xab\xdf\x07\xb3\xb0\x55\xdd\x30\x65\x1d\x2f\x40\x04\x10\x1f\xe1\x60\xd2\xef\xce\x03\x10\x4a\xaa\x16\x85\xf1\xa4\x91\x41\xfc\x0c\xc2\x66\xc9\x9f\x6e\xdf\xfd\xa9\x66\xc6\x83\x50\x7b\xc2\x12\x9e\x04\xba\x64\x45\xb6\x40\xf9\x15\x96\xa4\x2e\x25\x81\x93\xad\xc5\xf8\x86\x44\x94\x33\x8b\x79\xb5\x55\x55\xbd\x79\xef\xc4\x9b\x69\x35\x3d\xa9\xde\x08\xc6\xb0\xc4\x20\x36\xef\x59\x65\x98\xb6\x2f\x80\xd2\x50\x55\xa3\xb4\x59\x77\xbb\x45\x1f\x5c\xce\xe5\xd9\xa9\xe9\x97\xed\xee\xfc\xbc\x38\x5b\xeb\xf3\xad\xb5\x35\x4e\x5e\x3c\x29\x7f\x06\x00\x00\xff\xff\x67\xe9\xcc\x24\xc6\x07\x00\x00")

func kioskImageCreate_custom_imageBytes() ([]byte, error) {
	return bindataRead(
		_kioskImageCreate_custom_image,
		"kiosk-image/create_custom_image",
	)
}

func kioskImageCreate_custom_image() (*asset, error) {
	bytes, err := kioskImageCreate_custom_imageBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kiosk-image/create_custom_image", size: 1990, mode: os.FileMode(509), modTime: time.Unix(1624020807, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kioskImageDownload_ubuntu_image = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x4f\x4a\x2c\xce\xe0\x72\x72\x0c\x76\x8d\x0f\x75\x0a\xf5\x0b\x09\x8d\xf7\xf4\x75\x74\x77\x8d\xf7\x73\xf4\x75\xb5\x2d\x4d\x2a\xcd\x2b\x29\xd5\x35\x32\xd0\x33\x30\xd1\x33\xd2\xcd\xc9\x2c\x4b\xd5\x2d\x4e\x2d\x2a\x4b\x2d\xd2\x4d\xcc\x4d\x31\x33\xd1\xcb\x2c\xce\xe7\x4a\x2e\x2d\xca\x51\x50\xca\x28\x29\x29\x28\xb6\xd2\xd7\x2f\x4a\xcd\x49\x4d\x2c\x4e\x2d\xd6\x83\xe8\xd5\x4b\xce\xcf\xd5\x57\xa9\xc6\x34\x3f\xcc\x35\x28\xd8\xd3\xdf\xaf\x16\xab\x24\xc8\xf2\x5a\x25\x05\x3b\x05\x25\x7c\x5a\x95\x00\x01\x00\x00\xff\xff\xb8\x3b\x6e\x37\xc2\x00\x00\x00")

func kioskImageDownload_ubuntu_imageBytes() ([]byte, error) {
	return bindataRead(
		_kioskImageDownload_ubuntu_image,
		"kiosk-image/download_ubuntu_image",
	)
}

func kioskImageDownload_ubuntu_image() (*asset, error) {
	bytes, err := kioskImageDownload_ubuntu_imageBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kiosk-image/download_ubuntu_image", size: 194, mode: os.FileMode(509), modTime: time.Unix(1624042292, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kioskImagePrepare_kiosk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xef\x6f\x22\x37\x10\xfd\xee\xbf\x62\x4a\x39\x04\x1f\x8c\x41\xa9\x2e\xd5\xe9\x16\x95\x0a\x7a\x42\x25\x77\x27\x48\x94\x54\xd9\x28\x32\xeb\x01\xac\x78\xed\x95\xed\x05\x36\x4d\xff\xf7\x6a\x7f\x95\x4d\x0a\xd1\x7d\x81\xf5\xcc\xf3\x9b\x37\xe3\x67\xff\xfc\x13\x5b\x49\xcd\x56\xdc\x6d\x89\x4b\x85\x01\x9e\x78\x48\x13\xc1\x3d\x42\xa7\x03\x61\x33\xb8\xb1\x5c\x20\xd0\xec\x4d\x42\x6a\xe7\xb9\x52\x79\x42\xc9\x15\x6e\xd4\x90\xc6\xe8\x78\xbd\x38\x8f\x3e\xec\x79\xa6\xb8\x16\xef\x20\x84\xf1\xc6\xa8\xf3\x80\x58\xea\x0d\x7a\xff\x8e\xa4\xb5\xb4\xb8\x36\x87\x06\x20\x7e\x12\xd2\x02\x43\x1f\x31\x97\x39\x8f\xb1\xa8\xfe\x59\x41\xf5\x9b\xf7\xd9\xb0\xef\xd0\xee\x64\x84\xfd\xa6\xb8\xd4\xa1\xe5\x42\x00\x8d\x81\x7e\x81\x22\x44\xa9\xdb\xa2\x52\xf0\xdf\x10\xe1\x49\x1a\xf7\xd4\xd8\x94\x70\xe7\xf6\xa2\x0c\x7f\xfe\x3c\xfd\xf6\x07\x79\x1e\x5e\x3c\xb3\x3f\xf3\x75\xf3\x33\xcf\x14\x1b\x30\xda\x1a\xa0\xa8\xa1\x15\xea\x50\x97\x74\xe3\xf9\x3c\xe8\x8e\xe7\xf3\x1e\x7c\xfd\xf6\x7d\xbc\x5c\xde\x4e\x3e\xe5\xb1\xd6\x68\x54\xb6\x91\x0a\x83\xd6\x35\xaa\x46\xdc\xc3\xe8\x47\x5b\x64\x66\x87\xd6\x4a\x81\xfd\xc8\xe8\x35\x94\x2a\xef\x97\x65\xfa\x81\x4c\x0f\x18\x2d\x3d\xb7\x3e\x68\x7c\x52\xe6\xf2\x96\x79\x39\x7d\x4a\xb5\x91\xce\xa5\x08\x94\xf2\xd4\x1b\x65\x36\x52\x57\xa3\xf8\x30\x83\xb0\x7d\x3d\x5d\x5c\x91\xeb\x2c\xc1\x40\x0a\x85\xc7\x66\xcf\xea\x2c\xf6\xd6\x12\x2b\x45\x37\x5a\xfa\x07\x32\x41\x17\x59\x99\x78\x69\x74\x70\x65\x52\xed\x21\xd5\xd2\x43\x39\xc5\xf1\xda\xa3\x0d\x9c\xe6\x89\xa8\x37\x93\x93\xad\xb0\xd4\x59\xa6\x4c\xc4\x55\x71\x74\x55\xb9\x2d\xb9\x71\x68\x83\x62\x45\x16\xe8\x0a\x28\x57\x7b\x9e\xb9\x7a\xb9\xc4\x28\xb8\x20\xe4\x7e\x56\x5a\xec\x81\xdc\x72\xed\x51\xfc\x9e\x05\x02\xd7\x3c\x55\xbe\xef\xb9\xdd\xa0\x6f\xf4\xb8\x8d\x8d\x80\x8f\x1f\x7f\x39\xe9\xb9\x57\x8d\xfe\xef\x04\x4f\xab\x2c\xe7\xf1\xea\xea\xa2\xcf\x6d\xd9\xe9\x00\x1e\x12\x63\x3d\xdc\x4d\xbe\x3c\x2e\x6e\xbe\x5e\xcf\xae\xa6\x8f\x93\xd9\x22\x60\x36\xd5\x2c\x37\x30\x1b\x0e\x06\x83\x1c\xe8\x14\x62\x02\xc3\x8b\xb2\x66\xb7\x28\x9a\x0f\x0e\x36\xe8\x21\x96\x96\x56\x4e\x1e\x31\x81\x3b\xa6\x53\xa5\xe0\xe5\x05\x8e\xb0\xfa\x8e\x1d\xa1\x94\x0a\xdc\xc5\x46\x60\xaf\xe2\x64\x39\xb0\x10\x79\x9a\xaf\xd3\x3b\x0a\xf9\xb5\x21\x7e\x32\x5b\x7e\x9f\x8f\xff\x0a\x3e\x55\x8f\x47\xf7\xae\x7e\x29\xe8\xde\x42\xd8\xae\xf2\xe7\xb9\x48\x45\xc4\x93\xe4\x51\xf3\x18\x83\xea\x15\x20\xc5\xe5\x6a\x15\x07\x29\xf5\x06\xc2\xf6\xdf\x35\xe4\x9f\x16\xe9\x36\x47\x52\x3f\x3d\x7b\xa9\x85\xd9\x3b\xf9\x8c\x10\xb6\xbb\x75\xd4\x21\xb7\xd1\x16\x28\x35\x5a\x65\x3b\xe9\xe4\x4a\xe5\xde\xcf\x89\x20\x6c\xd7\x9c\x3d\x18\x0e\x06\x1f\x8a\x9f\x1e\x74\xc8\x31\xf1\xd6\x1c\x97\x97\x97\x70\xe6\xac\x1b\x96\x28\x1d\x13\x79\x05\xa8\x79\x5e\xf0\x95\x77\xc8\xbf\x01\x00\x00\xff\xff\x5c\xe3\x7c\x96\xcb\x05\x00\x00")

func kioskImagePrepare_kioskBytes() ([]byte, error) {
	return bindataRead(
		_kioskImagePrepare_kiosk,
		"kiosk-image/prepare_kiosk",
	)
}

func kioskImagePrepare_kiosk() (*asset, error) {
	bytes, err := kioskImagePrepare_kioskBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kiosk-image/prepare_kiosk", size: 1483, mode: os.FileMode(509), modTime: time.Unix(1623493567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"kiosk-image/Makefile":              kioskImageMakefile,
	"kiosk-image/chroot_commands":       kioskImageChroot_commands,
	"kiosk-image/create_custom_image":   kioskImageCreate_custom_image,
	"kiosk-image/download_ubuntu_image": kioskImageDownload_ubuntu_image,
	"kiosk-image/prepare_kiosk":         kioskImagePrepare_kiosk,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"kiosk-image": &bintree{nil, map[string]*bintree{
		"Makefile":              &bintree{kioskImageMakefile, map[string]*bintree{}},
		"chroot_commands":       &bintree{kioskImageChroot_commands, map[string]*bintree{}},
		"create_custom_image":   &bintree{kioskImageCreate_custom_image, map[string]*bintree{}},
		"download_ubuntu_image": &bintree{kioskImageDownload_ubuntu_image, map[string]*bintree{}},
		"prepare_kiosk":         &bintree{kioskImagePrepare_kiosk, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

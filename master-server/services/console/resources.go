// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// prepare_kiosk
package console

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _prepare_kiosk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xef\x6f\x22\x37\x10\xfd\xee\xbf\x62\x4a\x39\x04\x1f\x8c\x41\xa9\x2e\xd5\xe9\x16\x95\x0a\x7a\x42\x25\x77\x27\x48\x94\x54\xd9\x28\x32\xeb\x01\xac\x78\xed\x95\xed\x05\x36\x4d\xff\xf7\x6a\x7f\x95\x4d\x0a\xd1\x7d\x81\xf5\xcc\xf3\x9b\x37\xe3\x67\xff\xfc\x13\x5b\x49\xcd\x56\xdc\x6d\x89\x4b\x85\x01\x9e\x78\x48\x13\xc1\x3d\x42\xa7\x03\x61\x33\xb8\xb1\x5c\x20\xd0\xec\x4d\x42\x6a\xe7\xb9\x52\x79\x42\xc9\x15\x6e\xd4\x90\xc6\xe8\x78\xbd\x38\x8f\x3e\xec\x79\xa6\xb8\x16\xef\x20\x84\xf1\xc6\xa8\xf3\x80\x58\xea\x0d\x7a\xff\x8e\xa4\xb5\xb4\xb8\x36\x87\x06\x20\x7e\x12\xd2\x02\x43\x1f\x31\x97\x39\x8f\xb1\xa8\xfe\x59\x41\xf5\x9b\xf7\xd9\xb0\xef\xd0\xee\x64\x84\xfd\xa6\xb8\xd4\xa1\xe5\x42\x00\x8d\x81\x7e\x81\x22\x44\xa9\xdb\xa2\x52\xf0\xdf\x10\xe1\x49\x1a\xf7\xd4\xd8\x94\x70\xe7\xf6\xa2\x0c\x7f\xfe\x3c\xfd\xf6\x07\x79\x1e\x5e\x3c\xb3\x3f\xf3\x75\xf3\x33\xcf\x14\x1b\x30\xda\x1a\xa0\xa8\xa1\x15\xea\x50\x97\x74\xe3\xf9\x3c\xe8\x8e\xe7\xf3\x1e\x7c\xfd\xf6\x7d\xbc\x5c\xde\x4e\x3e\xe5\xb1\xd6\x68\x54\xb6\x91\x0a\x83\xd6\x35\xaa\x46\xdc\xc3\xe8\x47\x5b\x64\x66\x87\xd6\x4a\x81\xfd\xc8\xe8\x35\x94\x2a\xef\x97\x65\xfa\x81\x4c\x0f\x18\x2d\x3d\xb7\x3e\x68\x7c\x52\xe6\xf2\x96\x79\x39\x7d\x4a\xb5\x91\xce\xa5\x08\x94\xf2\xd4\x1b\x65\x36\x52\x57\xa3\xf8\x30\x83\xb0\x7d\x3d\x5d\x5c\x91\xeb\x2c\xc1\x40\x0a\x85\xc7\x66\xcf\xea\x2c\xf6\xd6\x12\x2b\x45\x37\x5a\xfa\x07\x32\x41\x17\x59\x99\x78\x69\x74\x70\x65\x52\xed\x21\xd5\xd2\x43\x39\xc5\xf1\xda\xa3\x0d\x9c\xe6\x89\xa8\x37\x93\x93\xad\xb0\xd4\x59\xa6\x4c\xc4\x55\x71\x74\x55\xb9\x2d\xb9\x71\x68\x83\x62\x45\x16\xe8\x0a\x28\x57\x7b\x9e\xb9\x7a\xb9\xc4\x28\xb8\x20\xe4\x7e\x56\x5a\xec\x81\xdc\x72\xed\x51\xfc\x9e\x05\x02\xd7\x3c\x55\xbe\xef\xb9\xdd\xa0\x6f\xf4\xb8\x8d\x8d\x80\x8f\x1f\x7f\x39\xe9\xb9\x57\x8d\xfe\xef\x04\x4f\xab\x2c\xe7\xf1\xea\xea\xa2\xcf\x6d\xd9\xe9\x00\x1e\x12\x63\x3d\xdc\x4d\xbe\x3c\x2e\x6e\xbe\x5e\xcf\xae\xa6\x8f\x93\xd9\x22\x60\x36\xd5\x2c\x37\x30\x1b\x0e\x06\x83\x1c\xe8\x14\x62\x02\xc3\x8b\xb2\x66\xb7\x28\x9a\x0f\x0e\x36\xe8\x21\x96\x96\x56\x4e\x1e\x31\x81\x3b\xa6\x53\xa5\xe0\xe5\x05\x8e\xb0\xfa\x8e\x1d\xa1\x94\x0a\xdc\xc5\x46\x60\xaf\xe2\x64\x39\xb0\x10\x79\x9a\xaf\xd3\x3b\x0a\xf9\xb5\x21\x7e\x32\x5b\x7e\x9f\x8f\xff\x0a\x3e\x55\x8f\x47\xf7\xae\x7e\x29\xe8\xde\x42\xd8\xae\xf2\xe7\xb9\x48\x45\xc4\x93\xe4\x51\xf3\x18\x83\xea\x15\x20\xc5\xe5\x6a\x15\x07\x29\xf5\x06\xc2\xf6\xdf\x35\xe4\x9f\x16\xe9\x36\x47\x52\x3f\x3d\x7b\xa9\x85\xd9\x3b\xf9\x8c\x10\xb6\xbb\x75\xd4\x21\xb7\xd1\x16\x28\x35\x5a\x65\x3b\xe9\xe4\x4a\xe5\xde\xcf\x89\x20\x6c\xd7\x9c\x3d\x18\x0e\x06\x1f\x8a\x9f\x1e\x74\xc8\x31\xf1\xd6\x1c\x97\x97\x97\x70\xe6\xac\x1b\x96\x28\x1d\x13\x79\x05\xa8\x79\x5e\xf0\x95\x77\xc8\xbf\x01\x00\x00\xff\xff\x5c\xe3\x7c\x96\xcb\x05\x00\x00")

func prepare_kioskBytes() ([]byte, error) {
	return bindataRead(
		_prepare_kiosk,
		"prepare_kiosk",
	)
}

func prepare_kiosk() (*asset, error) {
	bytes, err := prepare_kioskBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prepare_kiosk", size: 1483, mode: os.FileMode(509), modTime: time.Unix(1608228601, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"prepare_kiosk": prepare_kiosk,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"prepare_kiosk": &bintree{prepare_kiosk, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}

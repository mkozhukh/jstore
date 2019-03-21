// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// store/dataitem.go
// store/store.go
package main

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _storeDataitemGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x2e\xc9\x2f\x4a\xe5\xe2\x2a\xa9\x2c\x48\x55\x70\x49\x2c\x49\xf4\x2c\x49\xcd\x55\x28\x2e\x29\x2a\x4d\x2e\x51\xa8\xe6\xe2\xf4\x74\x51\x28\xcd\xcc\x2b\xe1\xaa\xe5\x02\x04\x00\x00\xff\xff\x5e\x88\x02\x25\x31\x00\x00\x00")

func storeDataitemGoBytes() ([]byte, error) {
	return bindataRead(
		_storeDataitemGo,
		"store/dataitem.go",
	)
}

func storeDataitemGo() (*asset, error) {
	bytes, err := storeDataitemGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "store/dataitem.go", size: 49, mode: os.FileMode(420), modTime: time.Unix(1553210624, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storeStoreGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x4d\x6f\xe3\x36\x10\x3d\x93\xbf\x62\x36\x87\x40\x4a\xbc\xd2\x06\x58\xf4\x60\xc0\x87\x22\x69\x8a\x05\xda\x2e\xd0\x0f\xf4\x60\x18\x05\x2d\x8d\x6d\xda\x12\x29\x50\xa3\xdd\x4d\x0d\xff\xf7\x62\x48\xea\xc3\x49\xbc\xdd\xbd\x04\x0e\x87\xf3\xe6\xbd\x37\xc3\x51\x9e\xc3\x16\x0d\x3a\x45\x58\xc2\x67\x4d\x3b\xd8\x11\x35\xed\x3c\xcf\xb7\x9a\x76\xdd\x3a\x2b\x6c\x9d\xd7\x07\xfb\xef\xae\x3b\xec\xf2\x7d\x4b\xd6\xa1\xcc\x73\x08\xbf\xe0\x41\x91\xfa\x40\x58\x83\xff\x37\x27\xe5\xb6\x48\xd9\xd6\x4a\xd9\xa8\xe2\xa0\xb6\x18\x02\x52\xea\xba\xb1\x8e\x20\x91\xe2\x0a\x4d\x61\x4b\x6d\xb6\xf9\xbe\xb5\xe6\x4a\x8a\x2b\x6d\x73\x6d\x3b\xd2\xd5\x95\x4c\xa5\xdc\x74\xa6\x80\xdf\xf0\xf3\xbd\xad\x2a\x2c\x48\x5b\x93\x34\x8a\x76\xd0\x92\xd3\x66\x9b\x42\x72\x33\x46\x66\x80\xce\x59\x97\xc2\x51\x8a\x52\x91\x82\xf9\x02\xc6\xe8\x91\xf3\xe6\xc0\x7f\x4f\x52\xa0\x73\x1c\xe6\x6b\x59\x65\x55\xf9\xe8\x6c\xfd\xa8\x2b\x4c\x52\x29\x85\x43\xea\x9c\x81\x6b\x8e\x7a\x50\x79\x92\x92\x9e\x1a\x9c\xc0\x31\x83\xae\x20\xae\xe5\x09\x45\x46\x52\x34\x5d\x55\x01\xd4\xaa\x59\x76\xda\xd0\xea\xa6\x37\x45\x0a\xeb\x4a\x74\xb0\x5c\x0d\x27\x52\x14\xb6\x33\x84\x0e\xf8\xea\x50\xa4\x6b\xd1\x3d\x5a\x57\x2b\x9a\x14\xf9\xe8\x93\x61\x9a\x2e\xee\x9f\x65\x7b\xb3\x92\x02\x26\x9e\xa4\xf0\x33\x52\xa2\x4b\x7f\x27\x85\x81\x0d\x43\x6a\xc2\x7a\x06\xf6\xc0\x46\x14\x19\xf3\x5e\xea\x72\x25\x85\xde\xc0\x1b\x7b\xe0\x1b\x83\x13\x7d\xda\xf1\x24\xc5\x69\x74\x88\x11\x2e\x16\xfe\xe9\x8b\x6e\xa9\x1d\x6b\xaf\xad\xad\x18\xf4\x9f\xd7\x6a\x46\x40\x7b\xf8\x9a\x8e\x1f\xab\x2a\x49\x27\x0e\x30\x5a\x4c\x2c\x32\xef\xee\xe0\x61\x7f\xe5\x17\x5b\x28\xb2\x0e\x18\x32\x19\xd4\x07\x32\x17\x0a\x3d\x6a\xd7\x52\x52\xd9\xe2\x39\xc8\x33\xfb\x36\xd6\x81\x66\x1d\x4e\x99\x2d\xf6\x0c\xbc\x6d\x7a\x03\x95\x2d\x92\xeb\x78\xb6\xd4\x2b\x3f\x94\xa3\x9f\x63\x40\x0a\x71\xee\xe9\xd4\xeb\x4b\x5e\xfc\xa1\x3e\x61\x62\xd7\x7b\x98\x28\xf2\xb3\xef\xdb\x6a\x4a\xfc\xc2\xbc\xde\xde\xf9\x5e\xda\xf5\x3e\xfb\xf0\x00\x8b\x05\xbc\xf3\x24\x8a\x2c\x4e\xdd\xed\xad\x14\xa2\x8f\xc2\x70\x2c\xc5\x09\xb0\x6a\x31\x28\xf1\x60\x1c\xf5\xbf\x3e\x6e\x92\x90\x90\x06\xce\x7a\x03\xf1\x06\x97\x8b\xf0\xc1\x87\x05\xa8\xa6\x41\x53\x26\xf1\x60\x06\x37\x76\xbd\x4f\xcf\xd0\x07\x1b\x18\x63\x05\x0b\x7f\x25\x20\xc7\xe1\x08\xd5\x38\xc4\x91\x49\xbb\x5b\xf5\x09\xff\xb4\xf1\xc9\x5e\xf2\xe9\x01\x2b\x24\x1c\x47\xf0\xa5\x47\xa3\x2e\x5d\xa6\xaf\xea\x89\x15\x8d\xae\x7a\x62\xaf\xeb\x5b\xce\x83\x8a\x19\x9c\xa9\xba\xbd\x9b\xaf\xb2\x2c\x4b\xa5\x28\x03\x99\x20\x6c\x06\x5c\xef\x7b\xf5\xf4\x5c\x0f\xf8\x14\x15\x69\x43\xff\x3f\x8b\xe3\xb4\xc5\x41\xe0\xfc\xe9\x3c\xea\x17\x53\xf8\xf6\xee\x22\x89\xf3\x75\x39\x7a\x1a\x84\xc1\x02\x6a\x75\xc0\xe4\xe5\x02\x64\xb9\xeb\x27\xc2\xd6\xaf\x54\x26\x1b\xf6\x7c\xf6\x3b\xaa\xd2\x83\x15\x19\xaf\xd2\xd0\x06\xbe\xf2\x66\xc1\xb6\x4f\xbb\xc0\xbb\xd8\xd3\x24\xac\x1b\x86\x18\x77\xe5\x31\x6e\xf5\x05\xf0\x97\x24\xfb\xcb\xd4\xca\xb5\x3b\x55\x25\xb1\xe6\x35\xa7\x7c\x13\xf6\xd8\x61\x4e\xc9\xfc\xe6\x95\xe3\xb3\xe9\xcf\xef\xfb\xe7\xf2\x15\xf3\xe3\x14\x9f\x35\x80\x87\xf9\x6c\x01\x4c\x6c\xe7\x29\xbb\xe4\xfb\x74\x42\x46\xd7\xcf\x1d\xf5\xd2\x7f\x8d\xc2\x27\xde\x48\x11\x3e\x20\x73\x80\x9e\xdf\x4c\x8a\xfe\x03\x32\x1f\x1f\xff\x4c\x8a\x53\x2a\xbf\xc5\xa5\x7e\x74\x42\x0f\xff\x76\x9a\x70\xd2\xc4\x19\x44\x62\xef\x7e\x78\xff\x9e\x07\xfa\xbf\x00\x00\x00\xff\xff\xf8\x9b\x57\x66\x5b\x08\x00\x00")

func storeStoreGoBytes() ([]byte, error) {
	return bindataRead(
		_storeStoreGo,
		"store/store.go",
	)
}

func storeStoreGo() (*asset, error) {
	bytes, err := storeStoreGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "store/store.go", size: 2139, mode: os.FileMode(420), modTime: time.Unix(1553210852, 0)}
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
	"store/dataitem.go": storeDataitemGo,
	"store/store.go":    storeStoreGo,
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
	"store": &bintree{nil, map[string]*bintree{
		"dataitem.go": &bintree{storeDataitemGo, map[string]*bintree{}},
		"store.go":    &bintree{storeStoreGo, map[string]*bintree{}},
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

// Code generated by go-bindata.
// sources:
// static/config.json
// static/css/docs.css
// static/template.html
// DO NOT EDIT!

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

var _configJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x4f\x6a\xf3\x30\x10\xc5\xd7\xf6\x29\x06\x61\xc8\x26\xf8\x00\xd9\x85\xd8\xf0\x05\xbe\xfc\x21\x55\x0f\xa0\xca\x93\x58\xad\x2d\x09\x69\x4c\x0a\x21\x77\x2f\x63\xd5\x34\x49\x03\xdd\x49\xf3\xe6\x3d\xff\x9e\x75\xc9\x33\xe1\x3b\xa5\xb1\x75\x5d\x83\x41\x2c\xe0\x92\x67\x99\xd0\xce\x12\x5a\x12\x0b\x10\x45\xb1\xda\x6d\x65\xbd\x95\x45\x21\xe6\xac\x91\xa1\x0e\x93\x22\xd7\xf2\x7f\x3d\xcd\xd5\x40\xad\x0b\x49\x58\xbe\xca\x7f\xbb\xc3\xa4\x34\x18\x75\x30\x9e\x8c\xb3\x49\xae\xea\x97\xd5\x61\xbd\x97\xeb\xdd\xb6\x28\x44\x9e\x5d\xe7\x39\x6f\x1d\xd5\xd0\xd1\x84\x60\xac\x1f\x46\x80\x43\xbd\xac\x36\x75\xd9\x37\x29\xcc\x2b\x6a\x79\x1c\x30\xf2\xf6\x3d\x52\xe5\x74\x7c\xa4\xd9\x07\xf7\x8e\x9a\x60\x99\x26\xcf\x88\x2a\xa7\x87\x1e\x2d\x29\x1e\x80\x3b\x82\xb2\xa0\xce\x18\x5d\x8f\xe0\x93\xbd\x9c\x30\x5b\xec\xfc\x2f\x46\xd9\x22\x8c\x17\xd8\xa8\xf0\xd1\xb8\xb3\x85\xa3\xe9\x10\xc8\xc1\xdb\x18\xa1\x31\x46\x6c\xca\xfb\x0a\xec\xe2\x33\xaf\xe1\xa7\x77\x81\x80\x5a\x84\xd4\xcc\xd8\xd3\x98\x11\xcb\x87\x8e\x92\x0f\x0c\xc9\xbb\xcd\x1d\xb9\x57\x27\x2c\x1f\xfb\xa7\xde\x93\xe1\xbb\x0e\x7f\x72\x7c\x77\x30\x76\x9c\xcf\x92\x61\x06\x3d\x92\x02\x52\xa7\xf2\xf9\x9f\xfa\xb9\xfe\x99\x78\x63\xbd\x8d\xcd\xb3\x6b\x7e\xfd\x0a\x00\x00\xff\xff\x35\x9b\x3e\xfd\x78\x02\x00\x00")

func configJsonBytes() ([]byte, error) {
	return bindataRead(
		_configJson,
		"config.json",
	)
}

func configJson() (*asset, error) {
	bytes, err := configJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.json", size: 632, mode: os.FileMode(420), modTime: time.Unix(1504554452, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cssDocsCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcb\xcd\xaa\x83\x40\x0c\xc5\xf1\xfd\x3c\x45\xe0\xae\x47\xbc\x20\xad\xc4\xa7\x19\x6c\x46\x03\x63\x32\xc4\x94\xda\x16\xdf\xbd\xf4\x63\x21\x67\xf7\xe3\xfc\x9b\x51\xc5\x13\x0b\x19\x3c\x03\x00\xc0\x92\x6c\x62\x89\xae\x15\xa1\xeb\xeb\x36\x84\x3d\x84\xe6\x42\x39\x5d\x8b\xc3\xe1\x3e\xff\x63\x66\x5b\x3d\x6a\x8e\x7e\xaf\xf4\xeb\x9d\x36\x8f\xa9\xf0\x24\x08\x23\x89\x93\x0d\x1f\xcf\x2a\x1e\x6f\xc4\xd3\xec\x08\xe7\xb6\x3d\xe8\xca\x0f\x42\xe8\x8d\x96\x2f\x8e\x5a\xd4\x10\xfe\xba\xd3\x7b\x43\xd8\x5f\x01\x00\x00\xff\xff\xba\x03\x6d\x8d\xa5\x00\x00\x00")

func cssDocsCssBytes() ([]byte, error) {
	return bindataRead(
		_cssDocsCss,
		"css/docs.css",
	)
}

func cssDocsCss() (*asset, error) {
	bytes, err := cssDocsCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "css/docs.css", size: 165, mode: os.FileMode(420), modTime: time.Unix(1504555564, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4d\x6f\xdb\x30\x0c\xbd\xe7\x57\x68\x82\x0e\x1b\xd0\x48\x69\xb1\xc3\x3e\xec\x00\x43\x9a\xa1\x01\x8a\xb8\xe8\xdc\xc3\x8e\x8a\xc4\x44\x72\x65\x29\x10\xe9\xb4\xfd\xf7\x83\xed\xa6\x4d\xd1\xd3\xb0\xf9\x62\x3c\x92\xef\x3d\x4a\xa4\x8a\x0f\x97\xd5\xa2\xfe\x7d\xb3\x64\x8e\xda\x30\x9f\x14\xfd\x8f\x05\x1d\x77\x25\x87\xc8\xe7\x93\x49\xe1\x40\xdb\xf9\x84\x31\xc6\x8a\x16\x48\x33\xe3\x74\x46\xa0\x92\xdf\xd5\x3f\xa7\x5f\xf8\x73\x8a\x3c\x05\x98\x0b\x51\xaf\xea\xeb\xa5\x10\x85\x1a\x03\x27\xbc\xa8\x5b\x28\xb9\x05\x34\xd9\xef\xc9\xa7\xc8\x99\x49\x91\x20\x52\xc9\x85\xb8\x5c\xfe\x5a\xdc\xae\x6e\xea\x55\xb5\x16\x82\xbf\xe7\xe9\x8e\x5c\xca\x6f\x28\x3f\xee\xea\xab\xea\x76\xa8\x7e\x57\x7e\xf0\xf0\xb0\x4f\x99\x4e\x08\x0f\xde\x92\x2b\x2d\x1c\xbc\x81\xe9\x00\xce\x98\x8f\x9e\xbc\x0e\x53\x34\x3a\x40\x79\x2e\x67\x67\xac\xd5\x8f\xbe\xed\xda\xd3\x50\x87\x90\x07\xac\x37\x01\xca\x98\x8e\xed\x05\x1f\xef\x59\x86\x50\x72\xa4\xa7\x00\xe8\x00\x88\x33\x97\x61\x5b\x72\x47\xb4\xc7\x6f\x4a\x19\x1b\x1b\x94\x26\xa4\xce\x6e\x83\xce\x20\x4d\x6a\x95\x6e\xf4\xa3\x0a\x7e\x83\x0a\xef\x21\x00\xa5\xa8\x2e\xe4\x4c\x7e\x7e\x81\xb2\xf5\x51\x1a\xc4\xff\xe8\xe4\xfc\xce\x05\xbf\x73\x24\x1b\x54\x5f\xe5\xf9\x85\x9c\xa9\x51\x4c\x59\xd8\xea\x2e\xd0\xdf\x98\x1a\x44\x65\x93\xc1\xd3\xfa\x71\xb0\x0c\xb3\xf9\xa7\xa6\x5e\x63\x7d\x3f\x0d\xf2\x79\xa1\x46\xe9\x37\x3e\x23\xe8\x3f\x17\x1a\x94\xfd\x24\xaf\x8e\x4c\x1f\x77\x55\xbc\x4e\xda\x7e\xfc\xf4\x7d\xe4\xbc\x28\x14\x6a\xdc\xe7\x49\xb1\x49\xf6\x89\x99\xa0\x11\xfb\xad\x1c\x6e\xe0\x78\x12\xeb\x0f\xc7\x4c\xbf\x3e\xda\x47\xc8\xfc\xd5\x50\x88\x45\xb5\xae\x97\xeb\x5a\x88\x67\x75\xeb\x0f\xbd\x74\x2f\xd9\x4b\xab\xe1\x31\xfd\x09\x00\x00\xff\xff\x55\xd4\xee\xdb\x5c\x03\x00\x00")

func templateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateHtml,
		"template.html",
	)
}

func templateHtml() (*asset, error) {
	bytes, err := templateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.html", size: 860, mode: os.FileMode(420), modTime: time.Unix(1504555936, 0)}
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
	"config.json": configJson,
	"css/docs.css": cssDocsCss,
	"template.html": templateHtml,
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
	"config.json": &bintree{configJson, map[string]*bintree{}},
	"css": &bintree{nil, map[string]*bintree{
		"docs.css": &bintree{cssDocsCss, map[string]*bintree{}},
	}},
	"template.html": &bintree{templateHtml, map[string]*bintree{}},
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


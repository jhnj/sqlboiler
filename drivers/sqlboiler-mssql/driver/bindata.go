// Code generated by go-bindata.
// sources:
// override/templates_test/singleton/mssql_main.tpl
// DO NOT EDIT!

package driver

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

var _templates_testSingletonMssql_mainTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x4d\x6f\xe3\x36\x13\x3e\x4b\xbf\x62\x5e\x01\xbb\x2b\x2d\xf4\x32\xdd\x1e\x7a\xf0\xc2\x07\xc7\x56\x76\x83\xc6\x71\x62\xb9\x48\x81\x20\xc8\xd2\xe2\x28\x21\x2a\x51\x36\x49\xc5\xeb\xa6\xf9\xef\x05\x49\xc9\x96\x13\x3b\xcd\xf6\xd2\x5c\x1c\xce\x3c\xf3\xf5\x8c\x66\x48\xbd\x5e\x20\x94\x4a\x2d\x8b\x19\x2a\x8d\x12\x94\x96\x75\xa6\xe1\xd1\xf7\xd8\x7c\x58\x09\x01\xe6\xef\xa3\x5a\x16\x64\x74\x6c\x64\xe7\xb4\x44\x2b\x53\x5a\x72\x71\xe7\x7b\xf7\x95\xd2\x00\x3b\xa2\x5a\xa1\x7c\x26\x5a\x50\xa5\x9e\x89\x94\x2a\xca\x8a\xe1\x0e\xaa\x92\xad\x2f\x2e\xb4\xef\x69\x54\x7a\x74\x6c\x43\x36\x90\x27\xdf\xcf\x6b\x91\x01\x17\x5c\x87\x91\x4b\x73\x4c\xb9\x80\x3e\xbc\xef\x94\xf1\xf8\xb4\x41\x86\x25\x7c\xec\x68\x22\x50\xa8\xeb\x45\x18\x01\x4a\x59\x49\xe3\xe1\x81\x4a\x73\x70\x02\xdf\x2b\x49\x53\x65\x1f\x1e\xf8\x02\x25\xf9\x82\x3a\xb5\xd1\xc3\xc0\x3a\x22\x6c\x2e\x68\x89\x41\x64\xb0\xb6\xfc\x83\x48\xa3\x75\x38\xcb\xc9\x41\x9c\xd1\x3a\x9c\x25\xea\x20\xce\x68\x1b\x9c\xa1\xaa\x83\x3b\x15\x7a\x03\xaa\x64\x13\xb4\xa5\xf8\xa0\xbf\x06\x60\xd0\x47\x47\x30\x94\x48\x35\x02\x05\x49\x05\xab\x4a\xfe\x27\x32\x60\x73\x30\xc5\x12\xe3\xae\xd3\x8e\xfe\x16\x43\x52\x4d\xe7\x05\x3a\x45\xd8\xb2\x17\xf9\xbe\xc7\x73\xcb\x6b\x1f\x4a\xc2\x64\xb5\x98\x59\xf3\x30\xfa\x6c\xa5\xff\xeb\x83\xe0\x85\xe1\xdf\x93\xa8\x6b\x29\x8c\xd4\xf7\x9e\x76\xcc\x32\x9b\xd1\x1b\x0d\x7d\xcf\xc1\x87\x25\x83\x5e\x1f\xf0\x3b\x66\x64\x58\x95\x25\x15\x2c\x0c\xd4\xb2\xc8\x4a\x16\xc4\x10\xfc\x3f\x0d\x62\x70\x8d\x33\xa7\xdf\xec\xc9\x34\xc0\x9c\x2e\xec\xc9\xd0\x6c\x4e\xcc\x9e\xb6\x65\x9b\xa2\xf2\xd8\x66\xd1\xeb\x43\xa5\xc8\x64\x81\x22\x0c\x6c\xfd\xea\x56\x65\xf7\x58\x52\xa2\x96\x85\xe1\xb3\xa9\x62\x6f\xb6\x95\x54\xe4\x4a\xd2\x45\x88\xd2\x44\xcd\x29\x2f\x90\x81\xae\xa0\x5a\xa0\x80\x17\xee\x20\xe7\x85\xed\x91\xa9\x91\x61\x8e\x12\x72\x32\x2c\x2a\x85\x61\xd4\xad\x9a\xa4\x9a\xd9\x51\x10\xb8\x3a\xf9\x15\xd7\x23\x54\x5a\x56\x6b\x94\xa1\xbc\xfb\x3e\x4e\xd3\xcb\xb3\x3f\x70\x1d\x43\xbe\xd3\x9b\xae\x35\x95\xfa\x55\x96\x0f\xe6\xad\x8c\x29\x38\x8e\x21\x73\x9c\xb7\x09\xef\x89\x74\x45\xf9\xbe\x40\x79\xa9\xc9\x85\xe4\x42\x17\xc2\x44\x88\xde\x16\x7b\x45\xb9\x86\xbc\x92\x07\xc2\x37\x1e\x04\x2f\x5e\x59\x09\xaa\x18\x57\x0c\x43\x3b\x2c\x6e\xd3\x44\xcd\xaf\x49\x4c\xad\xb8\xce\xee\xc1\x6a\x1f\x7d\x2f\xa3\x0a\x21\xd0\xb2\xc6\xa0\xb7\x4d\xd1\x09\x5a\x6d\x4e\x0b\xb5\xab\x76\x12\xdb\x3f\x5a\x17\xba\xab\x62\x5c\x99\x96\x07\x26\xe1\x83\x39\xee\x8e\xc2\x76\x7b\x99\xef\xa3\xd7\x07\xc3\x5d\xba\x30\xe4\xe5\xe1\x37\xdf\x1b\x4e\x93\xc1\x2c\x81\xd1\x60\x36\x38\x1e\xa4\x09\xbc\x53\x9f\x7d\xef\xcb\xc4\xf7\x06\x67\xb3\x64\xda\x91\x5f\x7f\xba\x51\xbe\x97\x26\x33\x98\x26\x83\xd1\xed\x70\x32\x1e\x9f\xce\x66\xc9\xe8\x36\x3d\x1f\x5c\xa4\x5f\x27\x33\x98\x9c\x5b\xd3\x6f\xcf\x47\xa1\x4d\xbf\x24\xb2\x16\xc3\x92\x85\x6a\x59\xc4\xf0\xe3\x83\x16\x1d\xae\xb9\xbb\x35\xb6\x15\x1f\x1d\x41\xca\x45\x86\x30\x4e\x21\xbd\x3c\x83\x9f\x7f\xfa\xf4\x0b\x70\x0d\x19\x15\x30\x47\x60\x95\x40\x58\x71\x7d\x6f\x91\xa3\xe9\xe4\x62\x5b\xee\x35\x9c\x9e\x40\xf2\xfb\x69\x3a\x4b\xe1\x06\x1e\x81\x51\x4d\xe7\x54\xe1\xad\xd9\x71\xf0\xd7\xf6\xac\x04\x5d\xa8\xfb\x4a\x3b\xc5\x13\x5c\x43\x4c\x08\x11\x70\x03\xd7\x9f\x6f\x0e\x91\xbe\xf1\x1d\xa6\xc9\x59\x32\x9c\xd9\xcd\x09\x27\xd3\xc9\x18\xd4\x5a\x91\xd6\xb9\x02\xdf\xf3\xae\xbe\x26\xd3\xc4\x01\xfa\xf0\xe1\x9d\xfa\x60\xbe\xf6\xdd\x64\xdf\xa9\x3d\xbc\xff\x07\x5d\xd0\x48\x25\xab\x56\xa2\xdb\x03\x9e\x9b\xa5\xee\x5e\x08\x9d\x09\x6e\x65\x9b\xed\xd4\xdd\x00\xbd\x1f\xbc\x08\xde\x36\xbb\x6d\xd9\x66\xf5\xc5\xed\xf0\x37\xc3\x1b\x03\x95\x77\x0a\x08\x21\xed\x50\x6f\x0a\xc8\xf6\x5c\x12\x8d\xb1\xb3\x22\x84\x44\x16\xb6\xd9\xaa\xce\x87\x22\xe7\xb8\x9a\x22\x65\x28\x5d\x50\xb3\x4d\x95\x66\x55\xad\x8d\xc3\xf7\xf3\xb5\x46\x45\x8e\xeb\x3c\xb7\xcf\x10\xa3\x6a\x8a\x7f\xa1\x6a\x9c\x1b\x4b\xe3\xdd\xfc\xb3\x11\xba\x95\xe9\x8c\xbb\x0c\x1a\xf5\xb4\x16\xff\xb0\x3d\xdb\xf5\x28\x6b\x21\xb8\xb8\xeb\x05\x1b\x66\x5c\x71\xd1\x33\xbc\x0b\x4e\x9a\xc7\x41\xb4\x47\x8d\x52\xee\xa8\xff\x4d\xab\xb2\x4a\x98\x8f\x28\x6c\x9e\x93\xb1\xeb\x46\xf4\xca\xf7\xb4\xf9\xb8\x9d\x2a\xb6\xfe\x6d\xbc\x3d\x8f\x36\x87\x68\x88\x5b\x16\xcd\xe5\x6c\x33\x08\x62\x60\x92\x3f\xa0\x24\xf6\x1a\x3c\xae\x79\xc1\x2e\x6b\x94\xeb\xa6\xa4\x76\x22\xda\xab\xff\xf9\xc4\xb9\xe9\x71\xaf\x2e\xf3\xdb\xbc\x9b\xa2\xd7\xee\x79\xc1\x8b\xf8\x05\x3f\xbb\x95\x3c\xf9\x7f\x07\x00\x00\xff\xff\xbe\x1b\xaf\x1a\x80\x0b\x00\x00")

func templates_testSingletonMssql_mainTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMssql_mainTpl,
		"templates_test/singleton/mssql_main.tpl",
	)
}

func templates_testSingletonMssql_mainTpl() (*asset, error) {
	bytes, err := templates_testSingletonMssql_mainTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mssql_main.tpl", size: 2944, mode: os.FileMode(420), modTime: time.Unix(1526677173, 0)}
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
	"templates_test/singleton/mssql_main.tpl": templates_testSingletonMssql_mainTpl,
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
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mssql_main.tpl": &bintree{templates_testSingletonMssql_mainTpl, map[string]*bintree{}},
		}},
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

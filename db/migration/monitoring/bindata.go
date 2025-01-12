// Code generated for package monitoring by go-bindata DO NOT EDIT. (@generated)
// sources:
// 001_init.down.sql
// 001_init.up.sql
// 002_indexes.down.sql
// 002_indexes.up.sql
// 003_add_net.down.sql
// 003_add_net.up.sql
package monitoring

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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\x4d\x4d\x2c\x2e\x2d\x4a\xcd\x4d\xcd\x2b\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\x3c\x83\x91\x54\x19\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 25, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\xcd\x4a\x03\x31\x14\x05\xe0\x7d\x9e\xe2\x30\xab\x16\x9a\x27\x70\x15\x35\xc2\xe0\xb4\xca\xf4\x0a\xed\x6a\x18\xe3\x55\x02\xcd\x24\xe4\x67\xe1\xdb\x4b\xa9\x2d\x23\x1d\xa7\x77\x95\x45\xbe\x73\x93\x23\x25\xe4\xcc\x08\x29\x41\xfd\xfb\x81\x91\x72\x2c\x26\x97\xc8\xf8\xf4\x11\x8e\xfb\x54\x22\x3b\x1e\x72\x12\xb7\x32\x1e\x5a\xad\x48\x83\xd4\x7d\xa3\x51\x3f\x61\xf3\x42\xd0\xbb\x7a\x4b\x5b\x54\xe3\xa0\x4a\x2c\x04\x00\x54\xe6\x60\x79\xc8\x9d\xfd\xa8\x30\x1e\xd2\x3b\x3a\x9f\x8f\x19\x9b\xb7\xa6\x59\x9d\x44\xb6\x8e\x53\xee\x5d\xf8\x2b\x1e\x15\x69\xaa\xd7\x7a\x2c\xf0\x4b\x4c\x28\x5d\x49\xfd\x17\x77\x81\xa3\xe1\x21\x9f\x68\xab\x55\xf3\xcf\x12\xc7\xce\xc7\xef\x2b\x34\x23\xac\x9f\x5a\x31\x27\x42\xf4\x86\x53\xe2\x74\xfd\xf5\xf3\x2b\x7c\x19\x72\xf0\xf6\x58\xd8\xe4\x8d\xd7\xb6\x5e\xab\x76\x8f\x67\xbd\xc7\xe2\x52\xe5\x0a\x97\x8e\x96\x62\x79\x27\x7e\x02\x00\x00\xff\xff\xce\xe5\xad\x53\xf9\x01\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 505, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_indexesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xc8\x4d\x4d\x2c\x2e\x2d\x4a\xcd\x4d\xcd\x2b\x29\x8e\x2f\xc9\xcc\x4d\x2d\x2e\x49\xcc\x2d\xb0\xe6\xc2\xa5\x24\x39\x27\x33\x35\xaf\x24\x3e\x33\xc5\x9a\x0b\x10\x00\x00\xff\xff\x74\x7a\xdb\x2d\x46\x00\x00\x00")

func _002_indexesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_indexesDownSql,
		"002_indexes.down.sql",
	)
}

func _002_indexesDownSql() (*asset, error) {
	bytes, err := _002_indexesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_indexes.down.sql", size: 70, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_indexesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xc5\x03\xb8\x74\x75\x15\x3c\xf3\x52\x52\x2b\x52\x8b\x15\xd2\xf2\x8b\x14\x4a\x12\x93\x72\x52\x15\x72\x53\x13\x8b\x4b\x8b\x52\x73\x53\xf3\x4a\x8a\xb9\x08\x99\xe0\x1c\xe4\xea\x18\xe2\xaa\xe0\xe9\xe7\xe2\x1a\xa1\xa0\x84\xac\x35\xbe\x24\x33\x37\xb5\xb8\x24\x31\xb7\x40\x49\xc1\xdf\x4f\x21\x01\x59\x2e\x41\x41\x83\x4b\x41\x41\x41\x41\x09\x49\x8d\x63\xb0\x33\x97\xa6\x35\x17\x3e\x13\x93\x73\x32\x53\xf3\x4a\xe2\x33\x53\xf0\x98\x88\xa4\x06\x6a\x22\x20\x00\x00\xff\xff\x1f\x9c\x57\xf4\x05\x01\x00\x00")

func _002_indexesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_indexesUpSql,
		"002_indexes.up.sql",
	)
}

func _002_indexesUpSql() (*asset, error) {
	bytes, err := _002_indexesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_indexes.up.sql", size: 261, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __003_add_netDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xc5\x03\xb8\x74\x75\x15\x52\x8a\xf2\x0b\x14\xf2\x52\x4b\x14\x92\xf3\x73\x4a\x73\xf3\x8a\xb9\x08\xe9\x71\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x50\xca\x4d\x4d\x2c\x2e\x2d\x4a\xcd\x4d\xcd\x2b\x29\x56\x52\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x50\xca\x4b\x2d\x89\xcf\x49\xcc\x8b\xcf\xcc\x53\xb2\x26\x59\x53\x7e\x69\x09\x89\xba\xca\xc9\xb1\xaa\x1c\x6e\x15\x20\x00\x00\xff\xff\x2a\x0e\x2f\x32\x2a\x01\x00\x00")

func _003_add_netDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__003_add_netDownSql,
		"003_add_net.down.sql",
	)
}

func _003_add_netDownSql() (*asset, error) {
	bytes, err := _003_add_netDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "003_add_net.down.sql", size: 298, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __003_add_netUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xc5\x03\xb8\x74\x75\x15\x12\x53\x52\x14\xf2\x52\x4b\x14\x92\xf3\x73\x4a\x73\xf3\x8a\xb9\x08\x69\x71\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x50\xca\x4d\x4d\x2c\x2e\x2d\x4a\xcd\x4d\xcd\x2b\x29\x56\x52\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x50\xca\x4b\x2d\x89\xcf\x49\xcc\x8b\xcf\xcc\x53\x52\xf0\xf4\x0b\x71\x75\x77\x0d\xb2\x26\x55\x6f\x7e\x69\x09\x79\x9a\xcb\x29\xb0\xb8\x1c\xc3\x62\x40\x00\x00\x00\xff\xff\x74\x55\x1b\x06\x45\x01\x00\x00")

func _003_add_netUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__003_add_netUpSql,
		"003_add_net.up.sql",
	)
}

func _003_add_netUpSql() (*asset, error) {
	bytes, err := _003_add_netUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "003_add_net.up.sql", size: 325, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
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
	"001_init.down.sql":    _001_initDownSql,
	"001_init.up.sql":      _001_initUpSql,
	"002_indexes.down.sql": _002_indexesDownSql,
	"002_indexes.up.sql":   _002_indexesUpSql,
	"003_add_net.down.sql": _003_add_netDownSql,
	"003_add_net.up.sql":   _003_add_netUpSql,
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
	"001_init.down.sql":    &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":      &bintree{_001_initUpSql, map[string]*bintree{}},
	"002_indexes.down.sql": &bintree{_002_indexesDownSql, map[string]*bintree{}},
	"002_indexes.up.sql":   &bintree{_002_indexesUpSql, map[string]*bintree{}},
	"003_add_net.down.sql": &bintree{_003_add_netDownSql, map[string]*bintree{}},
	"003_add_net.up.sql":   &bintree{_003_add_netUpSql, map[string]*bintree{}},
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

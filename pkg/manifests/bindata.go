// Code generated by go-bindata.
// sources:
// assets/tuned/01-service-account.yaml
// assets/tuned/02-cluster-role.yaml
// assets/tuned/03-cluster-role-binding.yaml
// assets/tuned/04-cm-tuned-profiles.yaml
// assets/tuned/05-cm-tuned-recommend.yaml
// assets/tuned/06-ds-tuned.yaml
// assets/tuned/default-cr-tuned.yaml
// DO NOT EDIT!

package manifests

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

var _assetsTuned01ServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc7\x31\x0e\x42\x31\x08\x06\xe0\xbd\xa7\xe0\x02\x1d\x5c\xd9\x3c\x83\x89\x3b\xa1\xbf\x4a\xf4\x41\x43\xe9\x3b\xbf\xcb\x1b\x3f\x99\xf6\x44\x2e\x0b\x67\x3a\x6f\xed\x6b\x3e\x98\x1e\xc8\xd3\x14\x77\xd5\xd8\x5e\xed\x40\xc9\x90\x12\x6e\x44\x2e\x07\x98\x6a\x3b\xc6\xa5\x35\x45\xc1\x14\x13\xbe\x3e\xf6\xaa\xae\xbf\xbd\x0a\xd9\x3d\x06\x7a\x6d\x37\x7f\xf7\x98\x48\xa9\xc8\xf6\x0f\x00\x00\xff\xff\x95\x01\xe5\x42\x70\x00\x00\x00")

func assetsTuned01ServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned01ServiceAccountYaml,
		"assets/tuned/01-service-account.yaml",
	)
}

func assetsTuned01ServiceAccountYaml() (*asset, error) {
	bytes, err := assetsTuned01ServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/01-service-account.yaml", size: 112, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned02ClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x3f\x4f\xc4\x30\x0c\xc5\xf7\x7c\x8a\x28\xf3\x5d\x11\x1b\xca\xca\xc0\xc6\xc0\xc0\x82\x3a\xf8\x12\xd3\x5a\xd7\xda\x91\xed\x94\x3f\x9f\x1e\x5d\x4f\x0c\x1c\x93\x25\xff\xde\x7b\x3f\x68\xf4\x8a\x6a\x24\x9c\xa3\x9e\xa0\x0c\xd0\x7d\x16\xa5\x6f\x70\x12\x1e\xce\x0f\x36\x90\xdc\x6d\xf7\xe1\x4c\x5c\x73\x7c\x5c\xba\x39\xea\x8b\x2c\x18\x56\x74\xa8\xe0\x90\x43\x8c\x0c\x2b\xe6\x58\xae\xf4\xc8\x52\xf1\xe8\x9d\x89\xa7\xec\x9d\xb1\x06\xed\x0b\x5a\x0e\xc7\x08\x8d\x9e\x54\x7a\xb3\x1c\xdf\x52\x1a\x43\x8c\x8a\x26\x5d\x0b\xee\x9f\x4b\xd3\xd2\x21\x35\xa9\xb6\xd3\x0d\xf5\xb4\x93\x09\x3d\x1d\xd2\x42\x76\x39\x1f\xe0\x65\x4e\xe3\xed\x9e\x61\xe9\x4a\xfe\x35\x48\x43\xb6\x99\xde\x7d\x20\xf9\x2f\xf9\x8d\x15\x61\xc7\x4f\x2f\xc2\xe6\x0a\xc4\x7e\xa3\xec\x86\x7f\xca\xcf\xb0\x5e\x07\x9a\xd2\x46\x0b\x4e\x58\xd3\x18\x7e\x02\x00\x00\xff\xff\x24\x0d\xec\x93\x42\x01\x00\x00")

func assetsTuned02ClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned02ClusterRoleYaml,
		"assets/tuned/02-cluster-role.yaml",
	)
}

func assetsTuned02ClusterRoleYaml() (*asset, error) {
	bytes, err := assetsTuned02ClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/02-cluster-role.yaml", size: 322, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned03ClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x45\x5e\x20\x87\xd8\x90\x37\x60\x67\x28\x12\xbb\x9b\xb8\xad\xe9\x9d\x7d\xb2\x9d\x4a\xf0\xf4\xe8\x94\x8a\x05\x86\x1b\x2d\xfd\xdf\xf7\x19\x57\xfe\x20\x73\x56\x81\x6c\x47\xac\x13\xf6\xb8\xa8\xf1\x37\x06\xab\x4c\xd7\x27\x9f\x58\x1f\x6e\x8f\xe9\xca\xd2\x20\xbf\xce\xdd\x83\xec\xa0\x33\xbd\xb0\x34\x96\x73\x5a\x28\xb0\x61\x20\xa4\x9c\x05\x17\x82\x5c\xc7\xa8\x88\x36\x2a\xd1\x85\xe5\x0c\xd1\x85\x5a\x32\x9d\xe9\x40\xa7\x6d\xfa\xc7\xb7\x03\xf7\x7e\xfc\xa4\x1a\x0e\xa9\xdc\xf9\x77\xb2\x1b\x57\x7a\xae\x55\xbb\xc4\xaf\x62\xcc\xc7\xe5\x2b\x56\x82\xac\x2b\x89\x5f\xf8\x14\xe5\x1f\x7f\xd1\x95\x0c\x43\x2d\x75\x27\x7b\xdb\xa8\xad\xe1\x5f\x1e\xb4\x80\x8f\x08\x8e\x08\xec\x33\xdd\x5f\xfe\x09\x00\x00\xff\xff\x51\x4a\xd8\x71\x60\x01\x00\x00")

func assetsTuned03ClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned03ClusterRoleBindingYaml,
		"assets/tuned/03-cluster-role-binding.yaml",
	)
}

func assetsTuned03ClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsTuned03ClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/03-cluster-role-binding.yaml", size: 352, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned04CmTunedProfilesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\xa1\x0e\xc2\x40\x0c\x06\x60\x7f\x4f\xd1\xcc\x57\x60\xcf\xa2\xb1\xf8\x66\xf7\x6f\x34\x6c\x6d\x73\xd7\xe3\xf9\x49\x08\x08\xec\x27\x3e\x09\xbd\xa3\x0f\x75\xab\xf4\xba\x94\xa7\x5a\xab\x74\x75\xdb\x74\xbf\x49\x94\x13\x29\x4d\x52\x6a\x21\x32\x39\x51\x29\xa7\xa1\x71\x74\xdf\xf4\xc0\xf8\xf2\x08\x59\x51\xc9\x03\x36\x1e\xba\x25\xaf\xc7\x1c\x89\xce\xe6\x0d\x9c\xd3\xd4\x76\xf6\x40\x97\xf4\x5e\x7e\xe1\x7f\xc5\x1f\xa6\x65\x29\xef\x00\x00\x00\xff\xff\x0a\xb4\xaa\xf4\x94\x00\x00\x00")

func assetsTuned04CmTunedProfilesYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned04CmTunedProfilesYaml,
		"assets/tuned/04-cm-tuned-profiles.yaml",
	)
}

func assetsTuned04CmTunedProfilesYaml() (*asset, error) {
	bytes, err := assetsTuned04CmTunedProfilesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/04-cm-tuned-profiles.yaml", size: 148, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned05CmTunedRecommendYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\xa1\xb2\x02\x31\x0c\x05\x50\xdf\xaf\xc8\xac\xaf\x78\x36\xf6\x69\x2c\x3e\xd3\xde\x5d\x32\xd0\x24\xd3\x66\xf9\x7e\x04\x30\xd8\x23\x8e\x84\x5e\x31\x97\xba\x31\x3d\xff\xca\x5d\xad\x33\xfd\xbb\xed\x7a\x5c\x24\xca\x40\x4a\x97\x14\x2e\x44\x26\x03\x4c\x79\x1a\x7a\x9d\x68\x3e\x06\xac\x7f\x7c\x85\x34\x30\x79\xc0\xd6\x4d\xf7\xac\xed\x71\xae\xc4\xac\xe6\x1d\x35\x4f\x53\x3b\xaa\x07\xa6\xa4\xcf\xf2\x1d\xdf\x97\xb7\xf8\x7d\x4c\xdb\x56\x5e\x01\x00\x00\xff\xff\x1b\xe2\xd5\x67\x95\x00\x00\x00")

func assetsTuned05CmTunedRecommendYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned05CmTunedRecommendYaml,
		"assets/tuned/05-cm-tuned-recommend.yaml",
	)
}

func assetsTuned05CmTunedRecommendYaml() (*asset, error) {
	bytes, err := assetsTuned05CmTunedRecommendYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/05-cm-tuned-recommend.yaml", size: 149, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTuned06DsTunedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x4d\x8f\x22\x37\x10\xbd\xf3\x2b\x4a\x9b\x5c\x9b\x66\x13\x25\x87\xbe\x8d\x80\x59\x8d\xb4\x30\x68\x98\x8d\x94\xd3\xc8\xb8\x0b\xb0\xf0\x47\xc7\x55\x66\x07\xad\xf2\xdf\x23\xd3\xdd\xd0\xa6\x81\x59\xa5\x2f\x33\x54\x3d\x3f\x97\x5d\xaf\xaa\x2c\x2a\xf5\x17\x7a\x52\xce\x16\x20\xaa\x8a\xf2\xfd\xe7\xc1\x4e\xd9\xb2\x80\x89\x40\xe3\xec\x12\x79\x60\x90\x45\x29\x58\x14\x03\x00\x2b\x0c\x16\xc0\xc1\x62\xd9\xfc\xa2\x4a\x48\x2c\xc0\x55\x68\x69\xab\xd6\x9c\x49\x1d\x88\xd1\x67\xd6\x95\x98\x71\xb0\xca\x6e\x32\x57\xa1\x17\xec\xfc\x00\x40\x8b\x15\x6a\x8a\x5c\xd0\x59\x24\xaa\xaa\xa5\xa5\x0a\x65\x74\x13\x6a\x94\xec\x7c\x0d\x35\x82\xe5\xf6\x6b\x67\xed\x8d\xd5\x00\x8c\xa6\xd2\x82\xb1\x59\xd7\x09\x3e\x7e\x3a\xa1\xb8\x49\x02\xd0\x86\x71\xfc\x1f\xfd\x5e\x49\x7c\x90\xd2\x05\xcb\xf3\xe4\x0e\xe2\x27\x02\x3b\x13\x5d\xcb\x04\xf8\xea\x76\x68\x0b\x60\x1f\xb0\x01\x4a\x67\x59\x28\x8b\xfe\x14\x41\x06\xd2\x19\x23\x6c\x79\x0e\x29\x83\x7c\x2f\x7c\xae\xd5\x2a\x3f\x6e\x92\xaf\x94\xcd\x7d\xb0\x27\x80\x47\x72\xc1\x4b\xec\x1c\x23\x1a\xff\x09\x48\x9c\xd8\x00\x64\x15\x0a\xf8\x3c\x32\x89\xd1\xa0\x71\xfe\x50\xc0\x6f\xa3\x99\x3a\x39\x94\x11\x1b\x2c\xe0\xd7\x1f\xe3\xaf\xdf\x96\xaf\xd3\x97\xb7\xf9\xf3\x64\xfa\xf6\xfa\x6d\x3e\x9d\xbc\x3d\xcd\x1e\xbe\x4c\xff\x4d\xa1\x8b\xa0\xf5\xc2\x69\x25\x0f\x05\x3c\xe8\xef\xe2\x40\x27\xbf\xed\x5d\x50\x37\x66\xf8\x71\x66\x22\x94\xc1\x2b\x3e\x8c\x9d\x65\x7c\xe7\x6e\xec\x95\x57\x7b\xa5\x71\x83\x65\x72\x83\x31\xbf\xde\x28\x2b\x58\x39\x3b\x43\xa2\x18\x8a\xe0\x6d\x01\x79\x89\xfb\xbc\xe3\xcc\xb4\xdb\xdc\x5b\xd4\xc4\xfe\xa8\xf4\x99\x7b\xef\x74\x30\x38\x8b\xc9\xa3\x6e\x42\x8e\xc9\x6d\xb6\x41\x96\x4d\x5e\x3c\xc6\xdc\xa1\x2d\x87\x65\x27\xf0\xfa\xf4\xc8\x32\x3b\xa2\xb2\x13\xea\x06\x5f\x9a\xeb\xca\xbb\xb5\xd2\x48\x59\x14\x6d\x8f\x75\x2f\x7c\xa6\xd5\xaa\x61\xbe\x8e\x4d\xd9\xa9\x93\x97\x96\xa5\x6b\xeb\xc7\xe2\x83\xcd\xcb\x55\xe8\x2f\x8b\x9b\xfb\x60\xb3\x0b\xa7\x47\x51\x3e\x5b\x7d\xb8\x48\x53\x4a\x1c\x49\xe9\x40\x8c\xa6\x6c\xfe\xf6\xe8\x23\x75\x03\xc9\x7a\x90\x9f\xda\x24\xde\xa2\x71\x65\xd0\xd8\x0f\x3e\xde\x5a\xdf\x77\x83\x16\xed\xbe\x2b\xc5\xac\xe1\x78\x1e\x2f\xea\xaa\x98\x3f\xcc\xa6\x49\x45\xed\x85\x0e\xf8\xe8\x9d\x49\xab\x0f\x60\xad\x50\x97\x2f\xb8\xbe\xb4\x37\x9e\x3a\xf0\xd8\x6a\x86\xb1\x5d\xc6\xce\x72\x65\xdf\x97\xe9\xf2\xef\xf9\xf8\x6d\x31\x7d\x79\x7a\x9e\xf4\xf7\x2d\xe0\xd3\x9f\xa3\x4f\x8d\xbd\x56\x70\xa7\xbd\x6c\x1d\xd5\xf7\xd3\xad\xad\xbe\x36\x2e\x95\x71\x77\xe1\x55\x99\xdc\x14\xc9\x5d\xaa\x3b\xc2\xf8\x40\x16\x77\x69\xaf\x49\xe1\x96\x10\x62\x0b\xb6\x6b\xb5\x99\x89\xaa\xcb\xa4\x18\x0d\xa5\x32\xd8\xe1\xa1\xe9\x6b\x99\x93\xd5\x95\xca\x3e\x47\xf0\xc7\x28\x3b\x8d\x96\x61\xe4\xef\x29\xf2\x56\x77\x88\x33\x29\xf6\x29\xa1\x2f\x54\xf9\x51\x63\xf9\x5f\xe7\xb8\xd5\x6f\xda\x73\xa4\xa8\xe1\x41\x18\x7d\xe3\x24\x2d\xe6\x67\x0f\xf2\x71\x2f\x2b\x2d\xb5\x5d\x7a\x5c\x3f\x2a\x1e\x95\x27\x6e\xbc\xb1\x60\x96\xc9\x1b\x21\x7e\x2b\x64\x31\xdc\x85\x15\x7a\x8b\x8c\x34\x54\x2e\x77\x14\xb3\x6e\xc3\x7b\x03\xaa\xbc\x72\xc7\xa1\xa3\x05\x51\x3d\xcc\x3f\xd5\xca\xaa\xdf\x2c\xd2\x2b\x56\x52\xe8\xb6\xa4\xa2\xcc\xe6\xc8\xdf\x9d\xdf\x25\x07\x39\xca\xef\x69\x92\xd8\x3c\x12\x0b\xcf\x57\x07\x23\xc9\x2d\x46\xdd\xf9\x7a\xcf\x12\xd7\x22\x68\xce\x4e\xe6\xd3\x63\x23\x9d\x89\xe7\x79\xd9\x19\x62\x5f\xbc\x90\xb8\x40\xaf\x5c\xb9\x44\xe9\x6c\x49\x05\xfc\x3e\x6a\x71\x4e\xc7\xe7\x96\x72\xf6\x94\xf7\x5f\x5a\x23\x82\xd0\x1a\xe2\x23\x84\x09\xc8\x01\x6f\x05\xd7\x09\x04\x45\x20\x8e\xf1\x42\xe5\x91\xd0\x32\x38\x7b\x44\xc7\x5b\x39\x17\x4b\xfb\x96\x2b\x60\xfa\xae\x88\x69\xf0\x5f\x00\x00\x00\xff\xff\x67\x5c\x5e\x98\x44\x0a\x00\x00")

func assetsTuned06DsTunedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTuned06DsTunedYaml,
		"assets/tuned/06-ds-tuned.yaml",
	)
}

func assetsTuned06DsTunedYaml() (*asset, error) {
	bytes, err := assetsTuned06DsTunedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/06-ds-tuned.yaml", size: 2628, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTunedDefaultCrTunedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4f\x6f\xdb\xc6\x13\xbd\xeb\x53\x0c\x94\xdf\xc1\xc6\x2f\xa4\x2c\x59\xb2\x63\x02\xec\x25\x0d\x90\x43\x8b\x04\x4d\xda\x4b\x51\x10\xeb\xe5\x50\x5c\x68\xff\x61\x67\x56\xb6\xd2\xf6\xbb\x17\x4b\x91\xb4\xad\xd8\x8e\x82\xba\x3a\x91\xdc\xd9\xf7\xe6\xcd\xdb\x19\xad\xf0\xea\x37\x0c\xa4\x9c\x2d\x80\xa3\xc5\x3a\x77\x1e\x2d\xb5\xaa\xe1\x5c\xb9\xd9\x76\x3e\xd9\x28\x5b\x17\xf0\x39\xad\x4d\x0c\xb2\xa8\x05\x8b\x62\x02\x60\x85\xc1\x02\x6a\x6c\x44\xd4\xdc\xbf\x93\x17\x12\x0b\x18\x21\x32\xa9\x23\x31\x86\xcc\xba\x1a\x33\x8e\x56\xd9\x75\xe6\x3c\x06\xc1\x2e\x4c\xc8\xa3\x4c\x48\x3e\xb8\x46\x69\x4c\x8f\x59\x0f\x3b\x1d\x21\xa6\x13\x00\x80\x8e\x14\xfe\xea\x9e\x01\x7e\x37\x42\xd9\x3f\xfa\x17\x8a\xc6\x88\xb0\x2b\x3f\x78\x56\x46\x7d\x41\xa0\x1d\x31\x1a\x82\x10\x6d\xe2\x83\x0f\x1e\xed\xa7\x04\x05\x27\x5e\x04\xb4\x3c\x10\x9e\xf6\x08\xca\x4a\x1d\x6b\x2c\xff\xf7\x67\x53\x6c\x55\xe0\x4a\xb6\x28\x37\xdd\x63\x14\x3a\x5b\x47\x24\x2e\xb8\x0d\x2e\xae\x5b\x1f\x39\xf3\x18\x1a\x17\x8c\xb0\x12\xff\x9e\x0c\x29\x11\x6a\x65\xe3\xed\x90\x95\xd8\xca\x4a\x0a\xd9\x62\xc5\x6d\x40\x6a\x9d\xae\xcb\x37\xf3\xab\xc5\x18\x6f\x91\x87\x58\xdb\x54\xd2\x59\xcb\x41\xc8\x4d\xd5\x0a\x6a\x49\x7d\xc1\x72\x7e\x3e\x3f\xbb\xbc\x8b\xa7\x1d\x49\xd6\xe3\x16\xe4\x5c\xf9\xed\x32\x57\xbe\x6a\x5c\xb8\x11\xa1\x2e\xe7\xfd\xda\x06\x83\x45\x9d\x7b\x55\x57\x46\xdc\x96\x3f\xf4\x40\x77\x1b\x2d\x72\xa3\x34\x63\xc8\x1f\x30\xa7\xe0\xf9\xd9\xf2\xcd\xea\xf2\xe2\x90\xc5\xa2\x5a\xb7\x79\xef\x75\xbe\x96\xbd\xa8\xf9\x5e\xd2\x71\xc1\x8b\xf2\x7c\x71\x79\xf1\xe6\xc8\xe8\xf3\xf2\x62\xb5\x3a\x3f\x48\xe4\xe2\x7b\x12\x79\x32\xf8\xd1\x44\x9e\x8c\x1e\x12\xb9\xe7\x43\x43\x83\x0d\x33\xda\xd1\xcc\xb8\x3a\x6a\x9c\xd9\xad\xc1\x4a\xba\x80\x33\x2f\x82\x30\xc8\x18\x68\xa6\x5c\xc5\xca\xa0\x8b\x5c\x2e\x17\x57\xcb\xab\x8b\xcb\xc5\xd5\xea\xd8\xcd\x46\xdc\x56\x01\x39\x28\xa4\x72\x7e\x36\x79\xb4\x3d\x32\xe9\x2c\x07\xa7\x33\xaf\x85\xc5\x97\x6c\x96\x1e\x18\x3a\xe0\x83\x4e\x19\xe9\x9f\x38\x9e\xaf\x60\x93\x86\x09\xec\xbf\x02\x21\xb3\xb2\x6b\x7a\x0d\x46\xdc\x26\xde\xc4\xa4\x66\x0e\xee\x9a\x6a\xd8\x38\x02\xfc\xac\xac\x32\x42\x83\x0f\x88\xc6\xb3\x72\x16\xd6\x41\xd8\xa8\x45\x50\xbc\x83\xc6\x05\x78\xfb\xf1\xd7\xec\xda\x45\x5b\x03\x0b\xda\x50\x31\xee\x3d\xe9\x5d\x2c\x60\x0e\x86\x50\xbe\x02\x38\x99\xc3\xff\x41\x69\xb7\x3e\xb1\xd2\x47\x3a\x3d\x7d\x0d\xd1\x2a\xa6\x02\xac\xb0\x8e\x50\x3a\x5b\xd3\xe9\xc3\x0e\x22\xd9\x62\x5d\x19\x65\xab\x7b\xcc\x95\x4d\x66\xec\x7f\x23\xe1\xe7\x16\x81\x1d\x0b\x0d\xc9\x6d\xe0\x16\xa1\xdb\x1c\x35\x06\xb8\x51\x5a\xa7\x6a\x92\xaa\x31\x80\x00\xa3\xd6\x41\x30\xd6\x69\x0c\x49\x24\x1a\x51\xa6\xdd\xc0\x80\xd6\xf1\x14\x44\x52\xd5\x46\x02\x8d\x44\xa0\xd5\x06\xf5\x0e\xd8\xc1\x35\x42\xc0\x6c\x80\xb8\x53\xbc\xf7\x71\x18\xc7\xa0\x08\x56\x5d\x8a\xaf\x41\xe5\x98\xc3\x59\xbe\x02\xf3\x94\xbe\x84\xa5\x9c\xad\xa4\x23\x4e\xf2\x56\x07\xea\x3e\xbd\x7d\xff\xee\xc7\xea\xc3\xe7\xf7\xef\x7e\x81\x1b\xb1\xc1\x2c\xfa\xfb\x5e\xe4\x5f\x79\xf7\xf1\x71\xcf\x6e\x5a\xb4\x7b\xab\x3a\x18\x88\x3e\x07\xf8\xc9\xdd\x60\xe8\x2a\xb6\x15\x3a\xa6\x32\x8e\x30\xca\xf8\xe0\xb6\x38\x72\x6a\xc1\x68\xe5\xae\xaf\xcd\x70\x72\xba\xa3\x30\x2c\xc9\xa0\x58\xc9\xe4\x43\xa2\xc9\x1f\xd3\x9b\xd0\xa2\x3f\xb4\x74\xd9\x6b\x7e\xbc\xc9\xd2\xdf\xd7\x4b\xf6\x56\xc2\xa3\xef\xec\xa9\x71\x62\xb2\xf4\x55\x23\x88\xd3\x86\xf2\xbc\x5f\x6d\x28\x57\xd6\xb1\x6a\x76\x79\x1a\x1b\x91\x30\x54\x37\x82\x65\x8b\x74\x37\xc0\xbe\x35\x3e\x32\xa4\x7f\xaf\xf2\xdd\x27\x70\xf6\xeb\x39\x92\x3d\x3b\x47\x1e\xe6\xf1\x44\x05\xb6\xa6\xd3\x66\x84\xaf\xa4\x8b\x96\xcb\xc5\xc5\x62\xbe\x5c\x3e\xe3\xd9\x7f\x22\xe8\x59\xf3\x3a\xda\xef\xcf\x3f\xa0\x74\xc6\xa0\xad\xf7\x97\xa0\xe1\x46\xf4\x6d\xa7\x7c\x50\x2e\x9d\xe1\x02\xe6\xfb\x86\x35\xc9\xf4\xfd\x28\xcc\x40\x8b\x6b\xd4\x05\x4c\x1f\xb9\xd9\xa1\x16\xc4\x4a\x12\x8a\x20\xdb\x69\x9f\x21\xef\x7c\x22\xf5\xae\x1e\xbe\xdc\x83\xbb\x0f\xd8\x15\x37\x38\x8d\xf9\x26\x5e\xa7\xee\x62\xa4\x84\x6a\x44\xba\xee\x4d\x8f\x8e\x57\xb6\x09\x62\x3a\x79\x52\xf4\x03\x13\xef\xb4\x2e\x5e\x58\xeb\x71\x45\x3f\xcc\xe2\xfc\x99\x2c\x8e\x29\xd0\xcb\x94\xe7\x30\xab\xe5\xd9\xe4\x9f\x00\x00\x00\xff\xff\x4a\xe0\xd1\x90\xd5\x0b\x00\x00")

func assetsTunedDefaultCrTunedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTunedDefaultCrTunedYaml,
		"assets/tuned/default-cr-tuned.yaml",
	)
}

func assetsTunedDefaultCrTunedYaml() (*asset, error) {
	bytes, err := assetsTunedDefaultCrTunedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tuned/default-cr-tuned.yaml", size: 3029, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"assets/tuned/01-service-account.yaml":      assetsTuned01ServiceAccountYaml,
	"assets/tuned/02-cluster-role.yaml":         assetsTuned02ClusterRoleYaml,
	"assets/tuned/03-cluster-role-binding.yaml": assetsTuned03ClusterRoleBindingYaml,
	"assets/tuned/04-cm-tuned-profiles.yaml":    assetsTuned04CmTunedProfilesYaml,
	"assets/tuned/05-cm-tuned-recommend.yaml":   assetsTuned05CmTunedRecommendYaml,
	"assets/tuned/06-ds-tuned.yaml":             assetsTuned06DsTunedYaml,
	"assets/tuned/default-cr-tuned.yaml":        assetsTunedDefaultCrTunedYaml,
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
	"assets": {nil, map[string]*bintree{
		"tuned": {nil, map[string]*bintree{
			"01-service-account.yaml":      {assetsTuned01ServiceAccountYaml, map[string]*bintree{}},
			"02-cluster-role.yaml":         {assetsTuned02ClusterRoleYaml, map[string]*bintree{}},
			"03-cluster-role-binding.yaml": {assetsTuned03ClusterRoleBindingYaml, map[string]*bintree{}},
			"04-cm-tuned-profiles.yaml":    {assetsTuned04CmTunedProfilesYaml, map[string]*bintree{}},
			"05-cm-tuned-recommend.yaml":   {assetsTuned05CmTunedRecommendYaml, map[string]*bintree{}},
			"06-ds-tuned.yaml":             {assetsTuned06DsTunedYaml, map[string]*bintree{}},
			"default-cr-tuned.yaml":        {assetsTunedDefaultCrTunedYaml, map[string]*bintree{}},
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

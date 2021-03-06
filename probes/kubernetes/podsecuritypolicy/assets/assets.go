// Code generated for package podsecuritypolicy by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/yaml/psp-azp-hostport-approved.yaml
// assets/yaml/psp-azp-hostport-unapproved.yaml
// assets/yaml/psp-azp-seccomp-approved.yaml
// assets/yaml/psp-azp-seccomp-unapproved.yaml
// assets/yaml/psp-azp-seccomp-undefined.yaml
// assets/yaml/psp-azp-volumetypes-approved.yaml
// assets/yaml/psp-azp-volumetypes-unapproved.yaml
package podsecuritypolicy

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

var _assetsYamlPspAzpHostportApprovedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xbb\x6e\xf3\x30\x0c\x85\x77\x03\x7e\x07\xc2\xff\xfa\xfb\xd6\x6e\xda\x8a\xb6\xe8\x54\x20\x4b\xbb\x04\x1d\x18\x99\x89\x85\x48\xa2\xa0\x8b\x9b\xa0\xe8\xbb\x17\x8a\xeb\x24\x43\xea\xc1\xc3\x47\xf2\xe8\xf0\x10\x9d\x7a\x27\x1f\x14\x5b\x01\x53\x5f\x16\x7b\x65\x07\x01\x2b\x1e\xca\xc2\x50\xc4\x01\x23\x8a\xb2\x00\xf8\x67\xd1\x90\x80\x40\x32\x79\x15\x8f\xb5\x64\x1b\xe9\x10\xeb\x81\x0c\xe7\x3a\x5a\xcb\x11\xa3\x62\x1b\x4e\xfd\x90\x5b\x25\x1b\xd7\x2c\x23\x0d\x6a\x37\x62\xb3\x4f\x1b\xf2\x96\x22\x85\x46\x71\xeb\x78\x10\x50\xf9\x64\xa3\x32\xd4\x0e\xb4\xc5\xa4\x63\x55\x16\xc1\x91\x14\x00\x59\x69\x99\x7f\x9c\x5f\xfc\x55\xf7\xc9\x3e\x84\xb7\x40\x5e\x40\xdf\x75\x5d\x46\x57\x85\x17\xcf\xc9\x09\xb8\xef\xba\x6e\xa6\xdb\x05\xdd\x9d\x51\x48\xce\x69\x32\x64\x23\xea\x53\x31\x08\x58\x43\x0f\x1f\x8b\xd6\xc4\x3a\x19\x9a\xd7\xa9\xe1\xbc\x7f\x2d\xe3\xa1\x9e\x58\xcf\x2a\x64\x5c\x3c\x3e\x29\x2f\xe0\xeb\x3b\x93\x9c\x0b\x2a\x4b\xfe\xf6\xdc\x12\x17\x80\x32\xb8\x23\x01\x9b\x14\x8e\x1b\x3e\xcc\x4c\xb2\x31\x98\xf3\x5f\x43\x15\xc6\xea\x3f\x54\xb5\xcc\xff\xa0\x89\x1c\xf4\x63\x75\x31\xb7\xd8\x7b\xe5\x64\xe3\x12\xf9\x9f\x2e\x01\x4c\xee\x5b\x61\x1c\x05\xb4\xf9\xa8\xed\xc5\xc9\xed\x80\x01\x50\x6b\xfe\x5c\x79\x35\x29\x4d\x3b\x7a\x0e\x12\xf5\xe9\xc0\x02\xb6\xa8\x03\xc1\xd5\xf7\x13\x00\x00\xff\xff\xbd\xec\xbb\x0d\x46\x02\x00\x00")

func assetsYamlPspAzpHostportApprovedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsYamlPspAzpHostportApprovedYaml,
		"assets/yaml/psp-azp-hostport-approved.yaml",
	)
}

func assetsYamlPspAzpHostportApprovedYaml() (*asset, error) {
	bytes, err := assetsYamlPspAzpHostportApprovedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/yaml/psp-azp-hostport-approved.yaml", size: 582, mode: os.FileMode(438), modTime: time.Unix(1600269942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsYamlPspAzpHostportUnapprovedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\xb1\x8e\x13\x31\x10\xed\x23\xe5\x1f\x9e\x96\x96\x4d\x36\x20\xa1\x93\x3b\x04\x88\x0a\x29\x0d\x34\x27\x8a\xc9\xee\xe4\xd6\x3a\xdb\x63\xd9\xe3\x25\x11\xe2\xdf\x91\xb3\xec\x25\x42\x39\x17\x2e\xde\x9b\x79\xe3\xf7\x3c\x14\xed\x0f\x4e\xd9\x4a\x30\x98\x76\xeb\xd5\xb3\x0d\x83\xc1\x5e\x86\xf5\xca\xb3\xd2\x40\x4a\x66\xbd\x02\xde\x04\xf2\x6c\x90\xb9\x2f\xc9\xea\xb9\xed\x25\x28\x9f\xb4\x1d\xd8\x4b\xe5\x29\x04\x51\x52\x2b\x21\x5f\xea\x51\x4b\x7b\xf1\x71\xb3\xb4\x6c\xc8\xc5\x91\x36\xcf\xe5\xc0\x29\xb0\x72\xde\x58\xd9\x46\x19\x0c\x9a\x54\x82\x5a\xcf\xdb\x81\x8f\x54\x9c\x36\xeb\x55\x8e\xdc\x1b\xa0\x2a\x2d\xfd\x9f\xe6\x89\xff\xd4\x53\x09\x1f\xf3\xf7\xcc\xc9\x60\xd7\x75\x5d\x85\x6e\x88\xaf\x49\x4a\x34\x78\xdf\x75\xdd\x8c\x1e\x17\xe8\xdd\x0b\x94\x4b\x8c\x8e\x3d\x07\x25\x77\x21\xb3\xc1\x23\x76\xf8\xb9\x68\x4d\xe2\x8a\xe7\xd9\x4e\x8b\x17\xff\x6d\xaf\xa7\x76\x12\x37\xab\xb0\x8f\x7a\xfe\x6c\x93\xc1\xef\x3f\x15\xa9\xb9\x90\x0d\x9c\xee\xf7\x2d\x71\x01\xd6\xd3\x13\x1b\x1c\x4a\x3e\x1f\xe4\x34\x63\xbd\x78\x4f\x35\xff\x47\x34\x79\x6c\xde\xa2\x69\xfb\x7a\x67\xc7\x1c\xb1\x1b\x9b\xeb\xe3\x80\x28\x49\x97\xac\xdb\xeb\xdc\xbd\x24\x35\x78\xe8\x1e\x3e\xcc\x14\x30\x4a\xd6\xff\xd1\xd9\xdb\x37\x29\xe1\x46\xe3\x15\x8b\x80\xaf\x75\x7b\xd2\xd1\x60\x5b\x37\x62\x7b\xb5\x71\xff\x77\x00\x72\x4e\x7e\xed\x93\x9d\xac\xe3\x27\xfe\x92\x7b\x72\x97\xed\x30\x38\x92\xcb\x8c\x9b\xf3\x37\x00\x00\xff\xff\x48\xd0\xb0\x8c\x83\x02\x00\x00")

func assetsYamlPspAzpHostportUnapprovedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsYamlPspAzpHostportUnapprovedYaml,
		"assets/yaml/psp-azp-hostport-unapproved.yaml",
	)
}

func assetsYamlPspAzpHostportUnapprovedYaml() (*asset, error) {
	bytes, err := assetsYamlPspAzpHostportUnapprovedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/yaml/psp-azp-hostport-unapproved.yaml", size: 643, mode: os.FileMode(438), modTime: time.Unix(1600269942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsYamlPspAzpSeccompApprovedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x4f\x8b\xdb\x30\x10\xc5\xef\x81\x7c\x87\xc1\xe7\x3a\x71\x5a\x7a\x51\x4f\xa5\x2d\xbd\x06\x4a\x7b\x59\xf6\x30\x2b\xbf\xac\x45\x24\x8d\xd0\x8c\xd3\xe4\xdb\x17\xc7\xeb\x24\x94\xd5\x41\x30\xbf\x37\x7f\x78\x23\x71\x09\x7f\x50\x35\x48\x76\x74\xda\xad\x57\xc7\x90\x7b\x47\x7b\xe9\xd7\xab\x04\xe3\x9e\x8d\xdd\x7a\x45\x94\x39\xc1\x91\xc2\x8f\x35\xd8\xa5\xf5\x92\x0d\x67\x6b\x7b\x24\x99\x64\xce\x59\x8c\x2d\x48\xd6\x6b\x3a\x4d\xa9\x5e\x52\xd9\x2c\x25\x1b\x8e\x65\xe0\xcd\x71\x7c\x41\xcd\x30\xe8\x26\xc8\xb6\x48\xef\xa8\xa9\x63\xb6\x90\xb0\xed\x71\xe0\x31\x5a\xb3\x5e\x69\x81\x77\x44\x53\xa7\xa5\xfe\xdb\x3c\xf1\xad\x7b\x1d\xf3\x57\xfd\xad\xa8\x8e\x76\x5d\xd7\x4d\xe8\x41\xf8\x59\x65\x2c\x8e\x3e\x75\x5d\x37\xd3\xc3\x82\x3e\xde\x90\x8e\xa5\x44\x24\x64\xe3\x78\x15\xd5\xd1\x13\xed\xe8\x99\x68\xe9\x36\x99\xe4\x90\x51\x67\x4f\xed\x7d\x09\xad\xb7\xf3\xcd\x3b\x51\x48\xfc\x0a\x47\xbd\xf8\x23\xea\xe4\x8b\x63\x09\x19\x2e\xb2\x41\x6d\xce\xf1\x92\x12\x4f\xbb\x7d\xa2\x46\x87\xe6\x03\x35\xad\x9f\x6e\x93\xd1\x0f\xb4\xb5\x54\xb6\x03\x38\xda\x70\xf9\x42\x1a\x81\x42\xbb\xa1\xa1\xe7\xb9\x38\x86\x13\x32\x54\xf7\x55\x5e\xf0\xb6\x02\x22\x9c\xe1\x6f\xc1\x7d\xc2\x9d\xb4\xe4\xd9\x1e\xc3\xc7\x31\x0b\x0f\x39\x58\xe0\xf8\x1d\x91\x2f\xbf\xe0\x25\xf7\xea\xe8\xf3\xa2\x16\xd4\x20\xfd\xff\xfc\xfd\x47\x21\xe2\x18\xe5\xef\xbe\x86\x53\x88\x78\xc5\x0f\xf5\x1c\xaf\x9f\xc2\xd1\x81\xa3\x82\x1e\xcf\xbf\x00\x00\x00\xff\xff\x92\x82\xce\x8b\x7a\x02\x00\x00")

func assetsYamlPspAzpSeccompApprovedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsYamlPspAzpSeccompApprovedYaml,
		"assets/yaml/psp-azp-seccomp-approved.yaml",
	)
}

func assetsYamlPspAzpSeccompApprovedYaml() (*asset, error) {
	bytes, err := assetsYamlPspAzpSeccompApprovedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/yaml/psp-azp-seccomp-approved.yaml", size: 634, mode: os.FileMode(438), modTime: time.Unix(1600269942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsYamlPspAzpSeccompUnapprovedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xb1\x6e\xf3\x30\x0c\x84\x77\x03\x7e\x07\xc2\xf3\x6f\xc7\xfe\xbb\x69\x2b\x8a\xa2\x6b\x96\x76\x09\x32\x30\x32\x13\xab\x91\x44\x41\x94\xd2\xe4\xed\x0b\xc5\xb5\x9b\xa1\x1a\x34\xdc\x51\x27\xde\x87\xc1\x7c\x50\x14\xc3\x5e\xc1\x65\xa8\xab\xb3\xf1\xa3\x82\x2d\x8f\x75\xe5\x28\xe1\x88\x09\x55\x5d\x01\x78\x74\xa4\x40\x48\xe7\x68\xd2\xad\xd5\xec\x13\x5d\x53\x3b\x92\xe3\x62\xa3\xf7\x9c\x30\x19\xf6\x72\x1f\x87\x32\xaa\xd9\x85\x6e\x79\xd2\xa1\x0d\x13\x76\xe7\x7c\xa0\xe8\x29\x91\x74\x86\x37\x81\x47\x05\x8d\x65\x8d\x76\x62\x49\x9b\xec\x31\x84\xc8\x17\x1a\xdb\x10\xf9\x68\x2c\x75\x9f\xc2\xbe\xa9\x2b\x09\xa4\x15\x40\xc9\x5e\x12\x5f\xe6\x1d\x7e\xfe\x8b\xd9\x3f\xcb\xbb\x50\x54\x30\xf4\x7d\x5f\xa4\x07\xe3\x2d\x72\x0e\x0a\x9e\xfa\xbe\x9f\xd5\xe3\x22\xfd\x5f\x25\xc9\x21\x58\x72\xe4\x13\xda\xbb\x29\x0a\x76\x30\xc0\x1e\x60\x49\x2b\xb5\xd1\x78\x8a\x73\xcb\xf6\x17\x4b\xab\xd3\x75\xa5\x01\x60\x1c\x9e\x48\xc1\x21\xcb\xed\xc0\xd7\x59\xd3\xec\x1c\x16\xba\x3b\x68\x64\x6a\xfe\x41\xd3\xea\x72\x8b\x25\x0a\x30\x4c\x0d\xec\x57\x74\x7f\x14\x04\x40\x6b\xf9\x6b\x1b\xcd\xc5\x58\x3a\xd1\xab\x68\xb4\x77\xe4\x0a\x8e\x68\x85\xe0\xe1\x7c\x07\x00\x00\xff\xff\x95\x72\x7e\x36\xd7\x01\x00\x00")

func assetsYamlPspAzpSeccompUnapprovedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsYamlPspAzpSeccompUnapprovedYaml,
		"assets/yaml/psp-azp-seccomp-unapproved.yaml",
	)
}

func assetsYamlPspAzpSeccompUnapprovedYaml() (*asset, error) {
	bytes, err := assetsYamlPspAzpSeccompUnapprovedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/yaml/psp-azp-seccomp-unapproved.yaml", size: 471, mode: os.FileMode(438), modTime: time.Unix(1600269942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsYamlPspAzpSeccompUndefinedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xb1\xae\xdb\x30\x0c\x45\x77\x03\xfe\x07\xc2\x6f\xad\x1d\xbb\xdd\xb4\x15\x45\xd1\x35\x4b\xbb\x04\x19\x68\x89\x4e\x84\x48\xa4\x20\x4a\x69\xf2\xf7\x85\xe3\xa4\xc8\xf0\x34\x68\x38\x24\x0f\xc8\x8b\xc9\xff\xa1\xac\x5e\xd8\xc0\x75\x6a\x9b\x8b\x67\x67\x60\x2f\xae\x6d\x22\x15\x74\x58\xd0\xb4\x0d\x00\x63\x24\x03\x29\xcb\x9c\x7b\x25\x6b\x25\xa6\xbe\xb2\xa3\xc5\x33\xb9\xb5\xe1\x03\x99\xa5\x60\xf1\xc2\xfa\x98\x80\x0f\x78\x36\x0e\x4a\xb6\x66\x5f\xee\x03\x86\x74\xc6\xe1\x52\x67\xca\x4c\x85\x74\xf0\xb2\x4b\xe2\x0c\x74\xb9\x72\xf1\x91\x76\x8e\x16\xac\xa1\x74\x6d\xa3\x89\xac\x01\x58\x55\xaf\xf9\x1f\xc2\x85\x6e\x65\xd3\x43\xae\xfc\x5d\x7f\x2b\x65\x03\xd3\x38\x8e\x2b\x7a\x2b\xfc\xca\x52\x93\x81\x6f\xe3\x38\x6e\x74\x79\xa1\xaf\xff\x91\xd6\x94\x02\x45\xe2\x82\xe1\x51\x54\x03\x07\x98\xe0\x08\xf0\xb2\x59\xe1\x82\x9e\x29\x6f\x47\xf5\xcf\x20\x94\x6c\x6f\xcb\xad\x77\x14\x65\x73\xf9\x88\x27\x32\x30\x57\xbd\xcf\x72\xdb\x98\x95\x18\x71\xcd\xf3\x00\x9d\x9e\xbb\x2f\xd0\xf5\x76\xfd\x35\x10\x25\x98\xce\x1d\x1c\x9f\x8b\x7c\x7a\x20\x00\x86\x20\x7f\xf7\xd9\x5f\x7d\xa0\x13\xfd\x54\x8b\xe1\x91\xb0\x81\x05\x83\x12\xbc\xbd\x7f\x01\x00\x00\xff\xff\x10\xfc\x21\x5f\xc9\x01\x00\x00")

func assetsYamlPspAzpSeccompUndefinedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsYamlPspAzpSeccompUndefinedYaml,
		"assets/yaml/psp-azp-seccomp-undefined.yaml",
	)
}

func assetsYamlPspAzpSeccompUndefinedYaml() (*asset, error) {
	bytes, err := assetsYamlPspAzpSeccompUndefinedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/yaml/psp-azp-seccomp-undefined.yaml", size: 457, mode: os.FileMode(438), modTime: time.Unix(1600269942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsYamlPspAzpVolumetypesApprovedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xb1\xae\xdb\x30\x0c\xdc\x03\xe4\x1f\x08\xcf\x55\x6c\xb7\x9b\xb6\xa2\x28\x3a\x3d\x20\x4b\xbb\x3c\x74\x60\x64\x3a\x16\x9e\x24\x0a\x12\xe5\xe6\xfd\x7d\xa1\xc8\x4e\x3a\x14\xcf\x83\x87\xbb\x13\xc5\xbb\x13\x46\xfb\x8b\x52\xb6\x1c\x34\xac\xe3\xf1\xf0\x66\xc3\xa4\xe1\xcc\xd3\xf1\xe0\x49\x70\x42\x41\x7d\x3c\x00\x04\xf4\xa4\x21\x93\x29\xc9\xca\xbb\x32\x1c\x84\x6e\xa2\x26\xf2\x5c\x69\x0c\x81\x05\xc5\x72\xc8\x77\x39\x54\xa9\x61\x1f\x4f\xfb\x91\x13\xba\xb8\xe0\xe9\xad\x5c\x28\x05\x12\xca\x27\xcb\x7d\xe4\x49\x43\x97\x4a\x10\xeb\xa9\x9f\x68\xc6\xe2\xa4\x3b\x1e\x72\x24\xa3\x01\xea\xa4\xfd\xfc\xb7\x76\xe3\x36\x3d\x95\xf0\x35\xff\xcc\x94\x34\x8c\xc3\x30\x54\xe8\x1f\xe2\x47\xe2\x12\x35\x7c\x19\x86\xa1\xa1\xf3\x0e\x7d\x7e\x40\xb9\xc4\xe8\xc8\x53\x10\x74\x77\x32\x6b\x78\x85\x11\x7e\xef\xb3\x56\x76\xc5\xd3\x66\x47\x6d\x01\x60\x8c\x89\x57\x9a\x54\x63\xdb\x28\x30\x1c\x66\x7b\x7d\xc1\xb8\x6d\x57\xbf\xa6\x17\xca\xa2\x1a\xad\x3c\xc6\x9d\xac\xb2\x1a\x21\xda\x40\xa9\x5d\xa1\x9e\x11\x2b\x23\xb7\x47\xb2\x00\xd6\xe3\x95\x34\x5c\x4a\x7e\xbf\xf0\xad\x61\x86\xbd\xc7\xda\xd4\x2b\x74\x79\xe9\x3e\x41\xa7\x4c\xfd\x67\x47\x14\x61\x5c\xba\xe6\xe3\x99\x4b\xdb\xf7\x85\x4b\x90\xbd\xa1\x8f\x3d\x81\xaf\xda\x33\xca\xa2\xa1\x17\x1f\xfb\x6a\x65\xe7\x1e\x15\xff\xa7\x1a\x00\x74\x8e\xff\x9c\x93\x5d\xad\xa3\x2b\x7d\xcf\x06\xdd\xfd\x69\x68\x98\xd1\x65\xfa\x1b\x00\x00\xff\xff\x99\x3d\x06\x23\x73\x02\x00\x00")

func assetsYamlPspAzpVolumetypesApprovedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsYamlPspAzpVolumetypesApprovedYaml,
		"assets/yaml/psp-azp-volumetypes-approved.yaml",
	)
}

func assetsYamlPspAzpVolumetypesApprovedYaml() (*asset, error) {
	bytes, err := assetsYamlPspAzpVolumetypesApprovedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/yaml/psp-azp-volumetypes-approved.yaml", size: 627, mode: os.FileMode(438), modTime: time.Unix(1600269942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsYamlPspAzpVolumetypesUnapprovedYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x3d\xab\xdc\x30\x10\xec\x0f\xfc\x1f\x16\xd7\xf1\x57\xd2\xa9\x0b\x21\xa4\x0a\xb8\x49\x9a\x47\x8a\x3d\x79\xef\x2c\x9e\xa4\x15\x5a\xc9\xb9\xf7\xef\x83\x4e\xf6\xbb\x2b\x02\xcf\x85\xc1\x33\xab\xf1\xec\x8c\x30\x98\xdf\x14\xc5\xb0\x57\xb0\x4d\xcd\xe9\xd5\xf8\x45\xc1\xcc\x4b\x73\x72\x94\x70\xc1\x84\xaa\x39\x01\x78\x74\xa4\x40\x48\xe7\x68\xd2\x5b\xa7\xd9\x27\xba\xa5\x6e\x21\xc7\x85\x46\xef\x39\x61\x32\xec\xe5\x3e\x0e\x65\x54\xb3\x0b\xfd\x71\xa4\x47\x1b\x56\xec\x5f\xf3\x99\xa2\xa7\x44\xd2\x1b\x1e\x02\x2f\x0a\xda\x98\x7d\x32\x8e\x86\x85\x2e\x98\x6d\x6a\x9b\x93\x04\xd2\x0a\xa0\x28\x1d\xe7\xbf\xd5\x3f\xee\xea\x31\xfb\xaf\xf2\x4b\x28\x2a\x98\xc6\x71\x2c\xd0\x13\xf1\x23\x72\x0e\x0a\xbe\x8c\xe3\x58\xd1\xcb\x01\x7d\x7e\x87\x24\x87\x60\xc9\x91\x4f\x68\xef\xa4\x28\x78\x81\x09\xfe\x1c\x5a\x1b\xdb\xec\x68\x5f\xa7\xdb\x03\xc8\x1e\x43\x88\xbc\xd1\xd2\x55\xbe\x8a\xc1\xca\x92\x66\x4c\xeb\x6e\x0f\x20\x94\x0f\x18\x36\x8c\x83\xe5\x6b\x59\x54\x0a\x55\x72\x43\xe3\x29\x56\xdd\xee\x91\x6b\xa7\xd3\xed\x3d\x4e\x00\xe3\xf0\x4a\x0a\xce\x59\xde\xce\x7c\xab\x98\x66\xe7\xb0\xd4\xf3\x02\xad\xac\xed\x27\x68\x3b\x5d\xde\x62\x89\x02\x4c\x6b\x5b\xcd\x3f\xc2\xa8\x16\x7f\x72\xf6\xe9\xa8\xe5\xa3\x45\xc0\x95\xe9\xb9\xba\x2f\xed\x0f\x0f\x4f\x87\xec\xff\x2b\x01\x40\x6b\xf9\xef\x1c\xcd\x66\x2c\x5d\xe9\xbb\x68\xb4\xf7\x2b\xa1\xe0\x82\x56\x08\x9e\x9e\x7f\x01\x00\x00\xff\xff\xd4\x7f\x20\xcb\x77\x02\x00\x00")

func assetsYamlPspAzpVolumetypesUnapprovedYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsYamlPspAzpVolumetypesUnapprovedYaml,
		"assets/yaml/psp-azp-volumetypes-unapproved.yaml",
	)
}

func assetsYamlPspAzpVolumetypesUnapprovedYaml() (*asset, error) {
	bytes, err := assetsYamlPspAzpVolumetypesUnapprovedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/yaml/psp-azp-volumetypes-unapproved.yaml", size: 631, mode: os.FileMode(438), modTime: time.Unix(1600269942, 0)}
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
	"assets/yaml/psp-azp-hostport-approved.yaml":      assetsYamlPspAzpHostportApprovedYaml,
	"assets/yaml/psp-azp-hostport-unapproved.yaml":    assetsYamlPspAzpHostportUnapprovedYaml,
	"assets/yaml/psp-azp-seccomp-approved.yaml":       assetsYamlPspAzpSeccompApprovedYaml,
	"assets/yaml/psp-azp-seccomp-unapproved.yaml":     assetsYamlPspAzpSeccompUnapprovedYaml,
	"assets/yaml/psp-azp-seccomp-undefined.yaml":      assetsYamlPspAzpSeccompUndefinedYaml,
	"assets/yaml/psp-azp-volumetypes-approved.yaml":   assetsYamlPspAzpVolumetypesApprovedYaml,
	"assets/yaml/psp-azp-volumetypes-unapproved.yaml": assetsYamlPspAzpVolumetypesUnapprovedYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"yaml": &bintree{nil, map[string]*bintree{
			"psp-azp-hostport-approved.yaml":      &bintree{assetsYamlPspAzpHostportApprovedYaml, map[string]*bintree{}},
			"psp-azp-hostport-unapproved.yaml":    &bintree{assetsYamlPspAzpHostportUnapprovedYaml, map[string]*bintree{}},
			"psp-azp-seccomp-approved.yaml":       &bintree{assetsYamlPspAzpSeccompApprovedYaml, map[string]*bintree{}},
			"psp-azp-seccomp-unapproved.yaml":     &bintree{assetsYamlPspAzpSeccompUnapprovedYaml, map[string]*bintree{}},
			"psp-azp-seccomp-undefined.yaml":      &bintree{assetsYamlPspAzpSeccompUndefinedYaml, map[string]*bintree{}},
			"psp-azp-volumetypes-approved.yaml":   &bintree{assetsYamlPspAzpVolumetypesApprovedYaml, map[string]*bintree{}},
			"psp-azp-volumetypes-unapproved.yaml": &bintree{assetsYamlPspAzpVolumetypesUnapprovedYaml, map[string]*bintree{}},
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

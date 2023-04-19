// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/ccm.yaml
// manifests/coredns.yaml
// manifests/local-storage.yaml
// manifests/metrics-server/aggregated-metrics-reader.yaml
// manifests/metrics-server/auth-delegator.yaml
// manifests/metrics-server/auth-reader.yaml
// manifests/metrics-server/metrics-apiservice.yaml
// manifests/metrics-server/metrics-server-deployment.yaml
// manifests/metrics-server/metrics-server-service.yaml
// manifests/metrics-server/resource-reader.yaml
// manifests/rolebindings.yaml
// manifests/traefik.yaml
//go:build !no_stage
// +build !no_stage

package deploy

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

var _ccmYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8f\x13\x31\x0c\x85\xef\xf3\x2b\xa2\x1e\x91\xd2\x15\xe2\x82\xe6\x08\x07\xee\x2b\xc1\xdd\x4d\x1e\xdd\xd0\x4c\x1c\xd9\x4e\x61\xf9\xf5\x68\x3a\x5d\x51\x3a\xdb\xd2\x16\x10\x7b\x8b\xac\xf8\xf3\xf3\x73\x62\xaa\xe9\x13\x44\x13\x97\xde\xc9\x8a\xc2\x92\x9a\x3d\xb0\xa4\xef\x64\x89\xcb\x72\xf3\x56\x97\x89\xef\xb6\xaf\xbb\x4d\x2a\xb1\x77\xef\x73\x53\x83\xdc\x73\x46\x37\xc0\x28\x92\x51\xdf\x39\x57\x68\x40\xef\x36\x6f\xd4\x87\xcc\x2d\xfa\xc0\xc5\x84\x73\x86\xf8\x81\x0a\xad\x21\x9d\xb4\x0c\xed\x3b\xef\xa8\xa6\x0f\xc2\xad\xea\x98\xe8\x5d\x60\x96\x98\xca\x61\xbd\xce\x39\x81\x72\x93\x80\xfd\xa5\x0c\x52\x68\xe7\xdc\x16\xb2\xda\xc7\xd6\xb0\x09\x20\x20\xc3\xee\xd8\x6a\x1c\x8f\xb3\x1a\x8b\xc5\x1c\x89\x2d\x8a\x1d\x21\x0f\x50\x95\x2c\x3c\x5c\x0d\x2d\x1c\x8f\x65\x2e\x5e\x2d\xae\xc8\xbd\x53\x23\x6b\xba\x0b\x28\x64\x9b\xc2\x61\xec\x00\x3b\xe9\xbb\x08\xfc\xc4\x99\xf2\x38\x9e\xf0\x31\x27\x9d\x0e\x5f\x6f\x42\xcf\xb4\x5d\xeb\xdd\x9e\x45\x21\x70\x3b\x37\x99\x51\xef\x65\x86\xd2\x00\xad\x34\x93\xf7\x3b\x16\xd5\xaa\x73\x5a\x24\x0c\x5c\x14\xc7\xca\x9e\x9f\x6f\x4c\x1a\x78\x0b\x79\xdc\x3f\xe9\xe7\x1e\x60\x89\x95\x53\x31\xcd\x73\x07\x4f\xcd\xc4\xfb\xee\xf6\x1f\xfb\x2e\x95\x98\xca\xfa\xea\x8f\xcb\x19\xf7\xf8\x3c\xde\x7e\xea\xf2\x4c\xe5\xce\xb9\xf9\xaa\xb8\xa8\x8e\xb6\xd5\x17\x04\xdb\xed\x88\x09\xf1\x51\x21\x97\xe5\xba\x9f\xc3\xee\xdd\xa6\xad\xe0\xf5\x51\x0d\xc3\x7f\x71\xcc\x8f\x7c\x1f\x91\xb1\x26\xe3\xbf\x6a\xe0\xd4\x55\x7f\x54\xe0\xa5\x38\xf7\x87\x96\xa1\x58\x0a\x3b\xb2\x17\x50\x3c\x27\xee\x46\x4b\x7f\xf1\x12\xdf\x0c\x65\xec\xcd\x53\x4d\xe3\xf2\x39\x29\xe3\x9f\xf8\xfb\x23\x00\x00\xff\xff\xde\xc0\x02\x82\x7a\x07\x00\x00")

func ccmYamlBytes() ([]byte, error) {
	return bindataRead(
		_ccmYaml,
		"ccm.yaml",
	)
}

func ccmYaml() (*asset, error) {
	bytes, err := ccmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ccm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x51\x6f\xdb\x38\x12\x7e\xf7\xaf\x20\x04\xe4\xe5\x70\x72\xe2\x0b\xda\xcb\xf1\x2d\x8d\xdd\x36\xb8\xc4\x35\x6c\xa7\x40\xb1\x58\x04\x34\x35\xb6\xb8\xa1\x38\x5c\x92\x72\xe2\xed\xe6\xbf\x2f\x28\xc9\x32\x69\x2b\x69\x92\xed\xfa\xc5\x92\x86\xf3\x0d\xf9\x71\xf8\xcd\x90\x69\xf1\x15\x8c\x15\xa8\x28\x59\x0f\x7a\x77\x42\x65\x94\xcc\xc0\xac\x05\x87\x73\xce\xb1\x54\xae\x57\x80\x63\x19\x73\x8c\xf6\x08\x51\xac\x00\x4a\x38\x1a\xc8\x94\x6d\xde\xad\x66\x1c\x28\xb9\x2b\x17\x90\xda\x8d\x75\x50\xf4\xd2\x34\xed\x85\xd0\x66\xc1\x78\x9f\x95\x2e\x47\x23\xfe\x60\x4e\xa0\xea\xdf\x9d\xd9\xbe\xc0\xe3\x36\xe8\x85\x2c\xad\x03\x33\x45\x09\x51\x44\xc9\x16\x20\xad\x7f\x22\x55\x08\xa3\xc0\x41\xe5\xba\x40\x74\xd6\x19\xa6\xb5\x50\xab\x3a\x46\x9a\xc1\x92\x95\xd2\xd9\x76\xaa\xf5\x84\xe8\x76\xc6\xa6\x94\x60\x69\x2f\x25\x4c\x8b\x4f\x06\x4b\x5d\x21\xa7\x24\x49\x7a\x84\x18\xb0\x58\x1a\x0e\xcd\x37\x50\x99\x46\xa1\x2a\xb0\x94\xd8\x9a\x94\xfa\x45\x63\x56\x3f\xb4\xeb\xf7\xaf\x6b\x30\x8b\xc6\x57\x0a\xeb\xaa\x87\x7b\xe6\x78\x7e\x18\x2f\x13\x96\xe3\x1a\xcc\xa6\xe1\xe1\x99\xe8\x52\xfc\x10\xfd\x6f\xb1\xfd\x41\xa8\x4c\xa8\x55\x44\x3a\x53\x0a\x5d\xe5\xd9\x30\xdf\x05\x19\x6d\x06\x2b\x1d\x96\x3a\x63\x0e\x28\x49\x9c\x29\x21\xf9\xf9\x7b\x87\x12\xa6\xb0\xac\xe6\xd7\xb0\xf9\xcc\x5a\x7b\x84\x1c\x26\xd6\x13\xc8\xb6\x5c\xfc\x06\xdc\x55\x89\xd1\x79\x04\xde\x9c\xf8\x3b\xc2\x51\x2d\xc5\xea\x9a\xe9\xb7\x1c\xa7\xed\xf0\x0b\x34\xb0\x14\x12\x28\xf9\xb3\xe2\xb4\x4f\xdf\x9d\x92\xef\xd5\xa3\xff\x81\x31\x68\x6c\xfb\x9a\x03\x93\x2e\x6f\x5f\x0d\xb0\x6c\xd3\xbe\xed\xb6\x83\x1c\x7d\xbf\xb8\xba\x99\xcd\x47\xd3\xdb\xe1\x97\xeb\xf3\xcb\xf1\xe3\x11\x11\x2a\x65\x59\x66\xfa\xcc\x68\x46\x84\x7e\x5f\x3f\xec\x22\x91\xea\x04\x10\xa1\x2c\xf0\xd2\x40\xf0\x7d\xc9\xa4\x74\xb9\xc1\x72\x95\x77\xa3\xb4\x63\x1f\x77\x13\x45\xeb\x2c\x39\x06\xc7\x8f\x1b\x2a\x8e\xc7\x98\xc1\xe7\xea\x73\x18\xd4\x39\x49\xde\x9f\x04\x1f\x0c\x48\x64\x19\x19\xbc\xb3\xdd\x53\xe8\x08\xa6\x0d\x16\xe0\x72\x28\x2d\xa1\xff\x1b\xbc\x3b\x6d\x0d\x4b\x34\xf7\xcc\x64\xa4\x5f\xcf\xc4\x1f\x47\xb9\xee\x73\x54\xcb\x76\x08\x67\x3c\x07\x72\xba\x9b\x81\x44\xd4\xbd\x78\x32\x81\x8d\x65\x0b\x26\x99\xe2\x35\x3f\xf5\x14\x44\xa1\xd1\xb8\x78\xb1\xbc\xb4\x0e\x8b\xe3\x7f\xf5\xbd\xc6\x80\x39\x48\x22\xa6\xb5\xdd\x1d\xdd\x21\x68\x89\x9b\x02\xde\xa6\xcc\x7b\x87\xf2\xcc\xa6\x4c\xeb\x66\x48\xed\xb8\x7f\x54\x6b\xe0\xc4\xe7\xde\x70\x3c\x4b\x7a\x56\x03\xa7\x95\x5e\xad\x85\x9f\xdf\x67\x61\x1d\x9a\xcd\x95\x28\x84\xa3\xc4\x73\xe3\x0f\xb6\x83\xd5\xa6\x8e\xe1\x36\x1a\x28\x99\xa2\x94\x42\xad\x6e\x2a\x89\xa8\x25\x25\xfc\x42\x1b\xda\x0a\xf6\x70\xa3\xd8\x9a\x09\xc9\x16\x3e\xcf\x07\x1e\x0e\x24\x70\x87\xa6\x1e\x53\x78\xc9\xbb\x0a\xd6\xd0\xbd\x0a\x07\x85\x96\x2d\x70\x48\x54\xb5\x37\x91\xff\x53\x3c\x6c\x57\x5a\xa7\x8d\x40\x23\xdc\xe6\x42\x32\x6b\xc7\x35\x25\x35\xa5\x29\xaf\x05\x26\xe5\x46\x38\xc1\x99\x4c\x1a\x17\x1b\x69\xc8\x78\x6f\x7f\x2a\x6a\x50\x82\x09\x65\xd6\xff\x52\x72\x07\x1b\x4f\x78\x03\x77\x9e\x65\xa8\xec\x17\x25\x37\x49\x90\xe4\xa8\xbd\x27\x1a\x4a\x92\xd1\x83\xb0\xce\x26\x07\x00\x0a\x33\x48\xbd\x68\xee\x49\x35\x47\xe5\x0c\xca\x54\x4b\xa6\xe0\x85\x98\x84\xc0\x72\x09\xdc\x51\x92\x8c\x71\xc6\x73\xc8\x4a\x09\x2f\x0f\x59\x30\xcf\xd0\xcf\x88\xe5\x23\xcc\xa2\x84\x38\xcc\x58\xb4\x94\x48\xa1\xca\x87\x96\x66\x8d\x12\x57\x9b\x99\xf6\x1a\x78\x81\xca\x27\xa8\x2f\xad\x21\xe9\x05\x7b\x98\xdd\xc1\x7d\x9d\x72\xdb\xdf\xd6\xf3\xff\x7e\x75\x71\x10\x2f\x5a\xfe\x68\x04\xa3\xef\x73\x50\x37\xca\x32\x27\xec\x52\xd4\xf9\x3b\xc4\x31\xba\xed\x1a\x82\xa1\x55\x02\x1e\xae\xe3\x89\x04\x7f\x3e\x4d\x09\xf1\x3b\xca\x84\x02\xd3\x7a\xa4\x07\x7a\x50\xff\x44\xc1\x56\x40\xc9\xd1\xf7\xd9\xb7\xd9\x7c\x74\x7d\x3b\x1c\x7d\x3c\xbf\xb9\x9a\xdf\x4e\x47\x9f\x2e\x67\xf3\xe9\xb7\xc7\x23\xc3\x14\xcf\xc1\x1c\x17\xc2\x57\x13\xc8\xd2\x06\x62\xfb\x4f\x07\xfd\xc1\x49\x7f\x10\x23\x4e\x4a\x29\x27\x28\x05\xdf\x50\x72\xb9\x1c\xa3\x9b\x18\xb0\x50\x15\xce\xfa\x17\x35\x37\x2d\x09\x5e\x32\xf6\x16\x59\x40\x81\x66\x43\xc9\xe0\xbf\x27\xd7\x22\x52\xfa\xdf\x4b\xb0\xfb\xa3\xb9\x2e\x29\x19\x9c\x9c\x14\x9d\x18\x11\x04\x33\x2b\x4b\xc9\x2f\x24\x49\xbd\xa4\x27\xff\x26\x49\xa4\xc1\xdb\xd2\x9a\x90\x5f\x5b\x97\x35\xca\xb2\x80\x6b\x7f\x7a\xa3\x54\xd9\x52\xeb\x2b\x7a\x5a\x0f\x0a\xe2\x17\x7e\xfc\x84\xb9\x9c\x46\x2a\x1f\xad\x85\x65\xfe\x3c\x53\xe2\x1b\xa5\x43\xe0\xaa\x1c\xa4\xaf\xc4\x6f\xaa\xc8\x8f\xc3\xf8\xfa\x13\x2d\xa7\xcd\x9e\x09\x1a\x47\x49\x50\x12\xb7\x55\x25\x9e\xbe\x36\xe8\x90\xa3\xa4\xe4\x66\x38\x79\x2d\x4e\xea\xb8\xee\xc4\x9a\x5f\x3c\x83\x15\x15\xea\x2d\x5a\x01\xce\x08\xde\x3d\xb3\x10\xad\xea\x51\xbc\x74\xa3\x72\xf0\xe0\xc2\x0c\x62\x52\xe2\xfd\xc4\x88\xb5\x90\xb0\x82\x91\xe5\x4c\x56\x72\x4c\x7d\x13\x61\x43\xd6\x39\xd3\x6c\x21\xa4\x70\x02\xf6\x72\x90\x65\x59\xfc\x21\x25\xe3\xd1\xfc\xf6\xc3\xe5\x78\x78\x3b\x1b\x4d\xbf\x5e\x5e\x8c\x22\x73\x66\x50\xef\x3b\x30\x29\x3b\x36\x6e\x8a\xe8\x3e\x0a\x09\x4d\xb7\x1a\x6f\xa3\x14\x6b\x50\x60\xed\xc4\xe0\x02\x42\xbc\xdc\x39\xfd\x09\x5c\x1c\x42\xd7\xf9\xb2\xd7\x12\x92\x26\x1d\x28\x39\x3b\x39\x3b\x89\x3e\x5b\x9e\x83\x27\xf9\xf3\x7c\x3e\x09\x0c\x42\x09\x27\x98\x1c\x82\x64\x9b\x19\x70\x54\x99\xa5\x71\x4b\xa6\xc1\x08\xcc\x5a\xdb\x20\xb4\x39\x51\x00\x96\x6e\x67\x0c\x6c\xb6\xe4\x1c\xac\x9d\xe7\x06\x6c\x8e\x32\x8b\xad\x4b\x26\x64\x69\x20\xb0\x9e\x46\x8d\xad\x78\x35\x15\x71\x3b\x1c\x30\x31\x38\x1b\xbc\x99\x89\x67\x88\xf8\xcf\x3f\xcc\x43\xa6\xec\x56\x81\x87\xf5\x45\xaa\x31\xd4\x02\xf2\x0a\x01\xe3\xdb\xab\x4a\xcc\x5b\x77\x41\xa9\xa8\x70\x50\xd8\xfd\x94\xae\x1a\x82\xad\xaa\xee\xd5\xb1\x7a\x0b\x3a\x8d\x8d\x63\xdb\xff\x77\x7a\x1e\x5a\x5f\xa8\x9d\x2f\x59\x5a\x7a\x20\xa4\xbe\x5b\xf1\xaa\xc0\x64\x73\x06\x9f\xbc\xe5\x35\xd7\xc6\x8e\xc6\x3c\xa8\xd8\x4f\x76\xe6\x07\xb7\xee\xdd\x5d\xc5\x77\x1c\x75\x7e\x26\x5e\x0b\x93\x0e\xb3\xe5\x86\xe9\x27\x6f\xdf\x2f\x68\xf4\xb7\x7d\x6c\xd3\xb7\x06\x48\x2f\xbd\x12\xc4\x9d\x7a\x57\xcc\x26\xc6\xe5\x84\x86\xd7\xce\xf1\xec\xf1\xa8\x17\x54\xa6\x74\xaf\xee\xe8\xb0\xa0\xec\x97\x9f\xb4\xa3\xb8\x3c\xe1\x50\x57\x85\xb4\xa3\x7e\xe8\xb8\xcc\xc4\x2e\x7f\x05\x00\x00\xff\xff\xfd\x41\xe7\x07\x25\x13\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localStorageYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5f\x6f\xdb\xb6\x16\x7f\xd7\xa7\x38\x57\xb7\x79\xb8\x17\xa5\x9d\x6c\x05\x32\xb0\xd8\x83\x9b\x38\x69\x80\xc4\x36\x6c\x77\x43\x51\x14\x06\x2d\x1d\xdb\x6c\x28\x92\x20\x29\xb7\x6a\x96\xef\x3e\x90\x94\x1d\xc9\x71\x13\x07\xdb\xde\xa6\x17\x81\x87\xe7\xef\xef\xfc\x23\xd3\xfc\x37\x34\x96\x2b\x49\x61\x7d\x92\xdc\x72\x99\x53\x98\xa0\x59\xf3\x0c\x7b\x59\xa6\x4a\xe9\x92\x02\x1d\xcb\x99\x63\x34\x01\x90\xac\x40\x0a\x42\x65\x4c\x10\xcd\xdc\x8a\x68\xa3\xd6\xdc\xcb\xa3\x21\x36\xca\x11\x56\x0b\x46\x76\xab\x59\x86\x14\x6e\xcb\x39\x12\x5b\x59\x87\x45\x42\x08\x49\x9a\x96\xcd\x9c\x65\x1d\x56\xba\x95\x32\xfc\x3b\x73\x5c\xc9\xce\xed\x2f\xb6\xc3\x55\x77\xeb\xd3\x99\x28\xad\x43\x33\x56\x02\x0f\x77\xc8\x78\x6e\x53\x0a\xb4\x34\x21\xc0\x34\xbf\x34\xaa\xd4\x96\xc2\xa7\x34\xfd\x9c\x00\x18\xb4\xaa\x34\x19\x06\x8a\x54\x39\xda\xf4\x35\xa4\xda\xbb\x65\x1d\x4a\xb7\x56\xa2\x2c\x30\x13\x8c\x17\xe1\x26\x53\x72\xc1\x97\x05\xd3\x36\x88\xaf\xd1\xcc\x83\xe8\x12\x9d\xbf\x16\xdc\x86\xff\x57\xe6\xb2\x55\xfa\xf9\x79\x93\x28\x73\xad\xb8\x74\x7b\xcd\x46\xa2\xca\x77\x6c\xfd\xff\x20\xc5\x6b\xf4\x5a\x5b\x82\x99\x41\xe6\x30\x28\xdd\xef\x9f\x75\xca\xb0\x25\xd6\xd0\x3f\x56\x5a\xdf\x67\x82\x59\x8b\x07\x22\xf0\x97\x12\xfd\x8e\xcb\x9c\xcb\xe5\xe1\xf9\x9e\x73\x99\x27\x3e\xe9\x63\x5c\x78\xe6\x4d\x78\x4f\x18\x4e\x00\x1e\x17\xd8\x21\x65\x65\xcb\xf9\x17\xcc\x5c\xa8\xac\xbd\x6d\xf3\x4f\x35\x0b\xd3\xda\x3e\xc0\x75\x8e\x5a\xa8\xaa\xc0\x17\xf4\xe9\x8f\x4d\x59\x8d\x19\x0d\x69\x8f\xbc\xef\xb9\xcf\x79\x75\xcd\x0b\xee\x28\x1c\x27\x00\xd6\x19\xe6\x70\x59\x79\x2e\x00\x57\x69\xa4\x30\x56\x42\x70\xb9\xfc\xa0\x73\xe6\x30\xd0\x4d\x93\x12\x59\x01\x0a\xf6\xed\x83\x64\x6b\xc6\x05\x9b\x0b\xa4\x70\xe2\xd5\xa1\xc0\xcc\x29\x13\x79\x0a\x5f\x35\xd7\x6c\x8e\xc2\x6e\x84\x98\xd6\x4f\x84\xe1\xb0\xd0\x62\x6b\xa2\x19\xbf\xff\x44\x4b\xd3\x73\xba\x00\x36\xd1\xfb\x4f\x1b\xae\x0c\x77\xd5\x99\x2f\xf6\x41\x00\x33\x8d\x20\x11\x3f\x27\x48\x66\xb8\xe3\x19\x13\x69\xcd\x6f\x5b\xb9\x1f\xbc\x2c\xf1\x01\x4a\x25\xd0\x84\xc2\x6c\x78\x0c\x40\xe0\x16\x2b\x0a\xe9\x59\x6d\xaf\x97\xe7\x4a\xda\xa1\x14\x55\xda\xe0\x02\x50\xda\x4b\x2b\x43\x21\xed\x7f\xe3\xd6\xd9\x74\x8f\x92\xe0\xb9\x2f\xde\x8e\x4f\xba\x91\xe8\x30\xf4\x5e\xa6\xa4\x33\x4a\x10\x2d\x98\xc4\x17\xe8\x05\xc0\xc5\x02\x33\x47\x21\x1d\xa8\x49\xb6\xc2\xbc\x14\xf8\x12\xc3\x05\xf3\x2d\xf7\x77\x59\xf4\x61\x30\x2e\xd1\x6c\x11\x24\xcf\xf5\x41\xfc\x78\xc1\x96\x48\xe1\xe8\x6e\xf2\x71\x32\xed\xdf\xcc\xce\xfb\x17\xbd\x0f\xd7\xd3\xd9\xb8\x7f\x79\x35\x99\x8e\x3f\xde\x1f\x19\x26\xb3\x15\x9a\xee\x7e\x45\x74\x7d\xdc\x39\xee\xfc\xf4\xa6\xad\x70\x54\x0a\x31\x52\x82\x67\x15\x85\xab\xc5\x40\xb9\x91\x41\x8b\xdb\x84\x7b\x7f\x8b\x82\xc9\xfc\x21\xdd\xe4\x39\x47\x09\x58\xc7\x8c\x6b\x9c\x09\x89\x3b\xa9\x41\xea\xa2\xcb\xba\x91\x5a\xff\x3a\x5f\xac\x92\x5b\x8e\xb8\x5d\x6e\x7c\xed\xd9\xa6\xed\x08\x55\x94\x20\x91\xa9\x81\x7c\xe1\xf9\x47\xcc\xad\x68\xcb\xc0\x96\x03\xe5\xfa\xb1\xb2\xd1\xf0\x7c\x36\xe8\xdd\xf4\x27\xa3\xde\x59\xbf\xa1\x6c\xcd\x44\x89\x17\x46\x15\xb4\x95\xdb\x05\x47\x91\xd7\xa3\xfb\x11\x3d\xda\xde\xf4\x78\x67\x3b\xc1\x92\x66\x54\x2f\x08\x28\xd2\x6f\x98\x6e\x5b\x7b\x54\x30\x35\xbe\xbb\x53\xb8\xbd\x2c\x1f\xe6\xf1\x24\xd2\xc3\xdc\x78\x72\x22\xfb\xf5\x24\xa5\x72\xcd\x9e\x6f\x6e\xd8\x9d\x56\xe1\x96\xe4\xb8\x60\xa5\x70\x24\x5c\x53\x48\x9d\x29\x31\x4d\x9a\x75\x08\x75\x9d\x7a\x81\x86\xa5\x18\x7b\xbd\x4d\x6f\x54\x8e\x14\x7e\x67\xdc\x5d\x28\x73\xc1\x8d\x75\x67\x4a\xda\xb2\x40\x93\x98\xf8\xd4\xd9\x14\xed\x39\x0a\x74\x18\x22\xaf\x57\xe4\x06\xb2\x64\xe7\xd9\xf8\xe4\xe6\xd9\x16\xe8\x0f\x96\xce\x46\xb0\x51\xab\x14\xfe\x20\x01\x90\xbb\x3a\x37\x61\x82\xf8\x0a\xb8\x61\x3a\xa5\x9f\x6a\xea\xdd\x36\x73\xe1\x3e\xa5\xe9\xa6\x73\x47\xbd\xe9\xfb\xd9\xc5\x70\x3c\x1b\x0c\x07\xb3\xeb\xab\xc9\xb4\x7f\x3e\x1b\x0c\xcf\xfb\x93\xf4\xf5\x83\x8c\xf7\xce\xa6\xf4\x53\x7a\x74\xb7\x91\xbb\x1e\x9e\xf5\xae\x67\x93\xe9\x70\xdc\xbb\xec\x07\x2d\xf7\x47\xe1\xa1\xe3\xbf\xfb\xfa\x1f\xcf\xf7\x61\x7d\x39\xff\xb8\xa8\x9d\xfd\xef\x7f\xba\x73\x2e\xbb\x76\x15\x4e\x5f\x57\x5c\x20\x2c\xd1\x29\xed\x2c\xa4\x05\xb5\x54\xd3\x14\x94\x8e\xed\x9b\xab\x87\x39\xc0\x2c\xc2\x2b\xa5\x1d\x70\xd9\xaa\x45\xfd\xbf\xd6\x91\xcd\xad\x12\xa5\x0b\x38\xfc\xfa\x6a\x38\x9a\xf6\xc6\x97\x2d\x86\xb7\x6f\x5b\x47\xdb\x16\xb7\xfc\x3b\x5e\xc9\x77\x95\x43\x7b\x88\x74\xd1\x96\x5e\x2b\xe1\x2b\xe7\x39\x49\xb4\x2c\xab\xe3\x93\xb1\xdb\x8a\xdb\x9c\x1b\x20\x05\x1c\x9f\x9e\x9e\x02\xd1\xf0\xea\xae\x19\x48\x04\x35\x5b\x15\x2a\x87\xd3\xe3\xe3\xdd\xdb\x6e\xa7\x13\xf6\x3c\x33\xb9\xfa\x2a\xff\x85\xfa\x49\xa8\x4d\x01\xc4\x2c\xf6\x00\xbc\x42\xa1\xd1\x8c\x54\xde\xa9\x58\x21\xb6\x28\xee\x74\xb1\x27\xc5\x46\x1f\xa9\x7c\xef\x8b\x2a\xf6\x76\xd4\x46\x74\xcd\xd4\x7c\x36\xfd\x78\x05\xef\x08\xc1\x8b\xd6\x6e\xc1\x8d\x51\x06\x73\x22\xf8\xdc\x30\x53\x91\x79\x69\xab\xb9\xfa\x46\x4f\x3a\x3f\xbf\xe9\x9c\x1c\xb8\x77\xff\x0c\x00\x00\xff\xff\x7c\x3e\x44\xe7\xec\x0e\x00\x00")

func localStorageYamlBytes() ([]byte, error) {
	return bindataRead(
		_localStorageYaml,
		"local-storage.yaml",
	)
}

func localStorageYaml() (*asset, error) {
	bytes, err := localStorageYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "local-storage.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAggregatedMetricsReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\x31\x6b\xf4\x30\x0c\xc6\xf1\xdd\x9f\x42\x78\x7e\x93\x97\x6e\xc5\x6b\x87\xee\x1d\xba\x94\x1b\x94\xf8\x21\x27\xce\xb1\x83\x24\xe7\x68\x3f\x7d\xb9\x70\xdc\x58\x68\x27\x0d\x7f\x7e\x0f\xe8\x22\x35\x27\x7a\x29\xdd\x1c\xfa\xd6\x0a\x02\x6f\xf2\x0e\x35\x69\x35\x91\x4e\x3c\x8f\xdc\xfd\xdc\x54\xbe\xd8\xa5\xd5\xf1\xf2\x6c\xa3\xb4\xff\xfb\x53\x58\xe1\x9c\xd9\x39\x05\xa2\xca\x2b\x12\xd9\xa7\x39\xd6\xc4\xcb\xa2\x58\xd8\x91\x87\x15\xae\x32\xdb\xa0\xe0\x0c\x0d\x44\x85\x27\x14\xbb\x11\xfa\x61\xfd\xb1\x30\x78\x1b\x76\xc1\x35\x51\x74\xed\x88\xbf\x71\xc8\xe2\x7f\x71\x9c\x57\xa9\x0f\xa8\xbd\xc0\x52\x18\x88\x37\x79\xd5\xd6\x37\x4b\xf4\x11\xef\x7f\xdd\x7d\x3c\x05\x22\x85\xb5\xae\x33\x8e\xbe\xb5\x6c\xf1\x1f\xc5\xda\x32\xec\xc8\x3b\x74\x3a\xd2\x02\xbf\x95\x22\x76\xdc\x2b\xfb\x7c\x8e\xa7\xf0\x1d\x00\x00\xff\xff\xe5\x1d\x7a\x17\x89\x01\x00\x00")

func metricsServerAggregatedMetricsReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAggregatedMetricsReaderYaml,
		"metrics-server/aggregated-metrics-reader.yaml",
	)
}

func metricsServerAggregatedMetricsReaderYaml() (*asset, error) {
	bytes, err := metricsServerAggregatedMetricsReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/aggregated-metrics-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthDelegatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x31\x4e\xc4\x40\x0c\x45\xfb\x39\xc5\x5c\x60\x82\xe8\x90\x3b\xa0\xa0\x5f\x24\x7a\x67\xf2\x59\x4c\x92\x71\x64\x7b\x22\xc1\xe9\xd1\x8a\x88\x06\xd8\xf6\x4b\xff\xbd\x57\x4a\x49\xbc\xc9\x0b\xcc\x45\x1b\x65\x1b\xb9\x0e\xdc\xe3\x4d\x4d\x3e\x39\x44\xdb\x30\xdf\xf9\x20\x7a\xb3\xdf\xa6\x59\xda\x44\xf9\x71\xe9\x1e\xb0\x93\x2e\x78\x90\x36\x49\x3b\xa7\x15\xc1\x13\x07\x53\xca\xb9\xf1\x0a\xca\x2b\xc2\xa4\x7a\x71\xd8\x0e\x23\xff\xf0\xc0\x4a\x17\x70\x99\xb0\xe0\xcc\xa1\x96\x4c\x17\x9c\xf0\x7a\x79\xf1\x26\x4f\xa6\x7d\xbb\x52\x90\x72\xfe\x15\xf0\xe3\xfb\x5b\xe0\x7d\x7c\x47\x0d\xa7\x54\x8e\xef\x33\x6c\x97\x8a\xfb\x5a\xb5\xb7\xf8\x27\xf7\x98\x7d\xe3\x0a\xca\x73\x1f\x51\xbe\xf9\xe9\x2b\x00\x00\xff\xff\xa5\xb5\x26\x22\x2f\x01\x00\x00")

func metricsServerAuthDelegatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthDelegatorYaml,
		"metrics-server/auth-delegator.yaml",
	)
}

func metricsServerAuthDelegatorYaml() (*asset, error) {
	bytes, err := metricsServerAuthDelegatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-delegator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x3d\x4e\x04\x31\x0c\x46\xfb\x9c\x22\x17\xf0\x22\x3a\x94\x0e\x1a\xfa\x45\xa2\xf7\x64\x3e\xc0\xcc\x8e\x13\xd9\xce\x08\x38\x3d\x1a\xb4\xfc\x34\x4b\x6f\xbf\xef\x3d\x22\x4a\xdc\xe5\x11\xe6\xd2\xb4\x64\x9b\xb8\x1e\x78\xc4\x4b\x33\xf9\xe0\x90\xa6\x87\xe5\xc6\x0f\xd2\xae\xb6\xeb\xb4\x88\xce\x25\x1f\xdb\x09\x77\xa2\xb3\xe8\x73\x5a\x11\x3c\x73\x70\x49\x39\x2b\xaf\x28\x79\x45\x98\x54\x27\x87\x6d\x30\xda\x51\x64\xe0\x19\x76\x3e\xf1\xce\x15\x25\x2f\x63\x02\xf9\xbb\x07\xd6\x64\xed\x84\x23\x9e\x76\x08\x77\xb9\xb7\x36\xfa\x3f\x26\x29\xe7\x5f\x91\x9f\x5d\xbc\x05\x74\x6f\x20\xee\xf2\x67\x1c\x1a\x52\xbf\xde\xbf\x35\x7c\x4c\xaf\xa8\xe1\x25\xd1\x19\xf4\x00\xdb\xa4\xe2\xb6\xd6\x36\x34\x2e\xa4\x5c\xd6\xff\x0c\x00\x00\xff\xff\x2a\x39\xe6\xe4\x44\x01\x00\x00")

func metricsServerAuthReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthReaderYaml,
		"metrics-server/auth-reader.yaml",
	)
}

func metricsServerAuthReaderYaml() (*asset, error) {
	bytes, err := metricsServerAuthReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsApiserviceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x4d\x6a\xc4\x30\x0c\x85\xf7\x3e\x85\x2e\xe0\x74\xb2\x2b\xde\x75\x59\x68\x61\x20\x65\xf6\x1a\x8f\x3a\x88\xe0\x1f\x24\xd9\x90\xdb\x97\x30\x71\xba\x13\x7a\xef\xfb\x24\xef\xbd\xc3\xca\x37\x12\xe5\x92\x03\x60\x65\xa1\x27\xab\x09\x1a\x97\x3c\xad\xef\x3a\x71\x79\xeb\xb3\x5b\x39\x3f\x02\x7c\x5c\x3f\x17\x92\xce\x91\x5c\x22\xc3\x07\x1a\x06\x07\x90\x31\x51\x80\x3e\xdf\xc9\x70\x9e\x12\x99\x70\xd4\x03\x76\x5a\x29\xee\x25\x7d\x81\xfb\x38\x88\xa3\xe9\xf7\x88\xe4\x0c\xb4\x62\xa4\x00\x6b\xbb\x93\xd7\x4d\x8d\x92\x03\x78\x4a\x69\xf5\x44\x86\x1c\xa0\x8f\xdf\x8f\xf3\x0e\x80\xb3\x52\x6c\x42\xcb\xca\xf5\xe7\x6b\xb9\x91\xf0\xef\x16\xc0\xa4\xd1\x10\x5d\x85\x8b\xb0\x6d\xdf\x9c\x39\xb5\x14\x60\xbe\x5c\xfe\x65\x23\x7d\xad\xff\x02\x00\x00\xff\xff\x14\x74\xa9\x1b\x25\x01\x00\x00")

func metricsServerMetricsApiserviceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsApiserviceYaml,
		"metrics-server/metrics-apiservice.yaml",
	)
}

func metricsServerMetricsApiserviceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsApiserviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-apiservice.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xdf\x4f\x1b\x47\x10\x7e\xf7\x5f\x31\x72\xc5\xe3\x61\x9b\x2a\x69\xb5\x12\x0f\x08\x9b\xa4\x12\x50\xcb\x67\x2a\xf1\x84\x96\xbd\x31\x5e\xb1\xbf\x3a\x33\xe7\x70\x45\xfc\xef\xd5\xfa\xc8\xe5\x8e\x40\x94\xaa\xc9\x3d\xce\x37\xf3\xcd\xb7\xdf\xee\xcc\x15\x45\x31\xd2\xc9\xfe\x85\xc4\x36\x06\x05\xbb\xd9\xe8\xde\x86\x4a\x41\x89\xb4\xb3\x06\x4f\x8c\x89\x75\x90\x91\x47\xd1\x95\x16\xad\x46\x00\x41\x7b\x54\xe0\x51\xc8\x1a\x2e\x18\x69\x87\xf4\x1c\xe6\xa4\x0d\x2a\xb8\xaf\x6f\xb1\xe0\x86\x05\xfd\xe8\x65\x07\x9d\x12\x4f\xba\x36\x73\x4c\x2e\x36\x1e\xff\x57\x0b\x00\xa7\x6f\xd1\x71\xae\x04\xb8\xff\x9d\x0b\x9d\xd2\x57\xe5\x9c\xd0\xe4\x0c\xc2\x9d\xcd\x52\x3e\x5a\x96\x48\xcd\xb9\xf5\x56\x14\x4c\x47\x00\x2c\xa4\x05\xef\x9a\x96\x47\x9a\x84\x0a\x56\xd1\x39\x1b\xee\xae\x52\xa5\x05\xf7\x71\xea\x47\xda\x54\x00\xaf\x1f\xae\x82\xde\x69\xeb\xf4\xad\x43\x05\xb3\x4c\x87\x0e\x8d\x44\x6a\x73\xbc\x16\xb3\x3d\xef\xe9\x7c\x5b\x29\x80\xa0\x4f\xae\xa3\xef\x3b\x93\xbf\x37\xdc\xc9\x9f\x1b\x34\xf8\x56\x0b\x80\xcf\x86\xe4\x2f\x91\x8d\x64\xa5\x39\x75\x9a\xf9\x72\xcf\x3f\x6e\xdd\x2d\x42\xac\xb0\x30\x64\xc5\x1a\xed\xc6\xcf\xf9\x3c\x78\x1e\x97\x6f\x0b\x92\xe8\x90\xb4\xd8\x18\x7a\xaa\x0a\xb8\xc7\x46\xc1\xf8\xf4\x99\xf5\xa4\xaa\x62\xe0\x3f\x83\x6b\xc6\x5d\x0e\x40\x4c\xb9\x32\x92\x82\xf1\xe2\xc1\xb2\xf0\xf8\x2b\x82\xbd\x36\x8a\x0e\x0f\xf3\x7b\xa0\x80\x82\x7c\x68\xe3\xc4\xc4\x20\x14\x5d\x91\x9c\x0e\xf8\x9d\x9c\x00\xb8\xd9\xa0\x11\x05\xe3\xcb\x58\x9a\x2d\x56\xb5\xc3\xef\x6f\xe9\x35\x0b\xd2\x8f\xe8\xb5\x8b\xae\xf6\xd8\xd9\xf5\x0b\xf8\xec\x31\xd8\x00\xe2\x13\x70\x84\x4f\x08\x46\x07\x60\xbd\x41\xd7\x40\xcd\x08\x1b\x8a\xbe\x60\x43\xf9\x8d\x81\xf5\xfa\x0e\x19\x74\xa8\x26\x91\x80\x50\x57\x45\x0c\xae\x81\x6c\x8a\xb6\x01\x89\x47\x9f\x8f\xd4\xbe\x24\xf1\xa9\xa8\x2c\x75\xea\xd0\x27\x69\xe6\x96\x14\x3c\x3e\x3d\x07\xbf\xd4\xaa\x17\xc5\xaf\xde\x3a\xb4\x22\x14\x1c\x3c\x96\xd7\xe5\x7a\x71\x71\x33\x5f\x9c\x9d\x5c\x9d\xaf\x6f\x56\x8b\x0f\x7f\x94\xeb\xd5\xf5\xd3\x01\xe9\x60\xb6\x48\x13\x6f\x89\x22\x61\x55\x0c\x99\xd4\x6e\x7a\xf8\xfe\xf0\xa8\x23\xd4\x74\x37\x78\x41\x45\x61\x90\x24\xeb\x3e\x9e\x88\x4f\x03\x84\xd1\xd4\x84\x45\x8a\x24\xc7\xb3\xe9\xd1\xbb\xe9\x00\xcd\xf7\xe6\x50\x8a\x44\xb8\x41\xca\x9d\x75\x55\x11\x32\x17\x79\xe4\xf9\xf8\xe0\x71\xb9\x5a\x9c\x2d\x56\xab\xc5\xfc\xe6\x64\x3e\x5f\x2d\xca\xf2\x66\x7d\xbd\x5c\x94\x4f\x07\xaf\xf2\xd4\x8c\xed\x90\xb0\x68\xa9\x79\xdf\x76\x90\xd8\x1e\xac\x20\xe4\xe8\xea\x3c\x0a\xc7\xb3\x77\xdc\x65\xe4\x70\x4d\x06\x7b\xa7\xcb\xc1\xbf\x6b\x64\x19\xc4\x00\x4c\xaa\x15\xcc\xa6\x53\x3f\x88\x7a\xf4\x91\x1a\x05\xbf\x4d\x2f\x6c\x07\x64\x11\x03\xbf\xda\xdb\xda\x8a\x24\xee\x55\x77\xf7\xba\x8c\x24\x99\xbb\x6f\x56\x5e\x0b\x51\xa2\x89\x4e\xc1\xfa\x74\xd9\x53\xac\x2b\x1b\x90\x79\x49\xf1\x16\xfb\x12\x33\xfd\x07\x94\xa1\xea\xa4\x65\xab\x60\x92\xab\x9a\x7f\x86\xc8\xbe\xe9\x4b\x4d\x00\x6c\xb6\x98\xd5\x7e\x5c\xaf\x97\x65\x0f\xb1\xc1\x8a\xd5\x6e\x8e\x4e\x37\x25\x9a\x18\x2a\x6e\x37\x77\x47\x88\x64\x63\xd5\x41\x47\x3d\x48\xac\xc7\x58\x4b\x87\xcd\x7a\x18\xd7\xc6\x20\xf3\x7a\x4b\xc8\xdb\xe8\xaa\x21\xba\xd1\xd6\xd5\x84\x3d\xf4\xd7\x0e\x75\x76\x87\xff\xd9\x89\x5c\xf4\x13\x8c\x78\xff\x0d\x27\x66\xd3\x9f\x6e\xc5\x7e\xe8\xf2\x2f\x24\x06\xc1\x07\x19\xbe\x66\x5d\xe5\xed\xbe\x8a\x51\xce\xac\xc3\xf6\xcf\xa2\x40\xa8\xc6\x7e\x5a\x1d\x4e\xf8\x32\x86\x9c\xf6\x3a\x78\xc5\x48\xfb\x09\xe8\x1f\x47\x3b\x17\x3f\x2d\xc9\xee\xac\xc3\x3b\x5c\xb0\xd1\x6e\xff\xc3\x51\xb0\xd1\x8e\xbf\x70\xb4\x7b\xf5\x22\x2f\xd3\x57\x26\xe3\xe5\x12\x84\x76\xed\x2e\xdb\x2b\xcb\x1b\xe6\xdf\x00\x00\x00\xff\xff\x14\xc8\x49\x02\x2c\x09\x00\x00")

func metricsServerMetricsServerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerDeploymentYaml,
		"metrics-server/metrics-server-deployment.yaml",
	)
}

func metricsServerMetricsServerDeploymentYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x3f\x4b\x04\x31\x10\xc5\xfb\x7c\x8a\x61\xfb\x28\xe2\x15\x92\xd6\x5a\x38\x50\xec\xe7\x72\x0f\x0d\x97\x4d\xc2\xcc\xec\x82\xdf\x5e\x76\xf6\x9a\x83\xed\x92\x37\xef\xcf\x2f\xc6\x18\x78\x94\x6f\x88\x96\xde\x12\xad\x2f\xe1\x56\xda\x35\xd1\x27\x64\x2d\x19\x61\x86\xf1\x95\x8d\x53\x20\x6a\x3c\x23\xd1\x0c\x93\x92\x35\x2a\x64\x85\xdc\x65\x1d\x9c\x91\xe8\xb6\x5c\x10\xf5\x4f\x0d\x73\x20\xaa\x7c\x41\xd5\x2d\x49\x7e\x91\x06\x83\x3e\x95\xfe\xbc\x37\x4d\x1f\x0f\x55\xd3\x81\x31\xd7\x45\x0d\xe2\x8e\xb2\x2d\x4c\x26\x0b\xa6\xa0\x03\x79\x2b\x56\x54\x64\xeb\x72\x1f\x79\xd3\xc8\x63\x1c\x30\x8e\x2e\xe6\x24\xd1\x9f\x89\x4e\xa7\x57\x8f\xec\x24\xbf\x66\x43\xfd\x3f\xa4\x5b\xcf\xbd\x26\xfa\x7a\x3f\xbb\x62\x2c\x3f\xb0\xb3\xa7\x76\xdf\x7f\x00\x00\x00\xff\xff\x7e\x3b\x1f\x83\x35\x01\x00\x00")

func metricsServerMetricsServerServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerServiceYaml,
		"metrics-server/metrics-server-service.yaml",
	)
}

func metricsServerMetricsServerServiceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerResourceReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x41\x4b\xc4\x30\x10\x85\xef\xf9\x15\xc3\xde\xd3\xc5\x9b\xe4\xa6\x1e\xbc\xaf\xe0\x3d\x4d\x9f\xbb\x63\xdb\xa4\xcc\x4c\x2a\xfa\xeb\xa5\xb6\x2a\xb8\xb0\x2c\x78\x4a\x98\xe4\x7d\x8f\xf9\xbc\xf7\x2e\x4e\xfc\x0c\x51\x2e\x39\x90\xb4\x31\x35\xb1\xda\xa9\x08\x7f\x44\xe3\x92\x9b\xfe\x56\x1b\x2e\xfb\xf9\xc6\xf5\x9c\xbb\x40\x0f\x43\x55\x83\x1c\xca\x00\x37\xc2\x62\x17\x2d\x06\x47\x94\xe3\x88\x40\xfa\xae\x86\x31\x8c\x30\xe1\xa4\x5e\x21\x33\xc4\x49\x1d\xa0\xc1\x79\x8a\x13\x3f\x4a\xa9\x93\x2e\x09\x4f\xbb\x9d\x23\x12\x68\xa9\x92\xb0\xcd\x72\xe9\xa0\xfb\x0d\xe0\x88\x66\x48\xbb\x3d\x1d\x61\xd7\x31\xa6\xd2\xe9\x2f\xec\x1c\xb2\x9c\x03\xeb\x7a\x79\x8b\x96\x4e\xee\x7f\x26\xee\x39\x77\x9c\x8f\xd7\x0b\x29\x03\x0e\x78\x59\xbe\x7d\xaf\x73\xa1\xd2\x11\x9d\xbb\xbf\x5c\xa0\xb5\x7d\x45\xb2\x2f\xe9\x6b\xf6\x09\x32\x73\xc2\x5d\x4a\xa5\x66\xfb\x89\xff\xc9\xad\x63\x9d\x62\x42\xa0\xbe\xb6\xf0\x2b\xdf\x7d\x06\x00\x00\xff\xff\xdb\x55\x9e\x61\x2a\x02\x00\x00")

func metricsServerResourceReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerResourceReaderYaml,
		"metrics-server/resource-reader.yaml",
	)
}

func metricsServerResourceReaderYaml() (*asset, error) {
	bytes, err := metricsServerResourceReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/resource-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rolebindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x21\xbb\x72\x38\xdc\x72\xf0\xd8\x0e\xdd\x03\xb4\x3b\x6d\xb3\x09\x6b\x59\x14\x48\x2a\x41\xfb\xeb\x0b\xa7\x6e\x82\xa4\x76\x90\xb4\xdd\x24\x41\x7c\x1f\x1f\xf9\x20\xd3\x13\x8a\x12\xa7\xca\x4b\x0d\xcd\x12\x8a\x6d\x58\xe8\x0d\x8c\x38\x2d\xbb\xff\xba\x24\xfe\xb3\xfd\xeb\x3a\x4a\x6d\xe5\xef\x63\x51\x43\x59\x71\xc4\x3b\x4a\x2d\xa5\xb5\xeb\xd1\xa0\x05\x83\xca\x79\x9f\xa0\xc7\xca\x77\xa5\xc6\x00\x99\x14\x65\x8b\x12\x86\x6b\x44\x0b\xd0\xf6\x94\x9c\x70\xc4\x15\x3e\x0f\xbf\x21\xd3\x83\x70\xc9\x17\xc8\xce\xfb\x2f\xe0\x03\x47\x5f\xd5\xb0\xaf\x0e\xfa\x99\x46\x86\x96\xfa\x05\x1b\xd3\xca\x85\x9b\x20\x8f\x8a\x32\xe3\xc2\xb9\x10\x82\xfb\xfe\xb4\x26\xc6\xf4\xd9\xfe\x3f\x0d\x0d\x27\x13\x8e\x11\xc5\x49\x89\x78\xd2\xb8\x0e\x15\xc1\x2f\x16\xce\x7b\x41\xe5\x22\x0d\x8e\x6f\x89\x5b\x54\xe7\xfd\x16\xa5\x1e\x9f\xd6\x68\x57\xd6\x42\x8f\x9a\xa1\x39\x17\x88\xa4\xb6\x3f\xec\xc0\x9a\xcd\x84\x56\x42\xdb\xb1\x74\x94\xd6\xa3\xdf\x29\xf1\x8f\x3f\x99\x23\x35\x74\x33\x61\x42\x10\x53\x9b\x99\x92\xe9\xfe\x96\xb9\x9d\xd3\x1c\xfc\x1f\xb5\x7f\xb8\xb4\xf9\x88\xcf\xec\xee\xf7\xb3\x7d\x0a\x38\x06\x7b\xf0\x78\x1d\xe3\x2c\xdc\x97\x01\xef\x01\x00\x00\xff\xff\x46\xd3\x6d\x9d\x0f\x04\x00\x00")

func rolebindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_rolebindingsYaml,
		"rolebindings.yaml",
	)
}

func rolebindingsYaml() (*asset, error) {
	bytes, err := rolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\x5f\x6f\xdb\x3a\x0c\xc5\xdf\xfd\x29\x08\x03\x79\xba\x90\xdd\xe4\xa9\xd7\x6f\xb9\xa9\x7b\x57\x6c\xeb\x8a\x38\xdd\xd0\xa7\x80\x91\x99\x58\x88\x2c\x09\x14\x1d\x2c\xeb\xfa\xdd\x07\x25\xe9\x3f\xa0\xc0\x86\x61\x7b\x13\x44\xf2\x77\xc8\x73\x94\x52\x19\x06\xf3\x99\x38\x1a\xef\x2a\xe8\xc8\xf6\x85\x46\x11\x4b\x85\xf1\xe5\x6e\x9c\x6d\x8d\x6b\x2b\x78\x47\xb6\x9f\x75\xc8\x92\xf5\x24\xd8\xa2\x60\x95\x01\x38\xec\xa9\x02\x61\xa4\xb5\xd9\x2a\xcd\xed\xe9\x2f\x06\xd4\x54\xc1\x76\x58\x91\x8a\xfb\x28\xd4\x67\x31\x90\x4e\x23\x3a\x41\x2a\xe8\x44\x42\xac\xca\x72\x74\xff\xfe\xf6\xbf\x7a\x7e\x5d\x2f\xea\x66\x39\xbd\xb9\x7a\x18\x95\x51\x50\x8c\x2e\x0f\x8d\xb1\x7c\x01\x57\x93\x71\x31\x29\xc6\xff\x0c\xe1\xf0\x38\x2b\x64\xf3\x2d\xfb\x83\x07\xfc\xbd\xe5\xdf\x5a\x1c\x20\x92\x24\x28\xc0\xc6\xfa\x15\xda\xe2\x28\x76\x41\x6b\x1c\xac\xcc\x69\x63\xa2\xf0\xbe\x82\x7c\x74\xdf\xdc\x35\x8b\xfa\xe3\xf2\xa2\xbe\x9c\xde\x7e\x58\x2c\xe7\xf5\xff\x57\xcd\x62\x7e\xb7\x9c\x4f\xbf\x3c\x8c\xf2\x0c\x60\x87\x76\xa0\x38\xf3\x4e\xc8\x49\x05\xdf\xd5\x81\x1b\x7c\x3b\x75\xce\xa7\x95\xbc\x8b\x47\x2d\x80\xc0\xbe\x27\xe9\x68\x88\xc9\xa0\xe0\xd3\x45\xf9\xf9\xd9\xf9\x24\x7f\xb3\x21\x6a\xc6\x40\x15\xe4\xc2\x03\x1d\x5b\x02\xfb\x9d\x69\x89\x9f\x90\xc9\x2b\x76\x24\x14\xaf\xdc\x86\x29\x3e\x15\x00\xc2\xb0\xb2\x26\x76\xd4\x36\xc4\x3b\xa3\xe9\xb9\x02\x40\x0e\x57\x96\xda\x14\xc0\x40\x27\xb2\xf1\x6c\x64\x3f\xb3\x18\xe3\xf5\x21\x9c\xfc\x68\x8b\xd2\x76\x88\x42\xac\x34\x1b\x31\x1a\xed\x71\x15\xd3\xe3\xe6\x89\xc9\x14\x7c\x34\xe2\x0f\xae\x31\x3a\xdd\x11\x97\xbd\x61\xf6\x4c\xad\xb2\x66\xc5\xc8\x7b\x75\x0a\xe5\xf1\x5a\xc1\x4d\x05\xf9\xa4\xf8\xb7\x18\x9f\x1d\xff\xc4\x5b\xe2\x97\x9e\x29\xd8\x52\x42\xce\x4e\xd2\xd3\xb6\xf5\x2e\x7e\x72\x76\xff\x08\xf1\x21\x4d\x78\xae\x20\xaf\xbf\x9a\x28\x31\x7f\x35\xe8\x7c\x4b\x8a\xbd\xa5\xe2\xd9\xa9\xe4\xad\xf6\x4e\xd8\x5b\x15\x2c\x3a\xfa\x09\x0b\x80\xd6\x6b\xd2\x29\xac\x6b\xdf\xe8\x8e\xda\xc1\xd2\xaf\xc9\xf4\x98\x9c\xfb\x7d\x7e\x7c\x1d\x9d\x09\x97\xd8\x1b\xbb\xbf\xf1\xd6\xe8\xa4\x7b\xc3\xb4\x26\xbe\x18\xd0\x36\x82\x7a\x9b\x67\x3f\x02\x00\x00\xff\xff\x12\x80\xc2\x85\x56\x04\x00\x00")

func traefikYamlBytes() ([]byte, error) {
	return bindataRead(
		_traefikYaml,
		"traefik.yaml",
	)
}

func traefikYaml() (*asset, error) {
	bytes, err := traefikYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "traefik.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"ccm.yaml":           ccmYaml,
	"coredns.yaml":       corednsYaml,
	"local-storage.yaml": localStorageYaml,
	"metrics-server/aggregated-metrics-reader.yaml": metricsServerAggregatedMetricsReaderYaml,
	"metrics-server/auth-delegator.yaml":            metricsServerAuthDelegatorYaml,
	"metrics-server/auth-reader.yaml":               metricsServerAuthReaderYaml,
	"metrics-server/metrics-apiservice.yaml":        metricsServerMetricsApiserviceYaml,
	"metrics-server/metrics-server-deployment.yaml": metricsServerMetricsServerDeploymentYaml,
	"metrics-server/metrics-server-service.yaml":    metricsServerMetricsServerServiceYaml,
	"metrics-server/resource-reader.yaml":           metricsServerResourceReaderYaml,
	"rolebindings.yaml":                             rolebindingsYaml,
	"traefik.yaml":                                  traefikYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"ccm.yaml":           &bintree{ccmYaml, map[string]*bintree{}},
	"coredns.yaml":       &bintree{corednsYaml, map[string]*bintree{}},
	"local-storage.yaml": &bintree{localStorageYaml, map[string]*bintree{}},
	"metrics-server": &bintree{nil, map[string]*bintree{
		"aggregated-metrics-reader.yaml": &bintree{metricsServerAggregatedMetricsReaderYaml, map[string]*bintree{}},
		"auth-delegator.yaml":            &bintree{metricsServerAuthDelegatorYaml, map[string]*bintree{}},
		"auth-reader.yaml":               &bintree{metricsServerAuthReaderYaml, map[string]*bintree{}},
		"metrics-apiservice.yaml":        &bintree{metricsServerMetricsApiserviceYaml, map[string]*bintree{}},
		"metrics-server-deployment.yaml": &bintree{metricsServerMetricsServerDeploymentYaml, map[string]*bintree{}},
		"metrics-server-service.yaml":    &bintree{metricsServerMetricsServerServiceYaml, map[string]*bintree{}},
		"resource-reader.yaml":           &bintree{metricsServerResourceReaderYaml, map[string]*bintree{}},
	}},
	"rolebindings.yaml": &bintree{rolebindingsYaml, map[string]*bintree{}},
	"traefik.yaml":      &bintree{traefikYaml, map[string]*bintree{}},
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

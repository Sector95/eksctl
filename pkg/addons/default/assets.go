// Code generated by go-bindata.
// sources:
// assets/aws-node-1.10.yaml
// assets/aws-node.yaml
// assets/coredns.yaml
// DO NOT EDIT!

package defaultaddons

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

var _awsNode110Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x6f\xdb\x36\x10\x7e\xd7\x5f\x71\x70\x1f\x0a\x0c\x95\x6c\x67\x59\x96\xe9\xcd\x75\xbc\xac\x58\xe2\x1a\xf1\x9a\x62\x28\x02\x83\xa6\x2e\x32\x67\x8a\xe4\xc8\xa3\x1d\xf7\xaf\x1f\x28\xf9\x87\x14\xcb\xee\x36\x0c\x18\x5f\x2c\xf3\xee\xbe\xe3\x7d\xfc\xc8\x63\x1c\xc7\x11\x33\xe2\x11\xad\x13\x5a\xa5\x60\xe7\x8c\x27\xcc\xd3\x42\x5b\xf1\x95\x91\xd0\x2a\x59\x5e\xbb\x44\xe8\xee\xaa\x1f\xbd\x81\xa5\x9f\xa3\x55\x48\xe8\x60\x55\x85\x38\x98\xe3\xb3\xb6\x08\xfd\xe4\x3a\xe9\x81\x5b\x68\x2f\x33\xf0\x0e\xcf\x42\xcd\x91\x58\x3f\x5a\x0a\x95\xa5\x30\x94\xde\x11\xda\x07\x2d\x31\x2a\x90\x58\xc6\x88\xa5\x11\x80\x62\x05\xa6\xc0\xd6\x2e\x56\x3a\xc3\xc8\x7a\x89\x2e\x8d\x62\x60\x46\xdc\x5a\xed\x8d\x0b\x4e\x31\x70\x9b\x95\xb8\xac\x60\x5f\xb5\x62\x6b\x97\x70\x5d\x44\x00\x16\x9d\xf6\x96\xe3\xd6\xad\xf3\x5d\xa7\xfc\x0d\xa8\xce\x30\x8e\x2e\x82\x50\xc3\xbc\x66\xaf\x63\xc3\x97\x4e\xe7\xe9\x18\xc6\xe8\xcc\x55\x38\x3a\x43\x77\x0a\x11\xbe\x74\xa4\x70\xd4\x79\x07\x9d\x35\x23\xbe\x08\x1f\x39\x52\xe7\xe9\x75\x0a\x7c\x21\x54\x25\x8d\x6d\xc9\x32\x86\x85\x56\x0e\xe9\x0c\xf2\x53\xf4\x7a\x0b\x57\x3b\x62\xa7\x68\x57\x82\xe3\x80\x73\xed\x15\x9d\xe3\x16\x0e\x45\xa4\xe5\x1e\xc7\x6e\xe3\x08\x8b\x23\xec\xff\x57\x1e\xef\x85\xca\x84\xca\xcf\xaa\x44\x4b\x7c\xc0\xe7\x60\xd9\x11\x7d\x66\xd5\x11\xc0\xb1\x06\x8f\x30\x9d\x9f\xff\x81\x9c\x4a\xf1\xb5\x32\xfb\xcf\xf8\xac\x20\x6e\xca\xbd\x9d\x22\x35\xf8\x65\xc6\xb8\xbf\x41\xe5\x4f\x4d\x2a\x0f\x2a\xda\x73\xf7\x6f\x36\x1b\x40\xb2\x39\xca\x52\x7c\x00\xcb\x6b\x17\x33\x63\xea\x3c\x18\xe4\xc1\xe6\x4d\xc6\x08\xa7\x64\x19\x61\xbe\xa9\xbc\x69\x63\x30\x85\x07\x2d\xa5\x50\xf9\xa7\xd2\x21\x02\x70\x28\x91\x93\xb6\x95\x4f\x11\x04\x7b\x57\x4b\xd1\x96\x04\x80\xb0\x30\x92\x11\x6e\x83\x6a\x85\x84\x21\x1b\xf1\xed\x08\x61\x30\xa5\x34\x95\x7b\x5d\x73\x76\x7c\x81\x99\x97\x68\x13\x26\xcd\x82\x25\x07\x92\x83\xee\xb8\x15\x24\x38\x93\xb1\xd1\x59\x0a\x6f\xdf\x96\x61\xbb\xa2\xcb\xef\xc6\xb6\x8f\x5f\xd3\x1a\xc6\x42\x3b\x1a\x23\xad\xb5\x5d\xa6\x40\xd6\xef\xe6\x49\x4b\xb4\xcd\xe5\xc4\xa0\x4d\x98\xd3\x36\x85\xd1\x8b\x70\xe5\x29\x0f\x83\x6b\x45\x4c\x28\xb4\x35\x57\x51\xb0\x1c\x53\xb8\xea\x5d\x5c\xf6\xfa\xfd\xcb\xef\x2f\x7f\xb8\x48\xb2\xa5\x4d\x90\xdb\xc4\xbb\x78\x8d\x8e\xe2\x8b\xe6\x1d\xd8\xad\xfe\xc5\x81\x21\xae\x44\xba\xea\x27\x97\x49\x6f\xcf\x45\x89\x38\xf1\x52\x4e\xb4\x14\x7c\x93\xc2\x40\xae\xd9\xc6\xed\xed\x46\x5b\xaa\x51\x17\x1f\x96\x35\xd1\x96\x52\xb8\xea\x5f\xfd\x78\xbd\x37\xef\x54\x56\x20\x59\xc1\x0f\x28\x47\xda\xab\x06\xaa\x55\x5a\x8b\x8d\xb7\x7e\x83\xcf\xd3\xd9\xe3\x64\x38\xfb\xf5\x7a\x3a\x1b\x8e\x3f\xcc\xee\x3e\xde\xde\x8d\x1e\x47\x77\x35\x57\x80\x15\x93\x1e\x53\xb8\x19\xbd\xff\x74\xdb\x82\x71\xff\xfb\x6c\xfc\xf1\x66\x34\x1b\x0f\xee\x47\xc7\x71\x3f\x5b\x5d\xa4\x8d\x69\x80\x67\x81\x32\xdb\x5e\x1a\x2d\x96\x09\xa3\x45\x5a\xea\x20\x09\x35\x84\x6d\x6f\x49\xfb\x79\xf0\xdb\xf0\x97\x32\xe9\x74\x32\x18\xfe\x97\x99\x77\x27\x20\xd9\x1f\xdb\xbd\x77\xa3\x5f\x1c\x26\xff\xf4\xe8\xc8\x35\x41\xb9\xf1\x29\xf4\x7b\xc5\xe1\x2c\x20\xf7\x56\xd0\x66\xa8\x15\xe1\x0b\xd5\xbd\x8d\x15\x2b\x21\x31\xc7\xac\xa1\x61\x80\x95\x96\xbe\xc0\xfb\xa0\xfe\x86\x34\x8a\x30\x53\xad\xb6\x1b\x4e\x40\x57\x1b\xea\x72\x25\xba\x73\xa1\x8e\x24\xc2\x95\x88\xe7\x42\xc5\x99\xb0\xe7\x20\x90\x78\x09\xa1\x90\x92\xac\x15\x44\x21\x7d\x0b\x64\xc5\x6c\x57\xea\xfc\x28\x5c\xea\xfc\x4c\x68\x88\xb2\x5e\x75\x33\xcd\x97\x68\x13\xa7\xf9\xf2\x08\xa1\xb2\xd5\x4c\x15\x37\xb5\x23\x7b\xba\xda\xb0\xb4\x32\x55\x9d\xf3\x2a\xf5\x31\x71\xf1\x99\x8a\xcf\x00\xb5\xd1\x17\x9f\xa8\xfe\x0c\x4c\x93\xc0\xf8\x54\xf1\xdf\xc4\x78\x4d\xe7\xeb\x87\x05\x33\xe2\xd0\xc5\x4e\x3c\x04\xbc\x23\x5d\x3c\x6c\x25\x7f\x83\xcf\x42\x89\x70\xa1\xb6\xf4\x3a\x54\x82\x6b\xf5\x2c\x72\x97\xb4\x3f\x0f\x77\xb7\xba\xe3\x3a\xf4\xad\x6d\xfb\x8f\x00\xf2\xea\xc5\x70\xea\x51\xf9\x06\xfa\x49\xbf\x17\x9a\xae\x83\xce\xb6\x2f\x77\xde\x01\x93\x12\x14\xae\xd1\x96\xed\x78\x67\x70\x9d\xea\xd9\xb6\x7b\x96\x95\x3d\xa7\xbf\xeb\xbf\x15\x51\x46\x7a\xcb\x64\x7d\xc5\x55\xd7\x11\x2a\xf7\x92\xd9\x9a\xa1\x6a\xca\x25\x13\xa3\xf1\x87\x61\x35\xf7\x57\x00\x00\x00\xff\xff\x41\x46\x08\xc0\xbf\x0b\x00\x00")

func awsNode110YamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNode110Yaml,
		"aws-node-1.10.yaml",
	)
}

func awsNode110Yaml() (*asset, error) {
	bytes, err := awsNode110YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node-1.10.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _awsNodeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x8f\x1a\x37\x10\x7e\xdf\xbf\x62\x44\x1f\x22\x55\xdd\xe5\x48\xaf\xe9\x75\xdf\x08\x47\xd3\xa8\x77\x04\x1d\xcd\x45\x55\x74\x42\xc6\x3b\x07\x2e\x5e\xdb\xb5\xc7\x70\xe4\xaf\xaf\xec\xe5\xc7\x2e\x2c\xa4\xad\x2a\xd5\x2f\xb0\x9e\x99\x6f\x3c\x9f\x3f\x7b\x9c\xa6\x69\xc2\x8c\x78\x44\xeb\x84\x56\x39\xd8\x19\xe3\x19\xf3\xb4\xd0\x56\x7c\x61\x24\xb4\xca\x96\x37\x2e\x13\xba\xbb\xea\x25\xdf\xc0\xd2\xcf\xd0\x2a\x24\x74\xb0\xaa\x42\x1c\xcc\xf0\x59\x5b\x84\x5e\x76\x93\x5d\x81\x5b\x68\x2f\x0b\xf0\x0e\x2f\x42\xcd\x90\x58\x2f\x59\x0a\x55\xe4\x30\x90\xde\x11\xda\x07\x2d\x31\x29\x91\x58\xc1\x88\xe5\x09\x80\x62\x25\xe6\xc0\xd6\x2e\x55\xba\xc0\xc4\x7a\x89\x2e\x4f\x52\x60\x46\xbc\xb3\xda\x1b\x17\x9c\x52\xe0\xb6\x88\xb8\xac\x64\x5f\xb4\x62\x6b\x97\x71\x5d\x26\x00\x16\x9d\xf6\x96\xe3\xd6\xad\xf3\x6d\x27\xfe\x06\x54\x67\x18\x47\x97\x40\xa8\x61\x56\xb3\xd7\xb1\xe1\x73\xa7\xf3\x74\x0a\x63\x74\xe1\x2a\x1c\x5d\xa0\x3b\x87\x08\x9f\x3b\x52\x38\xea\x7c\x07\x9d\x35\x23\xbe\x08\x7f\xe6\x48\x9d\xa7\xe3\x14\xf8\x42\xa8\x22\x8d\x6d\xc9\x0a\x86\xa5\x56\x0e\xe9\x02\xf2\x53\x72\xbc\x85\xab\x1d\xb1\x13\xb4\x2b\xc1\xb1\xcf\xb9\xf6\x8a\x2e\x71\x0b\x87\x22\xf2\xb8\xc7\xa9\xdb\x38\xc2\xf2\x04\xfb\xff\x95\xc7\x5b\xa1\x0a\xa1\xe6\x17\x55\xa2\x25\x3e\xe0\x73\xb0\xec\x88\xbe\xb0\xea\x04\xe0\x54\x83\x27\x98\xce\xcf\xfe\x40\x4e\x51\x7c\xad\xcc\xfe\x33\x3e\x2b\x88\xdb\xb8\xb7\x13\xa4\x06\xbf\xcc\x18\xf7\x37\xa8\xfc\xa9\x49\xe5\x41\x45\x7b\xee\xfe\xcd\x66\x03\x48\x36\x43\x19\xc5\x07\xb0\xbc\x71\x29\x33\xa6\xce\x83\x41\x1e\x6c\xde\x14\x8c\x70\x42\x96\x11\xce\x37\x95\x37\x6d\x0c\xe6\xf0\xa0\xa5\x14\x6a\xfe\x31\x3a\x24\x00\x0e\x25\x72\xd2\xb6\xf2\x29\x83\x60\xef\x6a\x29\xda\x92\x00\x10\x96\x46\x32\xc2\x6d\x50\xad\x90\x30\x64\x23\xbe\x1d\x21\x0c\xa6\x94\xa6\xb8\xd7\x35\x67\xc7\x17\x58\x78\x89\x36\x63\xd2\x2c\x58\x76\x20\x39\xe8\x8e\x5b\x41\x82\x33\x99\x1a\x5d\xe4\xf0\xea\x55\x0c\xdb\x15\x1d\x86\xb1\x42\x5b\x41\x9b\x81\x64\xce\x8d\x22\xab\x15\x75\x31\x71\xba\x8b\xdf\x7a\xbb\x86\x48\x46\xc7\x9b\x10\xc6\x42\x3b\x1a\x21\xad\xb5\x5d\xe6\x40\xd6\xef\xe6\x49\x4b\xb4\xcd\xc5\xa7\xa0\x4d\x98\xd3\x36\x87\xe1\x8b\x70\xf1\x4e\x08\x83\x6b\x45\x4c\x28\xb4\x35\x57\x51\xb2\x39\xe6\xf0\xe6\xea\xf5\xf5\x55\xaf\x77\xfd\xfd\xf5\x0f\xaf\xb3\x62\x69\x33\xe4\x36\xf3\x2e\x5d\xa3\xa3\xf4\x75\xf3\xc6\xec\x56\x5f\x69\xe0\x93\x2b\x91\xaf\x7a\xd9\x75\x76\xb5\x67\x2e\x22\x8e\xbd\x94\x63\x2d\x05\xdf\xe4\xd0\x97\x6b\xb6\x71\x7b\xbb\xd1\x96\x6a\x44\xa7\x87\x65\x8d\xb5\xa5\x1c\xde\xf4\xde\xfc\x78\xb3\x37\xef\x34\x59\x22\x59\xc1\x0f\x28\x27\x4a\xad\x06\xaa\x55\x5e\x8b\x4d\xb7\x7e\xfd\x4f\x93\xe9\xe3\x78\x30\xfd\xf5\x66\x32\x1d\x8c\xde\x4f\xef\x3e\xbc\xbb\x1b\x3e\x0e\xef\x6a\xae\x00\x2b\x26\x3d\xe6\x70\x3b\x7c\xfb\xf1\x5d\x0b\xc6\xfd\xef\xd3\xd1\x87\xdb\xe1\x74\xd4\xbf\x1f\x9e\xc6\xfd\x6c\x75\x99\x37\xa6\x01\x9e\x05\xca\x62\x7b\xc5\xb4\x58\xc6\x8c\x16\x79\x54\x4d\x16\x6a\x08\xdb\xde\x92\xf6\x53\xff\xb7\xc1\x2f\x31\xe9\x64\xdc\x1f\xfc\x97\x99\x77\xe7\x25\xdb\x1f\xf2\xbd\x77\xa3\xbb\x1c\x26\xff\xf4\xe8\xc8\x35\x41\xb9\xf1\x39\xf4\xae\xca\xc3\xc9\x41\xee\xa3\xf4\xb5\x22\x7c\xa1\xba\xb7\xb1\x62\x25\x24\xce\xb1\x68\x68\x18\x60\xa5\xa5\x2f\xf1\x3e\xa8\xbf\x21\x8d\x32\xcc\x54\xab\xed\x86\x13\xd0\xd5\x86\xba\x5c\x89\xee\x4c\xa8\x13\x89\x70\x25\xd2\x99\x50\x69\x21\xec\x25\x08\x24\x1e\x21\x14\x52\x56\xb4\x82\x28\xa4\xaf\x81\xac\x98\xed\x4a\x3d\x3f\x09\x97\x7a\x7e\x21\x34\x44\x59\xaf\xba\x85\xe6\x4b\xb4\x99\xd3\x7c\x79\x82\x50\xd9\x6a\xa6\x8a\x9b\xda\x91\x3d\x5f\x6d\x58\x5a\x4c\x55\xe7\xbc\x4a\x7d\x4a\x5c\x7a\xa1\xe2\x0b\x40\x6d\xf4\xa5\x67\xaa\xbf\x00\xd3\x24\x30\x3d\x57\xfc\x57\x31\x8e\xe9\x3c\x7e\x86\x30\x23\x0e\x3d\xef\xcc\xb3\xc1\x3b\xd2\xe5\xc3\x56\xf2\xb7\xf8\x2c\x94\x08\x17\x6a\x4b\x67\x44\x25\xb8\x56\xcf\x62\xee\xb2\xf6\xc7\xe4\xae\x07\x38\xae\x43\x97\xdb\x3e\x16\x12\x80\x79\xf5\xbe\x38\xf7\x04\xdd\x35\xee\xaa\xca\x1d\x1d\xab\x5e\x6c\x3e\xbd\x5a\x9b\x38\x3a\x3a\x8e\xb4\x8d\x17\xf8\x76\x2e\x1e\xe5\x0a\xc4\x48\x6f\x99\xac\xaf\xb9\xea\x52\x42\xcd\xbd\x64\xb6\x66\xa8\x9a\x78\xe4\x62\x38\x7a\x3f\xa8\xe6\xfe\x0a\x00\x00\xff\xff\x9f\x0c\x9d\x07\xef\x0b\x00\x00")

func awsNodeYamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNodeYaml,
		"aws-node.yaml",
	)
}

func awsNodeYaml() (*asset, error) {
	bytes, err := awsNodeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5b\x6f\xdb\x3a\x12\x7e\xf7\xaf\x20\xf4\xbc\x52\xec\x26\xe9\x66\xf9\xd6\xda\xde\xae\x81\xd6\x35\xe2\xa4\x2f\x8b\x45\x30\xa6\xc6\x16\xd7\x14\xc9\xe5\xc5\xb5\xbb\xe7\xfc\xf7\x03\xea\x66\x4a\x71\x8a\x14\x2d\x70\x8e\x9e\x24\x72\xf8\xcd\x70\x2e\xdf\x8c\x40\xf3\x2f\x68\x2c\x57\x92\x92\xc3\x64\xb4\xe7\x32\xa7\x64\x8d\xe6\xc0\x19\xbe\x63\x4c\x79\xe9\x46\x25\x3a\xc8\xc1\x01\x1d\x11\x22\xa1\x44\x4a\x98\x32\x98\x4b\xdb\x7c\x5b\x0d\x0c\x29\xd9\xfb\x0d\xa6\xf6\x64\x1d\x96\x23\x42\x04\x6c\x50\xd8\x70\x84\x10\xdc\xdb\x0c\x4a\xf8\xa6\x24\x7c\xb5\x19\x53\xe5\x15\x53\xa5\x56\x12\xa5\x8b\xb1\x08\xd9\xdf\xd9\x14\xb4\x6e\xb0\xc2\x6a\x9a\xa6\xa3\xd8\x46\xb3\x01\x96\x81\x77\x85\x32\xfc\x1b\x38\xae\x64\xb6\xbf\xb3\x19\x57\x57\x87\xc9\x06\x1d\xb4\x57\x98\x0a\x6f\x1d\x9a\x7b\x25\xb0\x67\x7f\x6c\x56\x50\x62\x24\x3a\xac\xce\x6f\x94\x72\xd6\x19\xd0\x9a\xcb\x5d\xad\x28\xcd\x71\x0b\x5e\x38\xfb\xb3\xb7\x68\xfd\x56\x7b\x87\xb6\xc2\xc6\x0b\xb4\x74\x94\x12\xd0\xfc\x83\x51\x5e\x57\x86\xa5\x24\x49\x46\x84\x18\xb4\xca\x1b\x86\xcd\x1a\xca\x5c\x2b\x2e\x2b\x5b\x52\x62\xeb\x08\xd5\x1f\x5a\xe5\xf5\x4b\x17\x8c\xf0\x79\x40\xb3\x69\xce\x0a\x6e\x5d\xf5\xf2\x15\x1c\x2b\x5e\xa7\x4f\xaa\x7c\x08\xb3\x43\xf7\x2b\xe2\xf1\x9e\xcb\x9c\xcb\x5d\x2f\x2c\x20\xa5\x72\xd5\xf1\x26\x36\x97\x70\x7b\xe1\x02\xef\x94\xd7\x39\x38\xa4\x24\x71\xc6\x63\xf2\x97\x8b\xae\x12\x78\x8f\xdb\xea\x7a\x8d\xbf\xbf\xe3\xaf\x11\x21\xcf\x33\xf7\x05\x64\xeb\x37\xff\x45\xe6\xaa\xd4\xb9\x58\xb1\xaf\xae\xd3\x61\x38\x3b\x0a\x98\x2a\xb9\xe5\xbb\x4f\xa0\xff\xd4\xea\x6f\xf5\x4e\x95\xc1\x2d\x17\x48\xc9\x6f\x95\x64\x46\x6f\xaf\xc9\xff\xab\xd7\x4a\x83\x31\xca\xd8\xee\xb3\x40\x10\xae\xe8\x3e\xcf\x89\x40\x58\xed\xdb\x4c\x28\x06\x22\x02\x20\x55\x0d\x11\x2e\x2d\x32\x6f\x30\x5a\xf7\xda\x3a\x83\x50\x46\x4b\x5b\x10\xc2\x15\x46\xf9\x5d\x41\xb8\x4c\x21\xcf\x4d\x06\x46\x03\xe1\xfa\x6d\xf5\xd2\xc9\xfe\xde\xbd\x69\xa3\x4a\x74\x05\x7a\x4b\xe8\x3f\x26\xb7\xd7\xf1\xc6\xf1\x44\x32\x72\x85\x8e\x5d\x85\x12\x14\x87\x8c\x29\xb9\xed\x04\x18\xb0\x02\xc9\xf5\x78\x54\x03\x0e\x03\x86\x47\x87\x32\xbc\xda\x41\xc1\xcd\x50\x0b\x75\x2a\xf1\x57\xf0\xf7\xa5\x94\x1f\x16\x58\x0d\x9c\x84\x48\xcd\x96\xeb\xe4\xf5\x91\xb7\x1a\x19\xad\xf8\x47\x0b\xce\xc0\x52\xf2\x66\x44\x48\xa8\x55\x87\xbb\x53\x6d\x80\x3b\x69\xa4\xe4\x5e\x09\xc1\xe5\xee\xb1\xaa\xfa\x9a\x25\xe2\x15\xda\xf8\xac\x84\xe3\xa3\x84\x03\x70\x01\x9b\x90\x32\x93\x00\x87\x02\x99\x53\xa6\x96\x29\x03\x0d\x7e\x8c\x2e\xf8\xd2\x15\x5f\x9d\xbc\x0e\x4b\x2d\x3a\x1b\x62\x87\x87\x47\xf4\x54\xbd\xac\xec\x07\x6a\xa5\xf5\x5a\xf5\xde\x2b\xfe\xe5\x20\xc2\x75\x96\x71\x65\xb8\x3b\x4d\x05\x58\xbb\x8c\x28\x25\x0d\x34\x9f\x32\xc3\x1d\x67\x20\x1a\x69\xd8\x6e\xb9\xe4\xee\x74\x36\x38\x48\xbd\x7b\xb6\x1a\x62\xf6\x3f\xcf\x0d\xe6\x33\x6f\xb8\xdc\xad\x59\x81\xb9\x0f\x01\x59\xec\xa4\xea\x96\xe7\x47\x64\x3e\x30\x5d\x7c\xb2\xc6\x5c\x37\x61\x79\x40\x53\xda\xfe\x76\x5a\x47\x69\x7e\xd4\x06\xad\x3d\x37\x86\x58\x62\x8f\x27\x4a\x92\x90\xf5\x83\xe6\xa0\x6c\x32\x10\x26\x44\x69\x34\x10\x52\x80\x2c\xe4\xb3\xcd\x03\x08\x8f\xcf\x34\xd4\xbd\x53\xfa\xe3\xeb\x35\x83\x61\xc5\x2f\xd3\x0d\x65\xfe\xf6\xa6\x59\x77\x4a\x04\x8c\xd8\x11\xad\x19\xd3\x26\x7c\xef\xf2\x5c\x49\xfb\x59\x8a\xd3\xd9\x82\xb3\xe6\x64\x7e\xe4\xd6\x75\x8e\x61\x4a\x3a\xe0\x12\x4d\x04\x37\x24\x87\xfa\xe1\x25\xec\x90\x92\xb7\xe3\x37\x37\xe3\xc9\xe4\xe6\xfa\xe6\xf6\x4d\x96\xef\x4d\x86\xcc\x64\xf7\xf3\x0f\x8b\xcf\xcb\x41\xca\xe2\xde\x5e\x35\x20\xf4\x30\xc9\x26\xd9\x75\x1f\x6b\xe5\x85\x58\x29\xc1\xd9\x89\x92\xc5\x76\xa9\xdc\xca\xa0\xc5\xaa\x6d\xd5\x4f\x6f\x14\x69\x1f\xc1\x4b\xee\x06\x6e\x2a\xb1\x54\xe6\x44\xc9\xe4\xef\xe3\x4f\x7c\x90\x97\x68\x87\xd2\x4c\x7b\x4a\x26\xe3\x71\x79\x11\xa3\x07\x01\x66\x67\x29\xf9\x37\x49\xd2\x40\xc6\xc9\xdf\x48\x52\x11\x74\x73\xab\xab\xb6\x1f\x25\xe4\x3f\xdd\x91\x83\x12\xbe\xc4\x4f\xa1\x04\x23\xbd\x67\xa7\x86\x7e\x9a\xd6\x42\x91\xfe\x32\xc8\xaf\xc0\x15\x94\xc4\x1a\x7a\x77\x81\x3c\xc4\x94\x92\x30\xe5\x9c\x1b\x87\x32\x7d\x3d\x5d\x40\x57\xca\x38\x4a\xa2\x1e\xd3\xb2\x7e\x1f\x57\x1b\xe5\x14\x53\x82\x92\xc7\xd9\xea\x47\x71\x52\xc7\xf4\x45\xac\x87\xe9\x77\xb0\x7a\x9d\xaf\x45\x2b\xd1\x19\xce\x2e\x5b\x16\xa3\x55\xad\x39\x70\x98\x92\x0e\x8f\x2e\x0e\x2d\x08\xa1\xbe\xae\x0c\x3f\x70\x81\x3b\x9c\x5b\x06\xa2\xaa\x14\x1a\x7a\xb5\x8d\xdd\xcd\x40\xc3\x86\x0b\xee\xf8\xb0\xe2\x20\xcf\x87\x04\xb4\x9c\x3f\x3c\xbd\x5f\x2c\x67\x4f\xeb\xf9\xfd\x97\xc5\x74\xde\xdb\xce\x8d\xd2\xc3\x03\x20\xc4\x85\xc0\xdd\x2b\xe5\xfe\xc9\x05\x36\x43\x5c\x3f\x8c\x82\x1f\x50\xa2\xb5\x2b\xa3\x36\x18\xe3\x15\xce\xe9\x0f\xe8\xfa\x2a\x74\x9d\x28\x83\x01\x87\x34\xe9\x40\xc9\xdd\xf8\x6e\xdc\x5b\xb6\xac\xc0\xe0\xe4\x7f\x3d\x3c\xac\xa2\x8d\x40\xe4\x1c\xc4\x0c\x05\x9c\xd6\xc8\x94\xcc\x6d\x28\xf0\x48\xc2\xf1\x12\x95\x77\xdd\xe6\x6d\xb4\x67\x3d\x63\x68\xed\x43\x61\xd0\x16\x4a\xe4\x75\x8b\x6d\x9f\x2d\x70\xe1\x0d\x46\xbb\xed\xd9\x5c\xda\xb6\xec\x67\xf5\xe8\xdd\x6c\xd4\x55\xf1\x03\x55\xc3\xda\xe9\x74\xd0\x52\x2e\xf2\x57\x75\x61\x87\xcf\x1b\x4c\xc5\x9e\x6d\x29\x0f\xe8\xb7\xf6\x74\xb7\xf9\xe2\x9c\xdc\x0c\xde\x17\x66\xac\xc1\xef\xc1\xc5\x21\xeb\xd9\x6f\xcf\x79\x4e\x0c\xcd\xa4\x0e\x6a\x12\xca\x26\xb9\xb0\x6d\x99\x01\xfd\xe2\xef\xcf\x2b\x66\xb6\x66\x1c\x4e\x9b\x01\x22\x42\xfa\xe9\xe9\xae\xd3\xda\x0e\x2a\xfd\x09\xec\x92\x75\x8d\x35\x8b\x15\x25\xb3\xe5\xfa\x69\xfa\xf1\x71\xfd\x30\xbf\x7f\x5a\x84\xc4\xed\xd8\x2e\x1d\x70\x99\x8e\x49\x6a\x48\x69\xe9\x05\xc2\x7a\xe1\x40\x60\x9a\x3f\x02\x00\x00\xff\xff\x17\xfe\xaa\x4f\x0c\x11\x00\x00")

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
	"aws-node-1.10.yaml": awsNode110Yaml,
	"aws-node.yaml": awsNodeYaml,
	"coredns.yaml": corednsYaml,
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
	"aws-node-1.10.yaml": &bintree{awsNode110Yaml, map[string]*bintree{}},
	"aws-node.yaml": &bintree{awsNodeYaml, map[string]*bintree{}},
	"coredns.yaml": &bintree{corednsYaml, map[string]*bintree{}},
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


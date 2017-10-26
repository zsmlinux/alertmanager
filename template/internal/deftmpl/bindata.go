// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x08\x0e\x56\xf1\x59\x02\xff\x31\x00\xed\x1b\x6b\x73\xda\xc6\xf6\xbb\x7e\xc5\xa9\x32\x77\x1a\xcf\x20\xb0\xf3\x9a\xfa\x81\xef\x10\x2c\xc7\x4c\x31\x78\x00\x27\xcd\x74\x3a\x1e\x21\x2d\xb0\x89\xd0\xaa\xda\x95\x31\xcd\xed\x7f\xbf\xe7\xac\xc4\x43\x20\x30\xf6\xa4\x36\xb9\x97\xa6\x69\xd9\xa3\x3d\xef\xd7\xae\x76\xf5\xed\x1b\x78\xac\xc7\x03\x06\xe6\xcd\x8d\xe3\xb3\x48\x0d\x9d\xc0\xe9\xb3\xc8\x84\xbf\xff\xae\xd0\xf8\x32\x19\x7f\xfb\x06\x2c\xf0\x10\x68\x7c\x5b\x85\x72\xdd\xaa\x13\x16\x3e\x2f\xda\x77\x8a\x45\x81\xe3\x23\x08\x21\xa5\x17\x25\x3d\x4f\xfe\x3b\x62\x2e\xe3\xb7\x2c\x2a\xd3\xa4\x56\x3a\x48\x70\x52\xea\x59\xf2\x32\xee\x7e\x61\xae\x22\xb2\xbf\x13\x4a\x5b\x39\x2a\x96\xf0\x1f\x50\xe2\x3a\x0c\x27\xa8\xbc\x07\xec\xcf\xe9\x43\xb3\xc7\x23\x1e\xf4\x09\xe7\x88\x70\xb4\x16\xb2\x78\xae\xa1\x88\xea\xb3\x60\x9e\xe3\x1f\x40\x93\x3e\x44\x22\x0e\xeb\x4e\x97\xf9\xb2\xd8\x16\x91\x62\xde\x95\xc3\x23\x59\xfc\xe8\xf8\x31\x23\x86\x5f\x04\x0f\xc0\x04\xa2\x0a\x09\xcb\xbe\x82\x97\x44\xab\x58\x15\xc3\xa1\x08\x12\xe4\xbd\x14\x36\x47\x6f\x0f\x51\x5e\x22\xca\x88\xab\x41\x76\x32\x5a\x60\x28\x6e\x59\x96\x7b\xc3\x19\x22\xc3\xc4\x8c\x79\xdc\xa7\x82\xef\x4d\x7f\xad\xf0\x8d\xc7\xa4\x1b\xf1\x50\x71\x11\x98\x6b\x6c\xac\xd8\x9d\x4a\xfc\x78\xe3\x73\xa9\xd2\xa9\x91\x13\xf4\x51\x32\x1c\x24\x72\x1d\x19\x33\xe0\xb2\x9d\xc8\x2a\x96\x36\x24\x89\x4f\xa3\x32\x4c\x15\x48\x05\x4b\x98\x57\x82\x40\xa0\x9f\x50\xa6\x0c\xc9\x39\xf0\xe3\xe8\xb6\x45\x1c\xb9\xec\x28\x71\x26\x0b\x58\xe4\x28\x11\x25\xe1\x67\xe4\x18\x2a\x63\x03\xe9\x3b\xee\xd7\x22\x8e\x9c\xd8\x57\x45\xc5\x95\xcf\x52\x2b\x28\x36\x0c\x7d\x47\x65\x63\xb1\xb8\xca\xe4\x59\x3a\xb1\xa4\x14\x18\xe6\x91\xca\x26\xda\x86\xf4\x7a\x8e\xef\x77\x11\xb0\x44\x2f\x57\x7c\x22\x8a\x81\x73\xdf\x44\x9f\x07\x5f\x37\x96\x20\x8c\x18\x05\x8b\xb9\xd9\xec\x39\xfa\x6b\x0d\xa0\xcb\xc6\x86\x12\x70\x57\x04\x98\x33\x5f\xb8\xb9\xf9\xfc\x38\xf2\x37\x95\x78\x49\xb9\x4c\x98\x0c\x78\xe8\x0e\x1c\x35\x73\x48\x24\x86\x8f\x77\xee\x22\x35\xcc\x7a\x89\x28\x9b\x07\x5e\x46\xb6\x90\xb8\x79\xb1\x1a\x4f\xe9\x2d\x67\xff\xc3\x82\x79\x99\xa2\xeb\x73\x16\xa8\xc7\x6b\xbc\x8a\xe2\xac\x6f\x3c\x2e\x44\x96\xe9\xf2\x40\x2a\x27\x70\x99\xcc\xa1\xbb\x54\xee\xd6\x58\x55\x84\xb2\xcf\x02\xce\x1e\xef\xa4\x75\xc4\x96\x3d\x94\x76\x87\x15\xc5\x30\xb7\x1d\x18\x0b\xcd\x28\xd3\xed\xf6\x60\x1f\x2c\x9c\x93\x00\x21\x01\xea\xb2\xbb\xde\x22\xd9\x96\xa9\x99\x58\x73\x1a\xe5\xf0\x6b\x31\x29\xfc\x5b\xe6\x2d\x70\x9c\x80\x37\xe7\x39\xc1\x58\xe2\x6a\x6d\x62\x52\xa9\xbb\xc0\xc3\xa3\x29\xe3\xf5\x11\x7b\x4c\x62\x1a\x3b\xff\xad\xf1\x5f\x65\xde\xfe\x91\xbf\x44\x2f\xd7\x3f\xf3\x04\xb2\x2e\xba\xe5\x2e\xb6\x77\x74\xff\xcc\xf3\x68\x70\x76\x93\xf5\xd5\xce\x1d\x0f\x4b\xa7\x65\xab\x62\x75\xe6\x6a\x7c\xe3\x71\x89\xac\xc6\x37\x2b\x96\x33\xf7\xd7\xbe\x65\xca\xe8\x17\x8e\x20\x34\xc8\x8d\x12\xc2\x7f\x60\x57\x99\xa7\xcd\x86\x0e\xf7\x67\x71\x30\xdb\x31\x3c\x58\xca\x2c\xa5\x81\x1a\x6a\xb1\x8c\x93\x9f\xce\x9a\xd5\xce\xe7\x2b\x1b\x08\x04\x57\xd7\xef\xeb\xb5\x2a\x98\x56\xa9\xf4\xe9\x75\xb5\x54\x3a\xeb\x9c\xc1\x6f\x17\x9d\xcb\x3a\x1c\x14\xf7\xa1\x83\x0b\x5a\xc9\x29\xd8\x1c\xbf\x54\xb2\x1b\x18\x56\x03\xa5\xc2\xa3\x52\x69\x34\x1a\x15\x47\xaf\x8b\x22\xea\x97\x3a\xad\xd2\x1d\xd1\x3a\x20\xe4\xf4\xa7\xa5\xe6\x30\x8b\x9e\xf2\xcc\x53\xe4\x6c\x59\x46\x5b\x8d\x7d\x06\x0e\x4a\xab\x99\x78\x2c\xe2\xe4\x50\x5a\x7d\x00\x91\x96\x48\xbb\x8f\x7b\x8b\xb8\x5b\x74\xc5\xb0\x44\x3a\xf4\xe3\xa0\xa4\xc9\x39\x6e\x42\xcf\xd2\xaa\x59\x13\x73\x48\xcc\xa6\xce\x80\xc1\x65\xad\x03\x75\xee\xb2\x40\x32\x78\x89\x83\x3d\xc3\xa8\x8a\x70\x1c\xf1\xfe\x00\x03\xd2\xdd\x83\x57\xfb\x07\x6f\xe0\x32\xa1\x68\x18\x57\x2c\x1a\x72\x29\x91\x22\x70\x09\x03\x16\xb1\xee\x18\xfa\xc8\x07\x53\xaa\x80\x02\x31\x06\xa2\x07\x58\x3c\xa3\x3e\x2b\xe0\x16\x0d\x85\x1e\x03\xee\xd2\x24\x22\x88\xae\x72\x78\x40\xf1\xef\x80\x8b\x3c\x0c\x9c\xa9\x06\x48\x46\x8a\x9e\x1a\x39\x51\xa2\xa1\x23\xa5\x70\x39\x4a\xe8\x81\x27\xdc\x78\x88\xf1\xa7\x13\x17\x7a\xdc\xc7\x54\x7d\xa9\x50\x68\xb3\x9d\x62\x98\x7b\x9a\x89\xc7\x1c\xdf\xc0\x04\xa6\x67\x93\x47\x7a\xb3\x25\x62\x05\x11\x93\x2a\xe2\xda\x0a\x05\xe0\x81\xeb\xc7\x1e\xc9\x30\x79\xec\xf3\x21\x4f\x39\x10\xba\x56\x5c\x1a\x48\x14\x17\xef\x05\x2d\x67\x01\x86\xc2\xe3\x3d\xfa\x3f\xd3\x6a\x85\x71\x17\x53\x6c\x50\x00\x4c\x0a\x24\xdd\x8d\x15\x02\x25\x01\xb5\x1d\x0b\xa4\x47\x49\x44\x20\x99\xef\x1b\x48\x81\xa3\xdc\x5a\xd7\x99\x74\x7a\x0e\x89\x1e\x92\x41\x55\x6a\x22\x49\x90\xd1\x00\xbd\x9a\xd1\x84\x4b\xa3\x17\x47\x01\xb2\x64\x1a\xc7\x13\x68\x32\xcd\x91\xa2\x99\x20\x34\xbd\x27\x7c\x5f\x8c\x48\x35\x5c\xf1\x7a\x3c\xdd\x5f\x69\x27\x3b\x5d\xda\x63\xba\x53\xbf\x62\x31\x44\x51\x13\x11\xc8\x01\xe1\xcc\xab\xe9\x23\x39\xc0\xad\x06\x74\x59\x6a\x30\xe4\x8b\xe6\x75\xe6\xd4\x89\x88\x3d\x2d\xb1\x14\x77\x7c\x08\xb1\xa6\x12\xbf\x45\x35\x8b\xc8\xff\xc2\x86\x76\xf3\xbc\xf3\xa9\xd2\xb2\xa1\xd6\x86\xab\x56\xf3\x63\xed\xcc\x3e\x03\xb3\xd2\xc6\xb1\x59\x80\x4f\xb5\xce\x45\xf3\xba\x03\x38\xa3\x55\x69\x74\x3e\x43\xf3\x1c\x2a\x8d\xcf\xf0\x6b\xad\x71\x56\x00\xfb\xb7\xab\x96\xdd\x6e\x43\xb3\x65\xd4\x2e\xaf\xea\x35\x1b\x61\xb5\x46\xb5\x7e\x7d\x56\x6b\x7c\x80\xf7\x88\xd7\x68\x62\x08\xd7\x30\x76\x91\x68\xa7\x09\xc4\x30\x25\x55\xb3\xdb\x44\xec\xd2\x6e\x55\x2f\x70\x58\x79\x5f\xab\xd7\x3a\x9f\x0b\xc6\x79\xad\xd3\x20\x9a\xe7\xcd\x16\x54\xe0\xaa\xd2\xea\xd4\xaa\xd7\xf5\x4a\x0b\x13\xbb\x75\xd5\x6c\xdb\xc8\xfe\x0c\xc9\x36\x6a\x8d\xf3\x16\x72\xb1\x2f\xed\x46\xa7\x88\x5c\x11\x06\xf6\x47\x1c\x40\xfb\xa2\x52\xaf\x13\x2b\xa3\x72\x8d\xd2\xb7\x48\x3e\xa8\x36\xaf\x3e\xb7\x6a\x1f\x2e\x3a\x70\xd1\xac\x9f\xd9\x08\x7c\x6f\xa3\x64\x95\xf7\x75\x3b\x61\x85\x4a\x55\xeb\x95\xda\x65\x01\xce\x2a\x97\x95\x0f\xb6\xc6\x6a\x22\x95\x96\x41\xd3\x12\xe9\xe0\xd3\x85\x4d\x20\xe2\x57\xc1\x7f\xab\x9d\x5a\xb3\x41\x6a\x54\x9b\x8d\x4e\x0b\x87\x05\xd4\xb2\xd5\x99\xa2\x7e\xaa\xb5\xed\x02\x54\x5a\xb5\x36\x19\xe4\xbc\xd5\xbc\x2c\x18\x64\x4e\xc4\x68\x6a\x22\x88\xd7\xb0\x13\x2a\x64\x6a\xc8\x78\x04\xa7\xd0\xf8\xba\x6d\x4f\x09\xc2\x99\x5d\xa9\x23\xad\x36\x21\x93\x8a\x93\xc9\x45\xc3\xb2\xb0\x22\xe9\x12\x78\x37\xf4\x03\x59\xce\x29\x6c\x07\x87\x87\x87\x49\x3d\x33\x37\x9b\x24\xa9\xb8\x95\xcd\x9e\x08\x94\xd5\x73\x86\xdc\x1f\x1f\xc1\xcf\x17\x0c\x5b\x16\x46\xa2\x03\x0d\x16\xb3\x9f\x0b\x30\x05\xa0\xaa\x11\x86\x1c\x86\x3f\x16\x37\x0b\x37\xd8\xbc\x77\x0c\x5d\x71\x67\x49\xfe\x17\xf5\x62\xfc\x1d\x61\x81\xb4\x10\x74\x0c\x9a\x28\x3e\x60\x47\x70\xf0\x26\x44\xc0\x10\x0b\x13\x0f\x8e\x60\xff\x98\x6a\xeb\x80\x39\xde\x73\xf2\x1f\x32\xe5\x00\x75\xd4\x32\xb6\x47\x36\xa2\x2c\x32\x29\x7b\x15\x16\xbd\xb2\x39\xe2\x9e\x1a\x94\x3d\x86\x9d\x93\x59\x7a\xf0\x7c\xc6\x82\xd2\x44\x5c\x72\xa6\xc5\xfe\x8c\xf9\x6d\xd9\xac\x26\xa2\x5a\x9d\x71\xc8\xe6\x04\xa7\xa5\x48\x89\x9c\x7b\xac\x3b\x81\x64\xaa\x7c\xdd\x39\xb7\x7e\x79\x66\xf1\xf5\xdb\x88\xe7\x73\xf7\xba\xb5\xc8\x49\x49\x0b\x77\x6a\x18\x27\x25\x0a\x4a\xfa\xd1\x15\xde\x18\x38\xa2\x48\xac\xb9\x28\xb1\xa9\x07\x6a\x4c\xbf\xd3\x8c\x92\xee\x00\xbb\xba\xce\x28\x9b\xba\xfb\xe5\x64\xed\xfb\xa4\x4a\x5a\x23\xd6\xfd\xca\x91\x91\x7e\x30\x14\x02\x7b\x0a\x21\x25\xbd\x81\x3b\x92\x79\xb3\x49\x14\x1b\x1a\xdb\x72\xbc\x2f\xb1\x54\x47\xd8\x71\x02\x76\x8c\x4b\x09\xea\x4c\x48\x72\x7f\xff\x5f\xc7\xd8\x94\x03\x66\x4d\x41\xc5\x77\x6c\x78\x0c\x3a\x03\x92\x09\xf0\x13\x1f\x52\xb2\x20\x07\x94\xd3\x71\xbf\xf6\x23\x11\x07\x9e\xe5\x0a\x5f\x44\x47\xf0\xa2\xf7\x8e\xfe\xcc\x9b\x1f\x42\xc7\xf3\xb4\x54\x14\x0d\xdd\xbe\x9e\x59\x36\xd3\x99\x26\xd9\x5b\x39\xdd\xa7\x0e\x8f\x39\x95\x36\xd4\x23\x57\x76\x80\x13\x15\x3d\x63\x1d\x03\x20\x09\x9e\xb8\x92\xde\xe2\xd6\x00\x89\xf8\x16\x86\x58\x1f\x25\x51\x22\xcc\x1a\xea\x56\x3f\xc0\x6a\x24\x42\xf3\x14\x13\xcc\x9b\x09\x9a\x54\x56\xf3\xdd\xfe\xbe\xb9\x05\x42\xa7\x5b\x2b\x44\xf5\x85\xfb\x35\x13\xdb\x43\xe7\xce\x4a\x83\x04\x85\x0d\xef\x32\x0f\x5d\x9f\x39\x11\x31\x54\x83\x0c\x7c\x55\xa2\x4c\x8d\x03\x4e\xac\xc4\x42\x4a\x64\xac\xa5\x0d\x85\xa6\xf2\xf8\xed\x53\x87\x55\x56\xdf\x45\xe3\xac\x57\x62\x22\x37\x39\x59\x27\x73\xea\x67\xb2\x04\xb6\x27\x5c\x8d\xa7\xb3\xcb\xe6\x7e\x32\x96\xa1\xe3\x4e\xc6\x4f\xaa\x68\xfa\x30\x72\x3c\x1e\xcb\x23\x78\xad\x61\x39\x05\xa0\xd7\xcb\x54\xb1\x04\x0d\x89\x60\x28\xe0\xae\x9e\x7b\xf0\x82\x1d\xd2\x9f\x6c\x61\xe8\xf5\xe6\x6c\xb1\x0d\xd5\x61\x26\xc9\xd3\x55\x89\x77\x2b\x13\x2e\x63\x5d\x8d\x32\x4a\x5b\xcd\xdb\x7d\x34\xb2\x6e\x51\xe9\x7c\xdc\xd0\x29\x16\xe5\xf9\x4b\xff\xdd\xd7\x4e\x59\xf6\x9b\xfd\xee\xed\xab\x57\xd5\xfc\x06\xf4\x8a\xe2\xda\x84\x34\xdf\x12\x06\xf3\xde\x4b\x70\xf3\x33\x72\xf2\xcf\xec\x50\x73\x7a\x9a\x09\xfa\x65\x49\xee\xbb\xa4\x3d\x38\xc0\x09\x72\xfa\xc2\x03\x75\x8e\x60\x76\xf0\xb6\xe2\xe0\x93\xde\x7b\x00\x2c\xf3\x4d\x8f\xe1\xca\x99\x43\xb8\xa5\x69\xe9\xab\x95\x8c\xf3\xa7\x35\x78\x3a\x8e\x76\x61\xba\x49\x33\x9b\x05\xcf\x41\x12\x3c\xeb\x62\x63\xeb\x6b\xdf\x4a\xb3\x6f\x57\x10\x6c\x7b\x28\x60\xed\x99\xd4\x92\x75\xe1\x90\xaa\x81\x1b\xb7\x88\xf5\xca\xe6\x26\x2f\xdd\x9f\x38\x1e\x26\x45\xf3\xfc\xfc\x3c\x2d\xbe\x1e\x73\x45\xa4\xdf\xc9\x4d\xb6\x07\x99\x0d\xc1\x2b\xda\x0e\x64\xea\x76\x57\xf8\x5e\x7e\xe1\x76\xe3\x48\x12\xf5\x50\xf0\x04\x30\x5d\x50\xf0\x40\x13\x4d\xd7\x15\x0b\x05\xfe\x2d\x09\xa6\xe9\xe9\x97\xa8\x58\x30\x87\x48\xd3\x09\xb9\x42\xfa\x7f\xb1\xdc\xa2\xff\xfa\xcd\x2f\xcc\x73\x72\xfa\xf5\xd2\x8c\x14\xac\xad\x7c\x94\x34\xf2\x29\x70\xba\x7a\xc3\xf6\x92\xb8\xf7\xf4\x23\x67\x23\x7a\xff\x76\xef\xdb\xf1\x93\x92\x93\x1b\xc3\x0b\x85\x37\xbf\xfc\x4e\x4b\xf7\xda\xc3\x8f\x9c\xa6\xb0\x4b\xd9\x7f\x26\x65\xa5\x8a\x44\xd0\x7f\x3e\xd3\xfe\xbe\xfa\xea\xd4\x1f\xe9\xc9\xd7\x49\x29\x11\xf2\x3b\x44\x5d\xce\x82\x21\x7d\x32\xb9\x1f\xb4\x78\x84\xb6\x8b\xc3\xff\x8f\x38\x4c\x96\xa6\xd3\x50\x3b\xe9\x46\xcf\xfa\x1e\x31\xcf\x46\xf7\x5c\x8c\x5b\x7d\x7b\xed\x99\x95\x59\x9d\x77\x79\xbd\x60\x76\x88\x9e\x74\x82\x67\x8f\x8c\x39\x89\xb6\x25\x3c\xee\xb5\xe8\xbd\xb7\x1d\x7f\xd0\x60\x99\x5f\x61\x2e\x5e\xbf\x7c\xa6\x05\xe5\x64\xb9\xb5\xb4\xa6\xc4\x55\x1b\x8b\x68\xf5\x97\x0d\xa7\xe4\x02\x29\x2d\xa2\xb6\xaf\xc6\x3c\xae\x9b\x6e\xb8\xbc\x9b\xbf\x6b\x92\xeb\xde\xdd\xaa\x70\x6b\xba\xf1\x16\x76\xbf\x93\xc1\x16\xca\xf4\x43\x67\xf0\xba\x15\xf1\x2e\xb1\xfe\xf7\xb7\x5b\xd3\x3b\x7b\xb3\x0d\xd7\x04\xf4\x0c\x5b\xae\xf9\x1b\x84\xbb\x68\xdc\x6d\xba\x76\x9b\xae\xdd\xa6\x6b\xb7\xe9\xda\x6d\xba\x76\x9b\xae\x0d\xfa\x29\xce\xa6\xf3\xb8\xd3\x07\x1c\x85\x4e\x51\x66\x90\x27\xbf\x89\x91\xb9\x9a\x34\x77\xd3\x64\xe6\xe8\xc3\xc3\xc3\x75\x07\xdc\xd9\x93\xdd\xe5\x23\xc9\x6d\x39\xe9\xdd\x9e\xe5\xcb\x53\x2e\x5d\x5e\xad\x5c\xba\xe4\x1e\xa2\xdd\xe7\xf2\xb9\xb5\xcd\xc2\xbd\x86\xec\x2d\xac\xf9\x72\x95\xfd\x40\xdc\x7c\x5a\xd5\x33\x1a\x6d\x5c\xaa\x50\x27\xe8\x8e\x37\x3b\x87\x5b\xae\x1d\x4b\xf7\x1d\x16\x2b\xc3\x49\x09\xd3\xfc\x34\xf9\xaf\x91\x2d\x13\x3f\xc8\xf5\xba\x44\xc5\x59\xfd\x3a\x29\xd1\x2d\x56\x82\xd0\x75\xe0\x53\xc3\xc8\xff\x7e\x27\x8c\xe5\x40\x20\xc7\xef\xf0\x01\xf6\x12\xa9\x7f\xfe\x7b\xb0\xef\xf3\x39\xd8\xe6\x5f\x83\x7d\xbf\x8f\xc1\xe6\x78\x6e\x60\xc9\xd9\x57\xd4\x0f\xf8\xac\xf2\xbf\xed\x58\x62\xaa\x5e\x42\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 16990, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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

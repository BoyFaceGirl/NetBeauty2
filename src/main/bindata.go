// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/runtime.compatibility.json
// assets/runtime.supported.json
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

var _assetsRuntimeCompatibilityJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xd1\x92\xa3\xba\x11\x7d\xdf\xaf\x98\xba\xcf\xac\xcb\x80\x41\x52\x7e\x25\x95\x87\xd9\xac\x93\x4c\xd5\x5c\xfb\xd6\xce\x4c\xc5\x37\xa9\xfc\x7b\x6a\x6c\x6c\xa4\xee\xd3\x12\xd8\x02\xc9\x63\xde\x86\x96\x91\x4e\x9f\xd3\xad\x16\x82\x81\xff\x7e\x7b\x7a\xfa\xed\xf9\xf5\x8f\x97\xdd\xf6\xb7\xbf\x3c\xfd\xf5\xdb\xd3\x53\x7f\x5c\x9c\x8e\x5e\x5f\x76\x1f\x87\xef\xbf\x7f\xbc\xbd\x3a\x96\xf3\xc1\xc7\xee\xe5\xf2\xf7\xf3\xee\xcf\xf3\x9f\x3f\x9e\xdf\xb6\xbf\x7d\x7b\x7a\xfa\x5b\xd1\x8f\xf0\xfd\xd0\x6e\xe8\x28\x47\x5b\xe1\x1f\xd7\xfe\x8d\x80\x86\xfd\xc4\x86\x67\x37\x8e\x80\xbb\xaa\x57\xe5\x9a\xe2\x3d\x19\x0b\x62\x32\xcc\xa2\x99\x45\x31\x4b\x1b\x72\xfc\x56\xc2\x8f\x60\x11\xeb\x7d\x03\xc5\x84\x9c\x83\x3f\x04\x2e\xc3\xdf\x01\x22\xe0\xef\x00\x3d\xf0\x77\x84\xb4\xcc\xe3\xa7\x44\xcc\x97\x83\x58\xcf\x37\xa4\x4a\x29\xa4\x4a\x1c\x52\xc8\xdf\x25\xf8\x26\x0f\xbe\x16\x48\x34\x47\x7c\xb4\x42\x78\x7c\x01\x4a\x15\x70\x2b\x4d\x12\x2a\x81\xe4\x07\x8a\x6f\x0d\xdc\xcf\x65\x8e\xd4\x82\x3c\xcb\x34\xc5\xa8\x32\x80\xa6\x7c\x8b\x9f\x11\x84\x5d\xea\xd4\xe8\x00\xd8\xfd\xfc\xb5\x7f\xf9\x69\x71\xd9\x19\xae\x17\xe9\xd4\xc1\xf7\xe7\x5f\xbf\xb3\x5e\x8f\xc6\x42\x1e\xc8\x6e\xe7\xbe\x59\x8d\x63\x91\x38\xd1\xe2\x98\x03\x68\x7c\x5c\x3b\xcd\x63\x10\xad\xaa\x92\xc1\xf9\xb4\x79\xb0\x5c\x39\x0c\x54\xe1\x6c\x2f\x02\x00\x92\xa8\xd5\x81\x03\x82\xf5\x2d\x43\x80\xcf\x2f\xed\x9f\x36\x62\xe9\x57\xfb\x77\xeb\x57\xfb\xf7\x60\xb7\xc7\xa3\xcb\x19\x6e\xdb\xdf\xb7\xbb\xf7\xfd\x5b\xdf\xda\x1d\x77\xfd\xfc\xfa\xd7\xf6\x86\xa9\xf6\xd4\x97\x3b\xcd\x5a\xb6\x42\x1a\xd1\x6e\x65\x08\xe2\x4e\x5d\xa7\xc1\xed\xd5\xe0\xc5\x22\xc2\xeb\xdb\x62\xd0\x43\xd7\x7f\x8e\xb5\xf0\x61\xa2\x34\xa1\x73\x64\xdc\xf3\x93\xac\x99\x93\xda\x03\x56\x47\x25\x59\x43\x92\x35\x20\x59\x07\x48\x06\xe7\xc8\xb8\x67\x24\xf9\xe7\xf6\xc7\xcb\xf3\xae\x77\xb1\x3b\xbe\x9a\xb8\xd3\xf9\xee\xdc\x6f\xd9\x0a\x71\x94\xb8\x33\x7a\x3f\xa4\xad\x9f\x63\xf5\x43\x89\x3b\x43\xf7\x03\x6f\x5f\x11\x9c\x5e\x62\x01\x8e\x27\x96\xdd\xe6\x31\x70\x0e\x80\x9b\x83\x9f\x99\x09\x62\xef\xfb\x41\xb7\x1c\x86\x6e\xbd\x30\xfa\x66\x00\xa3\x6f\x1c\x01\xc3\xd9\xea\xed\x4d\x32\x88\xeb\xc6\x40\x99\x71\x36\x17\xde\xc1\x13\xe4\x4f\x07\x8c\x87\x49\xdf\x10\xc6\x9c\x24\xd5\x3a\x80\x3c\xdb\xfa\x86\x30\xf2\x24\x59\xc9\x36\xec\x5d\x73\x00\xf5\xfc\xe9\x7b\x04\xc6\x33\xf8\x6c\x0e\xe1\x9d\x3f\xcf\x35\x83\xaa\x23\x67\xb9\x86\x49\xae\x41\x8e\xeb\xf4\x29\xae\x85\x0c\xd7\x30\xc1\x01\xe0\x34\xf9\xad\x85\xf4\xd6\x30\xbb\x01\xec\x34\xc9\xad\x61\x6e\x6b\x90\xda\x14\x72\x82\xcc\xd6\x30\xb1\x35\xc8\x6b\x06\x76\xfe\xb4\x36\x0c\xa8\x89\x9c\xd6\x06\xa6\xb5\x01\x69\x6d\xd2\xa7\xb5\x11\xd2\xda\xc0\xb4\x06\x80\xd3\xa4\xb5\x11\xd2\xda\xc0\xb4\x06\xb0\xd3\xa4\xb5\x81\x69\x6d\x40\x5a\x53\xc8\x09\xd2\xda\xc0\xb4\x36\x20\xad\x19\xd8\x19\xd3\xfa\x1f\xdb\x9f\xfb\x5f\xcf\x3d\xcc\xee\xf8\xea\x24\x3e\x9d\x4f\xb3\xc2\xb1\x16\xe2\x48\xf1\xe3\xbd\x1b\xf8\x00\xc0\x1c\xfc\x50\xa2\x86\xc4\x69\x80\x55\x55\x53\x14\x9f\x26\x19\xc4\x75\x63\x60\xf2\xfb\x86\xc2\x0b\x20\x99\x48\x9f\x00\x81\x4e\x67\x73\x00\x75\x0a\x35\x01\x56\x1f\x88\xeb\xc6\x90\xd4\xdc\x60\x35\x37\xd9\xa8\xb9\xc1\x6a\x6e\x90\x9a\x14\x75\x0a\x35\x1b\x8e\xb5\x89\xad\x66\x23\xa9\xd9\x60\x35\x9b\x6c\xd4\x6c\xb0\x9a\x0d\x52\x93\xa2\x4e\xa1\x66\xcb\xb1\xb6\xb1\xd5\x6c\x25\x35\x5b\xac\x66\x9b\x8d\x9a\x2d\x56\xb3\x45\x6a\x52\xd4\x29\xd4\x54\x1c\xab\x8a\xad\xa6\x92\xd4\x54\x58\x4d\x95\x8d\x9a\x0a\xab\xa9\x90\x9a\x14\x75\x0a\x35\x35\xc7\xaa\x63\xab\xc9\xf6\x5b\x68\x43\xe1\x05\x90\x50\x4d\x8d\xd5\xd4\x48\x4d\x8a\x3a\x85\x9a\x86\x63\x35\xb1\xd5\x64\x97\xd9\xb4\xa1\xf0\x02\x48\xa8\xa6\xc1\x6a\x1a\xa4\x26\x45\x9d\x40\xcd\x7a\xcd\xb0\xd6\xeb\xc8\x6a\xd6\xec\x6e\x07\x6d\x28\xbc\x00\xd2\xa9\x59\xaf\xa1\x9a\xf5\x1a\xa8\xc9\x50\xa7\x50\xb3\xe4\x58\xcb\xd8\x6a\xb2\x87\x79\x68\x43\xe1\x05\x90\x50\xcd\x12\xab\x59\x22\x35\x29\xea\x14\x6a\x56\x1c\x6b\x15\x5b\xcd\x4a\x52\xb3\xc2\x6a\x56\xd9\xa8\x59\x61\x35\x2b\xa4\x26\x45\x3d\xab\x9a\xbf\xb6\xdb\x1f\x6f\xd6\xf3\x98\x67\xc3\x15\x9d\x10\x97\x2d\x63\x21\xf6\x7d\x2d\x60\xe7\x9f\x6c\x2c\x9b\x3c\xd2\x88\x9e\xa1\x23\xe4\x1f\x5d\xc4\x31\xa7\x73\x78\xb5\x46\xa0\x56\xeb\x20\xa4\xab\xc7\x93\x88\x58\xad\x31\x15\x08\x4b\x6a\xd2\x60\x9c\xac\x28\x80\x69\x89\x14\x23\x6a\x25\xd0\x83\xf1\xdd\x0f\xe9\x15\x74\xb6\x1a\xea\xe8\x64\x42\xd0\x69\x99\xb6\x70\x2c\x10\xf3\xd7\x14\xad\x86\xc4\xd4\x43\x49\x99\x55\x48\xba\x3b\x4f\x5b\x38\x16\xe8\xc7\x83\x8a\xfe\xcf\xed\xee\x7d\xbf\xef\xe9\xeb\x8e\xa5\x55\xc5\xc0\xfe\x5c\x49\x2c\x5b\x21\x8e\x12\x77\x31\x73\xea\xe1\x82\x60\xbc\x1f\xfd\xad\x79\xb7\x93\xb8\x77\xeb\xed\xa5\x27\x1f\x27\xee\x7a\xd4\xbe\x05\xce\xc7\x8a\x7b\x5f\xfc\xd4\xed\xf1\x1f\x88\xc8\x50\xb7\xfd\x1b\x56\xdf\x07\x92\xe6\x62\x2f\xbc\x23\x4e\x21\xe1\x79\x68\xae\xa3\xd5\x12\x86\x35\x85\xe2\x67\x00\x5c\x76\xab\x25\x0c\x6d\xaa\x00\x71\xa7\x0a\x62\x0f\xc0\x8a\x3f\x65\x74\x43\xdb\x0f\x48\x10\x7b\x08\x52\xcc\x87\x20\x7a\x2f\x09\x9a\x09\x1c\x07\x3e\xc7\xf7\xe5\xf7\x97\xdd\xfb\xaa\x54\x64\xa0\xb3\xf5\xdc\xdd\x8f\x8f\xdd\xfb\xc7\xaa\xdc\xac\xd6\x1b\xd7\x76\x3e\xba\xf5\x89\x2e\x7b\x54\x40\xaf\xdd\x52\x0c\xc3\xe9\x90\x2e\xe2\xe7\xbf\x72\x3d\x9a\xf5\x71\x20\xdb\x25\xfb\xd2\x88\xd8\x07\x12\x30\xbd\x50\xf4\xd2\x09\xb4\x01\xac\xd8\x83\x47\x92\xb6\x12\x28\xab\x86\xd3\x95\x48\xf0\xca\x23\x78\x25\x0a\x2e\xf8\xb5\x04\x88\x44\x73\x2d\x50\x5c\x8f\xa1\x37\xa3\xb0\xa9\x3d\x61\x53\x8b\x61\x20\x78\xbb\x84\x59\x9c\x30\xd3\x50\x10\x4d\x1d\x6a\xe7\x08\x11\x2d\x06\x88\x16\x04\x81\x38\x91\x20\x1c\x7f\xb6\x82\x08\x35\x5f\xc3\x48\x4d\x24\x94\xa7\xe6\x6b\x31\xd7\x04\x0f\x1e\x49\x5a\x5c\xf3\x35\x9c\xb4\xf2\x12\x5c\xae\xf9\x5a\x9c\x8c\x45\xbf\x96\x00\x91\x68\xc6\x35\x5f\xc3\x2a\x78\x0f\x61\x23\xd7\x7c\x2d\xd6\x7c\xd1\xdb\x25\xcc\xe2\x84\x99\x81\x82\x18\xea\x90\x9e\x23\x44\x8c\x18\x20\x46\x10\x04\xe2\x44\x82\x70\xfc\xd9\x0a\x22\xd4\x7c\x03\x23\x35\x91\x50\x9e\x9a\x6f\xc4\x5c\x13\x3c\x78\x24\x69\x71\xcd\x37\x70\xd2\xca\x4b\x70\xb9\xe6\x1b\x71\x32\x16\xfd\x5a\x02\xc4\xa2\x79\x6f\xed\xf9\xef\x2f\x3b\xd5\xb7\xbd\x1b\x67\x4f\xb6\xec\xf7\xce\x56\xbd\x3b\x8a\xdd\x32\xf1\x9b\x6c\xf6\xaf\xf6\xfb\x98\xf6\xd6\x7b\x8b\x5c\x48\xb1\xde\xc3\xb4\x7f\xa5\xef\x60\xda\x93\x77\x29\xb9\x18\x28\x15\xf4\xb7\x32\xc6\x99\x49\xb4\x9f\x2d\xea\x8e\x81\x43\x27\x98\x76\xdb\x54\x24\xd3\x67\x8f\x2c\x5b\x21\xa1\xe4\x64\xf3\x33\x64\x4f\xee\x47\xaa\x92\xd0\x52\x62\x4a\x3a\x17\xcb\xc4\x42\x96\x40\xc8\x92\x09\x49\x7c\x00\x52\xf2\x73\xbc\x9e\x3e\x4a\x30\x54\x84\xda\x0a\xd3\xda\xb9\x58\x65\x1d\x2a\x15\x08\x95\x8a\xc9\x4e\x3c\x04\x42\xf3\x73\x02\x3c\x2c\xe1\x36\x54\xa2\x9a\xc8\x53\x63\x69\x3a\x17\xeb\x3b\x0e\xc6\x1a\x04\x63\xcd\x82\x82\xf8\x0f\xc2\x80\x9f\x13\x60\x69\x09\xe8\x39\x03\x9a\x4a\x2c\xc8\xdb\xb9\x38\x4a\xc8\xbb\x0a\xf7\x0d\x08\xf7\x0d\x0b\x19\xc2\x0e\x08\x12\x7e\x4e\x80\xc3\x25\x65\xee\x2d\x65\x1a\x12\x26\x0d\x0e\x91\xce\xc5\xe6\x41\x13\xaa\x01\x09\xd5\xb0\x80\x22\xdc\x81\x10\xe2\xe7\x04\x18\x5e\x92\xf2\x11\x93\xb2\x25\xa1\xd6\xe2\x30\xeb\x5c\x6c\x97\x94\x45\x1c\x82\x94\x6d\x59\xb8\x11\x66\x41\x80\xf1\x73\x02\xfc\x2f\x69\xbf\xa4\xfd\x35\x69\xaf\x9d\x60\xd5\x02\xc0\x48\xaf\x47\xdf\xbf\xd2\x37\x59\xee\xc9\x2b\xce\x5d\x0c\x8c\x44\xed\x21\x31\xd5\xab\xd0\x3f\x21\x93\xbd\x4f\xed\xec\x2a\x6a\x17\xa6\x38\x0f\x45\x24\x19\xec\x7d\x6a\x16\xbd\x04\x25\x27\x9b\x9f\x21\x7b\x72\x1f\x52\xfd\xb1\xdd\xbd\x7d\xd8\x9f\xa7\xb8\x58\xae\xa7\xbb\xeb\x81\x10\x6e\x5b\x0b\xcf\x58\xd3\x38\xb8\x2a\x6b\x67\x5b\xcf\x31\xfb\xe0\x5c\x3f\x1a\xf6\xbf\x6f\x2a\x06\x40\x49\xc6\x56\xe3\x24\xb0\x63\x9e\x80\xad\x86\xa5\x27\x6b\x62\x6c\x01\x28\x09\xd9\x2a\x31\xf6\x72\x1a\xb6\xe8\xfe\x3f\x6b\x02\x6c\x31\x28\xa9\xd8\xda\x54\x90\xad\xa3\x39\x3e\x5b\x9f\xdd\x0a\x6c\x5d\x9a\x28\x5b\x08\x4a\x42\xb6\xd0\xbc\x75\x34\x4f\xc2\x96\x34\x6f\x5d\x9a\x00\x5b\xd9\xcc\x5b\x9b\xca\xd9\x4d\x77\xcc\x93\xb0\x45\x77\xb4\x59\x13\x60\x8b\x41\x99\x9b\xad\x37\xeb\x1f\x7a\x3f\x0f\x46\x9e\x4c\x5c\x7e\x73\x10\x90\xfe\xae\x01\xb7\x2a\xd7\xce\x47\x44\x7a\x13\x1e\x63\x78\x9f\x0c\xb9\x65\x2e\xe4\xc1\xa6\xf0\xaf\xe4\x38\x4a\x1f\x86\xeb\x06\xc1\x0e\x97\xc8\x61\x3e\x7a\x0a\x5a\x2a\x8e\xb6\x0a\x23\xbd\x95\x27\x3a\xe7\x39\xe6\x22\x04\x27\x5f\x36\x6b\xee\x53\x1d\xf6\x27\x3a\xbd\x74\x92\x74\xcc\x45\x08\xdf\xbd\x8b\x00\x3c\x1f\xe0\xf5\xf4\xaa\xd0\xbb\x53\x8e\xb9\x08\x01\x7e\x0c\xed\x1a\xce\x4f\x13\xe6\x26\x81\x98\xf4\xce\x88\x63\x2e\x42\x1e\x2c\x92\x77\x3b\x2d\x17\x02\x6f\xdb\x71\x3a\xee\xe2\x90\x97\x58\x58\xb6\x42\x18\x23\xfe\x8b\x2b\x2e\xdb\x49\x2e\x8c\xf9\x36\x98\x8e\xbb\x5b\xad\x3b\x7c\xbf\x53\x7f\x3b\xcb\xf4\xfe\x82\x65\x2b\xa4\x11\xe7\x76\x5f\x11\x78\x91\x6e\xaf\x58\x9b\xe8\xa4\x7f\xbe\xd5\x9e\xd4\x7d\x7b\x2b\x89\xde\x17\x88\x4f\x08\xdd\x4a\x42\x77\x2a\x30\x8a\xdc\x68\x2b\x99\x13\xf4\x16\xcf\x94\x34\x96\x90\x46\x74\x2b\x8a\xa1\xba\x3f\xaa\x2b\xe6\x28\xbd\x23\x37\x27\xf5\x15\xa4\x1e\xdd\x49\x64\x28\xbf\xa6\x3c\x35\x23\x83\xde\x74\x4d\x29\x57\x0d\xe5\x42\x37\x8f\x19\xea\xc7\x95\x94\x13\x16\x24\x2b\xa1\xc4\x1b\x28\x31\x7a\xc6\x80\x79\xb1\x84\x81\x8f\xd8\x86\x91\x4a\x1f\xeb\xc8\x39\x2c\x1a\x18\x16\xe8\xf1\x15\xe6\xd5\x12\x3a\xb7\x86\x0e\xbd\xaa\x50\xee\x2a\xff\xce\x42\x09\x5d\xc7\x28\x70\x25\x03\xbc\x5c\xc2\x6d\xfa\x70\xd3\x44\x9a\x48\x8f\xa6\x9c\xfa\x42\x5b\x05\xf4\xbb\xc1\x7c\xd4\x24\x9b\x09\xf4\x71\x25\xcb\x26\x03\x9d\x5b\x2b\x76\xc1\xa9\x49\x70\x45\x55\x6f\xc5\xbe\x76\x42\xec\x85\x0f\x49\xa6\x2a\xc3\xab\x68\x0d\xf2\x96\x3b\x94\x59\x2c\xd0\xab\x68\x4d\x26\xa3\x69\x63\x83\x7d\x3b\x85\xd8\x0b\x3f\xb2\x7b\x8e\x20\xb4\x81\xa0\x41\x75\x40\x6e\xdf\x4f\x94\xbd\xbd\x6e\xdf\x7a\x27\x8f\x47\x57\xc7\xcc\xe7\xd9\x2e\x69\x17\x4b\x21\xf4\x1f\xdf\x19\xe7\x1e\xec\xd9\x20\x0d\x7f\x4d\xef\xdc\x43\x72\xd7\x00\x0c\x3a\x3f\x07\xf6\xbc\xd1\x9b\x64\x84\x37\xd3\x42\xf3\xc5\x35\x17\x7e\x24\xf9\xd1\xc7\x83\x68\x55\x85\xbc\x88\xcb\x27\x0e\xb4\x15\xe2\x0a\x40\xbb\x4f\xd6\x6b\xee\x6f\x1d\xf6\x75\x52\x19\x6a\x2c\x43\x8d\x08\xe6\x58\xbf\xb2\x58\x80\x95\x01\x8c\xcc\xab\xde\x06\xab\xb7\x41\x8c\x73\xf0\x8f\xae\x71\x43\x99\x6b\xc2\x9c\x25\x15\xbc\x41\x72\x37\x9c\x56\xee\xc7\x12\x10\x83\xf8\xe5\xcb\x8a\x86\x82\xce\x3e\x46\xf0\x4a\xa5\x41\x8a\x70\xe7\x96\x58\xba\x3d\x96\xde\x5f\xfe\xb3\xdd\xf5\xfc\x9f\x0e\xaf\x56\xf5\x78\x3a\xfd\xae\x89\x6d\x2c\xa4\x61\xe2\x7f\xcd\xe4\x34\xac\xf3\x05\x8d\xde\x24\xe3\x88\xfa\x7d\x8d\x63\xff\xab\xcd\x6a\x6d\x6f\x23\xd9\x46\x11\xc7\xb5\xc3\x40\xf2\x9d\xa6\xc2\x8f\x22\x91\x50\x1d\x44\x2e\x97\xd5\x10\x42\x3e\xbb\xb0\x0d\x12\xb6\x61\xc0\x26\xd0\xba\x91\xb5\x6e\x24\xad\x25\x60\x79\xc7\x45\x23\xc5\x45\x83\xe3\x42\xf4\x32\xb3\x08\xea\x5e\x43\x79\x71\x2a\xee\xab\x3c\xbb\xd7\x5e\x3a\xdf\xe3\xb2\x6c\x85\x3c\xa6\xdd\x8e\x5e\x85\x19\xf5\x3b\x5d\x3d\x24\x7b\x11\xe0\x58\x03\x50\xfb\x5f\x08\x60\xe3\xee\x47\x5a\xaf\x13\xa5\x70\x0f\x01\xb0\x07\x3f\xd4\xa8\x55\xfd\x0c\xc9\xce\x1a\xcb\xe6\x85\xd9\xb7\x43\x98\xf1\x73\xa0\xfb\x62\x06\x01\x3a\xe5\x97\x49\x9c\x2f\x78\x80\x1c\xb1\x5a\x8a\x10\xa6\x0c\xf3\xc9\xfa\x34\x09\x76\x8c\x85\x6a\xf6\x1f\x2d\x71\xe1\xf3\xb0\xb6\x5a\x06\x38\x96\x65\x0a\xd8\xff\xa6\xe6\x58\x27\x4e\x81\x72\x2d\xa5\x40\xd7\xc2\x08\xa5\x98\x32\x4d\x01\xfa\x3f\x7a\xac\x25\xec\x58\x9e\x29\x50\xae\xa5\x14\x28\xd7\x38\x05\xb8\x63\x19\xa6\x40\x03\xab\x40\x33\x79\x15\x68\xc4\x2a\xd0\x08\x55\x80\x63\xca\x32\x05\x1a\xb1\x0a\x34\x42\x15\x00\x8e\xe5\x98\x02\x8d\x58\x05\x1a\xa1\x0a\x20\xc7\xb2\x4c\x01\x54\x05\x9a\xc9\xab\x40\x23\x56\x81\x46\xa8\x02\x1c\x53\xa6\x29\x20\x55\x81\x46\xa8\x02\xc0\xb1\x3c\x53\x40\xaa\x02\x8d\x50\x05\x90\x63\x19\xa6\x40\x0b\xab\xc0\x74\x5f\x2c\x72\xbe\xec\x03\x53\xa0\x15\xaa\x00\xf8\xe6\x4f\x8e\x29\x70\x86\x8f\x92\xc0\x6a\x1b\xe4\x5c\xa6\xd7\xe3\xd6\x77\x99\xb0\x8b\x20\xcf\x33\xff\x62\x93\x0b\x1f\xe5\x79\x2b\x94\x3a\xe4\x58\x96\x79\x8e\x4a\x5d\x3b\x79\xa9\x6b\xc5\x52\xd7\x0a\xa5\x8e\x63\xca\x34\xcf\x4b\xf6\xf4\x28\x68\x1b\xe4\x5c\xc6\x79\x2e\xd5\xf3\x56\xa8\xe7\xc0\xc1\x3c\xf3\x5c\xaa\xe7\xad\x50\xcf\x91\x63\x19\xe6\xb9\x82\xf5\x5c\x4d\x5e\xcf\x95\x58\xcf\x95\x50\xcf\x39\xa6\x2c\xf3\x5c\x79\xea\xb9\x12\xeb\x39\x76\x2e\xdb\x3c\x57\x62\x3d\x57\x42\x3d\x07\x0e\xe6\x98\xe7\x4a\xac\xe7\x4a\xa8\xe7\xc8\xb1\x2c\xf3\x1c\xd5\x73\x35\x79\x3d\x57\x62\x3d\x57\x42\x3d\xe7\x98\x32\xcd\x73\xb9\x9e\x2b\xb1\x9e\x63\xe7\x32\xce\x73\xa9\x9e\x2b\xa1\x9e\x03\x07\xf3\xcc\x73\xa9\x9e\x2b\xa1\x9e\x23\xc7\x32\xcc\x73\x0d\xeb\xf9\x74\x1f\xa5\xb4\x47\xc0\x79\xae\x85\x7a\x0e\x3e\x03\x99\x63\x9e\x6b\x4f\x3d\xd7\x62\x3d\xc7\xce\x65\x9b\xe7\x5a\xac\xe7\x77\xfa\x11\x4f\x17\x3e\xca\x73\x2d\xd4\x73\xe4\x58\x96\x79\x8e\xea\xb9\x9e\xbc\x9e\x6b\xb1\x9e\x6b\xa1\x9e\x73\x4c\x99\xe6\xb9\x5c\xcf\xb5\x58\xcf\xb1\x73\x19\xe7\xb9\x54\xcf\xb5\x50\xcf\x81\x83\x79\xe6\xb9\x54\xcf\xb5\x50\xcf\x91\x63\x19\xe6\xb9\x81\xf5\xdc\x4c\x5e\xcf\x8d\x58\xcf\x8d\x50\xcf\x39\xa6\x2c\xf3\xdc\x78\xea\xb9\x11\xeb\x39\x76\x2e\xdb\x3c\x37\x62\x3d\x37\x42\x3d\x07\x0e\xe6\x98\xe7\x46\xac\xe7\x46\xa8\xe7\xc8\xb1\x2c\xf3\x1c\xd5\x73\x33\x79\x3d\x37\x62\x3d\x37\x42\x3d\xe7\x98\x32\xcd\x73\xb9\x9e\x1b\xb1\x9e\x63\xe7\x32\xce\x73\xa9\x9e\x1b\xa1\x9e\x03\x07\xf3\xcc\x73\xa9\x9e\x1b\xa1\x9e\x23\xc7\x72\xca\xf3\xcf\x1f\xf5\xce\x0c\x3d\x85\xa4\xe6\x55\x09\xd1\x87\x18\xeb\x68\xac\x3a\xfd\x7f\x0a\xb0\xae\x46\xfe\xf7\xc0\x25\x3a\xdc\x8e\x46\xc7\xcb\x59\x0b\xd2\xcd\x48\x75\xfe\xfd\x62\xfd\x93\xd5\xe7\xc1\x80\x13\xbe\x3f\xef\xdf\x9d\x93\x8e\x86\x82\xf7\xd1\x5b\xfd\xdd\xd9\x4a\x9f\x0d\xa0\xbb\x40\x1f\x00\x56\x67\x2c\x24\xa4\xd2\x48\xc3\x81\xdb\x3a\xf6\xa6\xf1\xe0\xdb\x0d\x84\xdf\x99\xbd\x0e\xe0\xf1\x06\xba\x70\xa0\x0e\x1c\xc6\xc3\x3f\x20\xf0\x87\x20\x74\x61\xa4\xa1\xc0\xed\xd0\x3f\x1b\xc6\x02\xd7\x2d\x00\xde\x19\x3d\xc0\xf1\x48\xc3\x80\xdb\x8b\x9d\xd3\x61\xdf\x95\x2e\xed\x03\xeb\x6f\x35\xd2\xb3\xcf\x22\x4f\xfc\xea\x4c\x85\x34\x32\x69\x75\xa1\xd0\x46\x1b\x1a\x69\x53\x32\x71\xa3\x99\x62\x33\x83\xbb\x2e\x43\x4e\x38\xad\xd4\x09\xb7\xd1\x75\xc2\x69\x53\x03\xe6\x88\x20\x72\x28\x01\x9b\x91\x80\x2e\x61\x1f\x99\x58\xc0\x10\x60\x82\x9e\xc1\x8f\xfd\x6c\x51\xd9\xf9\xb1\x97\xd1\x29\xe7\x65\xbe\xfc\xb5\x8d\x5e\x62\x9d\x76\x4e\x9a\xdb\x4c\xdd\x75\x5a\xa9\xc3\xa3\xa7\xd5\x33\x60\x21\x8c\xd8\xf4\x2a\x04\x52\xd8\xe3\x61\xc1\x14\x64\x66\x48\x40\x85\xf8\x1b\x12\x54\x01\x96\xa7\xad\x97\xf4\xa2\xa3\x37\xc9\x24\x1f\x7c\xd4\x1d\x3c\x94\x1c\x64\x57\xc7\x57\xe9\x13\x4c\x14\x4c\xbc\x52\xa3\x50\x0a\xfa\x38\x20\x8c\x42\x4c\x04\x43\x28\xc0\x56\x30\x7c\xfc\x8c\x4e\xb9\x5e\xa1\x17\x75\xbd\xc9\x43\xaa\xd3\x4a\xe9\x72\x1b\x5d\x37\x9d\x36\xd7\xcd\xd1\xab\xa4\x13\x4c\x18\x38\x6c\xa5\x04\x03\x27\xe4\xe3\x90\xc0\x09\x30\x11\x0e\x1c\x3f\x5b\xe1\xc0\xf1\x32\x3a\xe5\x7a\x51\x39\xac\x8f\x5d\x08\x2a\xa6\xdb\x74\xeb\x35\xc5\x96\x6b\x91\x97\x55\x0a\xae\xaa\x72\x5f\x8d\x28\xb0\x18\x99\x60\xb9\xa0\x84\xd5\xc2\xbd\x54\x56\xc5\x0a\x6b\xe4\x02\xa8\x60\xfd\xcb\xbd\x6e\x28\x56\x36\x22\x4f\xef\x0a\xce\xee\xb9\xcf\x8a\xda\x41\x7b\xcb\xa5\xb2\x66\xbe\xa7\xb9\xda\xd5\x6c\xf6\x9c\xf1\x82\x55\xc3\x99\xf5\xab\x5f\x27\xb2\x77\x92\xa7\xb8\xcc\xd3\xc2\xbc\xfd\x48\x57\x55\x9a\xcd\xfd\x33\x5e\x18\x69\x58\x17\xbe\xfa\xf5\x88\x66\x75\x65\xc6\x4b\x0a\x0d\x6b\xce\x57\x5f\xc9\xeb\xd2\x75\x37\xd6\x66\x6f\x77\xc1\xe4\xf6\x9d\xdf\x76\x6e\xb7\x27\x49\x61\x66\xb1\x61\x6b\xed\xaa\x72\x7c\xcb\x66\x6b\x98\x3a\x3a\x7d\xe7\xb6\x99\xea\xec\x75\x22\xa4\xcb\x26\xe8\x28\xc9\x59\xbd\xce\x65\x9b\xd3\xda\x8b\xe4\xf8\x96\x2d\xca\x10\x75\x74\x4d\x90\xcb\x26\xa4\xb5\x53\xc8\xf1\x2d\x1b\x88\x88\xba\x6f\xff\xfb\x7f\x00\x00\x00\xff\xff\xa4\x12\x39\x55\xaf\xfd\x00\x00")

func assetsRuntimeCompatibilityJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetsRuntimeCompatibilityJson,
		"assets/runtime.compatibility.json",
	)
}

func assetsRuntimeCompatibilityJson() (*asset, error) {
	bytes, err := assetsRuntimeCompatibilityJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/runtime.compatibility.json", size: 64943, mode: os.FileMode(420), modTime: time.Unix(1570344135, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRuntimeSupportedJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xe6\x52\x50\x50\x50\x50\x2a\xcf\xcc\xd3\xad\xb0\x30\x53\xd2\x41\xe2\x9a\x99\xc0\xb8\x39\x99\x79\xa5\x15\xc8\x02\xf9\xc5\x10\x2e\x57\x2c\x20\x00\x00\xff\xff\x6c\xfd\xc1\xad\x40\x00\x00\x00")

func assetsRuntimeSupportedJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetsRuntimeSupportedJson,
		"assets/runtime.supported.json",
	)
}

func assetsRuntimeSupportedJson() (*asset, error) {
	bytes, err := assetsRuntimeSupportedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/runtime.supported.json", size: 64, mode: os.FileMode(420), modTime: time.Unix(1570411694, 0)}
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
	"assets/runtime.compatibility.json": assetsRuntimeCompatibilityJson,
	"assets/runtime.supported.json":     assetsRuntimeSupportedJson,
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
		"runtime.compatibility.json": &bintree{assetsRuntimeCompatibilityJson, map[string]*bintree{}},
		"runtime.supported.json":     &bintree{assetsRuntimeSupportedJson, map[string]*bintree{}},
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

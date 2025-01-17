// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faucet.html (8.959kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x5a\x5b\x93\xdb\xb6\x92\x7e\x1e\x57\xf9\x3f\x74\xb8\x4e\x24\xed\x0c\x49\xc9\x93\x5b\x49\xa2\xb6\xbc\x4e\x36\xe5\xad\x3a\x49\x7c\x92\x94\xcf\xa9\x24\x0f\x10\xd1\x12\xe1\x01\x01\x06\x68\x4a\xa3\xa8\xf4\xdf\x4f\x01\xbc\x88\xa2\x34\x13\x3b\xb6\x1f\x34\x24\xd0\xe8\xfe\xd0\xdd\xe8\x0b\xe8\xf9\x27\xdf\xfc\xf0\xf2\xe7\x7f\xff\xf8\x2d\x64\x94\xcb\xc5\xd3\x27\x73\xf7\x17\x24\x53\xeb\x24\x40\x15\xf8\x11\x64\x7c\xf1\xf4\xc9\xd5\x3c\x47\x62\x90\x66\xcc\x58\xa4\x24\x28\x69\x15\x7e\x1d\x1c\x27\x32\xa2\x22\xc4\x3f\x4a\xb1\x49\x82\x7f\x85\xbf\xbc\x08\x5f\xea\xbc\x60\x24\x96\x12\x03\x48\xb5\x22\x54\x94\x04\xaf\xbe\x4d\x90\xaf\xb1\xb3\x4e\xb1\x1c\x93\x60\x23\x70\x5b\x68\x43\x1d\xd2\xad\xe0\x94\x25\x1c\x37\x22\xc5\xd0\xbf\xdc\x80\x50\x82\x04\x93\xa1\x4d\x99\xc4\x64\xe2\xd8\x38\x46\x24\x48\xe2\x62\xbf\x8f\xbe\x47\xda\x6a\x73\x77\x38\x4c\xe1\xff\x58\x99\x22\xcd\xe3\x6a\xae\xa2\x93\x42\xdd\x41\x66\x70\x95\x04\x0e\xae\x9d\xc6\x71\xca\xd5\x5b\x1b\xa5\x52\x97\x7c\x25\x99\xc1\x28\xd5\x79\xcc\xde\xb2\xfb\x58\x8a\xa5\x8d\x69\x2b\x88\xd0\x84\x4b\xad\xc9\x92\x61\x45\x7c\x1b\xdd\x46\x5f\xc5\xa9\xb5\x71\x3b\x16\xe5\x42\x45\xa9\xb5\x01\x18\x94\x49\x60\x69\x27\xd1\x66\x88\x14\x40\xbc\xf8\x7b\x72\x57\x5a\x51\xc8\xb6\x68\x75\x8e\xf1\xe7\xd1\x57\xd1\xd8\x8b\xec\x0e\xff\x85\x54\x27\xd7\xa6\x46\x14\x04\xd6\xa4\xef\x2c\xf8\xed\x1f\x25\x9a\x5d\x7c\x1b\x4d\xa2\x49\xfd\xe2\x05\xbd\xb5\xc1\x62\x1e\x57\x0c\x17\x1f\xc4\x3b\x54\x9a\x76\xf1\xf3\xe8\xf3\x68\x12\x17\x2c\xbd\x63\x6b\xe4\x8d\x24\x37\x15\x35\x83\x1f\x4d\xee\x43\x46\x7c\xdb\xb7\xe1\xc7\x10\x96\xeb\x1c\x15\x45\x6f\x6d\xfc\x3c\x9a\x7c\x1d\x8d\x9b\x81\x0b\xfc\xbd\x04\x67\x36\x27\xeb\x2a\xda\xa0\x21\x91\x32\x19\xa6\xa8\x08\x0d\xec\xdd\xe8\x55\x2e\x54\x98\xa1\x58\x67\x34\x85\xc9\x78\xfc\xe9\xec\xd2\xe8\x26\xab\x86\xb9\xb0\x85\x64\xbb\x29\xac\x24\xde\x57\x43\x4c\x8a\xb5\x0a\x05\x61\x6e\xa7\x50\x71\xf6\x13\x07\x2f\xb3\x30\x7a\x6d\xd0\xda\x5a\x58\xa1\xad\x20\xa1\xd5\xd4\xf9\x14\x23\xb1\xc1\x4b\xb4\xb6\x60\xea\x6c\x01\x5b\x5a\x2d\x4b\xc2\x1e\x90\xa5\xd4\xe9\x5d\x35\xe6\xcf\x70\x77\x13\xa9\x96\xda\x4c\x61\x9b\x09\xea\xc8\x29\x0c\x36\xcc\x19\xe7\x42\xad\xa7\xf0\x65\x51\x6f\x26\x67\x66\x2d\xd4\x14\xc6\x2d\xf9\x3c\x6e\x34\x38\x8f\xeb\x48\xf5\xf4\xc9\x7c\xa9\xf9\xce\x0d\x71\xb1\x81\x54\x32\x6b\x93\xa0\xa7\xdd\x2a\x02\x75\xe6\x5d\xdc\x61\x42\xd5\x33\x27\x53\x46\x6f\x03\xf0\x62\x92\xa0\x42\x10\x2e\x35\x91\xce\xa7\x30\x71\xd0\xaa\x15\x3d\x6e\x32\x94\xeb\x70\xf2\xbc\x9e\xbb\x9a\x67\x93\x86\x05\xe1\x3d\x85\xde\x2c\xad\x41\x82\xc5\x5c\x34\x4b\x57\x0c\x56\x2c\x5c\x32\xca\x02\x60\x46\xb0\x30\x13\x9c\xa3\x4a\x02\x32\x25\x3a\xff\x11\x0b\xe8\x86\xba\x36\xd2\x65\x93\x1a\x47\xcc\xc5\xa6\xda\xc4\xf1\xa9\xb7\x9d\x07\x11\x7f\x0d\xf5\x83\x5e\xad\x2c\x52\x78\xdc\x40\x87\x56\xa8\xa2\xa4\x70\x6d\x74\x59\x34\xd3\x57\x73\x3f\x08\x82\x27\x41\x69\x64\x50\xc7\x75\xff\x48\xbb\xa2\xde\x76\xd0\x6e\x52\x9b\x3c\x74\x2a\x37\x5a\x06\x50\x48\x96\x62\xa6\x25\x47\x93\x04\xaf\x3c\x9f\x9d\x2e\x0d\xbc\xc1\xe5\xed\x6b\x78\x99\x31\xa1\x80\x71\xee\x9c\x2f\x8a\xa2\xa3\x48\xef\x88\xe7\x98\xc2\x25\xa9\x96\xc8\xd1\x2d\x4b\x22\xdd\x52\x2e\x49\xc1\x92\x54\xc8\x71\xc5\x4a\x49\xc0\x8d\x2e\xb8\xde\xaa\x90\xf4\x7a\xed\xb2\x55\x85\xb7\x5a\x14\x00\x67\xc4\xea\xa9\x24\x68\x68\x1b\xd3\x30\x5b\xe8\xa2\x2c\x6a\xe3\x54\x83\x78\x5f\x30\xc5\x91\x3b\x53\x4a\x8b\xc1\xe2\x3b\xb1\x41\xc8\x11\xde\xdc\xbe\xbe\xea\xdb\x39\x65\x06\x29\xec\xb2\x3c\xb3\xf6\x3c\xae\xa0\xd4\x3b\x82\xfa\xdf\xbc\x94\x0d\xab\x76\x07\x39\xaa\x12\x4e\xde\x42\xe3\x82\x44\xb0\xd8\xef\x0d\x53\x6b\x84\x67\x82\xdf\xdf\xc0\x33\x96\xeb\x52\x11\x4c\x13\x88\x5e\xf8\x47\x7b\x38\x9c\xb2\x07\x98\x4b\xb1\x98\xb3\xc7\xbc\x16\xb4\x4a\xa5\x48\xef\x92\x80\x04\x9a\x64\xbf\x77\xdc\x0f\x87\x99\xdd\xe5\x4b\x2d\x93\xc1\x9b\xdb\xd7\x83\x19\xec\xf7\x62\x05\xcf\xa2\x7f\x62\xca\x0a\x4a\x33\x76\x38\xac\x4d\xf3\x1c\xe1\x3d\xa6\x25\xe1\x70\xb4\xdf\xa3\xb4\x78\x38\xd8\x72\x99\x0b\x1a\x36\xbc\xdc\xb8\xe2\x87\x83\xdb\x41\x8d\xfa\x70\x98\xc7\x6c\x31\x8f\xa5\x58\xd4\x93\x3d\xc5\xc4\xa5\x6c\x5d\x24\x76\x3e\xd2\xb8\xb0\x3f\x0e\x1e\x4e\x17\xcd\xb9\x7f\xaf\xc3\x16\x60\x6d\x7f\x2b\x08\xef\x70\x97\x04\xfb\x7d\x77\x69\x3d\x9b\x32\x29\x97\xcc\xe9\xa1\x42\xdf\x2e\xfa\x13\x9d\x63\x6e\x84\xf5\x65\xd0\xa2\x01\x70\x04\xfd\x2e\x67\xb5\x17\x7a\x48\x17\x53\xb8\x7d\xfe\x68\xdc\xf9\xb2\x77\x8a\x6f\x2f\x9d\xe2\x82\x29\x94\xe0\x7f\x43\x9b\x33\xd9\x3c\xd7\xe7\xe2\x78\xcc\xfa\x6b\x42\x17\x5e\x5b\x58\x6d\x94\x1e\xcf\x40\x6f\xd0\xac\xa4\xde\x4e\x81\x95\xa4\x67\x90\xb3\xfb\x36\x4d\xdd\x8e\xc7\x47\xcc\x8e\x2d\xb1\xa5\x44\x1f\x30\x0c\xfe\x51\xa2\x25\xdb\x86\x87\x6a\xca\xff\xba\x28\xc1\x51\x59\xe4\x3d\x45\x38\x81\x4e\xa5\x9e\xea\x68\xef\x46\x8b\x17\x81\xaf\xb4\x6e\x62\x7f\x17\x43\xcd\xb7\x93\xa0\x82\xc5\x9c\xcc\x31\x84\xcc\x89\xbf\x57\xf4\x36\xae\x2c\x7b\x28\x78\x57\x61\xcb\xed\xbb\x40\x34\x55\x49\xe0\x9c\x14\xfc\xeb\x3c\x26\xfe\xf7\x05\x3b\xbf\x5b\x32\x8b\xef\x22\xdd\xa7\xe6\xa3\x74\xff\xfa\x81\xe2\x33\x64\x86\x96\xc8\xe8\x5d\xe4\xaf\x4a\xc5\x3b\xbb\x7f\x73\xfb\xfa\x03\xa5\x97\x4a\x6c\xd0\x58\x41\xbb\x77\x15\x8f\xfc\x28\xbf\x7a\x3f\x41\x30\x8f\xc9\x3c\xe6\x61\x9d\xe7\x0b\x07\xb9\x79\x68\xfe\x3e\x7d\x32\x3f\x96\x96\x71\x0c\xdf\x49\xbd\x64\x12\x36\x0e\xea\x52\xa2\x05\xd2\xe0\x92\x20\x50\x86\x90\x96\xc6\xa0\x22\xb0\xc4\xa8\xb4\xa0\x57\x7e\x74\xe5\x53\xfd\xd3\x27\x57\x1b\x66\x80\x11\x61\x5e\x10\x24\x55\x39\xe4\x86\x2c\x9a\x4d\x55\xdf\xb9\x37\x17\x94\x4f\x66\xab\xb0\x1c\x04\xcd\x40\x73\xee\x20\x81\x5f\x7f\x9f\xf9\x82\x34\x8e\xe1\x1b\x5c\x09\x85\xc0\x9c\x46\x52\x57\xdc\x01\x65\x8c\x20\x35\xc8\x08\x2d\xa4\x52\xdb\xd2\x54\x70\x5d\xaa\x01\x07\xb9\xe1\x54\x33\x76\xe3\x85\x17\xde\xf0\x18\x66\xcc\x66\xa3\xaa\xb0\x33\x48\xa5\x51\xc7\xa9\x7a\xf8\x6a\xa5\x0d\x0c\xdd\x72\x91\x8c\x67\x20\xe6\x0d\xd3\x48\xa2\x5a\x53\x36\x03\x71\x7d\xdd\xd0\x5e\x89\x15\x0c\x1b\x82\x5f\xc5\xef\x11\xdd\x47\x4e\x04\x24\x09\x74\x44\x5d\x39\x69\x35\x17\x5b\x48\x91\xe2\x50\xdc\xc0\x64\x34\xab\x27\x97\x06\x59\x5d\x9e\x56\xd5\x67\xf5\xeb\x7e\x0e\xb3\x13\x75\x78\xf5\x9f\x28\xa4\x8a\xf5\x16\x18\xac\x85\x25\x28\x8d\x74\x2a\x71\x74\x95\x19\x1a\xad\x7b\xb2\xae\x2a\xce\x32\x50\xfd\x50\x27\x86\x1a\x7a\xc5\x24\xb2\xa8\xf8\xf0\xff\x7f\xfa\xe1\xfb\xc8\x92\x11\x6a\x2d\x56\xbb\xe1\xbe\x34\x72\x0a\xcf\x86\xc1\x7f\xb9\xfa\x6a\xf4\xeb\xf8\xf7\x68\xc3\x64\x89\x37\xb5\x89\xa7\xf5\xdf\x1b\xef\x02\x53\xff\x7b\x26\xf4\x06\xea\xc7\x29\x9c\xca\x3f\x8c\x46\xb3\x8b\x49\xb2\x93\xb6\x0d\x5a\xa4\xa1\xa3\x6b\x72\x59\x4f\x5d\x0c\x72\xa4\x4c\x73\xa7\x12\x83\xa9\x56\x0a\x53\x82\xb2\xd0\xaa\xd6\x0e\x48\x6d\x6d\xeb\x87\x0d\x41\x72\xe6\x15\x35\x75\x02\x0a\xb7\xae\x20\xfc\x49\xa7\x77\x48\xc3\xe1\x70\x2b\x14\xd7\xdb\x48\xea\x94\x39\x7a\xd7\x9e\x90\x4e\xb5\x84\x24\x49\xa0\xee\xd6\x82\x11\xfc\x0f\x04\x5b\xeb\xfa\xb6\x00\xa6\xee\xd1\x3d\x8d\xe0\x1a\xfa\xcb\x33\x6d\x09\xae\x21\x88\xab\x33\xe6\x32\xa2\xa1\x98\x15\x22\x18\x55\x27\xa3\x31\x88\x56\x39\x5a\xcb\xd6\xd8\x05\x8b\x1b\x54\xd4\xf8\x9c\xdb\x52\x6e\xd7\x90\x80\xb7\x5b\xc1\x8c\xc5\x8a\x22\x72\xf1\xb9\x76\x3e\xe7\xc0\x9e\x2a\x49\x40\x95\x52\xb6\x1e\x5b\x1d\x8f\x59\xe3\x8d\x5d\xe2\xc8\xc7\x4c\xf8\x24\x49\xc0\x45\x2b\xa7\x69\xde\xae\x73\xfe\x50\xc5\xd4\x51\xe4\xe2\xe5\x71\xc1\x68\x76\x74\xed\x13\x56\xc8\xff\x82\x17\xf2\x3e\x33\xe4\x17\xb9\xf9\xdc\xf5\x08\xb3\x2a\xd5\x75\x78\xf9\x81\x8b\xac\x54\x99\x2f\xd1\x3c\xc2\xab\x4a\x5c\x35\x2f\xaf\xdd\x57\x8a\x3a\x4b\x6f\x60\xf2\xe5\xe8\x22\x6b\x34\x46\x3f\xc4\x59\x69\xda\x0d\xf7\x92\xed\x74\x49\x53\x18\x90\x2e\x5e\xfa\x44\x33\xb8\x01\x27\x68\x0a\x2d\x83\x1b\xdf\x21\x4c\x61\xe0\xdf\xdc\xbc\xc8\xd1\xaf\xfa\x62\x3c\x1e\xdf\x40\xd3\x28\xff\x2f\x73\x67\xcf\x94\x78\xb8\x08\xc6\x96\x69\xea\xda\xe9\x0f\x80\x53\xb3\x68\x01\xd5\xef\x7f\x1b\x52\x9b\x10\x4e\x30\xc1\x67\x9f\xc1\xd9\xec\x89\xcb\xc6\x31\xfc\x83\x99\x3b\xf0\x45\xa4\xc1\x8d\xd0\xa5\x3d\x66\x97\x5c\x58\x2b\xd4\x1a\x98\x05\xae\x15\x56\x4b\xde\x2b\xd8\x9f\x01\xac\xa9\x60\x01\xe3\x3e\x3a\x17\x0d\x3b\xc9\xe0\x42\x8e\x38\xb2\x3d\x89\xff\xb5\x2e\x2e\xa4\x16\x91\x23\x7c\x92\x40\x10\x74\x56\x9e\x11\xb8\xf9\x86\xd3\x95\x45\xfa\xb9\xb2\xc0\xb0\xce\x85\x97\x92\xd5\xe8\xc6\x55\xc5\xe3\x51\x0f\xc0\xa1\x55\xea\x8b\xa2\x40\xc5\x81\xa9\x9d\x0f\x7f\xad\x46\x85\x22\x0d\xae\x43\x76\xe1\x4b\xba\x62\x5f\xa2\x0f\x44\xd5\x4a\xa7\xd6\x54\xe7\xb9\x56\x90\x40\x38\x99\x9d\x67\xcc\x8e\xfe\x8e\x7b\xea\x9b\xe4\x82\xc2\x7b\x66\x39\x55\x55\x8f\x36\x9c\x9c\x18\xe2\xc4\x46\x17\x8d\x71\xd5\x62\x16\xad\x22\x4f\x2d\xd4\x9a\xe8\x54\x53\x1d\xe0\x15\x8b\xeb\xc9\xbb\xe1\x6f\x67\x8b\xd2\x66\xc3\x1e\xc2\xd1\xac\x6f\x8e\x57\x84\x86\x11\xfa\x36\xc7\xab\x1f\x15\x09\x83\x67\x56\x00\xa6\x5c\x49\x14\x1a\x54\x1c\x4d\x53\x3a\xb8\x2e\xa9\xea\x69\xba\x56\xf2\xb7\xdd\x5d\xe7\xe9\x6c\xe6\x4c\xa1\x33\x10\xb0\x70\x15\x1d\x88\x30\x3c\x6e\xc3\x97\x5c\x5a\xa1\xeb\x7e\x7b\x1e\xef\x3d\xb3\xeb\x9a\x8e\x16\x25\x2b\x2c\x72\x48\xa0\xba\x9c\x1c\x8e\xa2\x52\x89\xfb\xe1\x28\xac\xdf\xfb\x2c\x9a\xf9\x3a\x13\x7a\x53\x55\xc0\xaf\x13\x08\xe6\x64\x5c\x45\x3d\x08\xe0\xfa\xd2\x71\x73\x59\x75\xb0\x68\x01\x74\x57\x02\xcc\x89\x2f\x7c\x9f\x56\x15\xfa\xbf\x05\xae\x8f\x5e\x1b\x5d\x2a\x3e\x75\x25\xd5\xf0\x8c\x2b\xdb\x30\x62\xc6\x33\x1d\xcd\xe0\x48\xee\xdb\xed\x29\xa4\xce\x38\x33\xa8\x7a\x39\xdf\x26\x43\xdb\x81\xfa\xb7\xa5\x36\x1c\x4d\x68\x18\x17\xa5\x9d\xc2\xe7\xc5\xfd\xec\xb7\xa6\x35\xf7\xd5\xff\x63\x48\x0b\x83\x8b\x33\x40\x69\xea\x2f\x53\xae\x21\x98\xc7\x8e\xe0\x2f\xb8\xb4\x5b\xed\x5e\x88\xc2\x85\x0e\x07\xda\x3b\xcb\x7a\x3c\x17\x9c\x4b\x74\x70\x5b\xee\xee\x00\x3a\xd3\x77\x8e\xd1\xa9\x40\xa8\x1b\x9b\x76\xc5\x01\x50\x5a\x7c\x98\xbc\xed\x90\x06\xce\xf2\xa1\xdb\xad\xf0\xda\xae\x9b\x2d\x3f\x6c\x06\x5e\x0d\xf5\xdd\x36\x2f\x8d\x2f\xa2\x86\x61\xed\x59\x37\x30\xb0\xae\xa6\xe3\x76\x30\x8a\xb2\x32\x67\x4a\xfc\x89\x43\x97\x77\x46\x95\x9a\x7c\xcb\x15\xf4\x03\xef\x19\x94\x63\x03\x3f\x68\x32\xd8\xa0\x56\xdf\xa0\xb1\xaa\x33\x20\x1c\xaf\x04\x06\xef\xa5\x9b\xcb\x32\xc2\x25\x33\xd0\x7d\x09\x9b\xc4\x0a\x46\x3b\xd9\xcd\xdc\x92\x99\x41\xd5\x6a\xfa\xfa\x5b\xe9\x6d\x32\xb8\x1d\xb7\x10\x2b\x03\x7b\xfb\x0e\x6a\x0f\xeb\x9b\xc1\x61\x6c\x8e\xe3\x02\x6e\xc7\x1f\x01\x2b\x67\x6a\x8d\x7d\xfc\x64\x44\x81\x1c\x58\x4a\x62\x83\x1f\x7f\x1b\x1f\xae\xe0\xf7\x06\xe8\xfc\xaf\xd1\x9c\x77\xcf\x13\xb4\x6e\xb6\x55\xec\x7f\xbb\x33\x06\xb1\x57\xef\x35\x04\x97\xb6\xf1\x90\x07\x9e\x92\xf5\xce\xf2\x83\xe7\xdc\x5f\x1c\x04\xa7\xf9\xc3\x55\xae\xed\x45\xd7\x28\xca\x28\x97\xc3\x60\x4e\xfe\x53\x85\x43\xdb\xae\xf7\xcb\xab\xe1\x6e\x91\x76\x38\xe9\x41\x5c\x17\x8e\xbd\x76\x09\x3a\x55\x47\xdb\x52\x35\x25\x06\xf8\x06\xed\x50\x35\x69\x3f\x11\x33\x04\x0c\x7e\x79\x05\x65\xc1\x19\xb9\x24\xa5\xc1\x25\x41\x9f\xac\xda\x2f\x3d\x4b\x66\x2c\xac\xb4\xd9\x32\xc3\xa1\x54\x24\xa4\x9b\xdf\x01\x33\xd8\x94\x72\x16\xe9\x95\x8b\x56\x1b\x26\x87\xfd\xde\xed\xd9\x70\x10\x75\xad\x3c\x18\x45\xc8\xd2\xec\x8c\xce\x27\xa5\x56\x68\x02\xdf\xfb\x32\x7e\xf8\x6c\x48\x99\xb0\xa3\x88\x11\x99\xe1\xe0\xc4\xfc\x83\x91\xb3\xe4\xe4\xd8\x49\xb5\xab\xe7\xdd\x43\xf4\x18\x87\x63\x51\xdc\x24\xfa\x86\x3a\xb5\x76\x58\xb9\xd1\xe0\xa6\xc3\xf9\xd4\x8b\x06\x9f\x0e\x1a\xeb\x1c\x4f\xf2\x71\x0f\xc9\x25\x18\x27\x8c\x07\xee\x40\x0d\xfa\xb2\x19\xe7\x2f\xdd\x59\x19\x06\x17\xce\xf4\xa9\x3f\x8c\x1a\x15\x57\x01\xf9\x31\xdd\x0a\xc5\xf1\xfe\x21\xc5\x0a\x3e\x18\x45\xb6\x5c\x56\xf7\x0b\xc3\x2f\x9a\xd6\xa9\xa1\xf2\x9e\xda\x8f\xf4\x67\x85\x82\x93\x70\x5a\x2c\x84\xbd\xe2\xe2\x91\xa4\x50\x49\xf4\x3b\x3a\xdc\x38\x35\x8f\x47\xed\x75\xd4\xb7\xd6\xd5\x4c\xc2\x66\xc0\x60\x8b\x4b\xeb\xfb\x7f\xa8\xbd\xdb\x5f\xc6\x54\x97\x2e\x2f\x7e\x7c\x75\xbc\x78\x69\xdd\xdf\xd7\x2c\xed\x37\xd7\x0b\xd7\x1a\x17\xbf\xf1\x6e\xb7\xdb\x68\xad\xf5\x5a\x56\x5f\x77\xdb\x7b\x8f\x98\x15\x22\x7a\x6b\x03\x60\x76\xa7\x52\xe0\xb8\x42\xb3\xe8\x70\xaf\x2f\x43\xe6\x71\xf3\x01\x32\xae\xff\x5f\xc5\x7f\x02\x00\x00\xff\xff\xde\x30\x0e\x5c\x69\x21\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x39, 0x4d, 0x5d, 0x6d, 0x7f, 0xf1, 0x26, 0x19, 0x7, 0xee, 0xd2, 0xb3, 0x2, 0xdf, 0xd2, 0x8c, 0x36, 0x87, 0x8e, 0x17, 0x50, 0xb6, 0x8f, 0xb0, 0xb5, 0xce, 0x73, 0xcc, 0x24, 0xd3, 0xd6, 0x8e}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"faucet.html": faucetHtml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}

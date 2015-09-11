package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _pr_review_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x59\x73\xe3\xc6\x11\x7e\xd7\xaf\x18\x79\xab\x4c\x72\x97\x0b\xae\xfc\x28\x46\xb1\x53\xeb\xb8\x2a\xc7\x1e\x65\xd9\x95\x07\x99\x91\x20\x60\x28\xc2\x22\x31\x0c\x06\x10\xa3\xac\xf9\xdf\xd3\x3d\xf7\x05\x92\x5a\xae\x2b\x76\x6c\x3e\x48\x22\xa6\xa7\xaf\xe9\xfe\xba\xa7\xa1\x79\x57\x17\x6d\xc5\x6a\xb2\x2c\x86\xbc\x6d\x46\xe4\xc3\x09\x21\x0d\x6d\xbb\xa6\x26\xf0\x9d\x7c\xfe\x39\xfe\xca\x5a\xf6\x77\xb6\xa1\xcd\xeb\x9c\x53\xf2\x65\xf8\x64\x38\x22\xe7\xf8\x6c\x7a\xb2\x3d\x39\x31\x0c\xef\x68\xfb\x2d\x7d\xa8\x28\x10\xf1\x6f\x58\xf3\x4d\xb5\xa4\xc3\x39\xfc\xe0\x63\xb2\xee\x96\xcb\x6f\xe9\xbf\x3a\xca\xdb\xef\x39\x6d\xe0\x41\x43\xe7\xb4\x69\x68\x69\x76\x48\x45\x1e\xf2\x06\x94\x91\x8f\x90\x01\x27\x17\xe4\xc3\x76\xaa\x56\x90\x9b\xd9\xa0\x56\x60\xe9\x3a\xa3\x79\xb1\x18\xc6\x3c\xc7\xc4\x28\xa7\x35\xd1\xcc\xa5\x38\xb3\x57\xad\xfa\xe4\x9a\x86\xf8\x82\xaf\xf0\xdb\x0c\xc4\xa7\x9e\xfe\xf4\x13\xb9\x9a\x4d\xfb\xb7\x65\xeb\x8e\x2f\x86\x46\x0b\x49\xb9\x15\xbf\xf1\xe7\xc9\x61\x1a\xb9\x6e\x42\x47\xc4\xa6\x0b\x69\x19\xfe\xa8\xf3\x95\xaf\x57\x35\x27\x46\x01\x9e\x01\xc1\x5d\xbb\x20\x7f\x24\xaf\xac\xb9\x4a\x85\x26\xe5\xc7\xd0\x81\xf8\xf1\x4e\xec\x4a\x7f\x43\x17\xf5\xad\x78\x6e\xea\x65\x90\x74\x96\x76\x17\xfc\x26\x74\x09\xe1\xa9\xf5\x40\xa7\xac\xf3\xb6\xa5\x10\xc9\x17\x60\xc3\xbc\xaa\xcb\x61\xc1\xea\x79\x75\x97\xa9\xe7\x91\x43\x5f\xc3\xb2\x6b\x89\xc3\x83\xab\x13\x46\x12\xbd\x7f\xea\x11\x76\x10\xca\xe8\x5d\x8f\xd2\x38\x4d\x1c\xa6\x25\x5e\xe5\x6d\xb1\xb0\x7a\x69\x21\x19\x5f\x2f\xab\x76\x38\x18\x0f\x46\xae\x6e\x6a\xd9\x55\x8d\x90\xc9\x04\xac\xe1\x0c\xce\x75\xc9\xee\x34\xc9\x98\x78\x27\x3d\x9a\x3a\x1b\x54\x5a\x0b\xd1\xe8\xda\xf7\x79\xbb\xd8\xbf\x4f\xc5\xa1\x96\x49\x74\x4c\x91\x39\x93\x29\x48\x8a\xbc\x6e\xc9\x2d\x25\xed\x82\x12\x0e\x9b\x49\xce\xc5\xdf\x98\xe6\x20\x54\xe4\x39\xe1\xb4\x2e\x69\x63\x18\x61\xd4\x49\x1f\x00\xc2\x0c\x8d\xeb\x32\x70\x06\xfd\xf7\xbb\xf9\x30\x80\x88\x11\xb9\xb8\xb8\x20\x2f\xcf\x30\x52\x2c\xb1\x09\xd6\xb3\x91\xef\x1a\x65\xe9\x75\xb6\xa9\xda\x05\xeb\x5a\x2b\x20\x02\x9f\x91\x13\xf2\x8e\xd1\x27\x91\xf1\xa8\xf1\x69\xe2\x20\xa4\x2c\x13\x8c\x9a\xdc\xcd\x47\xab\x87\xda\x9f\x39\xb9\x14\xea\x63\x04\xba\x59\xfd\x9a\x75\xe0\xe3\x0b\x12\xee\x97\xcf\xc1\x29\x67\xde\xb6\x7c\xd9\xd0\xbc\x7c\x94\x47\x55\xd5\x77\x42\x87\xaa\x86\xad\x9c\x8a\x90\x1a\x46\xb9\xc6\xc7\xe6\xd1\x75\x76\x4f\x1f\xf9\xd0\xcb\xc2\x91\x5a\xf5\xfd\x11\xca\x49\xc2\x87\x0a\x78\xd6\x18\xf0\xd3\xd6\xf8\x56\xbc\x8c\xd4\x56\xec\xa6\x7e\xd4\x44\x8c\x02\x69\x52\x5e\xcd\x5a\xdf\xfc\xb2\x9a\x03\x28\xd2\xba\xa0\x2e\x92\x85\x22\xbd\x84\xc1\x08\x1f\x22\xb3\x0a\x38\xbc\x9a\xc2\xaf\x3f\x24\xec\x80\x00\x76\x85\xb9\xf1\x44\xaa\x17\x2f\x7c\xd5\xd4\xb1\xe6\x75\xc9\x56\x26\x93\x2e\x7c\x06\x88\x01\xa0\xe5\x1b\x48\xcf\x6c\xbe\x64\xac\x91\x7f\xca\x4d\x50\x6c\x9f\xa7\xe4\x01\x5e\x9c\x8d\xae\x5e\xcd\xa6\x9e\xb0\xc8\xa3\x12\x42\x3d\xf1\x9e\xc9\xdb\x28\x05\xfc\x50\x0e\x19\x46\x71\x9f\xf2\x59\x13\xfa\x2b\x55\x6f\x22\x67\x25\x1d\x05\x6d\x8a\xdd\xbc\xdf\x51\xa1\x20\xed\xa5\x51\x6f\xb1\xf1\x04\x26\x6a\x56\xb0\xbe\xbb\x72\xf9\xc4\xd2\xf9\x69\x9c\xdd\xba\x69\xa5\xb3\x9c\x6d\x6a\xbf\xb2\xa2\x4b\xc4\x43\xe9\x09\x9f\xae\x4f\x0b\xb1\x1a\x5b\xa2\x1e\xef\x34\x40\xd2\xec\xd6\x5b\xff\xdc\x8e\x01\xf4\x2b\x2e\x01\x42\x01\xb0\xc7\x2d\x6a\x0f\xf1\xe9\xe5\x22\x7f\x93\xaf\x65\x7b\x63\x7b\xbe\xb9\x59\x31\x0d\xdf\x21\x7d\x90\xdd\x16\xf6\x3b\xb2\x22\x67\x7c\x91\xeb\xd6\xca\xe8\x68\x37\xf9\x0a\x3a\xb8\xfc\x6e\x4d\x6b\x5a\x0e\xe9\x03\xad\xdb\xaf\xf3\x36\xb7\x8a\x06\xe0\x2d\x8f\xc5\xd0\x65\xb8\x7c\xad\x2a\x60\x86\x35\x08\x4b\x75\x55\x4b\x1f\x69\x43\x31\xad\xc0\x1b\xef\x2d\x2b\xe1\x2e\x87\x4d\xdd\xad\x6e\xd5\x01\xa7\x3a\xe2\xa7\x74\xda\x1f\xb6\xa3\x69\xca\xc9\xa9\xe3\x10\x5a\x62\x3c\x26\x0b\x42\x02\xeb\x05\xce\x53\xce\xf3\x3b\x0a\x2c\x07\x7f\xda\x50\xce\xa0\x1d\xd8\xb0\xe6\xfe\x94\xbc\x65\x1b\xf2\x63\x87\x9d\x40\x05\xdd\x42\x5e\xdc\x13\xc8\x0e\xb2\xc9\xe1\x1b\x62\xc6\x77\x4d\xfe\x50\x41\xdf\xc0\xa0\xc6\x71\x2e\xd6\x18\x34\x11\x8d\x78\x24\x25\x93\x47\xd6\x35\xa4\x60\x25\xcd\x7e\xa8\x7f\xa8\x07\x32\x06\xb5\xc0\x17\x20\xf1\xd9\x33\xd3\x9e\x70\x41\xe0\x36\xf5\x9e\xfa\x07\xdc\x05\x7c\xd6\x5f\x0d\xc8\x0b\x43\x02\x7f\x0e\x8c\x02\x87\xde\x1a\x7c\x7e\x2f\xc9\x0d\x72\x14\x0d\x14\x70\xbb\x71\xd8\x6d\x4d\x76\xb9\x1b\x0c\x81\xe9\x49\xfc\xd5\x67\xc6\x78\x70\xd1\x6a\x05\x0e\xe4\x49\x1f\x41\x99\x2d\x0a\xba\x6e\xcf\xc9\xcd\xba\xb9\x56\x9e\x7d\xf7\xb7\x9b\x3e\xe2\xb2\x34\x66\x7b\x5b\x70\xe1\x4a\xf7\x55\x33\x72\xa5\xc0\x68\xd6\xc7\xa8\x58\xe4\x35\x7c\x49\xf2\x52\x6b\x57\x6c\x59\x5e\x3b\x2c\x6b\xba\xb1\x5f\xfb\xf8\xae\x1b\x68\x6c\xe0\xfa\x99\xb7\x1d\xf7\x98\xca\x47\x37\x36\x0e\x0a\x28\x5d\x2d\xfd\x0b\xe7\x1d\xf4\xe6\xab\x15\x64\x57\x94\x63\x63\xcd\x5c\x1d\x81\xdc\x72\x29\x38\xe9\x9e\xa9\x27\xbd\x17\x50\x17\x11\x60\x74\x1f\x35\x00\xd0\x28\xa1\x42\x0e\xc6\xbb\xf7\x75\xcd\xd2\x6c\x79\xef\x76\xcc\xd2\x8c\x73\x82\x61\x92\xce\xc0\x1f\x59\x55\xc3\x5d\x81\xc0\x65\xc1\x08\xd5\xe6\x0f\xc4\x13\x65\x06\xcf\x1f\x28\x4a\x8e\xec\xc5\xd0\x3b\xc7\x1c\x04\x45\xd5\xc6\xeb\x79\xc3\x56\xc0\xf4\xaf\x97\xef\xde\x66\x70\xab\x87\x95\x6a\xfe\x18\x88\x96\x37\x54\x74\xeb\x6e\xd6\x22\x25\x62\x66\x16\x7a\x90\x53\x0f\xea\x7e\xbf\x2e\xc1\xf7\x9f\x16\x76\x8f\x40\x5d\xab\xf3\x65\x2b\x24\xb2\xbc\xdc\x67\xb8\xd9\x0c\x81\x6d\x80\xe9\x90\xed\xa9\x23\x49\x03\xb7\x1d\x87\xb8\x32\xec\x34\xa4\x6d\x1e\xa3\xea\x08\xab\xe2\x40\xd6\x39\xdc\x07\x86\x9e\x61\xf2\x60\xe1\x4a\x07\x57\xb3\x21\x82\xd7\xd6\x67\x12\x08\x71\xd8\x04\x26\xa6\x19\x69\xcf\xe3\xb5\x77\x09\x70\xb1\x1b\x33\xb1\xf6\xf4\xd7\xf4\xd3\x0b\x5b\xd5\xfd\xd9\xcc\x21\x3d\x80\x6c\x76\x44\x07\xd0\x36\x1d\x9d\xda\x3e\xc6\x56\xe7\x63\x0a\xad\xeb\x8e\xa3\x2b\xe9\x25\x96\x51\x36\x17\xb7\x6a\xe9\x42\xd8\x4c\x49\x27\x33\xe4\xe0\x6a\xf8\xff\x5c\x0c\x27\x13\xf2\x86\x36\xb0\x0c\x8e\xe7\xca\x49\xd8\x65\xb8\x46\xbb\x67\x72\x90\xcd\x4d\xf2\xe2\x31\x9a\x06\xcb\x87\x0f\xba\xbe\x04\x55\xba\x1a\x2f\xe0\x81\xcc\x80\x10\x27\xaa\x73\xd9\x34\x7b\x56\x3e\xb1\x84\xfd\x5e\xc3\xfe\xc7\x35\x2c\x3a\x24\x5b\xc3\xaa\x12\xe2\xc4\x0a\xac\xf0\x4c\x95\x58\x8d\xe8\x72\x66\xf6\x16\x47\x6a\x41\x85\x93\x2b\x41\x59\xbb\x65\xe5\xa3\xc7\xb3\x90\xd2\x33\x5c\xc8\x40\x71\xb8\x13\x6b\xda\x62\x55\xba\xe3\x65\x41\xa1\x66\x8f\x90\x5b\xde\xf0\x51\x87\x93\x03\xca\xea\x91\x99\xd6\x39\x47\x26\x27\x75\xce\x40\x06\x24\xb9\x80\xef\x4c\x99\x56\x70\xd3\xd0\x9c\x78\x77\x2b\x5d\x3b\x3c\x7b\x35\xd2\x9a\x90\x70\x0a\xda\xb4\xc9\x39\x1f\x2e\x28\xfb\xb2\x86\xae\x97\x79\x41\x87\x93\x7f\x7e\x35\x81\xa8\x1a\x78\x93\x4c\x3b\xe3\xb2\x5c\x41\xbf\x24\xd3\xd3\x53\xe1\x23\xb3\x39\x1a\x80\xb9\xd3\xd7\x01\x5a\x89\x41\x0e\xbf\x55\x48\x83\xf2\x96\x96\x6f\x2a\xac\x83\xb0\x7a\xf5\x6a\xe6\xbf\x3b\x71\x85\x17\xf8\x7e\x65\xc0\xee\x07\xe7\xc1\x34\x0b\x7c\x83\x27\xea\xf6\x0d\x10\x3f\x3b\x3b\x05\xfd\xc1\xf3\x52\xdb\xc3\x51\x94\x2d\xeb\xbe\xb4\xb0\xee\x39\x55\x5e\x73\x9a\x06\xdb\x4a\xba\xa4\x2d\x0d\xe0\xcc\xc6\xef\x2c\xa4\x37\x29\xb7\xc3\x8c\xfd\xf9\x9b\xd0\x7d\xdd\x7c\x4d\xdb\xbc\x5a\xc6\xfd\x9d\x7a\x0e\x12\x23\xed\xf7\x96\xe6\x20\xa6\xed\xc7\x03\x57\x23\xdb\x02\x29\x19\xf0\x0e\xae\x5b\x1c\x01\xc4\x2e\x23\x5e\x26\x91\x92\x94\xac\xa6\x88\x86\x4e\x42\x85\xca\x26\xab\x40\x55\x8e\xc9\x67\xf6\xfa\xb7\xc6\xd3\x28\x4f\x3f\x8b\x36\x07\x2f\x49\x7a\x0c\x89\x56\x09\x49\xd8\x96\xa0\x8a\x0a\x46\x9a\x85\x53\x2f\xbc\xdd\x47\xd7\x0e\x5f\x17\xbf\x8e\xb8\x9f\xd8\x2f\x7e\x2c\xf9\x8d\xab\x4b\xa6\xff\xba\x05\x7f\xdd\x4f\x83\xcc\x85\x3b\xb1\x97\xba\x18\x57\x98\xf3\x67\xb3\xd4\x14\x18\xa8\x69\x69\xef\x2f\x8a\x70\x1a\xd1\xd9\xb7\x58\x48\xf2\xc5\x2c\xa6\x80\x8b\xb2\xce\xd5\xab\xc4\xf2\x11\xe8\xb1\x13\x3f\xd2\x08\xf2\x04\x0c\x89\x76\x3a\xd3\xcf\x54\xb6\x11\xd7\xd2\x6b\xac\x1f\xa6\xae\xa4\xee\x72\x90\xe9\xbb\x3a\x50\xf7\xf3\x71\xef\xc6\xec\x67\xfb\x54\x49\x1e\xcb\x34\xc7\xf8\xe9\x36\xf6\x98\x8f\xb8\x26\xa4\x66\xa2\xe6\xca\x56\xb3\x97\x44\xcc\x7e\xc7\xc6\xa7\x09\x81\x9f\x06\xa6\x43\xae\x1f\x8d\xd3\xbf\x79\x98\xea\xc5\x7e\x79\x25\xb3\x88\x02\x87\x25\xbf\xc5\xf5\xa3\x0f\xdb\x5c\x74\xdb\x81\x74\x93\xe7\x01\xe8\x35\x74\xc5\x1e\xe8\x13\x70\xef\x17\x0b\x47\xe1\xc6\x64\x53\x23\xed\x8a\x1a\x9a\x9f\x27\x57\x8e\x3c\xab\xe7\x93\xe0\xac\xe4\xa0\xb5\xe7\xac\xf0\xcd\x9c\xac\x30\xa9\x53\x43\xfd\x0f\x28\x56\x2d\xf3\x89\xbe\x48\x11\xfd\x62\x23\x20\x59\x90\xfc\xf3\xd7\x7e\x88\x9c\x24\x3f\x3e\xb1\xf4\x46\x3f\x18\xeb\x75\x85\xc4\x7d\x92\x52\xf5\x21\x19\x9c\x66\x4b\x6a\xc7\x11\xb0\xbb\x1f\x78\x0f\x85\xde\x3d\xe0\x7b\x00\xfc\x7e\x6a\x00\xde\x0d\xc1\x49\x10\xfe\xb9\x0a\xe3\x1e\x78\x37\x29\x88\xe8\x2e\x53\xb9\xc4\xd7\x63\x62\x51\x86\x52\xb2\x69\x08\xbe\x1f\x87\x29\x1a\xf4\x19\x1c\xa3\x03\x24\x02\x22\xf2\x7b\xfa\x67\x9c\x3e\xe0\xf8\x59\x8e\x32\xce\x09\xaa\xef\x0e\x8b\xce\xfb\xe3\x6e\x6b\xa5\xc4\x6f\x5f\x0d\x73\xc7\xc4\xa4\x6a\xf2\xbd\xcf\xaf\xef\x0a\x1d\xcc\x7c\xff\x21\xb5\xd2\x17\x43\xd4\xca\x99\xf3\xda\xcf\x47\x0f\x74\xed\xe7\xd0\xd1\x6e\x24\x74\xef\x90\x37\x2d\xe3\x25\x71\xa7\xbd\x49\x01\xa9\xee\x37\x39\x00\xf6\xb7\x04\x8f\x7a\xb2\x29\x18\x90\xda\xcf\xc7\xcf\x01\x3e\x55\x4b\xfb\xfb\x08\xe1\x57\xdb\x9b\x1f\x37\x42\x70\xfe\x91\x46\x8e\x93\x27\x13\xf2\xbe\x61\x78\xee\x64\xc1\xd8\x3d\x29\x01\x33\x4e\x30\x40\xc5\x7c\xf7\xbb\xc7\x35\x15\xd1\x38\x70\xb1\x75\x80\x2d\x9c\x9d\xff\xe6\x32\x2b\x05\x19\x13\x28\x3a\x90\xb1\xbb\xeb\x7f\x5b\x40\xb8\x3c\xe0\x84\x2c\x31\xa4\xbe\x56\x63\xe5\x1d\xc2\x64\x24\x94\x01\x85\x1c\x71\xbb\xfa\x46\xda\xc4\xf3\xf2\x5d\xea\x1c\x68\x3a\x7f\xac\x8b\x45\xc3\xea\xea\x3f\x34\xb6\x3f\x7e\xcb\x8c\xde\xff\x6f\x00\x00\x00\xff\xff\xb1\x59\xf1\x7a\x17\x2f\x00\x00")

func pr_review_js_bytes() ([]byte, error) {
	return bindata_read(
		_pr_review_js,
		"pr_review.js",
	)
}

func pr_review_js() (*asset, error) {
	bytes, err := pr_review_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "pr_review.js", size: 12055, mode: os.FileMode(420), modTime: time.Unix(1441863750, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _pr_review_peer_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x4b\x73\xdb\x36\x10\xbe\xeb\x57\xac\x26\x33\x01\x39\x91\xe9\xe4\x6a\x8d\xd3\xb4\xce\xa5\x8f\xb1\x3d\x71\x32\x3d\xd8\xae\x0c\x91\x2b\x89\x31\x45\xb0\x00\x68\x55\x4d\xfc\xdf\xbb\x00\xf8\x00\xa9\x87\xe5\x34\x97\xce\x94\x07\x5b\x04\x16\xbb\xdf\x2e\xf6\xc9\x59\x99\xc7\x3a\x15\x39\x64\x71\xa0\xb4\x0c\xe1\xcb\x00\x40\xa2\x2e\x65\x0e\xf4\x0e\x2f\x5f\x9a\x7f\x91\x16\xbf\x89\x15\xca\x33\xae\x10\x7e\xe8\xaf\x04\x21\x9c\x98\xb5\xf1\xe0\x71\x30\x68\x18\x16\x65\x96\x7d\xc0\x3f\x4b\x54\xfa\xa2\xc0\x1c\x93\x00\x1f\x30\xd7\xef\xb9\xe6\x23\x28\x8b\x84\x6b\x74\xd2\x1e\xb8\xf4\x89\x3f\x29\x94\x70\x6a\xf0\x34\xf4\x91\xd9\x9e\x48\xb7\x1f\x95\x44\x10\x65\x62\x9e\xe6\xe1\xb8\x3a\x2e\xf1\x21\x45\x83\x46\x94\xb9\xa6\xc3\xb1\xc8\x67\xe9\x3c\xea\x2c\xd7\xb4\x85\x3c\x2f\x97\x53\x2b\xa3\x15\x90\xdb\xa5\x9a\x64\x89\x4a\xf1\x39\x8e\x07\xf4\x9e\xce\x20\x18\xfa\x70\xa1\xde\x26\x06\xec\xc7\x15\x2a\xb1\x44\x58\x09\x79\x3f\x84\x73\xb1\x82\xcf\xa5\xd2\xa0\x52\x0d\x53\x1e\xdf\x03\xcf\x13\x58\x71\x7a\x9b\x09\x09\x1f\x25\x7f\x48\x15\x68\x01\x05\x57\xca\xee\xdd\x31\x78\xd5\x43\xff\x0a\xd8\x1d\x08\xbd\x40\x69\x49\xdd\x26\xac\x45\x29\x49\xad\x04\xa3\x9b\xfc\x26\x67\x16\x5a\x8b\xe4\x15\x41\xb9\xc9\x5f\xbc\x78\x01\x1f\x1c\x79\x2c\x96\x4b\xe2\xaf\x2c\x69\x9f\xf2\x08\x78\x1c\x63\xa1\x4f\xe0\xae\x90\x93\x4a\xc0\xc5\xaf\x77\x3b\x88\x0b\x99\x12\x2c\xa5\xb9\x2e\x55\xe7\x88\x5b\xaa\x8f\x3d\x02\x66\xe4\x1d\x5f\x36\x38\x5c\x19\x03\x89\x19\x90\x4a\x30\x4b\x33\x54\x40\xba\x62\xe5\x02\xc9\xf0\x29\x1b\xe4\x88\x89\x67\x08\xc3\x25\x5e\xf0\x7c\x8e\x2a\x72\x72\x8d\x29\x62\x89\xc4\xec\x67\xa5\x4a\x3c\x23\xd5\xe9\x5a\x83\xfa\x9e\x47\x35\x18\xeb\x2d\x8e\xf0\xca\x22\x0f\x2c\xd4\x1d\x5e\xb6\x40\x9e\x44\x6a\xc1\x47\x96\x88\x91\x07\x27\x69\x3e\x67\xa3\x7d\x67\x4a\x99\x55\xe4\x97\xb4\x0c\xd5\x72\x8d\xdc\xf8\x01\xb1\x80\xad\xea\x42\x81\xa4\x2c\xab\xa5\xd5\x36\x66\xf4\x6e\x71\x2b\xfe\x80\x46\x60\xa3\x96\x39\x74\x52\xb1\xac\x88\x27\xb1\xe1\xc6\x46\xf0\xcb\xd5\xc5\x79\x44\x01\x49\x5b\xe9\x6c\x1d\x74\x84\x85\x07\xb3\xc3\x64\x32\x5d\x6f\x72\xbb\xbe\x35\x2c\xc8\xea\x5b\x43\xbd\xb6\x7e\x63\x21\x2f\xc6\xb7\x05\x5e\x6a\xae\xac\x17\x7e\x8a\x6c\x8d\xf2\x9c\x2f\xb1\x9f\x07\xdc\x4e\x2f\xf8\xa7\x22\x59\x77\x78\xc6\x0e\x43\x64\x36\x22\xc2\xbd\x0c\x1a\xda\x78\x99\xd8\xd0\x99\x44\xc8\xe3\x45\x60\x29\x54\x91\xa5\x3a\xa0\xf8\x61\xe1\x08\x1a\x9d\x82\xda\x6b\x2a\x8f\x36\x79\xa0\x5a\x8a\x52\x42\xf1\xd7\xc5\x2c\xf0\xee\x29\x84\xd3\xd3\x53\x78\x5d\x53\x83\x91\x44\xa0\x26\x11\x79\xbc\x46\x19\x54\xab\x46\xf2\x92\x17\x0d\x27\x55\x4e\x9d\x65\x83\x37\xaf\xc3\x1a\x09\x74\x81\x14\x5c\xea\x96\xaf\x79\xaa\xfc\x6c\x36\x2a\xfd\x28\xd5\x15\x19\x8f\x31\x38\xfe\xe3\xdd\xf1\x08\x18\x0b\xc7\x0d\xfd\x63\x38\x6a\x7e\xb7\x5c\x09\xdf\x56\xa6\xc3\xa1\xb5\x51\x73\xb8\xfa\x15\x56\x19\x07\x4c\x6e\x55\x22\x43\x73\x09\x01\x33\x5a\x1a\x7f\xa6\xff\xd1\x67\x91\xe6\x16\x7c\x4b\xab\x56\xa9\x26\x33\xd3\xee\xf5\xeb\xdb\x6e\xc9\xf0\x85\xc7\xa6\xac\x30\x71\xcf\x4e\x06\x1e\x20\xe7\x33\xef\x51\xf3\x34\x53\x64\xcb\x39\xea\xcb\xd6\xd1\xaa\xf5\xc6\x81\x3d\x85\xf7\x94\x94\x86\x5f\xbf\x8a\xf8\x27\x3b\xf1\x72\xa5\xed\x51\xc1\x93\x03\xc3\x2f\x84\xaf\x5f\xb7\x16\x20\xd2\xff\xca\xdd\xf5\x2e\x89\xc9\x4f\xeb\xc3\xc5\xb9\xf0\xb4\xd2\xd8\xf5\x2d\x1b\xf7\x4d\xd7\xaf\x89\x36\x8c\xc9\x65\xc8\xf8\x7d\x05\xf7\xe0\xd9\x7a\xd0\xe1\x0c\x3b\x22\x4d\x7c\xf8\x81\x4b\xd1\xd0\xbb\x81\xae\xbb\x3d\x91\xb3\xd9\x19\xcf\x9b\xcc\x29\x56\x2e\xc7\xd4\x19\x95\x75\x00\xd7\xae\xeb\xaf\x3d\xf6\xa1\xb5\xc8\x9b\xe8\x6d\xd1\x86\xf0\x16\x8e\xde\xf4\xf1\x1d\x1f\x03\xcf\x08\x23\x65\x97\xfa\x34\x4c\xd7\x54\x80\xa8\x86\x1b\xef\x79\x86\x32\xef\x4c\x8c\x78\xd6\x31\x09\xbf\xcf\x7b\xf8\x4c\xad\x3c\x8d\x8a\x52\x2d\x7c\x75\xc6\x9b\x64\xee\xae\x8f\x8e\x36\xee\xac\xeb\x26\x6f\xfd\x14\xe6\x6b\xe6\x57\xcc\xf6\x69\xc3\xa9\x5b\x2e\xdb\xa7\x57\x38\xb7\x1d\x6d\xaa\xa6\x77\xea\xdf\xd5\x4f\x5f\xbc\x5f\x49\xdb\xa7\x63\xa2\x4e\xe7\xf2\x9d\xd4\x56\x25\xb5\x58\x6a\x13\xd0\xf3\xd5\x4e\x44\x8e\xdf\xa6\xd7\x53\x5e\xd9\xf6\x8a\x45\x86\x7a\xc3\x03\x3b\xde\xf6\x9d\xfb\x8f\xe7\xb0\xdd\xd1\x87\xb4\xfe\x1f\xfa\x99\x68\x4a\x2a\xdf\x8f\x7b\xa5\xc5\x75\xaa\x54\x5e\x5a\xc9\xff\x27\xfa\x6f\x4c\xf4\xde\x80\x64\x26\xa0\xdf\xab\xa8\xac\xbc\x75\x26\xc5\x72\x47\x47\xef\x42\xf4\xae\x1e\x60\x5a\x8e\xbb\xa6\x18\x9b\x6f\xdb\x61\xc7\x3d\x55\xe7\xd6\x02\xf4\x5b\xa5\x5a\x66\x3f\x87\xf5\x47\x1a\x1f\x5f\xc7\xe5\xbb\xaa\x1e\x3c\x55\x6c\x71\x3d\xd7\x38\x99\xbf\x8f\xe1\xce\xc1\xf8\x93\x9b\x81\x0e\xee\x96\xbb\x7d\xf2\xb3\x9a\xa3\xff\xec\xec\x63\x6c\x47\xa5\xf8\x52\x0a\x93\x50\x61\x21\xc4\x3d\x90\xd5\xf8\xc0\x54\x2f\x0b\xe6\xe3\xba\x70\x0d\x07\xf3\x11\x31\xf3\xed\xa2\x05\xcb\x9d\xf1\x2d\x99\xb0\xdf\x23\x98\x33\xf6\x9e\xef\x14\xe6\xe2\x5c\x6d\xd8\x22\xcb\x4e\x2e\x93\x6a\xd6\xd8\x23\xcc\x59\x3d\xe9\x51\xb8\xb9\xc7\xc7\xbb\x81\x66\x73\x94\xda\x07\xe7\x40\xd5\xd5\x3a\x8f\x17\x52\xe4\xe9\xdf\xf8\xb4\xfe\x23\xd0\xb2\x44\xeb\xbe\xff\x04\x00\x00\xff\xff\xfa\xcf\x50\x47\x28\x12\x00\x00")

func pr_review_peer_js_bytes() ([]byte, error) {
	return bindata_read(
		_pr_review_peer_js,
		"pr_review_peer.js",
	)
}

func pr_review_peer_js() (*asset, error) {
	bytes, err := pr_review_peer_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "pr_review_peer.js", size: 4648, mode: os.FileMode(420), modTime: time.Unix(1441955553, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"pr_review.js": pr_review_js,
	"pr_review_peer.js": pr_review_peer_js,
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

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"pr_review.js": &_bintree_t{pr_review_js, map[string]*_bintree_t{
	}},
	"pr_review_peer.js": &_bintree_t{pr_review_peer_js, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../transactions/destroy_nft.cdc (499B)
// ../../../transactions/mint_nft.cdc (1.304kB)
// ../../../transactions/scripts/borrow_nft.cdc (571B)
// ../../../transactions/scripts/get_collection_length.cdc (455B)
// ../../../transactions/scripts/get_nft_metadata.cdc (1.656kB)
// ../../../transactions/scripts/get_total_supply.cdc (118B)
// ../../../transactions/setup_account.cdc (888B)
// ../../../transactions/transfer_nft.cdc (1.133kB)

package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
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

var _destroy_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xc1\x4e\xc2\x40\x10\x86\xef\x7d\x8a\x3f\x3d\x68\x39\xd8\x5e\x8c\x87\x06\x25\x04\x25\xe1\xc2\xc1\xe0\x03\x2c\xdb\x29\x6c\x2c\x33\xcd\x74\x9a\x6a\x0c\xef\x6e\x0a\x52\xc0\x30\xa7\xa6\xf3\x7f\xdf\xcc\x6c\xd8\xd5\xa2\x86\xa5\xf0\xbc\xe5\x4d\x58\x57\xb4\x92\x4f\x62\x94\x2a\x3b\xc4\x69\x9a\x79\x61\x53\xe7\xad\xc9\xfe\x67\x52\x5f\xf8\x38\xfa\x13\xbc\x7d\xb9\x5d\x5d\xd1\x72\xbe\xba\x85\x9e\xbb\x47\x28\x32\x75\xdc\x38\x6f\x41\x38\x09\x45\x8e\x8f\x05\xdb\xd3\xe3\x08\x3f\x11\x00\xd4\x4a\xb5\x53\x4a\x9a\xb0\x61\xd2\x1c\xd3\xd6\xb6\x53\xef\xa5\x65\x3b\x45\xfa\xaa\xc8\xe0\xa5\xaa\xe8\xe0\x79\xa7\x12\xcf\x38\x22\xe9\x5a\x54\xa5\x1b\xdf\x5d\x0c\x9e\x0d\xc9\x97\xa4\x5f\x31\x47\xd6\x98\xa8\xdb\x50\xb6\x9c\xaf\xce\xdd\xd1\xe0\xef\x6b\x32\x41\xed\x38\xf8\x24\x9e\x49\x5b\x15\x60\x31\x1c\xe5\x70\x50\x2a\x49\x89\x3d\xc1\x04\xb6\x25\x48\xc7\xa4\xf7\xcd\xc5\x56\xf1\x28\x1a\x7c\x59\x86\x2e\xd8\xb6\x50\xd7\x1d\xd2\xc3\x63\xdd\x46\xaf\xee\xe4\xd2\x30\x7e\xb8\x3e\x37\x3d\xd9\x92\xd3\xc7\xe2\x35\x47\x28\x2e\x46\x16\xd4\x98\xca\x77\x8f\x1f\xfe\xed\xa3\x7d\xf4\x1b\x00\x00\xff\xff\x09\x2c\xf1\x05\xf3\x01\x00\x00"

func destroy_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_destroy_nftCdc,
		"destroy_nft.cdc",
	)
}

func destroy_nftCdc() (*asset, error) {
	bytes, err := destroy_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "destroy_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4a, 0x31, 0xab, 0x52, 0x78, 0xd1, 0x93, 0x72, 0xd5, 0xb8, 0x0, 0xaa, 0x41, 0xd4, 0xe1, 0xae, 0xb6, 0xa3, 0xbb, 0x52, 0xff, 0x2d, 0x23, 0xa7, 0x84, 0xed, 0x95, 0x47, 0x26, 0xb4, 0x2, 0xa1}}
	return a, nil
}

var _mint_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\xcd\x8e\xdb\x3c\x0c\xbc\xfb\x29\x88\x1c\xbe\xcf\x0b\x2c\xec\x7b\xd0\xee\x62\x1b\x34\x40\x0f\x0d\x0a\x34\x2f\x20\xcb\x8c\x4d\x54\xa6\x0c\x89\xda\x6c\xb1\xc8\xbb\x17\xb2\x1c\xff\xa4\x69\xeb\x43\x62\x4b\x9c\xe1\x70\x48\x52\xd7\x5b\x27\x70\xb0\xbc\x0f\xdc\x50\x65\xf0\x68\x7f\x20\xc3\xc9\xd9\x0e\x36\x45\x51\x6a\xcb\xe2\x94\x16\x5f\xde\xc6\x14\xba\xd6\x9b\x6c\x24\xf8\xfc\xa6\xba\xde\xe0\x61\x7f\xbc\x07\x9d\x6f\x13\x28\x2b\x4b\x38\xb6\xe4\xc1\x6b\x47\xbd\x40\xf0\xe8\x41\x5a\x84\xc3\xfe\xf8\x95\x58\xd0\x81\x43\x6f\x83\xd3\x08\x62\xa1\x23\x16\x50\xc0\x78\x8e\x01\x11\xfc\x45\xa0\x0b\x5e\xa0\x42\x70\x81\xe1\x4c\xd2\x0e\x78\xa5\xb5\x0d\x2c\x20\xad\x12\x68\x55\x22\xed\xd6\x8c\x11\xef\xc5\x3a\xac\x81\x18\xca\xf8\xaa\x1a\x2c\xa7\xd4\x59\x26\x4e\xb1\x57\x5a\xc8\x72\x9e\x01\x00\x38\xd4\xd4\x13\xb2\x6c\xe1\xa5\xae\x1d\x7a\xff\x38\x9c\xb3\xea\x70\x0b\xdf\xc5\x11\x37\xe9\xa4\xc6\x54\x12\x59\x5e\x5f\x48\x1b\xba\x8a\x15\x99\xf9\xf8\x01\xde\xb3\xe1\xae\x2c\xc1\x58\xad\x0c\xbc\x2a\x47\xaa\x32\x08\x27\xeb\x06\x8d\xc4\xcd\xba\x84\x13\x3a\x64\x8d\x03\xcc\xa0\x8c\x17\x5b\xf8\x6f\x61\xf1\xa2\x92\x18\xd6\x3b\xec\x95\xc3\xdc\x53\xc3\x31\xf4\x25\x48\xfb\x92\x7c\x8a\x0a\x60\x7c\xca\x12\x2a\xeb\x9c\x3d\x83\x9a\xf3\x44\xf7\xff\xd0\x17\x62\x18\xad\x9b\x28\x3c\x9a\x53\x31\x4a\xfd\x08\x29\x5f\x91\x48\x3f\xdc\x15\xf8\x94\xc7\x61\xd9\xde\x69\xc2\xc3\x44\x1a\x9f\xe7\x67\xe8\x15\x93\xce\x37\x3b\x1b\x4c\x0d\x6c\xe5\xef\x62\x47\x5f\x36\x89\xe6\x92\x8c\xc0\x37\xd4\x41\x70\x5d\xf3\xa7\x44\x13\x71\x53\x97\xff\xf7\xd0\x87\xca\x90\x1e\xa8\xb4\x35\x06\x87\x61\xb8\xf1\xff\xda\x03\x87\x1a\xe9\x75\xa8\xb9\x41\x19\xad\xcd\x27\xb6\x75\x29\x45\x83\xb2\x53\xbd\xaa\xc8\x90\xfc\xcc\xcb\x94\x28\x16\xbe\x9b\xf2\xdc\x20\xae\x0e\xbe\xff\xb6\x81\x33\xe4\xdb\x40\x73\x79\xca\xff\x6d\x5c\xb3\x94\x7c\xd7\xbc\x99\x76\xf3\x90\x2d\xcd\x8a\xbd\x99\xa2\x14\xd7\x50\x63\x6f\x3d\x09\x90\x5c\xf1\x4b\x13\x67\xe7\xee\x8d\xc8\xf0\x77\xd8\x1f\xf3\x95\xe2\xc5\xaa\x5d\x45\x3e\xae\x02\xd2\xce\xc5\xdf\xf5\xf9\x6a\xf3\x16\x1f\xeb\xa8\xc5\x1a\x4e\xaf\x73\xc4\x75\x5c\x2e\xd9\xaf\x00\x00\x00\xff\xff\xdb\xac\x03\xe6\x18\x05\x00\x00"

func mint_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_mint_nftCdc,
		"mint_nft.cdc",
	)
}

func mint_nftCdc() (*asset, error) {
	bytes, err := mint_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mint_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xaa, 0x4, 0x75, 0x3, 0xdd, 0x5c, 0x12, 0x12, 0xd5, 0x5d, 0x4c, 0xc8, 0x17, 0x16, 0x32, 0xf9, 0x5b, 0x85, 0xcf, 0xb8, 0x8a, 0x6a, 0xf2, 0xdc, 0xfd, 0x7, 0x16, 0xe4, 0x8f, 0x44, 0xbc, 0xd}}
	return a, nil
}

var _scriptsBorrow_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4f\x4b\xc3\x30\x18\xc6\xef\xfd\x14\x0f\x3d\x48\x0b\xd2\x5c\xc4\xc3\x70\x8e\x39\x2c\x78\x29\x22\xf5\x2c\x69\xfa\x76\x7b\x31\x4d\x42\x9a\xa2\x32\xf6\xdd\x65\xeb\xda\x4e\xc5\x9c\x12\xf2\xfc\xde\x3f\x3f\x6e\x9d\xf5\x01\x85\x35\x79\x6f\xb6\x5c\x69\x2a\xed\x3b\x19\x34\xde\xb6\x88\xb3\x4c\x64\x99\x50\xd6\x04\x2f\x55\xe8\xc4\xef\x58\xa6\x6a\x15\x47\xe7\x1a\x8f\x9f\xb2\x75\x9a\x8a\xbc\xfc\x87\x9e\x03\x03\x17\x09\x81\x72\xc7\x1d\x3a\xe5\xd9\x05\x54\xd6\x7b\xfb\xd1\x41\x1a\x4c\x45\x24\x94\xd5\x9a\x54\x60\x6b\x22\xd7\x57\x68\x7a\x83\x56\xb2\x49\x64\x5d\x7b\xea\xba\x05\xd6\xc3\xe5\x1a\x5c\x2f\xf0\xfa\x64\xc2\xed\x4d\x8a\x7d\x04\x00\x9a\x02\xa4\x52\xb6\x37\x01\x4b\x6c\x29\xac\x87\xc7\x08\xa7\xd1\x14\x9b\xdb\xbc\x50\x83\xe5\x88\x9d\xfe\x8f\x27\xdb\x52\xd8\x48\x27\x2b\xd6\x1c\xbe\x12\xe1\xfa\x4a\xb3\x12\x45\x5e\x6e\x26\x32\x9d\xd3\xc3\x2e\x77\x57\xfb\x3f\xca\xe6\xf8\xf3\xa9\xc4\xe1\x3e\x99\xb9\xd5\x0a\x4e\x1a\x56\x49\xbc\xb1\xbd\xae\x61\xec\xa8\x05\x6a\x6a\x3e\x98\x19\x06\xb8\x98\x3b\x3e\x6f\x23\x04\x1e\x06\x44\xc2\x53\x43\x9e\x8c\x22\x04\x0b\x89\xce\x91\xe2\x86\xd5\xc9\x2f\x1b\x84\x1d\x5d\xfa\x1d\x5d\xbc\x61\xf9\xd3\xc7\x79\x9d\x22\x2f\x93\xa3\x64\xae\xd3\xe8\x10\x7d\x07\x00\x00\xff\xff\xca\x1e\x3d\x11\x3b\x02\x00\x00"

func scriptsBorrow_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsBorrow_nftCdc,
		"scripts/borrow_nft.cdc",
	)
}

func scriptsBorrow_nftCdc() (*asset, error) {
	bytes, err := scriptsBorrow_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/borrow_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9, 0x1f, 0x4c, 0xe2, 0xa, 0xde, 0xe4, 0x18, 0xb8, 0x21, 0xaa, 0x29, 0x2a, 0x8f, 0x33, 0x67, 0xde, 0x8a, 0x13, 0x2b, 0x97, 0x99, 0xf7, 0x10, 0x95, 0x15, 0x33, 0xe8, 0x2d, 0xa, 0xc8, 0x19}}
	return a, nil
}

var _scriptsGet_collection_lengthCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xcd\x6a\xb4\x30\x14\x86\xf7\xb9\x8a\x83\x8b\x0f\xdd\xc4\xfd\xf0\x4d\x87\xc1\x56\x98\x8d\x94\xe2\x0d\xc4\x18\x6d\x68\x3c\x27\xc4\x13\xda\x32\xcc\xbd\x97\x31\x1d\xed\x0f\x75\x21\x09\x79\x9e\x37\xe7\x8d\x9d\x3c\x05\x86\x86\xb0\x8e\x38\xda\xce\x99\x96\x5e\x0c\xc2\x10\x68\x82\x4c\xca\x52\xca\x52\x13\x72\x50\x9a\xe7\xf2\x27\x26\x75\xaf\x33\xf1\x99\xf1\xf0\xa6\x26\xef\x4c\x53\xb7\x7f\xd8\x1b\x90\x3c\xe1\x63\x07\x43\x44\x98\x94\xc5\x5c\xf5\x7d\x30\xf3\xbc\x83\x63\x5a\x14\x3b\x38\x21\xc3\x59\x00\x00\x38\xc3\xa0\xb4\xa6\x88\x0c\x7b\x18\x0d\x1f\xd3\xe6\x66\x15\x62\xc5\x34\x39\x67\x34\x5b\xc2\x27\x33\xc0\xfe\xa6\x2d\xe7\xd7\x4f\x8e\x86\x2b\xe5\x55\x67\x9d\xe5\xf7\xbc\xf4\xb1\x73\x56\x97\x4d\xdd\x56\xab\x59\x6c\x74\x47\x21\xd0\xeb\xff\x7f\xe7\x5f\xdd\x37\xfc\x71\x89\xb8\xdc\xe5\x9b\x77\x38\x80\x57\x68\x75\x9e\x55\x14\x5d\x0f\x48\x0c\x29\x0a\xf4\x7a\x79\x7a\xa7\x34\xc0\x97\xb9\xb3\x14\xb3\xfc\x82\xe1\x18\xf0\x7b\xa9\x6b\x83\xd3\xfd\x9c\x17\xd2\x19\x1c\xf9\x59\x5c\xc4\x47\x00\x00\x00\xff\xff\x8a\x33\x85\x7b\xc7\x01\x00\x00"

func scriptsGet_collection_lengthCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_collection_lengthCdc,
		"scripts/get_collection_length.cdc",
	)
}

func scriptsGet_collection_lengthCdc() (*asset, error) {
	bytes, err := scriptsGet_collection_lengthCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_collection_length.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0x6b, 0x76, 0x1e, 0x72, 0x52, 0x12, 0xa4, 0x3d, 0x9e, 0xdc, 0x40, 0xd3, 0x47, 0x4e, 0xc3, 0xc9, 0x8b, 0x80, 0x35, 0x3c, 0xc9, 0x9, 0xc1, 0xfd, 0xa8, 0xee, 0xe, 0xfa, 0x56, 0x97, 0xc4}}
	return a, nil
}

var _scriptsGet_nft_metadataCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x5d\x6f\xdb\x3a\x0c\x7d\xf7\xaf\x60\xf3\x70\x61\x03\x17\x4e\xef\xdd\x27\x82\xa6\x45\xd1\x2d\x43\x1f\x16\x0c\x9b\xb7\xe7\xc8\x32\x95\x72\xb0\x25\x43\xa2\x93\x06\x45\xfe\xfb\x60\xf9\x3b\x4d\x5e\x96\x97\x58\x22\x0f\x79\x78\x70\x28\x2a\x4a\x63\x19\x3e\x3f\x8b\xa2\xcc\x71\xbd\x4a\x40\x59\x53\xc0\x2c\x8e\xe7\x71\x3c\x97\x46\xb3\x15\x92\xdd\x7c\x48\x88\x65\x26\x67\x41\x8b\xfb\x8a\x2c\x32\xc1\xe2\x17\xe1\xde\x5d\x80\x4e\x72\x1a\x74\x50\x56\x29\x38\xb6\x95\x64\x58\xaf\x92\xef\xe8\xaa\x9c\xe1\x25\x00\x00\x28\xab\x34\x74\xc8\x11\xec\x84\x05\x2d\x0a\x5c\xc0\x0f\xb6\xa4\xb7\xaf\xa3\x19\x3a\x69\xa9\x64\x32\xfa\x72\x12\x3f\x55\x45\xaa\x05\xe5\x97\x53\xcc\x5e\xa3\x5d\xc0\x7d\x96\x59\x74\xee\x4c\x89\x43\x39\xb0\xf0\x61\xd2\xc4\x61\xd4\x32\xae\x7f\x0e\x73\x15\xd7\x6c\x61\x09\xb3\xd9\xf4\x7a\x44\xf3\x4c\xb4\xe7\x77\x26\xe6\x89\xc1\x12\xae\x9f\xaf\x4f\x40\x87\x72\xd4\xe9\x18\x1c\x1b\x49\x55\xa5\xa1\x10\xa4\x43\xd1\x8c\xd2\xcf\xf4\x2f\x50\xb6\x80\x9f\x8f\x9a\xdf\xbf\x8d\x16\xaf\x34\xcf\x91\x41\x48\x69\x2a\xcd\xb0\x84\x2d\xf2\x7d\x73\xe8\xea\x44\x41\x9f\x26\x4d\x9e\xa3\x6c\x67\x69\x31\x3d\xb7\x78\x8b\xfc\x20\x4a\x91\x52\x4e\x7c\x08\xe7\x65\x95\xe6\x24\xe7\xeb\x55\xf2\xd0\xc3\xa2\x21\x3b\x35\xd6\x9a\xfd\xcd\x3f\x2f\x23\x77\x0d\x9f\x03\xe4\x9b\x2f\x73\xbc\x0d\x07\xec\xdd\x1d\x94\x42\x93\x0c\x67\x0f\xa6\xca\x33\xd0\x86\xa1\x29\x07\x02\x2c\x2a\xb4\xa8\x25\x02\x1b\xe0\x27\x1c\x91\x9e\x8d\x46\xd1\xaa\x9e\x76\x88\xb5\x7c\x06\x06\x61\x2d\x1a\x65\xd1\x55\x83\xf1\x96\x13\x2c\x60\x39\x08\x18\xb6\xf5\xe6\x73\xf8\x82\xec\x9b\xa5\xc2\x91\x84\x8c\x5c\x99\x8b\x03\x90\x56\xc6\x16\xc2\x0b\xa6\x4c\x6d\x47\x72\x35\xbc\xb1\x91\xf2\x44\x76\x84\x7b\x58\xd6\x7c\x62\x8b\xce\xe4\x3b\xac\x77\x25\x4c\x0e\x25\xde\x4c\xb7\xe7\x53\x53\xf5\x36\x8c\xc6\xe6\xab\x6b\x74\xfd\x96\x4d\x35\xe1\xae\xe0\x2c\x34\xe8\x51\x75\xa8\xb3\x6c\x8b\xf6\xc7\x69\xc2\xd4\xbc\x5d\xde\xe8\xb6\x75\xe0\xa9\x0a\x54\x88\x2d\x0e\xcb\x37\x99\x1d\x42\x52\x40\x0c\xf8\x4c\x8e\x5d\xf4\x77\x52\x24\x5d\xe9\x33\x62\x8c\x77\xea\x82\x1c\x3d\xfc\x44\x90\x31\xb4\xff\x8e\x2b\x4b\x27\x73\x26\x4f\xd8\xbc\x1b\x40\x0e\x1c\x1b\x8b\x19\x64\x64\x51\x72\x7e\x00\xa3\xbd\x06\xf5\xa8\x26\xfd\x8d\x92\x7b\xcf\x4d\x9f\x9a\x76\x52\x7f\x79\x15\xb7\xcb\xd6\xba\xcd\xb3\xe9\x1e\x00\xff\xdf\xf7\x7e\xd4\xae\x44\xd9\xe8\xec\x5f\x02\xa3\x06\x6d\xd9\xc0\x0e\x2d\xa9\x03\x10\x3b\x30\x96\xb6\xa4\xc7\x96\x4f\x9a\xa7\xa3\xee\xbb\x45\x7f\xea\x3c\xdc\xcc\xdf\x87\xeb\x50\x4c\x19\x6a\x26\x45\x68\xbb\xe6\x9b\x3e\x6d\x53\x8f\xbe\xb9\x8f\xd5\x1b\x25\xb3\xff\xe5\x7f\xe2\xc3\x47\xf5\x0e\x11\x47\x3b\x1c\xaf\x57\xc9\xa6\xa9\x6e\x91\x2b\xab\x7d\x93\xe0\x18\xfc\x09\x00\x00\xff\xff\x2d\x4c\x0b\x8e\x78\x06\x00\x00"

func scriptsGet_nft_metadataCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_nft_metadataCdc,
		"scripts/get_nft_metadata.cdc",
	)
}

func scriptsGet_nft_metadataCdc() (*asset, error) {
	bytes, err := scriptsGet_nft_metadataCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_nft_metadata.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xa5, 0xde, 0x8d, 0xa6, 0x12, 0xe5, 0x8b, 0x6b, 0xc6, 0x5d, 0x5d, 0xff, 0xcf, 0x46, 0x1, 0xe1, 0xc1, 0xbd, 0x18, 0x1d, 0xd5, 0xa8, 0x80, 0x4, 0x20, 0x5, 0xe1, 0xb0, 0x3a, 0xf0, 0x68}}
	return a, nil
}

var _scriptsGet_total_supplyCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\x31\x0a\x42\x31\x0c\x06\xe0\x3d\xa7\xf8\x79\x93\x2e\xe9\x22\x0e\xee\x0a\x2e\x2e\xea\x01\x6a\x7d\x0f\x0a\x6d\x1a\x62\x02\x8a\x78\x77\x47\xdd\x3f\xbe\xda\x75\x98\x63\xff\xcc\x5d\xdb\x7c\x3a\x5c\xb0\xd8\xe8\x98\x98\x13\x73\x2a\x43\xdc\x72\xf1\x47\xfa\x01\x2e\xf7\x32\x11\x69\xdc\xb0\x84\xa0\xe7\x2a\xab\xf5\x0e\xd7\xa3\xf8\x76\x83\x37\x01\x80\xcd\x1e\x26\x7f\x2b\xfb\xf0\xdc\xce\xa1\xda\x5e\xf4\xa1\x6f\x00\x00\x00\xff\xff\xab\xdd\xb2\x0f\x76\x00\x00\x00"

func scriptsGet_total_supplyCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_total_supplyCdc,
		"scripts/get_total_supply.cdc",
	)
}

func scriptsGet_total_supplyCdc() (*asset, error) {
	bytes, err := scriptsGet_total_supplyCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_total_supply.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x81, 0x9, 0x60, 0xa2, 0xa5, 0x58, 0x7b, 0xb8, 0xa2, 0x87, 0x3a, 0x50, 0x8b, 0x97, 0x82, 0xd3, 0xf7, 0x78, 0xfa, 0x17, 0x8a, 0xda, 0xc8, 0x54, 0x76, 0x3b, 0xe3, 0x9c, 0x92, 0x0, 0x29, 0x87}}
	return a, nil
}

var _setup_accountCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\xcd\x6a\xe3\x30\x10\xbe\xfb\x29\xbe\xed\xa1\x38\x90\x46\xf7\x92\x2d\x94\xd2\x1c\xc3\xb2\xf8\x05\x26\xca\x38\x16\x55\x24\x33\x1a\x37\x1b\x42\xde\x7d\xb1\x9d\xf5\x4f\x37\xd5\xc9\x48\xdf\xef\x8c\xdd\xb1\x8e\xa2\xd8\xc6\xb0\x69\xc2\xc1\xed\x3c\x17\xf1\x83\x03\x4a\x89\x47\x3c\xac\x56\xc6\xc6\xa0\x42\x56\x93\xf9\x8a\x59\xd9\xbd\x7d\xc8\x6e\x02\xef\x7f\xe8\x58\x7b\xde\x6e\x8a\x7b\xd4\xf1\xb5\x27\x65\xc6\xa0\xa8\x5c\x82\x0a\x85\x44\x56\x5d\x0c\x70\x09\xa7\x8a\x14\x14\x40\xd6\xc6\x26\x28\x4e\xb1\xf1\x7b\x48\x13\x5a\x82\x46\x24\x56\x38\x4d\xec\x4b\x34\x75\x7b\x21\x6c\xd9\x7d\x32\xb6\x9b\x22\x65\xd9\x54\xed\x92\x65\x00\x50\x0b\xd7\x24\x9c\x27\x77\x08\x2c\xcf\x78\x6d\xb4\x7a\xed\xd5\x17\xb8\x74\x90\xf6\x18\x83\xdf\xac\x8d\x04\x30\x89\x3f\xc3\x95\xd0\x8a\x87\x1c\xe4\x85\x69\x7f\x46\x45\x09\x04\x1b\xbd\xe7\xce\x65\xe0\xbb\x12\xbd\xc3\x6a\x17\x45\xe2\x69\xfd\x38\xa9\xfc\x36\xe0\x5f\xf2\x76\x38\xcf\x30\x49\xa3\xd0\x81\xcd\x76\x53\x8c\xaf\x0b\xfc\xf8\x89\xe0\xfc\x24\x57\x7b\xa4\x0b\x36\x5c\x5d\xb3\x69\xea\x37\x61\x52\x06\x21\xf0\x09\x7c\xac\xf5\x7c\x2f\x9e\x67\x9d\x5c\x63\xfd\x84\xe9\x46\x3a\x89\xf7\x96\x3b\x66\xc9\x17\x33\x9b\x44\x9f\x0c\xa7\xed\xc8\x27\x73\x19\x10\xb7\xee\x2d\x2a\x5f\x3f\x8d\x4e\x4b\x68\xfc\xb6\xed\xcc\xc0\xfe\xeb\x51\x37\x3b\xef\x2c\x2c\xd5\xb4\x73\xde\xe9\x19\x65\x94\xce\xf4\x4e\xaf\x9b\xaf\x77\xe1\x63\xfd\x78\xf9\xef\x07\x1d\xcd\x7e\x75\xaa\xcb\x69\xed\xf1\xf3\x2b\xec\xfa\x92\xcf\x16\x60\xfa\x4c\xf3\xf8\xcb\x19\x44\x49\x0e\xac\xdf\x55\x1d\xa0\x8b\xac\xdf\xe0\x35\xfb\x1b\x00\x00\xff\xff\x4c\xe9\xc4\x26\x78\x03\x00\x00"

func setup_accountCdcBytes() ([]byte, error) {
	return bindataRead(
		_setup_accountCdc,
		"setup_account.cdc",
	)
}

func setup_accountCdc() (*asset, error) {
	bytes, err := setup_accountCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setup_account.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6f, 0x47, 0xbd, 0xfc, 0x6c, 0xcb, 0xae, 0x32, 0xc8, 0x9e, 0x3d, 0xef, 0xb5, 0xfe, 0x84, 0xa3, 0x47, 0xb5, 0xa5, 0x93, 0xbc, 0xb2, 0x49, 0x23, 0x93, 0x34, 0x6e, 0x0, 0x8b, 0x62, 0x52, 0x53}}
	return a, nil
}

var _transfer_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x88\x1c\x36\x07\x58\xed\xcb\xb0\x43\x90\xb6\x08\xd2\x05\xe8\x25\x18\x86\xec\x03\x64\x99\xb6\xb5\x39\xa4\x41\xd1\xcb\x86\x22\xff\x3e\x38\x72\xec\x78\x4d\x73\xa9\x4e\x82\xc4\xf7\x1e\xdf\x13\xe5\xf6\x0d\x8b\xc2\x96\x69\xd3\x52\xe9\xb2\x1a\x77\xfc\x0b\x09\x0a\xe1\x3d\xcc\x92\x24\xb5\x4c\x2a\xc6\xaa\x4f\xff\xaf\x49\x6c\x6e\x67\x51\x4f\xf0\xf5\x8f\xd9\x37\x35\x6e\x37\xbb\x6b\xd0\xf1\x36\x80\xa2\x34\x85\x5d\xe5\x3c\xa8\x18\xf2\xc6\xaa\x63\x02\xe7\xa1\x60\x09\x47\x05\x8a\x38\x2a\xc1\x50\x0e\x67\xce\x0e\xc4\x84\x60\xac\xe5\x96\x14\x94\xc1\x10\x6b\x85\x12\x45\x17\x3c\xb1\xa0\x75\x8d\x43\xd2\x05\xac\xf2\x5c\xd0\xfb\x4f\x70\x70\x5a\xe5\x62\x0e\xcf\x4f\x0b\xf8\xf1\x4c\xfa\xe5\xf3\x1c\x5e\xa2\x08\x00\xa0\x11\x6c\x8c\x60\xec\x5d\x49\x28\x0b\x58\xb5\x5a\xad\x82\x44\x57\x03\xfd\x4a\x53\x28\x51\x41\x2b\x84\x41\xc0\x43\xd3\x66\xb5\xb3\x43\x4b\x9c\xfd\x44\xab\x03\xa6\x46\x1d\x8b\xe1\xbe\x23\xe8\x99\xc7\x26\xe7\xd1\xa5\x44\xc6\x22\x7c\x00\x03\x82\x05\x0a\x92\xc5\xce\x66\x27\x1a\xda\xfb\xe8\x4f\x71\x58\xae\x6b\x3c\xb9\x9d\x68\x8d\xc7\xdf\xb1\x80\xfb\x1e\x93\x04\xd2\xe5\x87\x8b\x57\x58\x0f\x95\x0f\x71\x97\xed\x02\x52\xaf\x2c\xa6\xc4\x74\xbb\xd9\x8d\xb7\xf3\x81\xbf\x5b\x8f\x8f\xd0\x18\x72\x36\x9e\xad\xb9\xad\x73\x20\xd6\x1b\x1d\xf3\x21\x34\x3c\x76\x35\x7b\xc3\x6c\x9f\xe2\x2b\x06\x41\x8b\xee\x37\x8a\x7f\xcb\x70\x8e\x0d\x7b\xa7\xc1\xed\x10\x69\x52\xa2\xae\x4d\x63\x32\x57\x3b\xfd\x1b\xa7\x81\xfe\x96\xb1\x21\xa2\x97\x57\x43\x3e\x42\xbe\x9d\x68\x8e\x0f\xf1\x7b\x42\x39\x5b\xba\x95\xcb\x79\x58\x4f\x80\xe1\x47\x5d\x8f\x74\x12\x07\x15\x0a\xcb\xbb\xe9\x18\x24\x67\xb6\xf8\xf2\x0f\x8c\xfb\xa9\xf4\x53\x08\x74\x50\x76\x34\x9d\xf8\xeb\xda\xe3\x33\x24\xfd\x36\xd6\x2e\xbc\x05\x2c\xef\xa8\xd0\x90\xd7\x31\x3a\xfe\x0b\x00\x00\xff\xff\x0b\x16\x45\xaa\x6d\x04\x00\x00"

func transfer_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_transfer_nftCdc,
		"transfer_nft.cdc",
	)
}

func transfer_nftCdc() (*asset, error) {
	bytes, err := transfer_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transfer_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x97, 0x4, 0x1, 0x4c, 0xec, 0xaf, 0xb9, 0xd3, 0x4, 0x9, 0xa, 0x93, 0xae, 0x2f, 0x53, 0xf5, 0x95, 0xcf, 0x9a, 0x68, 0xfe, 0xa8, 0x9d, 0x85, 0x72, 0x3e, 0xad, 0xa4, 0xab, 0x80, 0x96, 0x4a}}
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
	"destroy_nft.cdc":                   destroy_nftCdc,
	"mint_nft.cdc":                      mint_nftCdc,
	"scripts/borrow_nft.cdc":            scriptsBorrow_nftCdc,
	"scripts/get_collection_length.cdc": scriptsGet_collection_lengthCdc,
	"scripts/get_nft_metadata.cdc":      scriptsGet_nft_metadataCdc,
	"scripts/get_total_supply.cdc":      scriptsGet_total_supplyCdc,
	"setup_account.cdc":                 setup_accountCdc,
	"transfer_nft.cdc":                  transfer_nftCdc,
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
	"destroy_nft.cdc": {destroy_nftCdc, map[string]*bintree{}},
	"mint_nft.cdc": {mint_nftCdc, map[string]*bintree{}},
	"scripts": {nil, map[string]*bintree{
		"borrow_nft.cdc": {scriptsBorrow_nftCdc, map[string]*bintree{}},
		"get_collection_length.cdc": {scriptsGet_collection_lengthCdc, map[string]*bintree{}},
		"get_nft_metadata.cdc": {scriptsGet_nft_metadataCdc, map[string]*bintree{}},
		"get_total_supply.cdc": {scriptsGet_total_supplyCdc, map[string]*bintree{}},
	}},
	"setup_account.cdc": {setup_accountCdc, map[string]*bintree{}},
	"transfer_nft.cdc": {transfer_nftCdc, map[string]*bintree{}},
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

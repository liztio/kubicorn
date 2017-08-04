// Copyright © 2017 The Kubicorn Authors // // Licensed under the Apache License, Version 2.0 (the "License"); // you may not use this file except in compliance with the License. // You may obtain a copy of the License at // // http://www.apache.org/licenses/LICENSE-2.0 // // Unless required by applicable law or agreed to in writing, software // distributed under the License is distributed on an "AS IS" BASIS, // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. // See the License for the specific language governing permissions and // limitations under the License.

package bootstrap

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _bootstrap_readme_md = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\x41\x8f\xd4\x30\x0c\x85\xef\xf9\x15\x4f\x9a\x03\x97\x55\xf9\x09\x48\x1c\x90\xf6\x00\x27\xb4\x5c\x9b\x69\xdc\xad\x77\xd2\xb8\x4a\x9c\xad\x46\x88\xff\x8e\xdd\x59\x46\xc0\x9e\x32\x53\xdb\xef\x7d\xcf\x3e\xe1\xb3\x88\x36\xad\x71\x0b\xe1\xfb\x42\x8d\x10\x2b\x41\x17\xc2\xf9\x4f\x01\x6d\xaa\xbc\x69\xb3\xaf\x51\xd1\x16\xde\xb0\xb3\x2e\x47\x53\xa2\x39\xf6\xac\x18\x2f\xfd\xcc\x93\xd4\x32\x62\xab\x32\x73\xa6\x36\x84\xf0\x85\x28\x63\xae\x64\x82\x82\x98\x12\xae\xd2\x2b\x64\x2f\x0f\x90\x8a\x55\x12\xcf\x57\x97\x71\x57\x45\x2c\xf6\x87\x57\x1a\x0e\x92\xbb\xab\xf3\xd0\x3c\xd3\xa4\xfc\x4a\xf9\x8a\xdd\x29\x76\x42\xf7\xa9\x86\xd1\xde\x8a\x14\x35\x8e\xee\xc2\x85\x4d\x09\x4f\x5f\x43\x38\x9d\x4e\x78\x44\x21\x4a\x5e\x50\x5a\xb7\x1c\x95\x20\x5d\x21\xc5\xde\xf9\xcd\xfa\x5d\xd0\x10\xbe\x89\xa3\x22\x49\xf9\xa0\x03\x7e\x54\xb6\xb9\x73\x6c\x0b\x32\x5f\xcc\xd5\x33\x0e\xff\x18\xac\x62\x94\x0e\x61\x00\x56\xff\x5f\xf2\x06\xdd\x16\xe9\x39\xd9\x48\x92\x4f\x21\x3c\xce\x87\x47\xa5\x98\x2d\xd5\x14\x8b\x51\xd9\x8f\x67\x52\x58\x84\xb9\xca\xfa\xf7\x56\x63\x49\x28\xb2\x2f\xe4\xdb\xc8\x8d\x1e\x8e\x61\x9f\xf2\x3d\xf8\x29\xc6\xa7\x98\x3b\xb5\x9f\xbf\x46\x98\x75\x9f\xd4\x53\xdb\x7d\xd8\xb2\x4e\xbd\xa9\xe9\x5d\xe8\xfa\xf1\xd5\xbb\xb0\x45\xae\x6f\x07\xdd\x39\x67\x9c\xc9\xc0\x5f\x6c\xc7\x96\x85\x8b\xca\xed\x52\x37\xf8\xc1\xce\xc1\xed\xde\x17\x31\x49\x32\xc9\x25\x96\x67\xc3\x70\x30\xab\xda\x10\x95\x74\x5b\xb5\x35\xbd\x98\xe1\x21\x3f\xfc\x0e\x00\x00\xff\xff\x8c\x9a\x9e\x97\x62\x02\x00\x00")

func bootstrap_readme_md() ([]byte, error) {
	return bindata_read(
		_bootstrap_readme_md,
		"bootstrap/README.md",
	)
}

var _bootstrap_amazon_k8s_ubuntu_16_04_master_sh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x54\x5f\x6f\x23\x35\x10\x7f\xf7\xa7\x18\xd2\x93\x8e\x4a\xe7\x75\x5b\x10\x48\x95\x8a\x74\x77\xf4\x21\x1c\xba\xab\xaa\x00\x2f\xbc\x38\xf6\x64\xd7\xac\xd7\xb6\xfc\x27\x6d\x04\xe5\xb3\x33\xf6\x6e\xaf\xa9\x38\x21\xde\x9a\x48\x89\x67\x3c\xfe\xcd\xbf\xdf\xcc\xc9\x57\xa2\xa4\x28\xb6\xc6\x09\x74\x7b\xd8\xca\x34\xb0\x84\x19\x38\x32\xa5\xe1\x6f\xc6\x4e\x80\xbf\xd0\x87\x5c\x6f\x06\x4c\x08\x7b\x69\x0b\x26\x90\x11\xc1\xb8\x3f\x50\x65\xd4\x74\xc8\x1e\xf2\x80\x90\x54\x34\x21\x77\xf0\x1b\x36\x03\xbc\x0f\xd6\x28\x93\xed\x01\x9c\xcf\x50\x92\x71\x3d\x48\xc8\x38\x05\x2b\x73\x15\xac\x74\x7d\x91\x3d\x02\x01\xcc\x70\x0d\x67\x76\x42\x3e\x65\xaa\x37\xe8\x94\x2f\xb1\x99\xd1\x65\x49\x18\xab\xd6\x9a\xc9\x34\x73\x13\xab\x0e\xfc\xee\x19\xb2\xef\x8d\x22\xcc\x6a\x40\x97\x3b\x63\x31\x75\xf0\xee\x00\x1a\x93\xe9\x1d\x48\x6b\x9f\x12\x58\x92\x4a\x83\x2f\x56\x93\xdb\x2d\xc5\xbf\xb5\x2d\x2a\x3a\xd6\x16\xc8\x0c\xb1\xb8\x6c\x26\x7c\x03\xd2\xe9\x39\xdb\x01\x09\x64\xce\x19\x22\x4a\x0b\x77\x3e\x8e\x1d\xac\x77\x70\xf0\x05\x1c\x12\xb2\xf2\x4e\x9b\x6c\xbc\xa3\xdb\x16\xd2\x1b\xb8\x8b\x26\x53\xf1\x72\x0d\xae\xb5\xf8\x04\x7c\x84\x49\x8e\xe4\x94\xca\x34\x50\x76\xc7\xc8\x1d\x3b\xa1\xef\xe6\xd3\x87\xeb\x8f\x57\xab\xf5\xc7\x9f\xae\xdf\x6f\xae\x7f\x6c\xe2\x8a\xdd\x7c\xba\xdd\x3c\x29\xab\xb4\x7a\x49\x8e\x30\x55\xa2\x05\x9e\x60\xc8\x39\xa4\x4b\x21\x82\x54\x23\xb5\x2d\x75\xca\xfa\xa2\xbb\xde\xfb\xde\x62\xa7\xfc\x24\x64\xc8\x42\x7b\x55\xff\xf9\x88\x87\xae\x0f\x3d\xfc\x05\xa9\x68\x0f\x8b\x0a\xa4\xd6\xc0\x59\xf6\x45\x0d\x20\x30\x37\x5b\x91\x88\x08\x8a\x00\xad\x49\xb9\xd3\x62\x2c\x5b\x8c\x0e\xf3\xa2\x61\x69\x00\xae\xe0\x35\xaa\xc1\xc3\x4a\xe3\xb6\x45\x42\x81\xd0\xd3\xee\xc8\xd6\x78\x01\x4f\x22\xbf\x47\x67\xa8\x3f\x93\x34\x6e\x05\x3f\xfc\x5f\x67\xaf\x19\xab\xa1\xf6\xc4\x8e\x12\xb4\xa4\x9e\xf2\xc3\x67\x8d\x71\x29\x57\x8a\xf1\x03\xfc\xce\x80\x3e\xc9\x2b\xe2\xd0\x7c\xc6\x6d\xae\xf4\x4a\x8b\x48\x75\x18\x31\x52\x50\x8b\x5c\x31\x72\x94\x2e\x05\x1f\x33\x6f\xb5\x5c\x6e\x6a\x04\x16\xf3\x91\x24\xf5\x74\x75\xde\x7d\xdf\x9d\xf1\xb3\xb3\x45\xdd\x6a\xcd\x4b\x36\x36\x31\xc6\xd2\x21\xd1\x54\xa8\x6c\x69\x8a\x1a\xa7\x67\x6f\x47\x7a\x0a\x34\xe6\x47\x35\xbb\xf9\xe5\xdd\xcf\xeb\xf7\xeb\x9b\xab\x57\x5f\xa3\xba\x98\x30\x4b\x4a\x4d\x12\xa7\x42\xd9\xd2\x30\x73\x13\xf6\xdf\x52\xa7\x54\xa1\xb5\xa4\x61\x45\x5f\xbe\x83\x8b\x53\x76\x73\xbb\xfe\xf5\xed\xe6\xba\x3d\x34\x3b\xa2\xfd\xce\xd4\x8e\xf6\x11\x03\xf0\xb7\x70\x0e\x98\x87\xb3\x47\x85\xa1\x2a\x1e\x81\x5c\xce\x20\xff\x82\x3d\x27\xcd\xbd\x8c\x7d\x3a\x65\x6c\x49\x96\xe6\x8c\xc6\xf1\xb3\x64\x1c\xcd\x11\xe7\x32\x18\x5a\x0a\x7b\x8c\x9c\x76\xa7\xe6\xb5\x6e\xf0\xea\xcf\x3a\x0f\x0f\x74\x9b\xfd\x88\x8e\xe4\x36\x34\x0f\xf0\xcc\x5e\x6a\xfa\xcd\x24\xd1\x49\x13\x76\xaa\xef\x96\x12\x3c\x3c\xb3\x54\x64\xc7\xf1\x9e\x1a\xc3\x13\xf5\xe6\x99\x1d\x9d\x1f\xd3\x7f\x60\x6d\x55\x4a\x37\x26\xf8\x80\x36\x11\x91\x2f\x4f\x5b\xbc\xb5\xd6\x32\x04\x3b\x33\x82\xd2\x5b\xb8\x49\x95\x4f\x5d\x88\xbe\xee\x22\x25\xa9\xc6\xbe\xf3\xb1\x17\xfb\x8b\xee\x1b\x41\x54\xaa\xdb\x8c\xb7\x1e\xe1\x31\x05\xc5\xc2\x30\x59\x97\x8b\x18\x7c\x7a\xbc\xa6\xb2\x88\xf3\xee\x3b\xb1\x40\x1d\xe4\x64\x67\x87\xbc\x05\x31\x37\xa6\x51\xfc\x08\x8c\x1e\x19\xd7\xd5\x4b\xc6\xa6\x51\xd3\x46\xe5\x01\x08\x75\x42\x51\xb6\xb4\xf7\x8a\x68\xe3\xc3\x54\xf8\x8f\xa7\x5f\x78\x20\x66\x7f\x8c\xe6\xf1\xce\x01\xbf\x85\xf9\xf2\x72\xfe\xfb\xc2\x83\x7f\x02\x00\x00\xff\xff\x97\x66\x54\x00\x03\x07\x00\x00")

func bootstrap_amazon_k8s_ubuntu_16_04_master_sh() ([]byte, error) {
	return bindata_read(
		_bootstrap_amazon_k8s_ubuntu_16_04_master_sh,
		"bootstrap/amazon_k8s_ubuntu_16.04_master.sh",
	)
}

var _bootstrap_amazon_k8s_ubuntu_16_04_node_sh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x53\x4d\x6f\xdc\x38\x0c\xbd\xeb\x57\x70\x67\x16\xc8\x25\xb2\xb3\xa7\x05\x16\xc8\x02\xfd\x98\x43\x5a\x34\x05\xda\x01\x7a\xe9\x45\x96\x18\x5b\xb1\x2c\x19\x12\x95\xc4\x68\xd3\xdf\x5e\x4a\x76\x93\xe9\xad\xb7\x78\x0e\x23\x3e\xd2\x8f\x8f\xe2\xf3\xfe\xaf\x36\xa7\xd8\x76\xd6\xb7\xe8\xef\xa0\x53\x69\x10\x09\x09\x24\x0a\x6d\xe0\x87\x10\x7b\x90\x2f\xf4\x70\xeb\xe3\x80\x09\xe1\x4e\xb9\x8c\x09\x54\x44\xb0\xfe\x16\x35\xa1\xe1\x03\x05\xa0\x01\x21\xe9\x68\x67\x6a\xe0\x0b\xd6\x02\x7c\x98\x9d\xd5\x96\xdc\x02\x3e\x10\xe4\x64\x7d\x0f\x0a\x08\xa7\xd9\x29\x2a\x81\x53\xbe\xcf\xaa\x47\x60\x82\x95\xae\xf2\xac\x4d\xb8\xa7\x4a\x25\x83\x5e\x87\x1c\x6b\x19\x27\x73\xc2\x58\x50\x67\x27\x5b\xcb\x6d\x2c\x18\x84\x9b\xdf\x98\x43\x6f\x35\x73\x96\x02\x4e\xde\x58\x87\xa9\x81\xd7\x0b\x18\x4c\xb6\xf7\xa0\x9c\x7b\x1e\x60\x1b\x2a\x0d\x21\x3b\xc3\x6d\x3b\xd6\xdf\xb9\xaa\x8a\x8f\x65\x05\x8a\x20\x66\x4f\x76\xc2\x73\x50\xde\xac\xd3\x0e\xc8\x24\xeb\xcc\x10\x51\x39\xb8\x0f\x71\x6c\xe0\xea\x06\x96\x90\xc1\x23\x33\xeb\xe0\x8d\x25\x1b\x3c\x67\xab\xa4\x73\xb8\x8f\x96\xf8\xf2\xa8\x88\xab\x2b\xde\x43\x88\x30\xa9\x91\x9b\xf2\x35\x0d\x3c\xdd\x29\x73\x23\xf6\xfc\x3b\x7e\x7c\x7f\xb8\xbe\xdc\x5d\x5d\xbf\x3b\xbc\x39\x1e\xde\xd6\x70\x27\x3e\xbc\xfa\x7c\x3c\x7c\x7a\x86\xd7\x78\xf7\x92\x3e\x11\x3a\x47\x07\x32\xc1\x40\x34\xa7\xff\xda\x76\x56\x7a\xe4\xd5\xa5\x46\xbb\x90\x4d\xd3\x87\xd0\x3b\x6c\x74\x98\x5a\x35\x53\x6b\x82\x2e\xff\x72\xc4\xa5\xe9\xe7\x1e\xbe\x43\xca\x26\xc0\x06\x81\x32\x06\xa4\xa0\x90\xf5\x00\x2d\x52\xad\x6d\x13\x9b\x41\x33\xa1\xb3\x89\x1a\xd3\x8e\xb9\xc3\xe8\x91\x36\x44\xa4\x01\xa4\x86\x33\xd4\x43\x80\x9d\xc1\xae\x2a\x61\x21\xfc\x6a\x73\x52\x6b\x43\x0b\xcf\xa1\x7c\x40\x6f\x79\x47\x93\xb2\x7e\x07\xff\xff\x69\xb3\x33\x21\x8a\xd4\x9e\x1d\x92\x67\xa3\x78\xaf\x72\x79\x42\xac\x4f\x54\x6c\x26\x17\xf8\x2a\x80\x9f\x14\x34\xfb\x68\x3d\x63\x47\xc5\x62\x69\x0b\xf9\x1e\x46\x8c\x2c\x6a\x8b\x0b\x07\x45\xe5\xd3\x1c\x22\xc9\x7a\x97\x5b\xa6\x28\x70\x48\x27\x91\x32\xd3\xe5\x3f\xcd\xbf\xcd\x85\xbc\xb8\x10\x22\x2d\x89\x3f\x03\x4d\x8e\x3f\x9b\x6a\xe2\x95\xfa\x04\x67\x55\x91\x7e\xc1\x62\x63\x60\x03\xb3\xcf\x9f\xa2\xdb\xc0\xee\x94\x92\xc2\x88\x1e\xfe\xfe\x56\xfd\xf6\xc8\x87\xd5\x61\x8f\x3f\x03\x00\x00\xff\xff\x23\x59\x14\x94\xaa\x04\x00\x00")

func bootstrap_amazon_k8s_ubuntu_16_04_node_sh() ([]byte, error) {
	return bindata_read(
		_bootstrap_amazon_k8s_ubuntu_16_04_node_sh,
		"bootstrap/amazon_k8s_ubuntu_16.04_node.sh",
	)
}

var _bootstrap_digitalocean_k8s_ubuntu_16_04_master_sh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x55\x6d\x6b\x1c\x37\x10\xfe\xae\x5f\x31\x3d\x87\xa4\x81\x48\xb2\x53\xda\x42\xc0\x85\x38\x39\xe8\x35\x8d\x63\x5c\x37\xa5\x50\x28\x3a\x69\x6e\x57\x39\xad\xb4\x48\x5a\xdb\x47\xeb\xfc\xf6\x8e\xb4\x7b\xbe\x3d\x0a\xa5\xdf\x72\x07\xb6\xe6\x45\x33\xf3\xcc\x3c\x9a\x3b\xf9\x4a\x0e\x29\xca\xb5\xf5\x12\xfd\x2d\xac\x55\x6a\x59\xc2\x0c\x1c\x99\x36\xf0\x99\xb1\x13\xe0\x5f\xe8\x43\xa9\x6f\x5a\x4c\x08\xb7\xca\x0d\x98\x40\x45\x04\xeb\x3f\xa1\xce\x68\xe8\x90\x03\xe4\x16\x21\xe9\x68\xfb\x2c\xe0\x37\xac\x0e\x78\xdf\x3b\xab\x6d\x76\x3b\xf0\x21\xc3\x90\xac\x6f\x40\x41\xc6\xae\x77\x2a\x17\xc1\x29\xdf\x0c\xaa\x41\xa0\x00\x63\xb8\x1a\x67\x4c\x42\x39\x55\x2a\x16\xf4\x3a\x0c\xb1\xba\x91\x71\x48\x18\x8b\xd6\xd9\xce\x56\x77\x1b\x8b\x0e\xc2\xe6\x28\x72\x68\xac\xa6\x98\xc5\x81\x8c\x1b\xeb\x30\x09\xb8\xd8\x81\xc1\x64\x1b\x0f\xca\xb9\x03\x80\x09\x54\x6a\xc3\xe0\x0c\xa5\x5d\x53\xfd\x6b\x57\xab\xa2\x63\x19\x81\xca\x10\x07\x9f\x6d\x87\x2f\x40\x79\x33\xa2\x6d\x91\x82\x8c\x98\x21\xa2\x72\x70\x17\xe2\x56\xc0\x6a\x03\xbb\x30\x80\x47\x8a\xac\x83\x37\x36\xdb\xe0\xc9\x5a\x4b\x7a\x01\x77\xd1\x66\x6a\x5e\x2e\xc5\xd5\x11\x9f\x40\x88\xd0\xa9\x2d\x25\xa5\x36\xb5\x84\x6e\x1e\x59\xb0\x13\xfa\xde\x7c\x78\xb7\xbc\x3c\x5f\xac\x2e\x7f\x5a\xbe\xb9\x59\xbe\xad\xe2\x82\x5d\x7d\xb8\xbe\x39\x28\x8b\xb4\xa8\xde\x5f\x90\x27\x94\xfb\xe3\xd5\x25\xbc\x47\x02\xa6\x87\xe8\xa0\xcd\xb9\x7f\x25\x65\x47\x8a\xb5\x8d\x46\xe8\xd0\x49\xeb\x53\xa6\x01\x88\xd4\xc2\xdf\x04\x96\xed\x8d\xd4\xb3\x3b\x78\xfa\x03\xc8\xdc\xf5\x92\xda\xb5\xc5\x1d\x7b\xbf\xfc\xe5\xc7\x8b\xd5\xf5\xdb\x3f\xdf\x2d\x7f\x3f\x7f\xf2\xb5\xa6\x49\xcc\xcc\x74\x5f\x0f\xf4\x42\x0c\x2c\xe8\xcb\x37\xf0\xed\x73\x86\xba\x0d\xf0\x64\x7e\x0f\xa6\x90\xa2\x04\x24\x56\x86\x98\xe1\x38\xee\x5c\x3a\x94\xf3\x29\xd0\x8c\x4a\x3d\xb7\x2a\x96\x84\x07\x14\x24\xc0\x53\x36\x22\xe4\xa9\x82\x4c\x84\xb2\x57\x7a\x4b\x3c\x4d\x42\xbb\x30\x18\xd1\x84\xd0\x38\xac\x90\x55\x9f\xa5\x09\xba\xfc\xe7\x54\x85\x68\xfa\xa6\x60\x1f\x4c\x80\x49\x05\xca\x18\xe0\x2c\x87\x41\xb7\x20\x31\x57\x5f\x99\x88\xf9\x9a\x02\x3a\x9b\xb2\x30\x72\x3b\xac\x31\x7a\xcc\x93\x86\x51\x07\xb9\x86\x67\x15\xf2\xc2\xe0\x7a\xdf\x6e\xba\x2a\x66\xbe\x36\x48\x38\x88\xfc\x1e\xbd\x25\x42\x76\xca\xfa\x45\xe9\xcd\xff\x4b\xf6\x8c\xb1\x52\x6a\x43\xcf\x61\xe8\x8d\x22\x12\xf3\xdd\xa3\x66\x1a\x29\xa9\xe0\x0f\x06\xf4\x49\xa1\x8c\x6a\x3c\xe3\x3a\x97\xf7\x94\x26\x91\xfa\xb0\xc5\x48\x45\x4d\x72\x89\x91\xa3\xf2\xa9\x0c\x86\xd7\x5e\x4e\x96\x52\x81\xc3\x3c\x93\x94\xe9\xce\xcf\xc4\xf7\xe2\x94\x9f\x9e\x4e\xea\xda\x6b\x3e\x64\xeb\x12\x63\x2c\xed\x12\xad\x01\x9d\x1d\xad\x8d\xfa\x88\xc7\x6c\x33\x3d\x15\x4a\xf3\x9f\xd4\xec\xea\x7a\xf5\xf1\xf5\xcd\x72\x75\x45\xe4\xb2\x1b\x7a\xaf\x1b\x5b\x26\xd3\x44\xec\x81\xbf\x86\x33\x58\xe4\xc1\x9f\x2e\xf6\x2a\x4b\xfd\x80\x19\xeb\x5e\x55\xd6\xbd\xfc\x17\x0f\xcf\xf6\x3c\x7c\x8c\xff\x48\x42\xdb\xb3\xab\x5f\x2f\x7e\x5e\xbd\xa9\x39\x2b\x85\xf6\x89\x45\x87\xcf\x09\xc3\x04\x94\x96\x0a\xed\x9e\x47\xc9\x7a\x5a\x1a\x9c\xab\xde\xd2\x06\xbc\xc5\xc8\xe9\x87\xc2\xf0\x4a\xe6\x27\x7f\x95\xc7\xff\x40\xd6\x1c\xb6\xe8\x49\xae\x1b\xe2\x01\x8e\xfc\x95\xa1\xbf\x99\x24\x3a\x19\x8a\x9d\xca\xbd\xa9\x92\x87\x23\x4f\x4d\x7e\x1c\xef\x69\x28\x3c\xd1\x5c\x8e\xfc\xe8\xbc\x87\xf4\x30\xd5\x5a\xba\xaa\xfa\xde\x8d\xb3\x27\xf4\x13\x0b\xa9\xc7\x49\xf4\x31\x94\x35\xab\x15\xfd\x16\x04\x11\x62\x23\x6f\x5f\x8a\x6f\x24\x91\xa6\x2c\x6a\x5e\xa7\x81\x73\xb2\xed\xd7\x83\x2a\x7b\x53\xb6\x21\xed\xcd\xd4\x04\x79\x26\xbe\x93\x53\xa8\x9d\xea\xdc\x98\x90\xd7\x22\xc6\xd1\x55\x32\xcf\x82\xd1\x25\xeb\x45\x31\x96\xdd\x74\x1d\x42\x66\xdd\xd6\xd0\x6f\x06\xef\xe1\xb3\xac\x8f\x84\xe9\xfe\x3f\xae\xed\xbd\xe4\x98\x80\xb1\x7f\x02\x00\x00\xff\xff\x92\x57\xec\x6e\xa7\x07\x00\x00")

func bootstrap_digitalocean_k8s_ubuntu_16_04_master_sh() ([]byte, error) {
	return bindata_read(
		_bootstrap_digitalocean_k8s_ubuntu_16_04_master_sh,
		"bootstrap/digitalocean_k8s_ubuntu_16.04_master.sh",
	)
}

var _bootstrap_digitalocean_k8s_ubuntu_16_04_node_sh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x53\x41\x4f\xdc\x3a\x10\xbe\xe7\x57\xcc\xdb\x45\x70\x21\x09\xef\xf4\xa4\x27\x81\x54\xca\x4a\xa5\x08\x5a\xc1\xaa\x55\xa5\x4a\x95\x63\x0f\x89\x89\x63\x47\xf6\x18\x58\xb5\xf4\xb7\x77\xec\x04\xed\xd2\x5e\x7a\x23\x7b\x58\xcf\x37\x9e\x6f\xbe\x19\xcf\x2c\xff\xa9\x63\xf0\x75\xa3\x6d\x8d\xf6\x1e\x1a\x11\xba\x22\x20\x41\x89\x85\x54\xf0\xb3\x28\x96\x50\xbe\xd2\xc7\xa9\xd7\x1d\x06\x84\x7b\x61\x22\x06\x10\x1e\x41\xdb\x3b\x94\x84\x8a\x0f\xe4\x80\x3a\x84\x20\xbd\x1e\xa9\x82\xcf\x98\x2f\xe0\xe3\x68\xb4\xd4\x64\x36\x60\x1d\x41\x0c\xda\xb6\x20\x80\x70\x18\x8d\xa0\x64\x18\x61\xdb\x28\x5a\x04\x26\x98\xe8\x32\xcf\x94\x84\x73\x8a\x90\x3c\x68\xa5\x8b\x3e\x5f\x63\x67\x0c\xe8\x13\x6a\xf4\xa0\xf3\x75\xed\x13\x06\xee\xf6\x05\xb3\x6b\xb5\x64\xce\x74\x81\x9d\xb7\xda\x60\xa8\xe0\x74\x03\x0a\x83\x6e\x2d\x08\x63\xb6\x05\xcc\x45\x85\xce\x45\xa3\x38\x6d\xc3\xfa\x1b\x93\x55\xf1\x31\x3d\x81\x20\xf0\xd1\x92\x1e\xf0\x10\x84\x55\x53\xb5\x1d\x32\xc9\x54\x33\x78\x14\x06\x1e\x9c\xef\x2b\x38\xbf\x85\x8d\x8b\x60\x91\x99\xa5\xb3\x4a\x93\x76\x96\xbd\x59\xd2\x21\x3c\x78\x4d\xdc\x3c\x4a\xe2\xf2\x13\x2f\xc1\x79\x18\x44\xcf\x49\xb9\x4d\x1d\x57\xb7\xcb\x5c\x15\x4b\xfe\xad\x3f\x5c\xac\xae\x8e\x17\xe7\x57\xef\x57\x6f\xd7\xab\xb3\x6c\x2e\x8a\xcb\x37\x37\xeb\xd5\xf5\x16\x9e\x6c\xc6\x57\x37\xef\x4e\xcf\xaf\xcf\xbe\x5d\xac\xbe\xec\x78\x19\x65\x60\xf1\x9a\x63\x94\x66\xf8\xd3\xc7\x2b\xb8\x44\x2e\x5c\x46\x6f\xa0\x23\x1a\xff\xaf\xeb\x81\x81\x46\x7b\x55\x49\x37\xd4\xda\x06\xe2\x07\xaa\x42\x07\x3f\xb8\x19\x05\x0f\x92\xf3\x04\x2f\xca\xda\xdb\xb5\x8a\xe7\x70\xb8\x73\xdc\xd6\xfd\x13\xa8\xef\x85\xaf\xb9\xe3\x5b\x62\x36\x60\xbf\x28\x42\x54\x0e\x72\xe6\x32\xe4\xe4\x81\xb3\x8f\x42\xf6\x3c\x5f\xa1\x92\xc6\x45\x55\xb5\xce\xb5\x06\xb3\x14\x31\x52\xad\x9c\x4c\xff\x65\x8f\x9b\xaa\x1d\xdb\xa4\x29\x91\xcc\x10\x08\xa5\xa0\x9c\x78\xc9\x45\xd9\x41\x8d\x94\x03\xea\xc0\x63\x2b\x99\xd5\xe8\x40\x95\xaa\xfb\xd8\xa0\xb7\x48\x33\x32\x85\x70\x8d\xa5\x84\x03\x94\x9d\x83\x85\xc2\xe6\xb9\x21\x1c\x5f\xed\x04\x68\x57\xc3\xd6\x2c\x1f\xd1\x6a\x1e\xa9\x41\x68\xbb\x80\x93\xbf\xcd\x78\x30\x97\x9f\x94\xb7\x3c\xd5\x71\x54\x82\x67\xb1\xdc\xbc\x84\xe7\xf6\x33\x0e\x5f\x0b\xe0\x2f\x38\xc9\x0b\x30\x9d\xb1\xa1\xb4\x1b\x61\x36\xb9\x37\x3d\x7a\x96\x37\xdb\x89\x83\xbc\xb0\x21\xbd\x58\x99\xfb\x3b\x7b\x92\x16\x83\xb4\x63\x09\x35\x1c\xff\x5b\xfd\x57\x1d\x95\x47\x47\xb3\xb2\xb0\x09\xbc\xc4\x92\x0c\x2f\x7d\x5e\xc1\x89\xff\x77\x27\xeb\xe3\x79\x98\x7d\x93\xb3\x5c\x3d\x73\xf2\x2e\xf2\xca\xfe\x81\xe6\xc9\x28\x4b\x72\x3d\x5a\xd8\xfb\x9e\x57\xe8\x89\x0f\xd3\xd2\x3c\x15\xbf\x02\x00\x00\xff\xff\x2f\xc5\xa4\x9b\x7e\x05\x00\x00")

func bootstrap_digitalocean_k8s_ubuntu_16_04_node_sh() ([]byte, error) {
	return bindata_read(
		_bootstrap_digitalocean_k8s_ubuntu_16_04_node_sh,
		"bootstrap/digitalocean_k8s_ubuntu_16.04_node.sh",
	)
}

var _bootstrap_inject_go = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x93\xd1\x6a\xdb\x30\x14\x86\xaf\xab\xa7\x38\xf8\x2a\x01\xcf\x6e\x7b\x33\xe8\xd8\x85\xd7\x66\xcc\xb4\x38\x10\xa7\x2b\xa5\xf4\x42\x96\x8f\x1d\xad\x8e\xa4\x49\x72\xdd\x30\xfa\x40\x7b\x8d\x3d\xd9\x8e\x6c\x67\xb4\x2c\x04\x84\x7c\xfe\xf3\xe9\xff\x8f\xec\x34\x85\x4b\x6d\x0e\x56\xb6\x3b\x0f\x7f\x7e\xc3\xf9\xe9\xd9\x47\xd8\xee\x10\xae\xfb\x4a\x0a\x6d\x15\x64\xbd\xdf\x69\xeb\x58\x9a\xd2\x1f\x6e\xa4\x40\xe5\xb0\x86\x5e\xd5\x68\xc1\x93\x32\x33\x5c\xd0\x32\x57\x62\xf8\x8e\xd6\x49\xad\xe0\x3c\x39\x85\x45\x10\x44\x73\x29\x5a\x7e\x0a\x88\x83\xee\x61\xcf\x0f\xa0\xb4\x87\xde\x21\x31\xa4\x83\x46\x76\x08\xf8\x22\xd0\x78\x90\x0a\x84\xde\x9b\x4e\x72\x25\x10\x06\xe9\x77\xe3\x39\x33\x25\x09\x8c\xfb\x99\xa1\x2b\xcf\x49\xce\xa9\xc1\xd0\xae\x79\x2b\x04\xee\x67\xd3\xe1\xb7\xf3\xde\x5c\xa4\xe9\x30\x0c\x09\x1f\x0d\x27\xda\xb6\x69\x37\x49\x5d\x7a\x93\x5f\xae\x8a\x72\xf5\x81\x4c\xcf\x4d\xb7\xaa\x43\xe7\xc0\xe2\xcf\x5e\x5a\x0a\x5c\x1d\x80\x1b\x32\x25\x78\x45\x56\x3b\x3e\x80\xb6\xc0\x5b\x8b\x54\xf3\x3a\x98\x1e\xac\xf4\x52\xb5\x31\x38\xdd\xf8\x81\x5b\x0c\x98\x5a\x3a\x6f\x65\xd5\xfb\x77\x33\x3b\x5a\xa4\xe4\x6f\x05\x34\x35\xae\x20\xca\x4a\xc8\xcb\x08\xbe\x64\x65\x5e\xc6\x01\x72\x97\x6f\xbf\xad\x6f\xb7\x70\x97\x6d\x36\x59\xb1\xcd\x57\x25\xac\x37\x70\xb9\x2e\xae\xf2\x6d\xbe\x2e\x68\xf7\x15\xb2\xe2\x1e\xae\xf3\xe2\x2a\x06\xa4\x89\xd1\x39\xf8\x62\x6c\x48\x40\x36\x65\x98\x26\xd6\xe3\xe8\x4a\xc4\x77\x16\x1a\x3d\x59\x72\x06\x85\x6c\xa4\xa0\x68\xaa\xed\x79\x8b\xd0\xea\x67\xb4\x8a\x12\x81\x41\xbb\x97\x2e\xdc\xaa\x23\x83\x75\xc0\x74\x72\x2f\x3d\xf7\xe3\xa3\xff\x72\x25\x8c\xd1\x8c\x9f\x02\xa4\xd2\xda\x53\x40\x6e\x18\x23\x17\xda\x7a\x88\x42\x5e\xd5\xba\x88\xb1\xa6\x57\x02\x72\xf5\x03\x85\x5f\xd4\xdc\x73\x78\x78\xac\x0e\x9e\xde\xa1\x67\xde\xf5\xe8\xe8\x8a\xcd\xc3\xa4\x7e\x9c\x96\x25\x2c\x8e\x12\xb4\x56\xdb\x25\xfc\x62\x27\x54\xba\x0a\xcd\x17\x9f\x61\x52\x8d\xac\x25\x3b\x09\xd1\x9e\x08\x16\x2a\x96\x52\xe1\x91\x4b\x4d\xff\xba\x8e\x4d\x2e\xd9\xa0\xe9\xb8\xc0\xc5\x5c\x89\xc7\xde\x18\xce\x88\xf4\xca\x4e\x2c\xfa\x9e\x3e\x87\xe9\xf8\xa3\x66\x19\x83\x92\x1d\x7b\x65\x7f\x03\x00\x00\xff\xff\xbc\x34\x80\x62\x48\x03\x00\x00")

func bootstrap_inject_go() ([]byte, error) {
	return bindata_read(
		_bootstrap_inject_go,
		"bootstrap/inject.go",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"bootstrap/README.md": bootstrap_readme_md,
	"bootstrap/amazon_k8s_ubuntu_16.04_master.sh": bootstrap_amazon_k8s_ubuntu_16_04_master_sh,
	"bootstrap/amazon_k8s_ubuntu_16.04_node.sh": bootstrap_amazon_k8s_ubuntu_16_04_node_sh,
	"bootstrap/digitalocean_k8s_ubuntu_16.04_master.sh": bootstrap_digitalocean_k8s_ubuntu_16_04_master_sh,
	"bootstrap/digitalocean_k8s_ubuntu_16.04_node.sh": bootstrap_digitalocean_k8s_ubuntu_16_04_node_sh,
	"bootstrap/inject.go": bootstrap_inject_go,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bootstrap": &_bintree_t{nil, map[string]*_bintree_t{
		"README.md": &_bintree_t{bootstrap_readme_md, map[string]*_bintree_t{
		}},
		"amazon_k8s_ubuntu_16.04_master.sh": &_bintree_t{bootstrap_amazon_k8s_ubuntu_16_04_master_sh, map[string]*_bintree_t{
		}},
		"amazon_k8s_ubuntu_16.04_node.sh": &_bintree_t{bootstrap_amazon_k8s_ubuntu_16_04_node_sh, map[string]*_bintree_t{
		}},
		"digitalocean_k8s_ubuntu_16.04_master.sh": &_bintree_t{bootstrap_digitalocean_k8s_ubuntu_16_04_master_sh, map[string]*_bintree_t{
		}},
		"digitalocean_k8s_ubuntu_16.04_node.sh": &_bintree_t{bootstrap_digitalocean_k8s_ubuntu_16_04_node_sh, map[string]*_bintree_t{
		}},
		"inject.go": &_bintree_t{bootstrap_inject_go, map[string]*_bintree_t{
		}},
	}},
}}

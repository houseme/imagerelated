/*
 *  Copyright image-related Author(https://houseme.github.io/imagerelated/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/imagerelated.
 */

package imagerelated

import (
	"errors"
	"net/http"
	"os"
	"strings"
)

// ImageType is a function that returns the image type
func ImageType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buff)
	ext := ImageExtList()
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:]) {
			return filetype, nil
		}
	}
	return "", errors.New("invalid image type")
}

// ImageExt is a function that returns the image extension
func ImageExt(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buff)
	ext := ImageExtList()
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:]) {
			return filetype[6:], nil
		}
	}
	return "", errors.New("invalid image ext")
}

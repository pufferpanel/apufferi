/*
 Copyright 2016 Padduck, LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 	http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package common

import (
	"errors"
	"github.com/pufferpanel/apufferi/logging"
	"os"
	"path/filepath"
	"strings"
)

const MaxRecursivePath = 256

func JoinPath(paths ...string) string {
	result, _ := filepath.Abs(filepath.Join(paths...))
	return result
}

func EnsureAccess(source string, prefix string) bool {
	logging.Devel("Checking " + source)
	replacement, err := findFullPath(source)
	if err != nil && !os.IsNotExist(err) {
		logging.Devel("Error: " + err.Error())
		return false
	} else if os.IsNotExist(err) {
		replacement, err = filepath.Abs(source)
		if err != nil {
			logging.Devel("Error on ABS conversion for path: " + source)
			logging.Devel(err.Error())
			return false
		}
	}
	logging.Devel("Result: " + replacement)

	return strings.HasPrefix(replacement, prefix)
}

func RemoveInvalidSymlinks(files []os.FileInfo, sourceFolder, prefix string) []os.FileInfo{
	i := 0
	for _, v := range files {
		if v.Mode() & os.ModeSymlink != 0{
			if !EnsureAccess(sourceFolder + string(os.PathSeparator) + v.Name(), prefix) {
				logging.Develf("Removing file as symlink with invalid path: %s", sourceFolder + string(os.PathSeparator) + v.Name())
				continue
			}
		}
		files[i] = v
		i++
	}

	return files[:i]
}

func findFullPath(source string) (string, error) {
	fullPath, err := filepath.EvalSymlinks(source)

	if err == nil {
		return fullPath, err
	}

	//if file doesn't exist, then filepath doesn't resolve properly, so check backwards
	if os.IsNotExist(err) {
		var updatePath string
		dir, filename := filepath.Split(source)
		suffix := string(os.PathSeparator) + filename

		i := 0
		for i < MaxRecursivePath && dir != "" {
			dirFullPath, err := filepath.EvalSymlinks(dir)
			if err != nil && os.IsNotExist(err) {
				//update our mapping to look farther down
				suffix = filepath.Base(dir) + string(os.PathSeparator) + suffix
				dir = filepath.Dir(dir)
			} else if err != nil {
				return "", err
			} else {
				//we found a good path!
				updatePath = dirFullPath + string(os.PathSeparator) + suffix
				break
			}
			i++

			if i == MaxRecursivePath {
				return "", errors.New("path too recursive")
			}
		}

		updatePath, err := filepath.Abs(updatePath)
		return updatePath, err

	}

	return "", err
}
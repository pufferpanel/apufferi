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
	"github.com/pufferpanel/apufferi/logging"
	"os"
	"path/filepath"
	"strings"
)

func JoinPath(paths ...string) string {
	result, _ := filepath.Abs(filepath.Join(paths...))
	return result
}

func EnsureAccess(source string, prefix string) bool {
	logging.Devel("Checking " + source)
	replacement, err := filepath.EvalSymlinks(source)
	if err != nil && !os.IsNotExist(err) {
		logging.Devel("Error: " + err.Error())
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

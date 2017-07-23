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

package config

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"strconv"

	"github.com/pufferpanel/apufferi/logging"
)

var config map[string]interface{}

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logging.Error("Error loading config", err)
		config = make(map[string]interface{})
		return
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		logging.Error("Error loading config", err)
	}
}

func Get(key string) string {
	val := config[key]
	if val == nil {
		return ""
	}

	switch val.(type) {
	case string:
		return val.(string)
	case int:
		return strconv.Itoa(val.(int))
	case bool:
		if val.(bool) == true {
			return "true"
		} else {
			return "false"
		}
	default:
		return fmt.Sprintf("%v", val)
	}
}

func GetOrDefault(key string, def string) string {
	val := Get(key)
	if val == "" {
		return def
	}
	return val
}

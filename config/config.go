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

	"github.com/pufferpanel/apufferi/logging"
	"strings"
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

func GetString(key string) string {
	return GetStringOrDefault(key, "")
}

func GetStringOrDefault(key string, def string) string {
	res := get(key)

	switch res.(type) {
	case string:
		return res.(string)
	}
	return def
}

func GetInt(key string) int {
	return GetIntOrDefault(key, 0)
}

func GetIntOrDefault(key string, def int) int {
	res := get(key)

	switch res.(type) {
	case int:
		return res.(int)
	}
	return def
}

func GetBool(key string) bool {
	return GetBoolOrDefault(key, false)
}

func GetBoolOrDefault(key string, def bool) bool {
	val := get(key)

	switch val.(type) {
	case string:
		if val.(string) == "true" {
			return true
		}
		return def
	case bool:
		return val.(bool)
	default:
		return def
	}
}

func get(key string) interface{} {
	value := config[key]
	if value == nil {
		value = config[strings.ToLower(key)]
		if value == nil {
			return nil
		} else {
			return value
		}
		return nil
	}
	return value
}

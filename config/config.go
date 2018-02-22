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

//Deprecated: Use the correct Get### function to get the type needed
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

func GetString(key string) string {
	val := get(key)

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

func GetStringOrDefault(key string, def string) string {
	res := GetString(key)
	if res == "" {
		return def
	} else {
		return res
	}
}

func GetInt(key string) int {
	val := get(key)

	cast, ok := val.(int)
	if ok {
		return cast
	} else {
		return 0
	}
}

func GetIntOrDefault(key string, def int) int {
	res := GetInt(key)

	if res == 0 {
		return def
	} else {
		return res
	}
}

func GetBool(key string) bool {
	val := get(key)

	cast, ok := val.(bool)
	if ok {
		return cast
	} else {
		return false
	}
}

func GetBoolOrDefault(key string, def bool) bool {
	val := get(key)

	switch val.(type) {
	case string:
		return def
	case bool:
		return val.(bool)
	default:
		return def
	}
}

func get(key string) interface{}{
	value := config[key]
	if value == nil {
		return ""
	}
	return value
}
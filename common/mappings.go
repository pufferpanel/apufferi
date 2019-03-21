/*
 Copyright 2019 Padduck, LLC
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

func GetStringOrDefault(data map[string]interface{}, key string, def string) string {
	if data == nil {
		return def
	}
	var section = data[key]
	if section == nil {
		return def
	} else {
		return section.(string)
	}
}

func GetBooleanOrDefault(data map[string]interface{}, key string, def bool) bool {
	if data == nil {
		return def
	}
	var section = data[key]
	if section == nil {
		return def
	} else {
		return section.(bool)
	}
}

func GetMapOrNull(data map[string]interface{}, key string) map[string]interface{} {
	if data == nil {
		return (map[string]interface{})(nil)
	}
	var section = data[key]
	if section == nil {
		return (map[string]interface{})(nil)
	} else {
		return section.(map[string]interface{})
	}
}

func GetObjectArrayOrNull(data map[string]interface{}, key string) []interface{} {
	if data == nil {
		return ([]interface{})(nil)
	}
	var section = data[key]
	if section == nil {
		return ([]interface{})(nil)
	} else {
		return section.([]interface{})
	}
}

func GetStringArrayOrNull(data map[string]interface{}, key string) []string {
	if data == nil {
		return ([]string)(nil)
	}
	var section = data[key]
	if section == nil {
		return ([]string)(nil)
	} else {
		v, k := section.([]string)
		if k {
			return v
		} else {
			var sec = section.([]interface{})
			var newArr = make([]string, len(sec))
			for i := 0; i < len(sec); i++ {
				newArr[i] = sec[i].(string)
			}
			return newArr
		}
	}
}

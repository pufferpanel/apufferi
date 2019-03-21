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

package http

type response struct {
	Success  bool        `json:"success"`
	Message  string      `json:"msg,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Status   int         `json:"-"`
	Code     Code        `json:"code,omitempty"`
	Metadata *metadata   `json:"metadata,omitempty"`
}

type metadata struct {
	Paging *paging `json:"paging"`
}

type paging struct {
	Page    uint `json:"page,omitempty"`
	Size    uint `json:"pageSize,omitempty"`
	MaxSize uint `json:"maxSize,omitempty"`
	Total   uint `json:"total,omitempty"`
}

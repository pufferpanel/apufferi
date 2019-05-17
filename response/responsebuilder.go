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

package response

import "github.com/pufferpanel/apufferi"

type responseBuilder struct {
	response response
	context  Context
}

type Builder interface {
	Send()

	Status(status int) Builder
	WithStatus(status int) Builder

	Message(message string) Builder
	WithMessage(message string) Builder

	Data(data interface{}) Builder
	WithData(data interface{}) Builder

	Error(err error) Builder
	WithError(err error) Builder

	Fail() Builder
	Success() Builder
	WithSuccess(success bool) Builder

	PageInfo(page, pageSize, maxSize, total uint) Builder
	WithPageInfo(page, pageSize, maxSize, total uint) Builder
}

func Respond(c Context) Builder {
	return &responseBuilder{
		response: response{
			Success: true,
			Status:  200,
		},
		context: c,
	}
}

func (rb *responseBuilder) Status(status int) Builder {
	return rb.WithStatus(status)
}

func (rb *responseBuilder) WithStatus(status int) Builder {
	rb.response.Status = status
	if status > 299 || status < 200 {
		return rb.Fail()
	}
	return rb
}

func (rb *responseBuilder) Message(message string) Builder {
	return rb.WithMessage(message)
}

func (rb *responseBuilder) WithMessage(message string) Builder {
	rb.response.Message = message
	return rb
}

func (rb *responseBuilder) Data(data interface{}) Builder {
	return rb.WithData(data)
}

func (rb *responseBuilder) WithData(data interface{}) Builder {
	rb.response.Data = data
	return rb
}

func (rb *responseBuilder) Fail() Builder {
	return rb.WithSuccess(false)
}

func (rb *responseBuilder) Success() Builder {
	return rb.WithSuccess(true)
}

func (rb *responseBuilder) WithSuccess(success bool) Builder {
	rb.response.Success = success
	return rb
}

func (rb *responseBuilder) Send() {
	rb.context.JSON(rb.response.Status, rb.response)
}

func (rb *responseBuilder) PageInfo(page, pageSize, maxSize, total uint) Builder {
	return rb.WithPageInfo(page, pageSize, maxSize, total)
}

func (rb *responseBuilder) WithPageInfo(page, pageSize, maxSize, total uint) Builder {
	rb.response.Metadata = &metadata{
		Paging: &paging{
			Page:    page,
			Size:    pageSize,
			MaxSize: maxSize,
			Total:   total,
		}}
	return rb
}

func (rb *responseBuilder) Error(err error) Builder {
	return rb.WithError(err)
}

func (rb *responseBuilder) WithError(err error) Builder {
	rb.response.Error = apufferi.FromError(err)

	return rb.Fail()
}

type response struct {
	Success  bool           `json:"success"`
	Message  string         `json:"msg,omitempty"`
	Data     interface{}    `json:"data,omitempty"`
	Status   int            `json:"-"`
	Metadata *metadata      `json:"metadata,omitempty"`
	Error    apufferi.Error `json:"error,omitempty"`
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

type Context interface {
	JSON(code int, body interface{})
}

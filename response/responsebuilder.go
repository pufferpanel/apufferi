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

import (
	"github.com/gin-gonic/gin"
	"github.com/pufferpanel/apufferi/v3"
)

type Builder struct {
	Response    *Response
	Context     *gin.Context
	IsDiscarded bool
	IsSent      bool
}

func Respond(c *gin.Context) *Builder {
	return From(c)
}

func From(c *gin.Context) *Builder {
	val := c.Value("response")
	if val == nil {
		return &Builder{
			Response: &Response{
				Success: true,
				Status:  200,
			},
			Context: c,
		}
	} else {
		return val.(*Builder)
	}
}

func (rb *Builder) Status(status int) *Builder {
	return rb.WithStatus(status)
}

func (rb *Builder) WithStatus(status int) *Builder {
	rb.Response.Status = status
	if status > 299 || status < 200 {
		return rb.Fail()
	}
	return rb
}

func (rb *Builder) Message(message string) *Builder {
	return rb.WithMessage(message)
}

func (rb *Builder) WithMessage(message string) *Builder {
	rb.Response.Message = message
	return rb
}

func (rb *Builder) Data(data interface{}) *Builder {
	return rb.WithData(data)
}

func (rb *Builder) WithData(data interface{}) *Builder {
	rb.Response.Data = data
	return rb
}

func (rb *Builder) Fail() *Builder {
	return rb.WithSuccess(false)
}

func (rb *Builder) Success() *Builder {
	return rb.WithSuccess(true)
}

func (rb *Builder) WithSuccess(success bool) *Builder {
	rb.Response.Success = success
	return rb
}

func (rb *Builder) Send() {
	if rb.IsDiscarded || rb.IsSent {
		return
	}
	rb.IsSent = true
	rb.Context.JSON(rb.Response.Status, rb.Response)
}

func (rb *Builder) PageInfo(page, pageSize, maxSize, total uint) *Builder {
	return rb.WithPageInfo(page, pageSize, maxSize, total)
}

func (rb *Builder) WithPageInfo(page, pageSize, maxSize, total uint) *Builder {
	rb.Response.Metadata = &Metadata{
		Paging: &Paging{
			Page:    page,
			Size:    pageSize,
			MaxSize: maxSize,
			Total:   total,
		}}
	return rb
}

func (rb *Builder) Error(err error) *Builder {
	return rb.WithError(err)
}

func (rb *Builder) WithError(err error) *Builder {
	rb.Response.Error = apufferi.FromError(err)

	return rb.Fail()
}

func (rb *Builder) Discard() *Builder {
	rb.IsDiscarded = true
	return rb
}

type Response struct {
	Success  bool           `json:"success"`
	Message  string         `json:"msg,omitempty"`
	Data     interface{}    `json:"data,omitempty"`
	Status   int            `json:"-"`
	Metadata *Metadata      `json:"metadata,omitempty"`
	Error    apufferi.Error `json:"error,omitempty"`
}

type Metadata struct {
	Paging *Paging `json:"paging"`
}

type Paging struct {
	Page    uint `json:"page,omitempty"`
	Size    uint `json:"pageSize,omitempty"`
	MaxSize uint `json:"maxSize,omitempty"`
	Total   uint `json:"total,omitempty"`
}

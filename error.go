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

package apufferi

import "fmt"

type Error interface {
	error

	GetMessage() string

	GetCode() string

	Is(Error) bool

	Set(data ...interface{}) Error

	Metadata(metadata map[string]interface{}) Error
}

type genericError struct {
	Message string                 `json:"msg,omitempty"`
	Code    string                 `json:"code,omitempty"`
	Args    []interface{}          `json:"-"`
	Meta    map[string]interface{} `json:"metadata,omitempty"`
}

func (ge genericError) GetMessage() string {
	return fmt.Sprintf(ge.Message, ge.Args...)
}

func (ge genericError) GetCode() string {
	return ge.Code
}

func (ge genericError) Error() string {
	return ge.GetMessage()
}

func (ge genericError) Is(err Error) bool {
	return ge.GetCode() == err.GetCode()
}

func (ge genericError) Set(machine ...interface{}) Error {
	cp := ge
	cp.Args = machine
	return cp
}

func (ge genericError) Metadata(metadata map[string]interface{}) Error {
	cp := ge
	cp.Meta = metadata
	return cp
}

func CreateError(msg, code string) Error {
	return genericError{
		Message: msg,
		Code:    code,
	}
}

func FromError(err error) Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(Error); e != nil && ok {
		return e
	}
	return genericError{
		Message: err.Error(),
		Code:    "ErrGeneric",
	}
}

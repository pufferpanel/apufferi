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

	GetCode() int

	Is(Error) bool

	Set(data ...interface{}) Error
}

type genericError struct {
	message   string
	code      int
	data      []interface{}
}

func (ge genericError) GetMessage() string {
	return fmt.Sprintf(ge.message, ge.data...)
}

func (ge genericError) GetCode() int {
	return ge.code
}

func (ge genericError) Error() string {
	return fmt.Sprintf(ge.message, ge.data...)
}

func (ge genericError) Is(err Error) bool {
	return ge.GetCode() == err.GetCode()
}

func (ge genericError) Set(machine ...interface{}) Error {
	cp := ge
	cp.data = machine
	return cp
}

func CreateError(msg string, code int) Error {
	return genericError{
		message: msg,
		code:    code,
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
		message: err.Error(),
		code:    0,
	}
}

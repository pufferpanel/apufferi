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

type Error interface {
	error

	GetMessage() string

	GetHumanMessage() string

	GetCode() int
}

type genericError struct {
	message string
	human string
	code int
}

func (ge genericError) GetMessage() string {
	return ge.message
}

func (ge genericError) GetHumanMessage() string {
	return ge.human
}

func (ge genericError) GetCode() int {
	return ge.code
}

func (ge genericError) Error() string {
	if ge.human != "" {
		return ge.human
	} else {
		return ge.message
	}
}

func CreateError(msg, humanMsg string, code int) Error {
	return genericError{
		message: msg,
		human: humanMsg,
		code: code,
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
		human: err.Error(),
		code: 0,
	}
}
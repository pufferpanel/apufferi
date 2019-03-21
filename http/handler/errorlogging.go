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

package handler

import (
	"runtime/debug"

	"github.com/pufferpanel/apufferi/http"
	"github.com/pufferpanel/apufferi/logging"
)

func Recovery() func(Middleware) {
	return func(c Middleware) {
		defer func() {
			if err := recover(); err != nil {
				http.Respond(c).Fail().Status(500).Code(http.UNKNOWN).Message("unexpected error").Data(err).Send()
				logging.Error("Error handling route\n%+v\n%s", err, debug.Stack())
				c.Abort()
			}
		}()

		c.Next()
	}
}

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

package logging

import (
	"fmt"
	"github.com/pufferpanel/apufferi/common"
	"io"
	"os"
	"time"
)

const (
	formatNoData   = "[%s] [%s] %s\n"
	formatWithData = "[%s] [%s] %s\n%v\n"
	timeFormat     = "15:04:05"
)

type logWriter struct {
	writer io.Writer
	level  *Level
	ignore *Level
}

type message struct {
	level   *Level
	message string
	data    []interface{}
}

var (
	writers = make([]*logWriter, 0)
	input   = make(chan *message, 10)
)

func init() {
	//add stdout and stderr to the logging
	//errors go to stderr, regular goes to stdout
	WithWriterIgnore(os.Stdout, INFO, ERROR)
	WithWriter(os.Stderr, ERROR)

	go func() {
		for {
			select {
			case msg := <-input:
				runLogMessage(msg)
			}
		}
	}()
}

func WithWriter(writer io.Writer, lvl *Level) {
	WithWriterIgnore(writer, lvl, nil)
}

func WithWriterIgnore(writer io.Writer, lvl *Level, ignored *Level) {
	writers = append(writers, &logWriter{writer: writer, level: lvl, ignore: ignored})
}

func Close() {
	for _, v := range writers {
		if closer, ok := v.writer.(io.WriteCloser); ok {
			common.Close(closer)
		}
	}
}

func Info(msg string, data ...interface{}) {
	Log(INFO, msg, data)
}

func Warn(msg string, data ...interface{}) {
	Log(WARN, msg, data)
}

func Debug(msg string, data ...interface{}) {
	Log(DEBUG, msg, data)
}

func Error(msg string, data ...interface{}) {
	Log(ERROR, msg, data)
}

func Critical(msg string, data ...interface{}) {
	Log(CRITICAL, msg, data)
}

func Devel(msg string, data ...interface{}) {
	Log(DEVEL, msg, data)
}

func Log(lvl *Level, msg string, data ...interface{}) {
	logMsg := &message{
		level:   lvl,
		message: msg,
		data:    data,
	}
	input <- logMsg
}

func runLogMessage(message *message) {
	dataLength := 0
	if message.data != nil {
		dataLength = len(message.data[0].([]interface{}))
	}
	if message.data == nil || dataLength == 0 {
		var output = fmt.Sprintf(formatNoData, getTimestamp(), message.level.GetName(), message.message)
		logString(message.level, output)
	} else {
		cast := make([]interface{}, 4)
		cast[0] = getTimestamp()
		cast[1] = message.level.GetName()
		cast[2] = message.message
		if dataLength == 1 {
			cast[3] = message.data[0].([]interface{})[0]
		} else {
			cast[3] = message.data[0].([]interface{})
		}
		var output = fmt.Sprintf(formatWithData, cast...)
		logString(message.level, output)
	}
}

func logString(lvl *Level, output string) {
	for _, v := range writers {
		//log to this writer messages which are either the same level or higher, but not over the max
		if lvl.GetScale() >= v.level.GetScale() && (v.ignore == nil || lvl.GetScale() < v.ignore.GetScale()) {
			fmt.Fprint(v.writer, output)
		}
	}
}

func getTimestamp() string {
	return time.Now().Format(timeFormat)
}

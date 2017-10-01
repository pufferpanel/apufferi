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

package logging

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type level struct {
	scale   byte
	display string
}

var (
	DEBUG        level = level{scale: 7, display: "DEBUG"}
	INFO         level = level{scale: 31, display: "INFO"}
	WARN         level = level{scale: 63, display: "WARN"}
	ERROR        level = level{scale: 127, display: "ERROR"}
	CRITICAL     level = level{scale: 255, display: "CRITICAL"}
	logLevel       = INFO
	logFile      *os.File
)

const (
	logFileName    = "2006-01-02T15-04-05.log"
	formatNoData   = "[%s] [%s] %s\n"
	formatWithData = "[%s] [%s] %s\n%v\n"
)

var (
	logFileFolder = "logs"
)

func SetLogFolder(path string) {
	logFileFolder = path
}

func Init() {
	var err error
	os.Mkdir(logFileFolder, 0755)
	logFile, err = os.OpenFile(path.Join(logFileFolder, time.Now().Format(logFileName)), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		Critical("Could not create log file", err)
	}
}

func Info(msg string, data ...interface{}) {
	log(INFO, msg, data)
}

func Warn(msg string, data ...interface{}) {
	log(WARN, msg, data)
}

func Debug(msg string, data ...interface{}) {
	log(DEBUG, msg, data)
}

func Error(msg string, data ...interface{}) {
	log(ERROR, msg, data)
}

func Critical(msg string, data ...interface{}) {
	log(CRITICAL, msg, data)
}

func Infof(msg string, data ...interface{}) {
	logf(INFO, msg, data)
}

func Warnf(msg string, data ...interface{}) {
	logf(WARN, msg, data)
}

func Debugf(msg string, data ...interface{}) {
	logf(DEBUG, msg, data)
}

func Errorf(msg string, data ...interface{}) {
	logf(ERROR, msg, data)
}

func Criticalf(msg string, data ...interface{}) {
	logf(CRITICAL, msg, data)
}

func SetLevel(lvl level) {
	logLevel = lvl
}

func SetLevelByString(lvl string) {
	switch strings.ToUpper(lvl) {
	case "DEBUG":
		SetLevel(DEBUG)
	case "INFO":
		SetLevel(INFO)
	case "WARN":
		SetLevel(WARN)
	case "ERROR":
		SetLevel(ERROR)
	case "CRITICAL":
		SetLevel(CRITICAL)
	}
}

func log(lvl level, msg string, data ...interface{}) {
	var dataLength = len(data[0].([]interface{}))
	if data == nil || dataLength == 0 {
		var output = fmt.Sprintf(formatNoData, getTimestamp(), lvl.display, msg)
		logString(lvl, output)
	} else {
		cast := make([]interface{}, 4)
		cast[0] = getTimestamp()
		cast[1] = lvl.display
		cast[2] = msg
		if dataLength == 1 {
			cast[3] = data[0].([]interface{})[0]
		} else {
			cast[3] = data[0].([]interface{})
		}
		var output = fmt.Sprintf(formatWithData, cast...)
		logString(lvl, output)
	}
}

func logf(lvl level, msg string, data ...interface{}) {
	if data == nil || len(data[0].([]interface{})) == 0 {
		var output = fmt.Sprintf(formatNoData, getTimestamp(), lvl.display, msg)
		logString(lvl, output)
	} else {
		var output = fmt.Sprintf(formatNoData, getTimestamp(), lvl.display, fmt.Sprintf(msg, data[0].([]interface{})...))
		logString(lvl, output)
	}
}

func logString(lvl level, output string) {
	if lvl.scale >= logLevel.scale {
		fmt.Print(output)
	}
	logFile.WriteString(output)
	logFile.Sync()
}

func getTimestamp() string {
	return time.Now().Format("15:04:05")
}

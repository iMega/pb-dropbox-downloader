// Copyright © 2020 Dmitry Stoletov <info@imega.ru>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"io"
	"log"
	"strings"
)

const (
	errorLog int = iota
	infoLog
	debugLog
)

type Logger struct {
	loggers []*log.Logger
}

func New(logLevel string, w io.Writer) *Logger {
	label := []string{
		errorLog: "ERROR",
		infoLog:  "INFO",
		debugLog: "DEBUG",
	}
	errorW := io.Discard
	infoW := io.Discard
	debugW := io.Discard

	if w != nil {
		switch logLevel {
		case "", label[errorLog], strings.ToLower(label[errorLog]):
			infoW = w
			errorW = w
		case label[debugLog], strings.ToLower(label[debugLog]):
			debugW = w
			infoW = w
			errorW = w
		}
	}

	return &Logger{
		loggers: []*log.Logger{
			log.New(errorW, label[errorLog]+": ", log.LstdFlags),
			log.New(infoW, label[infoLog]+": ", log.LstdFlags),
			log.New(debugW, label[debugLog]+": ", log.LstdFlags),
		},
	}
}

func (l *Logger) GetWriter() io.Writer {
	return l.loggers[infoLog].Writer()
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.loggers[infoLog].Printf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.loggers[errorLog].Printf(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.loggers[debugLog].Printf(format, args...)
}

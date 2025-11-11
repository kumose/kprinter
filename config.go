// Copyright (C) Kumo inc. and its affiliates.
// Author: Jeff.li lijippy@163.com
// All rights reserved.
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package kprinter

import "go.uber.org/zap/zapcore"

type PrintMode int

const (
	NormalMode PrintMode = iota
	ProductionMode
	FileMode
)

type LogLevel int

const (
	DebugLevel LogLevel = iota + 10
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type Config struct {
	Mode      PrintMode
	Level     LogLevel
	LogFile   string
	WarnToErr bool
	AllToErr  bool
	Silence   bool
}

var PrintConfig = Config{
	NormalMode,
	InfoLevel,
	"",
	false,
	false,
	false,
}

func IsConsole() bool {
	return PrintConfig.Mode != FileMode
}

func zapLevel(l LogLevel) zapcore.Level {
	switch l {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	}
	return zapcore.InfoLevel
}

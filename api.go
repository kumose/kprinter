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

import (
	"fmt"
	"os"

	perrs "github.com/kumose/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Printer interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Debug(format string)
	Info(format string)
	Warn(format string)
	Error(format string)
	Fatal(format string)
}

var (
	colorConsole = NewConsoleColorPrinter()
)
var GlobalPrinter Printer = colorConsole

// Verbose logs debug message with formatting
func Verbose(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func Formatf(cf func(string, ...any) string, format string, args ...any) string {
	return cf(format, args...)
}

func Printf(cf func(string, ...any) string, format string, args ...any) {
	fmt.Fprint(os.Stderr, cf(format, args...))
}

// PrinterInit initializes the global logger based on PrintConfig.Mode
func PrinterInit(conf Config) error {
	var core zapcore.Core
	var err error

	PrintConfig = conf
	switch PrintConfig.Mode {
	case NormalMode:
		return initConsoleColorPrinter()
	case ProductionMode:
		// Production also uses stderr but with different config
		return initConsoleRawPrinter()
	case FileMode:
		core, err = newFileCore()
		if err != nil {
			return perrs.Errorf("Failed to initialize file logging: %s", err.Error())
		}
		logger := zap.New(core)
		zap.ReplaceGlobals(logger)
	}
	return perrs.Errorf("unknown mode %v", PrintConfig.Mode)
}

// Debugf logs debug message with formatting
func Debugf(format string, args ...any) {
	if PrintConfig.Level > DebugLevel {
		return
	}
	GlobalPrinter.Debugf(format, args...)
}

// Debug logs debug message with formatting
func Debug(format string) {
	if PrintConfig.Level > DebugLevel {
		return
	}
	GlobalPrinter.Debug(format)
}

// Debugf logs debug message with formatting
func Infof(format string, args ...any) {
	if PrintConfig.Level > InfoLevel {
		return
	}
	GlobalPrinter.Infof(format, args...)
}

// Debug logs debug message with formatting
func Info(format string) {
	if PrintConfig.Level > DebugLevel {
		return
	}
	GlobalPrinter.Info(format)
}

// Debugf logs debug message with formatting
func Warnf(format string, args ...any) {
	GlobalPrinter.Warnf(format, args...)
}

// Debugf logs debug message with formatting
func Warn(format string) {
	GlobalPrinter.Warn(format)
}

// Debugf logs debug message with formatting
func Errorf(format string, args ...any) {
	GlobalPrinter.Errorf(format, args...)
}

// Debugf logs debug message with formatting
func Error(format string) {
	GlobalPrinter.Error(format)
}

// Debugf logs debug message with formatting
func Fatalf(format string, args ...any) {
	GlobalPrinter.Fatalf(format, args...)
}

// Debugf logs debug message with formatting
func Fatal(format string) {
	GlobalPrinter.Fatal(format)
}

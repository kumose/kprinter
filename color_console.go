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
	"io"
	"os"
)

type consoleColorPrinter struct {
	normalWriter io.Writer
	warnWriter   io.Writer
	errorWriter  io.Writer
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Debugf(format string, args ...any) {
	fmt.Fprint(l.normalWriter, Yellowf(format, args...))
}

// Debug logs debug message with formatting
func (l *consoleColorPrinter) Debug(format string) {
	fmt.Fprint(l.normalWriter, Yellow(format))
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Infof(format string, args ...any) {
	fmt.Fprint(l.normalWriter, Greenf(format, args...))
}

// Debug logs debug message with formatting
func (l *consoleColorPrinter) Info(format string) {
	fmt.Fprint(l.normalWriter, Green(format))
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Warnf(format string, args ...any) {
	fmt.Fprint(l.warnWriter, Magentaf(format, args...))
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Warn(format string) {
	fmt.Fprint(l.warnWriter, Magenta(format))
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Errorf(format string, args ...any) {
	fmt.Fprint(l.errorWriter, Redf(format, args...))
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Error(format string) {
	fmt.Fprint(l.errorWriter, Red(format))
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Fatalf(format string, args ...any) {
	fmt.Fprint(l.errorWriter, RedOnWhitef(format, args...))
}

// Debugf logs debug message with formatting
func (l *consoleColorPrinter) Fatal(format string) {
	fmt.Fprint(l.errorWriter, RedOnWhite(format))
}

func NewConsoleColorPrinter() Printer {
	return &consoleColorPrinter{
		os.Stdout,
		os.Stdout,
		os.Stderr,
	}
}

func initConsoleColorPrinter() error {
	c := &consoleRawPrinter{
		os.Stdout,
		os.Stdout,
		os.Stderr,
	}
	if PrintConfig.WarnToErr {
		c.warnWriter = os.Stderr
	}
	if PrintConfig.AllToErr {
		c.warnWriter = os.Stderr
		c.normalWriter = os.Stderr
		c.errorWriter = os.Stderr
	}
	return nil
}

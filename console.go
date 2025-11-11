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

type consoleRawPrinter struct {
	normalWriter io.Writer
	warnWriter   io.Writer
	errorWriter  io.Writer
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Debugf(format string, args ...any) {
	fmt.Fprintf(l.normalWriter, format, args...)
}

// Debug logs debug message with formatting
func (l *consoleRawPrinter) Debug(format string) {
	fmt.Fprint(l.normalWriter, format)
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Infof(format string, args ...any) {
	fmt.Fprintf(l.normalWriter, format, args...)
}

// Debug logs debug message with formatting
func (l *consoleRawPrinter) Info(format string) {
	fmt.Fprint(l.normalWriter, format)
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Warnf(format string, args ...any) {
	fmt.Fprintf(l.warnWriter, format, args...)
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Warn(format string) {
	fmt.Fprint(l.warnWriter, format)
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Errorf(format string, args ...any) {
	fmt.Fprintf(l.errorWriter, format, args...)
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Error(format string) {
	fmt.Fprint(l.errorWriter, format)
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Fatalf(format string, args ...any) {
	fmt.Fprintf(l.errorWriter, format, args...)
}

// Debugf logs debug message with formatting
func (l *consoleRawPrinter) Fatal(format string) {
	fmt.Fprint(l.errorWriter, format)
}

func NewConsoleRawPrinter() Printer {
	return &consoleRawPrinter{
		os.Stdout,
		os.Stdout,
		os.Stderr,
	}
}

func initConsoleRawPrinter() error {
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

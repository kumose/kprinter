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
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newFileCore creates core that writes to file
func newFileCore() (zapcore.Core, error) {
	logDir := filepath.Base(PrintConfig.LogFile)

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}
	logFile, err := os.OpenFile(PrintConfig.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(logFile),
		zapLevel(PrintConfig.Level),
	)

	return core, nil
}

type loggerPrinter struct{}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Debugf(format string, args ...any) {
	zap.L().Debug(fmt.Sprintf(format, args...))
}

// Debug logs debug message with formatting
func (l *loggerPrinter) Debug(format string) {
	zap.L().Debug(format)
}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Infof(format string, args ...any) {
	zap.L().Info(fmt.Sprintf(format, args...))
}

// Debug logs debug message with formatting
func (l *loggerPrinter) Info(format string) {
	zap.L().Info(format)
}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Warnf(format string, args ...any) {
	zap.L().Warn(fmt.Sprintf(format, args...))
}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Warn(format string) {
	zap.L().Warn(format)
}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Errorf(format string, args ...any) {
	zap.L().Error(fmt.Sprintf(format, args...))
}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Error(format string) {
	zap.L().Error(format)
}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Fatalf(format string, args ...any) {
	zap.L().Fatal(fmt.Sprintf(format, args...))
}

// Debugf logs debug message with formatting
func (l *loggerPrinter) Fatal(format string) {
	zap.L().Fatal(format)
}

func NewLoggerPrinter() Printer {
	return &loggerPrinter{}
}

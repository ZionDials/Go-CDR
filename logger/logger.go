// Copyright (c) 2023 Zion Dials <me@ziondials.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package logger

import (
	"fmt"
	"os"

	"github.com/ziondials/go-cdr/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger() {

	conf := config.GetLoggerFromGlobalConfig()

	lumber := &lumberjack.Logger{
		Filename: conf.Path + "/" + conf.Name,
		MaxSize:  int(conf.MaxSize),
		MaxAge:   int(conf.MaxAge),
		Compress: conf.Compress,
	}

	writer := zapcore.AddSync(lumber)

	logLevel := translateLogLevel(conf.Level)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		// write to stdout as well as log files
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer),
		logLevel,
	)

	if logLevel.Level() == zap.DebugLevel {
		Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		Logger = zap.New(core)
	}
}

func Info(format string, a ...any) {
	Logger.Info(fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	Logger.Error(fmt.Sprintf(format, a...))
}

func Debug(format string, a ...any) {
	Logger.Debug(fmt.Sprintf(format, a...))
}

func Panic(format string, a ...any) {
	Logger.Panic(fmt.Sprintf(format, a...))
}

func Fatal(format string, a ...any) {
	Logger.Fatal(fmt.Sprintf(format, a...))
}

// Translates a string log level to a zap.AtomicLevel
func translateLogLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	}
}

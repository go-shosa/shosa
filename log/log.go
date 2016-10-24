package log

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var lg *logrus.Logger

func init() {
	lg = logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	// default use color (TTY).
	lg.Formatter = &logrus.TextFormatter{DisableColors: true}

	// Output to stderr instead of stdout, could also be a file.
	lg.Out = os.Stderr

	lv := os.Getenv("LOG_LEVEL")
	if lv == "" {
		lv = "info"
	}
	level, err := logrus.ParseLevel(lv)
	if err != nil {
		level = logrus.InfoLevel
	}
	lg.Level = level
}

func Log() *logrus.Logger {
	return lg
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return lg.WithFields(fields)
}

func Debugf(format string, args ...interface{}) {
	lg.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	lg.Infof(format, args)
}

func Printf(format string, args ...interface{}) {
	lg.Printf(format, args)
}

func Warnf(format string, args ...interface{}) {
	lg.Warnf(format, args)
}

func Warningf(format string, args ...interface{}) {
	lg.Warningf(format, args)
}

func Errorf(format string, args ...interface{}) {
	lg.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	lg.Fatalf(format, args)
}

func Panicf(format string, args ...interface{}) {
	lg.Panicf(format, args)
}

func Debug(args ...interface{}) {
	lg.Debug(args)
}

func Info(args ...interface{}) {
	lg.Info(args)
}

func Print(args ...interface{}) {
	lg.Print(args)
}

func Warn(args ...interface{}) {
	lg.Warn(args)
}

func Warning(args ...interface{}) {
	lg.Warning(args)
}

func Error(args ...interface{}) {
	lg.Error(args)
}

func Fatal(args ...interface{}) {
	lg.Fatal(args)
}

func Panic(args ...interface{}) {
	lg.Panic(args)
}

func Debugln(args ...interface{}) {
	lg.Debugln(args)
}

func Infoln(args ...interface{}) {
	lg.Infoln(args)
}

func Println(args ...interface{}) {
	lg.Println(args)
}

func Warnln(args ...interface{}) {
	lg.Warnln(args)
}

func Warningln(args ...interface{}) {
	lg.Warningln(args)
}

func Errorln(args ...interface{}) {
	lg.Errorln(args)
}

func Fatalln(args ...interface{}) {
	lg.Fatalln(args)
}

func Panicln(args ...interface{}) {
	lg.Panicln(args)
}

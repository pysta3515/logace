package logace

import (
	"log"
	"os"
)

// Logger log interface
type Logger interface {
	SysLogger

	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
}

// SysLogger sys log interface
type SysLogger interface {
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicln(args ...interface{})
	Panicf(format string, args ...interface{})
}

var entity interface{} = log.New(os.Stderr, "", log.LstdFlags)

// SetLogEntity set a log entity
func SetLogEntity(e interface{}) {
	entity = e
}

// Print ...
func Print(args ...interface{}) {
	entity.(SysLogger).Print(args...)
}

// Println ...
func Println(args ...interface{}) {
	entity.(SysLogger).Println(args...)
}

// Printf ...
func Printf(format string, args ...interface{}) {
	entity.(SysLogger).Printf(format, args...)
}

// Faltal ...
func Faltal(args ...interface{}) {
	entity.(SysLogger).Fatal(args...)
}

// Fatalln ...
func Fatalln(args ...interface{}) {
	entity.(SysLogger).Fatalln(args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	entity.(SysLogger).Fatalf(format, args...)
}

// Panic ...
func Panic(args ...interface{}) {
	entity.(SysLogger).Panic(args...)
}

// Panicln ...
func Panicln(args ...interface{}) {
	entity.(SysLogger).Panic(args...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	entity.(SysLogger).Panicf(format, args...)
}

// by default funcs below not impleted! DO NOT USE THEM

// Info ...
func Info(args ...interface{}) {
	entity.(Logger).Info(args...)
}

// Infoln ...
func Infoln(args ...interface{}) {
	entity.(Logger).Infoln(args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	entity.(Logger).Infof(format, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	entity.(Logger).Warn(args...)
}

// Warnln ...
func Warnln(args ...interface{}) {
	entity.(Logger).Warnln(args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	entity.(Logger).Warnf(format, args...)
}

// Error ...
func Error(args ...interface{}) {
	entity.(Logger).Error(args...)
}

// Errorln ...
func Errorln(args ...interface{}) {
	entity.(Logger).Errorln(args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	entity.(Logger).Errorf(format, args...)
}

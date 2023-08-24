package mylog

import (
	"io"
	"log"
	"os"
)

// Logging context

type LoggingContext struct {
	outputPath string

	InfoLogger    *log.Logger
	WarnLogger    *log.Logger
	ErrLogger     *log.Logger
	ConsoleLogger *log.Logger
}

type ILoggingContext interface {
	Setup(string) *LoggingContext
	Console(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

func (ctx *LoggingContext) Setup(filePath string) *LoggingContext {
	var lCtx *LoggingContext
	lCtx = &LoggingContext{outputPath: filePath}

	file, err := os.OpenFile(lCtx.outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error setting up file logger: %s", err)
	}

	w := io.MultiWriter(file, os.Stdout)

	lCtx.InfoLogger = log.New(w, "INFO | ", log.Lshortfile|log.Ldate|log.Ltime)
	lCtx.WarnLogger = log.New(w, "WARN | ", log.Lshortfile|log.Ldate|log.Ltime)
	lCtx.ErrLogger = log.New(w, "ERR  | ", log.Lshortfile|log.Ldate|log.Ltime)
	lCtx.ConsoleLogger = log.New(os.Stdout, "CONSOLE | ", log.Lshortfile|log.Ldate|log.Ltime)

	return lCtx
}

// General functions.

func (ctx *LoggingContext) Console(output string) {
	ctx.ConsoleLogger.Println(output)
}

func (ctx *LoggingContext) Info(output string) {
	ctx.InfoLogger.Println(output)
}

func (ctx *LoggingContext) Warn(output string) {
	ctx.WarnLogger.Println(output)
}

func (ctx *LoggingContext) Error(output string) {
	ctx.ErrLogger.Println(output)
}

func (ctx *LoggingContext) Fatal(output string) {
	ctx.ErrLogger.Fatal(output)
}

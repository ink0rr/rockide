package rockide

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type ctxLogger struct{}

func getLogger() *log.Logger {
	var home string
	if runtime.GOOS == "windows" {
		home = os.Getenv("UserProfile")
	} else {
		home = os.Getenv("HOME")
	}
	fileName := filepath.Join(home, ".rockide", "log.txt")
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file")
	}
	return log.New(logFile, "[rockide]", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogger(ctx context.Context) *log.Logger {
	logger, ok := ctx.Value(ctxLogger{}).(*log.Logger)
	if !ok {
		panic("Failed to get logger")
	}
	return logger
}

func WithLogger(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, ctxLogger{}, getLogger())
	return ctx
}

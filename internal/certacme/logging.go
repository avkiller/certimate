package certacme

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	legolog "github.com/go-acme/lego/v4/log"
)

type legoLogger struct {
	callLogger *slog.Logger
	legoLogger legolog.StdLogger
}

func (l *legoLogger) Fatal(args ...any) {
	l.callLogger.Error("go-acme/lego: " + fmt.Sprint(args...))
	l.legoLogger.Fatal(args...)
}

func (l *legoLogger) Fatalln(args ...any) {
	l.Fatal(fmt.Sprintln(args...))
}

func (l *legoLogger) Fatalf(format string, args ...any) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *legoLogger) Print(args ...any) {
	message := fmt.Sprint(args...)
	print := l.callLogger.Debug
	if strings.HasPrefix(message, "[WARN] ") {
		message = strings.TrimPrefix(message, "[WARN] ")
		print = l.callLogger.Warn
	} else if strings.HasPrefix(message, "[INFO] ") {
		message = strings.TrimPrefix(message, "[INFO] ")
		print = l.callLogger.Info
	}

	print("go-acme/lego: " + message)
	l.legoLogger.Print(message)
}

func (l *legoLogger) Println(args ...any) {
	l.Print(fmt.Sprintln(args...))
}

func (l *legoLogger) Printf(format string, args ...any) {
	l.Print(fmt.Sprintf(format, args...))
}

func NewLegoLogger(logger *slog.Logger) legolog.StdLogger {
	return &legoLogger{
		callLogger: logger,

		// https://github.com/go-acme/lego/blob/master/log/logger.go
		legoLogger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

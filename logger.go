package ctxlog

import (
	"log"
	"os"

	formatter "github.com/nolleh/caption_json_formatter"
	"github.com/sirupsen/logrus"
)

var (
	Logger = NewLogger()
)

func NewLogger() *logrus.Logger {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.SetFormatter(&formatter.Formatter{PrettyPrint: true})
	return logger
}

func Log() *formatter.Entry {
	return &formatter.Entry{Entry: logrus.NewEntry(Logger)}
}
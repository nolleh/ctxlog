package ctxlog

import (
	"log"
	"os"

	formatter "github.com/nolleh/caption_json_formatter"
	"github.com/sirupsen/logrus"
)

var (
	logger, console = newLogger()
)

func newLogger() (*logrus.Logger, *formatter.Formatter) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	logger := logrus.New()
	logger.Level = logrus.TraceLevel

	consoleFormat := formatter.Console()
	logger.SetFormatter(consoleFormat)
	return logger, consoleFormat
}

func Logger() *logrus.Logger {
	return logger
}

func CaptionJsonFormatter() *formatter.Formatter {
	return console
}

/*
Log makes extended logrus entry.
@example ctxlog.Log().Debug(Message) where Message is struct or string, whatever
*/
func Log() *formatter.Entry {
	return &formatter.Entry{Entry: logrus.NewEntry(logger)}
}

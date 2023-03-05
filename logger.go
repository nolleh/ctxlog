package ctxlog

import (
	"log"
	"os"

	formatter "github.com/nolleh/caption_json_formatter"
	"github.com/sirupsen/logrus"
)

var (
	logger = newLogger()
)

func newLogger() *logrus.Logger {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.SetFormatter(&formatter.Formatter{PrettyPrint: false})
	return logger
}

/*
Log makes extended logrus entry.
@example ctxlog.Log().Debug(Message) where Message is struct or string, whatever
*/
func Log() *formatter.Entry {
	return &formatter.Entry{Entry: logrus.NewEntry(logger)}
}

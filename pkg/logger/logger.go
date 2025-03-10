package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = "\033[36m" // Cyan
	case logrus.InfoLevel:
		levelColor = "\033[32m" // Green
	case logrus.WarnLevel:
		levelColor = "\033[33m" // Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = "\033[31m" // Red
	default:
		levelColor = "\033[0m" // Reset
	}

	return []byte(fmt.Sprintf(
		"%s[%s]%s \033[1m%s\033[0m \033[2m(%s:%d)\033[0m: %s\n",
		levelColor,
		entry.Time.Format("2006-01-02 15:04:05"),
		levelColor,
		entry.Level.String(),
		entry.Caller.File,
		entry.Caller.Line,
		entry.Message,
	)), nil
}

func Init() *Logger {
	logger := logrus.New()

	stderrOutput := os.Stderr
	logger.SetOutput(stderrOutput)
	logger.SetFormatter(&CustomFormatter{})

	fileOutput, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(stderrOutput, fileOutput)
	logger.SetOutput(multiWriter)

	logger.SetFormatter(&CustomFormatter{})

	logger.SetReportCaller(true)

	logger.SetLevel(logrus.DebugLevel)

	return &Logger{logger}
}
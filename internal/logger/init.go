package logger

import (
	"log"
	"os"
)

type logger struct {
	level int
}

var ProbeLog *logger

var WriteLog func(int) *log.Logger

func init() {
	WriteLog = func(logType int) *log.Logger {
		return log.New(os.Stdout, getLogTypeName(logType), log.Ltime)
	}

	ProbeLog = newLogger(NOINFO)
}

func newLogger(logType int) *logger {
	return &logger{
		level: logType,
	}
}

func getLogTypeName(logType int) string {
	switch logType {
	case INFO:
		return "INFO\t"
	case WARNING:
		return "WARNING\t"
	case ERROR:
		return "ERROR\t"
	case INFO1:
		return "INFO1\t"
	case STAT1:
		return "STAT1\t"
	default:
		return "DEBUG\t"
	}
}

func (l *logger) Write(logType int, log interface{}) {
	if l.Validate(logType) {
		WriteLog(logType).Println(log)
	}
	return
}

func (l *logger) Validate(logType int) bool {

	ruleList := []ApplyLogLevel{
		ShowAllRule{},
		NotShowRule{},
		GeneralRule{},
		ErrorWarningRule{},
	}

	for _, rule := range ruleList {
		if rule.apply(logType, l.level) {
			return true
		}
	}

	return false
}

func (l *logger) SetLevel(level int) {
	l.level = level
}

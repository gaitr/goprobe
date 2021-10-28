package logger

type LogType struct {
	logType int
}

const (
	NOINFO = iota
	ERROR
	WARNING
	INFO
	DEBUG
	INFO1
	STAT1
)

package logger

type ApplyLogLevel interface {
	apply(logType, level int) bool
}

type GeneralRule struct{}

func (rule GeneralRule) apply(logType, level int) bool {
	return logType == level
}

type ShowAllRule struct{}

func (rule ShowAllRule) apply(_, level int) bool {
	return level == DEBUG
}

type NotShowRule struct{}

func (rule NotShowRule) apply(logType, _ int) bool {
	return logType == NOINFO
}

type ErrorWarningRule struct{}

func (rule ErrorWarningRule) apply(logType, level int) bool {
	return logType == WARNING && level == ERROR
}

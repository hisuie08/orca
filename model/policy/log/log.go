package log

type LogLevel int

const (
	LogSilent LogLevel = iota
	LogNormal
	LogDetail
)

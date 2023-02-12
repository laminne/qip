package logger

type Logger interface {
	Debug(s string)
	Info(s string)
	Warn(s string)
	Error(s string)
	Panic(s string)
	Fatal(s string)
}

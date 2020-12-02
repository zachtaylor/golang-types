package types

// Logger is a logging interface different from std log
type Logger interface {
	New() Logger
	Add(string, interface{}) Logger
	With(Dict) Logger
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Out(...interface{})
}

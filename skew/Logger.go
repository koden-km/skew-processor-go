package skewd

// TODO:
// "Logger" should be the interface and something like "Log" should be the implementation in Go style?
// Also check if Go has some base logging classes.
// I think standard Go style is lower case file names also?

/* type Logger interface { */
/* 	Log(level string, message string, context [string]string) */
/* 	Progress(percentage float32) */
/* } */

type Logger struct {
	processor *Processor
}

func NewLogger(processor *Processor) *Logger {
	return &Logger{processor}
}

// Logs with an arbitrary level.
func (l *Logger) Log(level string, message string, context [string]string) {
	l.processor.SendJobLog(level, message)
}

func (l *Logger) Progress(percentage float32) {
	l.processor.SendJobProgress(percentage)
}

package logger

type Logger struct{}

func NewLogger() (*Logger, error) {
	return &Logger{}, nil
}

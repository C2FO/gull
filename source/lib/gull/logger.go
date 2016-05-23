package gull

import "fmt"

type ILogger interface {
	Info(message string, args ...interface{})
	Debug(message string, args ...interface{})
}

func NewLogger(verbose bool) ILogger {
	if verbose {
		return NewVerboseLogger()
	} else {
		return NewNullLogger()
	}
}

type NullLogger struct{}

func NewNullLogger() ILogger {
	return &NullLogger{}
}

func (nl *NullLogger) Info(message string, args ...interface{}) {

}

func (nl *NullLogger) Debug(message string, args ...interface{}) {

}

type VerboseLogger struct{}

func NewVerboseLogger() ILogger {
	return &VerboseLogger{}
}

func (vl *VerboseLogger) Info(message string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf("%v\n", fmt.Sprintf(message, args...))
	} else {
		fmt.Printf("%v\n", message)
	}

}

func (vl *VerboseLogger) Debug(message string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf("%v\n", fmt.Sprintf(message, args...))
	} else {
		fmt.Printf("%v\n", message)
	}
}

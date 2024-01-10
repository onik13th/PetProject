package errors

import (
	"fmt"
	"runtime"
)

type Err struct {
	Error string `json:"error"`
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	pc, file, line, ok := runtime.Caller(1) // 1 means one level up the call stack
	if !ok {
		return fmt.Errorf("%s: %v", message, err)
	}

	funcName := runtime.FuncForPC(pc).Name()
	stackTrace := fmt.Sprintf("[%s:%d %s]", file, line, funcName)

	// Return an error that shows the message, the original error, and the stacktrace
	return fmt.Errorf("%s: %v - Trace: %s", message, err, stackTrace)
}

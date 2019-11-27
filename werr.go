package werr

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Error struct {
	Err   error
	Stack []runtime.Frame
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) Error() string {
	return e.Err.Error()
}

// remove lines before skip in suffix
// ex : ServeHTTP
// and path of current file
func PrintSkip(err error, skip string) string {
	s := ""
	var e Error
	if errors.As(err, &e) {
		if len(e.Stack) > 0 {
			for i := len(e.Stack) - 1; i >= 0; i-- {
				frame := e.Stack[i]
				if skip != "" && strings.HasSuffix(frame.Function, skip) {
					s = ""
					continue
				}
				file := frame.File
				s += fmt.Sprintf("--- %s\n", frame.Function)
				s += fmt.Sprintf("%s:%d\n", file, frame.Line)
			}
		}
	} else {
		return fmt.Sprintf("%+v", err) // pkg.errors stack
	}
	return s + err.Error()
}

func Print(err error) string {
	s := ""
	var e Error
	if errors.As(err, &e) {
		if len(e.Stack) > 0 {
			for i := len(e.Stack) - 1; i >= 0; i-- {
				frame := e.Stack[i]
				s += fmt.Sprintf("--- %s\n", frame.Function)
				s += fmt.Sprintf("%s:%d\n", frame.File, frame.Line)
			}
		}
	}
	return s + err.Error()
}

// wrap error with stack only if not already
// error is wrapped with fmt.Errorf(msg + " : %w",err)
func Wrapf(err error, msg string, args ...interface{}) error {
	s := fmt.Sprintf(msg, args...)
	err = fmt.Errorf(s+" : %w", err)
	var es Error
	if errors.As(err, &es) {
		return err
	}
	stk := getStackTrace(3)
	return Error{Err: err, Stack: stk}
}

// add stack trace to an error if it's not
func Stack(e error) error {
	var es Error
	if errors.As(e, &es) {
		return e
	}
	stk := getStackTrace(3)
	return Error{Err: e, Stack: stk}
}

// format stack trace to be printed line by line
// in reverse order
func getStackTrace(nb int) []runtime.Frame {
	stackBuf := make([]uintptr, 1024)
	length := runtime.Callers(nb, stackBuf[:])
	stack := stackBuf[:length]

	trace := []runtime.Frame{}
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "runtime") {
			trace = append(trace, frame)
		}
		if !more {
			break
		}
	}
	return trace
}

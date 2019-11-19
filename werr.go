package werr

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Error struct {
	Err   error
	Stack []string
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) Error() string {
	return e.Err.Error()
}
func Print(err error) string {
	s := ""
	var e Error
	if errors.As(err, &e) {
		if len(e.Stack) > 0 {
			path, _ := os.Getwd()
			for i := len(e.Stack) - 1; i >= 0; i-- {
				fl := e.Stack[i]
				if strings.HasPrefix(fl, path) {
					fl = fl[len(path)+1 : len(fl)]
				}
				s += fmt.Sprintf("%s\n", fl)
			}
		}
	}
	return s + err.Error()
}

func Wrapf(e error, msg string, args ...interface{}) error {
	s := fmt.Sprintf(msg, args...)
	e = fmt.Errorf(s+" : %w", e)
	var es Error
	if errors.As(e, &es) {
		return e
	}
	stk := getStackTrace(3)
	return Error{Err: e, Stack: stk} //fmt.Sprintf("%s line %d", file, line)}
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
func getStackTrace(nb int) []string {
	stackBuf := make([]uintptr, 50)
	length := runtime.Callers(nb, stackBuf[:])
	stack := stackBuf[:length]

	trace := []string{}
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			trace = append(trace, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
		}
		if !more {
			break
		}
	}
	return trace
}

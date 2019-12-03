package werr

import (
	"errors"
	"fmt"
	"io"
	"os"
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

// FprintSkip write SprintSkip to writer
func FprintSkip(w io.Writer, err error, skip string) {
	w.Write([]byte(SprintSkip(err, skip)))
}

// PrintSkip write SprintSkip to stdout
func PrintSkip(err error, skip string) {
	FprintSkip(os.Stdout, err, skip)
}

// PrintSkip remove lines before skip in suffix
// ex : ServeHTTP
// return as string
func SprintSkip(err error, skip string) string {
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

// Fprint write traceback in f
func Fprint(f io.Writer, err error) {
	f.Write([]byte(Sprint(err)))
}

// Sprint return traceback as string
func Sprint(err error) string {
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

// Print print traceback to stdout
func Print(err error) {
	Fprint(os.Stdout, err)
}

// Wrapf wrap error with stack only if not already
// error is wrapped with fmt.Errorf(msg + " : %w",err)
func Wrapf(err error, msg string, args ...interface{}) error {
	if err == nil {
		return err
	}
	s := fmt.Sprintf(msg, args...)
	err = fmt.Errorf(s+" : %w", err)
	var es Error
	if errors.As(err, &es) {
		return err
	}
	stk := getStackTrace(3)
	return Error{Err: err, Stack: stk}
}

// Stack add stack trace to an error if it's not
func Stack(e error) error {
	if e == nil {
		return e
	}
	var es Error
	if errors.As(e, &es) {
		return e
	}
	stk := getStackTrace(3)
	return Error{Err: e, Stack: stk}
}

// getStackTrace return Frames after nb
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

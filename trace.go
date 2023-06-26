package werr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// Allow trace or disable for the project
var WithTrace = true

// Wrapf return formated error with trace and : %w
func Wrapf(err error, s string, vals ...any) error {
	vals = append(vals, err)
	return tracef(2, s+": %w", vals...)
}

// Wrap return error with trace and : %w
func Wrap(err error) error {
	return Wrapf(err, "")
}

// Errorf like fmt.Errorf with trace
func Errorf(s string, vals ...any) error {
	return tracef(2, s, vals...)
}

// New like errors.New with trace
func New(s string) error {
	return tracef(2, s)
}

// Trace add trace to the error
func Trace(err error) error {
	if err == nil {
		return err
	}
	return tracef(2, err.Error())
}

// tracef add trace before calling fmt.Errorf
func tracef(skip int, s string, vals ...any) error {
	pc, file, line, ok := runtime.Caller(skip)
	if ok && WithTrace {
		splt := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		name := strings.Join(splt[1:], "/")
		info := fmt.Sprintf("\n> %s() %s:%d\n",
			name,
			filepath.Base(file), line)
		s = info + strings.TrimSpace(s)
	}
	return fmt.Errorf(s, vals...)
}

// Copyright (c) 2023 William Dode. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package werr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// Allow trace or disable for the project
var WithTrace = true

// Show full path instead of basename
var WithFullPath = false

// Show full name instead of [1:]
var WithFullName = false

var WithShortName = true

// Show only package/method():line
var WithJustMethod = true

// Wrapf returns formated error with trace and : %w
func Wrapf(err error, s string, vals ...any) error {
	if err == nil {
		return nil
	}
	vals = append(vals, err)
	return tracef(2, s+": %w", vals...)
}

// Wrap returns error with trace and %w
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return tracef(2, "%w", err)
}

// Errorf is like fmt.Errorf with trace
// wrap if %w
// not wrap if %v
func Errorf(s string, vals ...any) error {
	return tracef(2, s, vals...)
}

// New is like errors.New with trace
func New(s string) error {
	return tracef(2, s)
}

// tracef add trace before calling fmt.Errorf
func tracef(skip int, s string, vals ...any) error {
	pc, file, line, ok := runtime.Caller(skip)
	if ok && WithTrace {
		name := runtime.FuncForPC(pc).Name()
		if !WithFullName {
			splt := strings.Split(runtime.FuncForPC(pc).Name(), "/")
			name = strings.Join(splt[1:], "/")
		}
		if WithShortName {
			splt := strings.Split(runtime.FuncForPC(pc).Name(), "/")
			name = splt[len(splt)-1]
		}
		path := file
		if !WithFullPath {
			path = filepath.Base(file)
		}
		info := ""
		if WithJustMethod {
			info = fmt.Sprintf("\n> %s(%d): ", name, line)
		} else {
			info = fmt.Sprintf("\n> %s() %s:%d\n",
				name, path, line)
		}
		s = info + strings.TrimSpace(s)
	}
	return fmt.Errorf(s, vals...)
}

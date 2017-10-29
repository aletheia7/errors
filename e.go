// Copyright 2016 aletheia7. All rights reserved. Use of this source code is
// governed by a BSD-2-Clause license that can be found in the LICENSE file.

// Package errors augments an error with the file and line number of where it
// occurred.

package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type e_stack struct {
	err  error
	file string
	line int
}

func (o *e_stack) Error() string {
	return fmt.Sprintf("%v %v:%v", o.err.Error(), o.file, o.line)
}

type causer interface {
	Cause() error
}

// Cause returns the original error without the file and line number
//
func Cause(err error) error {
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		return cause.Cause()
	}
	return err
}

// Implements causer
//
func (o *e_stack) Cause() error {
	return o.err
}

// Adds file/line number to the error
//
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	r := &e_stack{err: err}
	r.file, r.line = file_line()
	return r
}

// New returns a new error
//
func New(msg string) error {
	r := &e_stack{err: fmt.Errorf(msg)}
	r.file, r.line = file_line()
	return r
}

// Errorf returns a new error with format
//
func Errorf(format string, args ...interface{}) error {
	r := &e_stack{err: fmt.Errorf(format, args...)}
	r.file, r.line = file_line()
	return r
}

func file_line() (file string, line int) {
	pc := make([]uintptr, 1)
	n := runtime.Callers(3, pc)
	if n == 0 {
		return ``, 0
	}
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return trim_go_path(frame.Function, frame.File), frame.Line
}

func trim_go_path(name, file string) string {
	// Here we want to get the source file path relative to the compile time
	// GOPATH. As of Go 1.6.x there is no direct way to know the compiled
	// GOPATH at runtime, but we can infer the number of path segments in the
	// GOPATH. We note that fn.Name() returns the function name qualified by
	// the import path, which does not include the GOPATH. Thus we can trim
	// segments from the beginning of the file path until the number of path
	// separators remaining is one more than the number of path separators in
	// the function name. For example, given:
	//
	//    GOPATH     /home/user
	//    file       /home/user/src/pkg/sub/file.go
	//    fn.Name()  pkg/sub.Type.Method
	//
	// We want to produce:
	//
	//    pkg/sub/file.go
	//
	// From this we can easily see that fn.Name() has one less path separator
	// than our desired output. We count separators from the end of the file
	// path until it finds two more than in the function name and then move
	// one character forward to preserve the initial path segment without a
	// leading separator.
	const sep = "/"
	goal := strings.Count(name, sep) + 2
	i := len(file)
	for n := 0; n < goal; n++ {
		i = strings.LastIndex(file[:i], sep)
		if i == -1 {
			// not enough separators found, set i so that the slice expression
			// below leaves file unmodified
			i = -len(sep)
			break
		}
	}
	// get back to 0 or trim the leading separator
	file = file[i+len(sep):]
	return file
}

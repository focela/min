// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package debug

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const (
	maxCallerDepth = 1000
	stackFilterKey = "/debug/debug"
)

var (
	goRootForFilter  = runtime.GOROOT()
	binaryVersion    = "" // Version of the current running binary (uint64 hex).
	binaryVersionMd5 = "" // Version of the current running binary (MD5).
	selfPath         = "" // Absolute path of the current running binary.
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// Caller returns the function name and the absolute file path along with its line
// number of the caller.
func Caller(skip ...int) (function string, path string, line int) {
	return CallerWithFilter(nil, skip...)
}

// CallerWithFilter returns the function name and the absolute file path along with
// its line number of the caller.
//
// The parameter `filters` is used to filter the path of the caller.
func CallerWithFilter(filters []string, skip ...int) (function string, path string, line int) {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}
	pc, file, line, start := callerFromIndex(filters)
	if start != -1 {
		for i := start + number; i < maxCallerDepth; i++ {
			if i != start {
				pc, file, line, _ = runtime.Caller(i)
			}
			if file == "" || filterFileByFilters(file, filters) {
				continue
			}
			function = "unknown"
			if fn := runtime.FuncForPC(pc); fn != nil {
				function = fn.Name()
			}
			return function, file, line
		}
	}
	return "", "", -1
}

// callerFromIndex returns the caller position and according information exclusive of the
// debug package.
func callerFromIndex(filters []string) (pc uintptr, file string, line int, index int) {
	for index = 0; index < maxCallerDepth; index++ {
		pc, file, line, _ = runtime.Caller(index)
		if file != "" && !filterFileByFilters(file, filters) {
			if index > 0 {
				index--
			}
			return
		}
	}
	return 0, "", -1, -1
}

func filterFileByFilters(file string, filters []string) bool {
	if file == "" || strings.Contains(file, stackFilterKey) {
		return true
	}
	for _, filter := range filters {
		if strings.Contains(file, filter) {
			return true
		}
	}
	if goRootForFilter != "" && strings.HasPrefix(file, goRootForFilter) {
		fileSeparator := file[len(goRootForFilter)]
		if fileSeparator == filepath.Separator || fileSeparator == '\\' || fileSeparator == '/' {
			return true
		}
	}
	return false
}

// CallerPackage returns the package name of the caller.
func CallerPackage() string {
	function, _, _ := Caller()
	indexSplit := strings.LastIndexByte(function, '/')
	if indexSplit == -1 {
		return function[:strings.IndexByte(function, '.')]
	}
	leftPart := function[:indexSplit+1]
	rightPart := function[indexSplit+1:]
	indexDot := strings.IndexByte(rightPart, '.')
	return leftPart + rightPart[:indexDot]
}

// CallerFunction returns the function name of the caller.
func CallerFunction() string {
	function, _, _ := Caller()
	function = filepath.Base(function)
	return function[strings.IndexByte(function, '.')+1:]
}

// CallerFilePath returns the file path of the caller.
func CallerFilePath() string {
	_, path, _ := Caller()
	return path
}

// CallerDirectory returns the directory of the caller.
func CallerDirectory() string {
	_, path, _ := Caller()
	return filepath.Dir(path)
}

// CallerFileLine returns the file path along with the line number of the caller.
func CallerFileLine() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, path, line)
}

// CallerFileLineShort returns the file name along with the line number of the caller.
func CallerFileLineShort() string {
	_, path, line := Caller()
	return fmt.Sprintf(`%s:%d`, filepath.Base(path), line)
}

// FuncPath returns the complete function path of given `f`.
func FuncPath(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// FuncName returns the function name of given `f`.
func FuncName(f interface{}) string {
	return filepath.Base(FuncPath(f))
}

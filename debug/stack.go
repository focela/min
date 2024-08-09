// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package debug

import (
	"bytes"
	"fmt"
	"runtime"
)

// PrintStack prints to standard error the stack trace returned by runtime.Stack.
func PrintStack(skip ...int) {
	fmt.Print(Stack(skip...))
}

// Stack returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
func Stack(skip ...int) string {
	return stackWithFilters(nil, skip...)
}

// StackWithFilter returns a formatted stack trace of the goroutine that calls it.
// The parameter `filter` is used to filter the path of the caller.
func StackWithFilter(filter string, skip ...int) string {
	var filters []string
	if filter != "" {
		filters = append(filters, filter)
	}
	return stackWithFilters(filters, skip...)
}

// StackWithFilters returns a formatted stack trace of the goroutine that calls it.
// The parameter `filters` is a slice of strings used to filter the path of the caller.
//
// TODO Improve the performance using debug.Stack.
func StackWithFilters(filters []string, skip ...int) string {
	return stackWithFilters(filters, skip...)
}

// stackWithFilters is the internal function that implements stack trace formatting with filtering.
func stackWithFilters(filters []string, skip ...int) string {
	number := 0
	if len(skip) > 0 {
		number = skip[0]
	}
	var (
		name                  string
		space                 = "  "
		index                 = 1
		buffer                = bytes.NewBuffer(nil)
		pc, file, line, start = callerFromIndex(filters)
		ok                    bool
	)
	for i := start + number; i < maxCallerDepth; i++ {
		if i != start {
			pc, file, line, ok = runtime.Caller(i)
		}
		if !ok {
			break
		}
		if filterFileByFilters(file, filters) {
			continue
		}
		if fn := runtime.FuncForPC(pc); fn == nil {
			name = "unknown"
		} else {
			name = fn.Name()
		}
		if index > 9 {
			space = " "
		}
		buffer.WriteString(fmt.Sprintf("%d.%s%s\n    %s:%d\n", index, space, name, file, line))
		index++
	}
	return buffer.String()
}

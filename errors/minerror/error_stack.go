// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/focela/min/internal/consts"
	"github.com/focela/min/internal/errors"
)

// stackInfo manages stack info of a certain error.
type stackInfo struct {
	Index   int          // Index of the current error in the whole error stack.
	Message string       // Error information string.
	Lines   []*stackLine // Slice contains all error stack lines of the current error stack in sequence.
}

// stackLine manages each line info of the stack.
type stackLine struct {
	Function string // Function name, which contains its full package path.
	FileLine string // FileLine is the source file name and its line number of the Function.
}

// Stack returns the error stack information as string.
func (err *Error) Stack() string {
	if err == nil {
		return ""
	}
	var (
		loop             = err
		index            = 1
		infos            []*stackInfo
		isStackModeBrief = errors.IsStackModeBrief()
	)
	for loop != nil {
		info := &stackInfo{
			Index:   index,
			Message: fmt.Sprintf("%-v", loop),
		}
		index++
		infos = append(infos, info)
		loopLinesOfStackInfo(loop.stack, info, isStackModeBrief)
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				loop = e
			} else {
				infos = append(infos, &stackInfo{
					Index:   index,
					Message: loop.error.Error(),
				})
				index++
				break
			}
		} else {
			break
		}
	}
	filterLinesOfStackInfos(infos)
	return formatStackInfos(infos)
}

// filterLinesOfStackInfos removes repeated lines, which exist in subsequent stacks, from top errors.
func filterLinesOfStackInfos(infos []*stackInfo) {
	var (
		set     = make(map[string]struct{})
		removes []int
	)
	for i := len(infos) - 1; i >= 0; i-- {
		info := infos[i]
		for j, line := range info.Lines {
			if _, ok := set[line.FileLine]; ok {
				removes = append(removes, j)
			} else {
				set[line.FileLine] = struct{}{}
			}
		}
		for _, j := range removes {
			info.Lines = append(info.Lines[:j], info.Lines[j+1:]...)
		}
		removes = removes[:0]
	}
}

// formatStackInfos formats and returns error stack information as string.
func formatStackInfos(infos []*stackInfo) string {
	var buffer = bytes.NewBuffer(nil)
	for i, info := range infos {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, info.Message))
		if len(info.Lines) > 0 {
			formatStackLines(buffer, info.Lines)
		}
	}
	return buffer.String()
}

// formatStackLines formats and returns error stack lines as string.
func formatStackLines(buffer *bytes.Buffer, lines []*stackLine) {
	for i, line := range lines {
		space := "  "
		if i >= 9 {
			space = " "
		}
		buffer.WriteString(fmt.Sprintf(
			"   %d).%s%s\n        %s\n",
			i+1, space, line.Function, line.FileLine,
		))
	}
}

// loopLinesOfStackInfo iterates the stack info lines and produces the stack line info.
func loopLinesOfStackInfo(st stack, info *stackInfo, isStackModeBrief bool) {
	if st == nil {
		return
	}
	for _, p := range st {
		if fn := runtime.FuncForPC(p - 1); fn != nil {
			file, line := fn.FileLine(p - 1)
			if isStackModeBrief {
				if strings.Contains(file, consts.StackFilterKey) {
					continue
				}
			} else {
				if strings.Contains(file, stackFilterModulePath) {
					continue
				}
			}
			if strings.Contains(file, "<") {
				continue
			}
			if goRootForFilter != "" && strings.HasPrefix(file, goRootForFilter) {
				continue
			}
			info.Lines = append(info.Lines, &stackLine{
				Function: fn.Name(),
				FileLine: fmt.Sprintf(`%s:%d`, file, line),
			})
		}
	}
}

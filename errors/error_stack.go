// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"bytes"
	"container/list"
	"fmt"
	"runtime"
	"strings"

	"github.com/focela/aid/internal/consts"
	"github.com/focela/aid/internal/errors"
)

// stackInfo manages stack info of certain error.
type stackInfo struct {
	Index   int        // Index is the index of current error in whole error stacks.
	Message string     // Error information string.
	Lines   *list.List // Lines contains all error stack lines of current error stack in sequence.
}

// stackLine manages each line info of stack.
type stackLine struct {
	Function string // Function name, which contains its full package path.
	FileLine string // FileLine is the source file name and its line number of Function.
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
		ok      bool
		set     = make(map[string]struct{})
		removes []*list.Element
	)
	for i := len(infos) - 1; i >= 0; i-- {
		info := infos[i]
		if info.Lines == nil {
			continue
		}
		for e := info.Lines.Front(); e != nil; e = e.Next() {
			line := e.Value.(*stackLine)
			if _, ok = set[line.FileLine]; ok {
				removes = append(removes, e)
			} else {
				set[line.FileLine] = struct{}{}
			}
		}
		for _, e := range removes {
			info.Lines.Remove(e)
		}
		removes = removes[:0]
	}
}

// formatStackInfos formats and returns error stack information as string.
func formatStackInfos(infos []*stackInfo) string {
	buffer := bytes.NewBuffer(nil)
	for i, info := range infos {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, info.Message))
		if info.Lines != nil && info.Lines.Len() > 0 {
			formatStackLines(buffer, info.Lines)
		}
	}
	return buffer.String()
}

// formatStackLines formats and returns error stack lines as string.
func formatStackLines(buffer *bytes.Buffer, lines *list.List) {
	var (
		space  = "  "
		length = lines.Len()
	)
	for i, e := 0, lines.Front(); i < length; i, e = i+1, e.Next() {
		line := e.Value.(*stackLine)
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
			if isStackModeBrief && strings.Contains(file, consts.StackFilterKeyForAid) {
				continue
			}
			if !isStackModeBrief && strings.Contains(file, stackFilterKeyLocal) {
				continue
			}
			if strings.Contains(file, "<") {
				continue
			}
			if goRootForFilter != "" &&
				len(file) >= len(goRootForFilter) &&
				file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if info.Lines == nil {
				info.Lines = list.New()
			}
			info.Lines.PushBack(&stackLine{
				Function: fn.Name(),
				FileLine: fmt.Sprintf(`%s:%d`, file, line),
			})
		}
	}
}

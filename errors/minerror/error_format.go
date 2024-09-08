// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import (
	"fmt"
	"io"
)

// Format formats the error according to the fmt.Formatter interface.
//
// %v, %s   : Print all the error string;
// %-v, %-s : Print current level error string;
// %+s      : Print full stack error list;
// %+v      : Print the error string and full stack error list.
func (err *Error) Format(s fmt.State, verb rune) {
	var output string

	switch verb {
	case 's', 'v':
		switch {
		case s.Flag('-'):
			if err.text != "" {
				output = err.text
			} else {
				output = err.Error()
			}
		case s.Flag('+'):
			if verb == 's' {
				output = err.Stack()
			} else {
				output = err.Error() + "\n" + err.Stack()
			}
		default:
			output = err.Error()
		}
	}

	_, _ = io.WriteString(s, output)
}

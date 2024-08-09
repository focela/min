// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"io"
)

// ReadCloser implements the io.ReadCloser interface
// used for reading the request body content multiple times.
//
// Note that it cannot be closed.
type ReadCloser struct {
	index      int    // Current read position.
	content    []byte // Content to be read.
	repeatable bool   // Indicates whether the content can be read repeatedly.
}

// NewReadCloser creates and returns a ReadCloser object.
func NewReadCloser(content []byte, repeatable bool) io.ReadCloser {
	return &ReadCloser{
		content:    content,
		repeatable: repeatable,
	}
}

// Read implements the io.ReadCloser interface.
func (b *ReadCloser) Read(p []byte) (n int, err error) {
	if b.index >= len(b.content) {
		if !b.repeatable {
			return 0, io.EOF
		}
		// Make it repeatable reading.
		b.index = 0
	}
	n = copy(p, b.content[b.index:])
	b.index += n
	if b.index >= len(b.content) && !b.repeatable {
		err = io.EOF
	}
	return n, err
}

// Close implements the io.ReadCloser interface.
func (b *ReadCloser) Close() error {
	return nil
}

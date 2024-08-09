// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package rand

import (
	"crypto/rand"

	"github.com/focela/aid/errors"
	"github.com/focela/aid/errors/code"
)

const (
	// Buffer size for uint32 random number.
	bufferChanSize = 10000
)

var (
	// bufferChan is the buffer for random bytes,
	// every item storing 4 bytes.
	bufferChan = make(chan []byte, bufferChanSize)
)

func init() {
	go produceRandomBufferBytes()
}

// produceRandomBufferBytes is a named goroutine, which uses an asynchronous goroutine
// to produce the random bytes, and a buffer chan to store the random bytes.
// So it has high performance to generate random numbers.
func produceRandomBufferBytes() {
	var step int
	for {
		buffer := make([]byte, 1024)
		n, err := rand.Read(buffer)
		if err != nil {
			panic(errors.WrapCode(code.CodeInternalError, err, `error reading random buffer from system`))
		}

		// The random buffer from system is very expensive,
		// so fully reuse the random buffer by changing
		// the step with a different number can
		// improve the performance a lot.
		// for _, step = range []int{4, 5, 6, 7} {
		for _, step = range []int{4} {
			for i := 0; i <= n-4; i += step {
				bufferChan <- buffer[i : i+4]
			}
		}
	}
}

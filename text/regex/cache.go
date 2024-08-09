// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package regex

import (
	"regexp"
	"sync"

	"github.com/focela/aid/errors"
)

var (
	regexMu sync.RWMutex

	// Cache for regex object.
	// Note that:
	// 1. It uses sync.RWMutex ensuring concurrent safety.
	// 2. There's no expiring logic for this map.
	regexMap = make(map[string]*regexp.Regexp)
)

// getRegexp returns *regexp.Regexp object with the given `pattern`.
// It uses cache to enhance the performance for compiling regular expression patterns,
// meaning it will return the same *regexp.Regexp object for the same pattern.
//
// It is concurrent-safe for multiple goroutines.
func getRegexp(pattern string) (*regexp.Regexp, error) {
	// Retrieve the regular expression object using reading lock.
	regexMu.RLock()
	regex, exists := regexMap[pattern]
	regexMu.RUnlock()

	// If the pattern is already cached, return it.
	if exists {
		return regex, nil
	}

	// If it does not exist in the cache, compile the pattern and create one.
	regexMu.Lock()
	defer regexMu.Unlock()

	// Double check to prevent duplicate compilation if multiple goroutines are accessing the function.
	if regex, exists = regexMap[pattern]; exists {
		return regex, nil
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.Wrapf(err, `regexp.Compile failed for pattern "%s"`, pattern)
	}

	// Cache the compiled regex object.
	regexMap[pattern] = regex
	return regex, nil
}

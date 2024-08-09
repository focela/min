// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package regex provides high performance API for regular expression functionality.
package regex

import (
	"regexp"
)

// Quote quotes `s` by replacing special chars in `s`
// to match the rules of regular expression pattern.
// It returns the quoted string.
//
// Example: Quote(`[foo]`) returns `\[foo\]`.
func Quote(s string) string {
	return regexp.QuoteMeta(s)
}

// Validate checks whether the given regular expression pattern `pattern` is valid.
func Validate(pattern string) error {
	_, err := getRegexp(pattern)
	return err
}

// IsMatch checks whether the given bytes `src` match the `pattern`.
func IsMatch(pattern string, src []byte) bool {
	r, err := getRegexp(pattern)
	if err != nil {
		return false
	}
	return r.Match(src)
}

// IsMatchString checks whether the given string `src` matches the `pattern`.
func IsMatchString(pattern string, src string) bool {
	return IsMatch(pattern, []byte(src))
}

// Match returns the byte slices that match the `pattern`.
func Match(pattern string, src []byte) ([][]byte, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil, err
	}
	return r.FindSubmatch(src), nil
}

// MatchString returns the strings that match the `pattern`.
func MatchString(pattern string, src string) ([]string, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil, err
	}
	return r.FindStringSubmatch(src), nil
}

// MatchAll returns all byte slices that match the `pattern`.
func MatchAll(pattern string, src []byte) ([][][]byte, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil, err
	}
	return r.FindAllSubmatch(src, -1), nil
}

// MatchAllString returns all strings that match the `pattern`.
func MatchAllString(pattern string, src string) ([][]string, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil, err
	}
	return r.FindAllStringSubmatch(src, -1), nil
}

// Replace replaces all matches of `pattern` in bytes `src` with bytes `replace`.
func Replace(pattern string, replace, src []byte) ([]byte, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil, err
	}
	return r.ReplaceAll(src, replace), nil
}

// ReplaceString replaces all matches of `pattern` in string `src` with string `replace`.
func ReplaceString(pattern, replace, src string) (string, error) {
	r, err := Replace(pattern, []byte(replace), []byte(src))
	return string(r), err
}

// ReplaceFunc replaces all matches of `pattern` in bytes `src`
// using a custom replacement function `replaceFunc`.
func ReplaceFunc(pattern string, src []byte, replaceFunc func(b []byte) []byte) ([]byte, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil, err
	}
	return r.ReplaceAllFunc(src, replaceFunc), nil
}

// ReplaceFuncMatch replaces all matches of `pattern` in bytes `src`
// using a custom replacement function `replaceFunc`.
// The parameter `match` for `replaceFunc` is of type [][]byte,
// which is the result containing all sub-patterns of `pattern` using the Match function.
func ReplaceFuncMatch(pattern string, src []byte, replaceFunc func(match [][]byte) []byte) ([]byte, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil, err
	}
	return r.ReplaceAllFunc(src, func(bytes []byte) []byte {
		match, _ := Match(pattern, bytes)
		return replaceFunc(match)
	}), nil
}

// ReplaceStringFunc replaces all matches of `pattern` in string `src`
// using a custom replacement function `replaceFunc`.
func ReplaceStringFunc(pattern string, src string, replaceFunc func(s string) string) (string, error) {
	bytes, err := ReplaceFunc(pattern, []byte(src), func(bytes []byte) []byte {
		return []byte(replaceFunc(string(bytes)))
	})
	return string(bytes), err
}

// ReplaceStringFuncMatch replaces all matches of `pattern` in string `src`
// using a custom replacement function `replaceFunc`.
// The parameter `match` for `replaceFunc` is of type []string,
// which is the result containing all sub-patterns of `pattern` using the MatchString function.
func ReplaceStringFuncMatch(pattern string, src string, replaceFunc func(match []string) string) (string, error) {
	r, err := getRegexp(pattern)
	if err != nil {
		return "", err
	}
	return string(r.ReplaceAllFunc([]byte(src), func(bytes []byte) []byte {
		match, _ := MatchString(pattern, string(bytes))
		return []byte(replaceFunc(match))
	})), nil
}

// Split slices `src` into substrings separated by the expression and returns a slice of
// the substrings between those expression matches.
func Split(pattern string, src string) []string {
	r, err := getRegexp(pattern)
	if err != nil {
		return nil
	}
	return r.Split(src, -1)
}

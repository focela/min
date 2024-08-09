// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package times provides functionality for measuring and displaying time.
//
// This package should keep much less dependencies with other packages.
package times

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/focela/aid/errors"
	"github.com/focela/aid/errors/code"
	"github.com/focela/aid/internal/intlog"
	"github.com/focela/aid/internal/utils"
	"github.com/focela/aid/text/regex"
)

// Time is a wrapper for time.Time for additional features.
type Time struct {
	wrapper
}

// UnixNanoProvider is an interface for custom time.Time wrappers that provide UnixNano functionality.
type UnixNanoProvider interface {
	UnixNano() int64
}

const (
	// Short writes for common usage durations.
	D  = 24 * time.Hour
	H  = time.Hour
	M  = time.Minute
	S  = time.Second
	MS = time.Millisecond
	US = time.Microsecond
	NS = time.Nanosecond

	// Regular expression1(datetime separator supports '-', '/', '.').
	timeRegexPattern1 = `(\d{4}[-/\.]\d{1,2}[-/\.]\d{1,2})[:\sT-]*(\d{0,2}:{0,1}\d{0,2}:{0,1}\d{0,2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`

	// Regular expression2(datetime separator supports '-', '/', '.').
	timeRegexPattern2 = `(\d{1,2}[-/\.][A-Za-z]{3,}[-/\.]\d{4})[:\sT-]*(\d{0,2}:{0,1}\d{0,2}:{0,1}\d{0,2}){0,1}\.{0,1}(\d{0,9})([\sZ]{0,1})([\+-]{0,1})([:\d]*)`

	// Regular expression3(time).
	timeRegexPattern3 = `(\d{2}):(\d{2}):(\d{2})\.{0,1}(\d{0,9})`
)

var (
	timeRegex1 = regexp.MustCompile(timeRegexPattern1)
	timeRegex2 = regexp.MustCompile(timeRegexPattern2)
	timeRegex3 = regexp.MustCompile(timeRegexPattern3)

	monthMap = map[string]int{
		"jan":       1,
		"feb":       2,
		"mar":       3,
		"apr":       4,
		"may":       5,
		"jun":       6,
		"jul":       7,
		"aug":       8,
		"sep":       9,
		"sept":      9,
		"oct":       10,
		"nov":       11,
		"dec":       12,
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
)

// Timestamp retrieves and returns the timestamp in seconds.
func Timestamp() int64 {
	return Now().Timestamp()
}

// TimestampMilli retrieves and returns the timestamp in milliseconds.
func TimestampMilli() int64 {
	return Now().TimestampMilli()
}

// TimestampMicro retrieves and returns the timestamp in microseconds.
func TimestampMicro() int64 {
	return Now().TimestampMicro()
}

// TimestampNano retrieves and returns the timestamp in nanoseconds.
func TimestampNano() int64 {
	return Now().TimestampNano()
}

// TimestampStr is a convenience method which retrieves and returns
// the timestamp in seconds as string.
func TimestampStr() string {
	return Now().TimestampStr()
}

// TimestampMilliStr is a convenience method which retrieves and returns
// the timestamp in milliseconds as string.
func TimestampMilliStr() string {
	return Now().TimestampMilliStr()
}

// TimestampMicroStr is a convenience method which retrieves and returns
// the timestamp in microseconds as string.
func TimestampMicroStr() string {
	return Now().TimestampMicroStr()
}

// TimestampNanoStr is a convenience method which retrieves and returns
// the timestamp in nanoseconds as string.
func TimestampNanoStr() string {
	return Now().TimestampNanoStr()
}

// Date returns current date in string like "2006-01-02".
func Date() string {
	return time.Now().Format("2006-01-02")
}

// Datetime returns current datetime in string like "2006-01-02 15:04:05".
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// ISO8601 returns current datetime in ISO8601 format like "2006-01-02T15:04:05-07:00".
func ISO8601() string {
	return time.Now().Format("2006-01-02T15:04:05-07:00")
}

// RFC822 returns current datetime in RFC822 format like "Mon, 02 Jan 06 15:04 MST".
func RFC822() string {
	return time.Now().Format("Mon, 02 Jan 06 15:04 MST")
}

// parseDateStr parses the string to year, month and day numbers.
func parseDateStr(s string) (year, month, day int) {
	array := strings.Split(s, "-")
	if len(array) < 3 {
		array = strings.Split(s, "/")
	}
	if len(array) < 3 {
		array = strings.Split(s, ".")
	}
	if len(array) < 3 {
		return
	}
	if utils.IsNumeric(array[1]) {
		year, _ = strconv.Atoi(array[0])
		month, _ = strconv.Atoi(array[1])
		day, _ = strconv.Atoi(array[2])
	} else {
		if v, ok := monthMap[strings.ToLower(array[1])]; ok {
			month = v
		} else {
			return
		}
		year, _ = strconv.Atoi(array[2])
		day, _ = strconv.Atoi(array[0])
	}
	return
}

// StrToTime converts string to *Time object. It also supports timestamp string.
func StrToTime(str string, format ...string) (*Time, error) {
	if str == "" {
		return &Time{wrapper{time.Time{}}}, nil
	}
	if len(format) > 0 {
		return StrToTimeFormat(str, format[0])
	}
	if isTimestampStr(str) {
		timestamp, _ := strconv.ParseInt(str, 10, 64)
		return NewFromTimeStamp(timestamp), nil
	}
	var (
		year, month, day     int
		hour, min, sec, nsec int
		match                []string
		local                = time.Local
	)
	if match = timeRegex1.FindStringSubmatch(str); len(match) > 0 && match[1] != "" {
		year, month, day = parseDateStr(match[1])
	} else if match = timeRegex2.FindStringSubmatch(str); len(match) > 0 && match[1] != "" {
		year, month, day = parseDateStr(match[1])
	} else if match = timeRegex3.FindStringSubmatch(str); len(match) > 0 && match[1] != "" {
		hour, _ = strconv.Atoi(match[1])
		min, _ = strconv.Atoi(match[2])
		sec, _ = strconv.Atoi(match[3])
		nsec, _ = strconv.Atoi(match[4])
		for len(match[4]) < 9 {
			nsec *= 10
			match[4] += "0"
		}
		return NewFromTime(time.Date(0, time.Month(1), 1, hour, min, sec, nsec, local)), nil
	} else {
		return nil, errors.NewCodef(code.CodeInvalidParameter, `unsupported time converting for string "%s"`, str)
	}

	if len(match[2]) > 0 {
		hour, _ = strconv.Atoi(match[2][0:2])
		min, _ = strconv.Atoi(match[2][2:4])
		sec, _ = strconv.Atoi(match[2][4:6])
	}
	if len(match[3]) > 0 {
		nsec, _ = strconv.Atoi(match[3])
		for len(match[3]) < 9 {
			nsec *= 10
			match[3] += "0"
		}
	}
	if match[4] != "" && match[6] == "" {
		match[6] = "000000"
	}
	if match[6] != "" {
		zone := strings.ReplaceAll(match[6], ":", "")
		zone = strings.TrimLeft(zone, "+-")
		if len(zone) <= 6 {
			zone += strings.Repeat("0", 6-len(zone))
			h, _ := strconv.Atoi(zone[0:2])
			m, _ := strconv.Atoi(zone[2:4])
			s, _ := strconv.Atoi(zone[4:6])
			if h > 24 || m > 59 || s > 59 {
				return nil, errors.NewCodef(code.CodeInvalidParameter, `invalid zone string "%s"`, match[6])
			}
			operation := match[5]
			if operation != "+" && operation != "-" {
				operation = "-"
			}
			_, localOffset := time.Now().Zone()
			if (h*3600+m*60+s) != localOffset ||
				(localOffset > 0 && operation == "-") ||
				(localOffset < 0 && operation == "+") {
				local = time.UTC
				switch operation {
				case "+":
					hour -= h
					min -= m
					sec -= s
				case "-":
					hour += h
					min += m
					sec += s
				}
			}
		}
	}
	if month <= 0 || day <= 0 {
		return nil, errors.NewCodef(code.CodeInvalidParameter, `invalid time string "%s"`, str)
	}
	return NewFromTime(time.Date(year, time.Month(month), day, hour, min, sec, nsec, local)), nil
}

// ConvertZone converts time in string `strTime` from `fromZone` to `toZone`.
func ConvertZone(strTime string, toZone string, fromZone ...string) (*Time, error) {
	t, err := StrToTime(strTime)
	if err != nil {
		return nil, err
	}
	var l *time.Location
	if len(fromZone) > 0 {
		if l, err = time.LoadLocation(fromZone[0]); err != nil {
			return nil, errors.WrapCodef(code.CodeInvalidParameter, err, `time.LoadLocation failed for name "%s"`, fromZone[0])
		}
		t.Time = time.Date(t.Year(), time.Month(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), l)
	}
	if l, err = time.LoadLocation(toZone); err != nil {
		return nil, errors.WrapCodef(code.CodeInvalidParameter, err, `time.LoadLocation failed for name "%s"`, toZone)
	}
	return t.ToLocation(l), nil
}

// StrToTimeFormat parses string `str` to *Time object with given format `format`.
func StrToTimeFormat(str string, format string) (*Time, error) {
	return StrToTimeLayout(str, formatToStdLayout(format))
}

// StrToTimeLayout parses string `str` to *Time object with given format `layout`.
func StrToTimeLayout(str string, layout string) (*Time, error) {
	t, err := time.ParseInLocation(layout, str, time.Local)
	if err != nil {
		return nil, errors.WrapCodef(
			code.CodeInvalidParameter, err,
			`time.ParseInLocation failed for layout "%s" and value "%s"`,
			layout, str,
		)
	}
	return NewFromTime(t), nil
}

// ParseTimeFromContent retrieves time information for content string.
func ParseTimeFromContent(content string, format ...string) *Time {
	var (
		err   error
		match []string
	)
	if len(format) > 0 {
		for _, item := range format {
			match, err = regex.MatchString(formatToRegexPattern(item), content)
			if err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
			if len(match) > 0 {
				return NewFromStrFormat(match[0], item)
			}
		}
	} else {
		if match = timeRegex1.FindStringSubmatch(content); len(match) >= 1 {
			return NewFromStr(strings.Trim(match[0], "./_- \n\r"))
		} else if match = timeRegex2.FindStringSubmatch(content); len(match) >= 1 {
			return NewFromStr(strings.Trim(match[0], "./_- \n\r"))
		} else if match = timeRegex3.FindStringSubmatch(content); len(match) >= 1 {
			return NewFromStr(strings.Trim(match[0], "./_- \n\r"))
		}
	}
	return nil
}

// ParseDuration parses a duration string.
func ParseDuration(s string) (duration time.Duration, err error) {
	if utils.IsNumeric(s) {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, errors.WrapCodef(code.CodeInvalidParameter, err, `strconv.ParseInt failed for string "%s"`, s)
		}
		return time.Duration(num), nil
	}
	match, err := regex.MatchString(`^([\-\d]+)[dD](.*)$`, s)
	if err != nil {
		return 0, err
	}
	if len(match) == 3 {
		num, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			return 0, errors.WrapCodef(code.CodeInvalidParameter, err, `strconv.ParseInt failed for string "%s"`, match[1])
		}
		s = fmt.Sprintf(`%dh%s`, num*24, match[2])
		duration, err = time.ParseDuration(s)
		if err != nil {
			return 0, errors.WrapCodef(code.CodeInvalidParameter, err, `time.ParseDuration failed for string "%s"`, s)
		}
		return
	}
	duration, err = time.ParseDuration(s)
	return duration, errors.WrapCodef(code.CodeInvalidParameter, err, `time.ParseDuration failed for string "%s"`, s)
}

// FuncCost calculates the cost time of function `f` in nanoseconds.
func FuncCost(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
}

// isTimestampStr checks and returns whether given string a timestamp string.
func isTimestampStr(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return len(s) > 0
}

// New creates and returns a Time object with given parameter.
// The optional parameter can be type of: time.Time/*time.Time, string, or integer.
func New(param ...interface{}) *Time {
	if len(param) > 0 {
		switch r := param[0].(type) {
		case time.Time:
			return NewFromTime(r)
		case *time.Time:
			return NewFromTime(*r)
		case Time:
			return &r
		case *Time:
			return r
		case string:
			if len(param) > 1 {
				switch t := param[1].(type) {
				case string:
					return NewFromStrFormat(r, t)
				case []byte:
					return NewFromStrFormat(r, string(t))
				}
			}
			return NewFromStr(r)
		case []byte:
			return NewFromStr(string(r))
		case int:
			return NewFromTimeStamp(int64(r))
		case int64:
			return NewFromTimeStamp(r)
		default:
			if v, ok := r.(UnixNanoProvider); ok {
				return NewFromTimeStamp(v.UnixNano())
			}
		}
	}
	return &Time{
		wrapper{time.Time{}},
	}
}

// Now creates and returns a time object of now.
func Now() *Time {
	return &Time{
		wrapper{time.Now()},
	}
}

// NewFromTime creates and returns a Time object with given time.Time object.
func NewFromTime(t time.Time) *Time {
	return &Time{
		wrapper{t},
	}
}

// NewFromStr creates and returns a Time object with given string.
// Note that it returns nil if there's error occurs.
func NewFromStr(str string) *Time {
	if t, err := StrToTime(str); err == nil {
		return t
	}
	return nil
}

// NewFromStrFormat creates and returns a Time object with given string and
// custom format like: Y-m-d H:i:s.
// Note that it returns nil if there's error occurs.
func NewFromStrFormat(str string, format string) *Time {
	if t, err := StrToTimeFormat(str, format); err == nil {
		return t
	}
	return nil
}

// NewFromStrLayout creates and returns a Time object with given string and
// stdlib layout like: 2006-01-02 15:04:05.
// Note that it returns nil if there's error occurs.
func NewFromStrLayout(str string, layout string) *Time {
	if t, err := StrToTimeLayout(str, layout); err == nil {
		return t
	}
	return nil
}

// NewFromTimeStamp creates and returns a Time object with given timestamp,
// which can be in seconds to nanoseconds.
// Eg: 1600443866 and 1600443866199266000 are both considered as valid timestamp number.
func NewFromTimeStamp(timestamp int64) *Time {
	if timestamp == 0 {
		return &Time{}
	}
	if timestamp < 1e9 {
		return &Time{wrapper{time.Unix(timestamp, 0)}}
	}
	for timestamp < 1e18 {
		timestamp *= 10
	}
	sec := timestamp / 1e9
	nano := timestamp % 1e9
	return &Time{
		wrapper{time.Unix(sec, nano)},
	}
}

// Timestamp returns the timestamp in seconds.
func (t *Time) Timestamp() int64 {
	return t.UnixNano() / 1e9
}

// TimestampMilli returns the timestamp in milliseconds.
func (t *Time) TimestampMilli() int64 {
	return t.UnixNano() / 1e6
}

// TimestampMicro returns the timestamp in microseconds.
func (t *Time) TimestampMicro() int64 {
	return t.UnixNano() / 1e3
}

// TimestampNano returns the timestamp in nanoseconds.
func (t *Time) TimestampNano() int64 {
	return t.UnixNano()
}

// TimestampStr is a convenience method which retrieves and returns
// the timestamp in seconds as string.
func (t *Time) TimestampStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.Timestamp(), 10)
}

// TimestampMilliStr is a convenience method which retrieves and returns
// the timestamp in milliseconds as string.
func (t *Time) TimestampMilliStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampMilli(), 10)
}

// TimestampMicroStr is a convenience method which retrieves and returns
// the timestamp in microseconds as string.
func (t *Time) TimestampMicroStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampMicro(), 10)
}

// TimestampNanoStr is a convenience method which retrieves and returns
// the timestamp in nanoseconds as string.
func (t *Time) TimestampNanoStr() string {
	if t.IsZero() {
		return ""
	}
	return strconv.FormatInt(t.TimestampNano(), 10)
}

// Month returns the month of the year specified by t.
func (t *Time) Month() int {
	if t.IsZero() {
		return 0
	}
	return int(t.Time.Month())
}

// Second returns the second offset within the minute specified by t,
// in the range [0, 59].
func (t *Time) Second() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Second()
}

// Millisecond returns the millisecond offset within the second specified by t,
// in the range [0, 999].
func (t *Time) Millisecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e6
}

// Microsecond returns the microsecond offset within the second specified by t,
// in the range [0, 999999].
func (t *Time) Microsecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond() / 1e3
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (t *Time) Nanosecond() int {
	if t.IsZero() {
		return 0
	}
	return t.Time.Nanosecond()
}

// String returns current time object as string.
func (t *Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.wrapper.String()
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t *Time) IsZero() bool {
	if t == nil {
		return true
	}
	return t.Time.IsZero()
}

// Clone returns a new Time object which is a clone of current time object.
func (t *Time) Clone() *Time {
	return New(t.Time)
}

// Add adds the duration to current time.
func (t *Time) Add(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Add(d)
	return newTime
}

// AddStr parses the given duration as string and adds it to current time.
func (t *Time) AddStr(duration string) (*Time, error) {
	d, err := time.ParseDuration(duration)
	if err != nil {
		return nil, errors.Wrapf(err, `time.ParseDuration failed for string "%s"`, duration)
	}
	return t.Add(d), nil
}

// UTC converts current time to UTC timezone.
func (t *Time) UTC() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.UTC()
	return newTime
}

// ISO8601 formats the time as ISO8601 and returns it as string.
func (t *Time) ISO8601() string {
	return t.Layout("2006-01-02T15:04:05-07:00")
}

// RFC822 formats the time as RFC822 and returns it as string.
func (t *Time) RFC822() string {
	return t.Layout("Mon, 02 Jan 06 15:04 MST")
}

// AddDate adds year, month and day to the time.
func (t *Time) AddDate(years int, months int, days int) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.AddDate(years, months, days)
	return newTime
}

// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t *Time) Round(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Round(d)
	return newTime
}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t *Time) Truncate(d time.Duration) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Truncate(d)
	return newTime
}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 CEST and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
func (t *Time) Equal(u *Time) bool {
	switch {
	case t == nil && u != nil:
		return false
	case t == nil && u == nil:
		return true
	case t != nil && u == nil:
		return false
	default:
		return t.Time.Equal(u.Time)
	}
}

// Before reports whether the time instant t is before u.
func (t *Time) Before(u *Time) bool {
	return t.Time.Before(u.Time)
}

// After reports whether the time instant t is after u.
func (t *Time) After(u *Time) bool {
	switch {
	case t == nil:
		return false
	case t != nil && u == nil:
		return true
	default:
		return t.Time.After(u.Time)
	}
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
func (t *Time) Sub(u *Time) time.Duration {
	if t == nil || u == nil {
		return 0
	}
	return t.Time.Sub(u.Time)
}

// StartOfMinute clones and returns a new time of which the seconds is set to 0.
func (t *Time) StartOfMinute() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Truncate(time.Minute)
	return newTime
}

// StartOfHour clones and returns a new time of which the hour, minutes and seconds are set to 0.
func (t *Time) StartOfHour() *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, d, newTime.Time.Hour(), 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfDay clones and returns a new time which is the start of day, its time is set to 00:00:00.
func (t *Time) StartOfDay() *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, d, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfWeek clones and returns a new time which is the first day of week and its time is set to
// 00:00:00.
func (t *Time) StartOfWeek() *Time {
	weekday := int(t.Weekday())
	return t.StartOfDay().AddDate(0, 0, -weekday)
}

// StartOfMonth clones and returns a new time which is the first day of the month and its is set to
// 00:00:00
func (t *Time) StartOfMonth() *Time {
	y, m, _ := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, m, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// StartOfQuarter clones and returns a new time which is the first day of the quarter and its time is set
// to 00:00:00.
func (t *Time) StartOfQuarter() *Time {
	month := t.StartOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// StartOfHalf clones and returns a new time which is the first day of the half year and its time is set
// to 00:00:00.
func (t *Time) StartOfHalf() *Time {
	month := t.StartOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// StartOfYear clones and returns a new time which is the first day of the year and its time is set to
// 00:00:00.
func (t *Time) StartOfYear() *Time {
	y, _, _ := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(y, time.January, 1, 0, 0, 0, 0, newTime.Time.Location())
	return newTime
}

// getPrecisionDelta returns the precision parameter for time calculation depending on `withNanoPrecision` option.
func getPrecisionDelta(withNanoPrecision ...bool) time.Duration {
	if len(withNanoPrecision) > 0 && withNanoPrecision[0] {
		return time.Nanosecond
	}
	return time.Second
}

// EndOfMinute clones and returns a new time of which the seconds is set to 59.
func (t *Time) EndOfMinute(withNanoPrecision ...bool) *Time {
	return t.StartOfMinute().Add(time.Minute - getPrecisionDelta(withNanoPrecision...))
}

// EndOfHour clones and returns a new time of which the minutes and seconds are both set to 59.
func (t *Time) EndOfHour(withNanoPrecision ...bool) *Time {
	return t.StartOfHour().Add(time.Hour - getPrecisionDelta(withNanoPrecision...))
}

// EndOfDay clones and returns a new time which is the end of day the and its time is set to 23:59:59.
func (t *Time) EndOfDay(withNanoPrecision ...bool) *Time {
	y, m, d := t.Date()
	newTime := t.Clone()
	newTime.Time = time.Date(
		y, m, d, 23, 59, 59, int(time.Second-getPrecisionDelta(withNanoPrecision...)), newTime.Time.Location(),
	)
	return newTime
}

// EndOfWeek clones and returns a new time which is the end of week and its time is set to 23:59:59.
func (t *Time) EndOfWeek(withNanoPrecision ...bool) *Time {
	return t.StartOfWeek().AddDate(0, 0, 7).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfMonth clones and returns a new time which is the end of the month and its time is set to 23:59:59.
func (t *Time) EndOfMonth(withNanoPrecision ...bool) *Time {
	return t.StartOfMonth().AddDate(0, 1, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfQuarter clones and returns a new time which is end of the quarter and its time is set to 23:59:59.
func (t *Time) EndOfQuarter(withNanoPrecision ...bool) *Time {
	return t.StartOfQuarter().AddDate(0, 3, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfHalf clones and returns a new time which is the end of the half year and its time is set to 23:59:59.
func (t *Time) EndOfHalf(withNanoPrecision ...bool) *Time {
	return t.StartOfHalf().AddDate(0, 6, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// EndOfYear clones and returns a new time which is the end of the year and its time is set to 23:59:59.
func (t *Time) EndOfYear(withNanoPrecision ...bool) *Time {
	return t.StartOfYear().AddDate(1, 0, 0).Add(-getPrecisionDelta(withNanoPrecision...))
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note that, DO NOT use `(t *Time) MarshalJSON() ([]byte, error)` as it looses interface
// implement of `MarshalJSON` for struct of Time.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		t.Time = time.Time{}
		return nil
	}
	newTime, err := StrToTime(string(bytes.Trim(b, `"`)))
	if err != nil {
		return err
	}
	t.Time = newTime.Time
	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// Note that it overwrites the same implementer of `time.Time`.
func (t *Time) UnmarshalText(data []byte) error {
	vTime := New(data)
	if vTime != nil {
		*t = *vTime
		return nil
	}
	return errors.NewCodef(code.CodeInvalidParameter, `invalid time value: %s`, data)
}

// NoValidation marks this struct object will not be validated by package gvalid.
func (t *Time) NoValidation() {}

// DeepCopy implements interface for deep copy of current type.
func (t *Time) DeepCopy() interface{} {
	if t == nil {
		return nil
	}
	return New(t.Time)
}

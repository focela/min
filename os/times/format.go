// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package times

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/focela/aid/text/regex"
)

var (
	// Refer: http://php.net/manual/en/function.date.php
	formats = map[byte]string{
		'd': "02",
		'D': "Mon",
		'w': "Monday",
		'N': "Monday",
		'j': "=j=02",
		'S': "02",
		'l': "Monday",
		'z': "",
		'W': "",
		'F': "January",
		'm': "01",
		'M': "Jan",
		'n': "1",
		't': "",
		'Y': "2006",
		'y': "06",
		'a': "pm",
		'A': "PM",
		'g': "3",
		'G': "=G=15",
		'h': "03",
		'H': "15",
		'i': "04",
		's': "05",
		'u': "=u=.000",
		'U': "",
		'O': "-0700",
		'P': "-07:00",
		'T': "MST",
		'c': "2006-01-02T15:04:05-07:00",
		'r': "Mon, 02 Jan 06 15:04 MST",
	}

	weekMap = map[string]string{
		"Sunday":    "0",
		"Monday":    "1",
		"Tuesday":   "2",
		"Wednesday": "3",
		"Thursday":  "4",
		"Friday":    "5",
		"Saturday":  "6",
	}

	dayOfMonth = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
)

func (t *Time) Format(format string) string {
	if t == nil {
		return ""
	}
	runes := []rune(format)
	var builder strings.Builder
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '\\':
			if i < len(runes)-1 {
				builder.WriteRune(runes[i+1])
				i++
			}
		case 'W':
			builder.WriteString(strconv.Itoa(t.WeeksOfYear()))
		case 'z':
			builder.WriteString(strconv.Itoa(t.DayOfYear()))
		case 't':
			builder.WriteString(strconv.Itoa(t.DaysInMonth()))
		case 'U':
			builder.WriteString(strconv.FormatInt(t.Unix(), 10))
		default:
			if f, ok := formats[byte(runes[i])]; ok {
				result := t.Time.Format(f)
				switch runes[i] {
				case 'j':
					result = strings.ReplaceAll(result, "=j=0", "")
					result = strings.ReplaceAll(result, "=j=", "")
				case 'G':
					result = strings.ReplaceAll(result, "=G=0", "")
					result = strings.ReplaceAll(result, "=G=", "")
				case 'u':
					result = strings.ReplaceAll(result, "=u=.", "")
				case 'w':
					result = weekMap[result]
				case 'N':
					result = strings.ReplaceAll(weekMap[result], "0", "7")
				case 'S':
					result = formatMonthDaySuffixMap(result)
				}
				builder.WriteString(result)
			} else {
				builder.WriteRune(runes[i])
			}
		}
	}
	return builder.String()
}

func (t *Time) FormatNew(format string) *Time {
	if t == nil {
		return nil
	}
	return NewFromStr(t.Format(format))
}

func (t *Time) FormatTo(format string) *Time {
	if t == nil {
		return nil
	}
	t.Time = NewFromStr(t.Format(format)).Time
	return t
}

func (t *Time) Layout(layout string) string {
	if t == nil {
		return ""
	}
	return t.Time.Format(layout)
}

func (t *Time) LayoutNew(layout string) *Time {
	if t == nil {
		return nil
	}
	newTime, err := StrToTimeLayout(t.Layout(layout), layout)
	if err != nil {
		panic(err)
	}
	return newTime
}

func (t *Time) LayoutTo(layout string) *Time {
	if t == nil {
		return nil
	}
	newTime, err := StrToTimeLayout(t.Layout(layout), layout)
	if err != nil {
		panic(err)
	}
	t.Time = newTime.Time
	return t
}

func (t *Time) IsLeapYear() bool {
	year := t.Year()
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func (t *Time) DayOfYear() int {
	day := t.Day()
	month := t.Month()
	if t.IsLeapYear() {
		if month > 2 {
			return dayOfMonth[month-1] + day
		}
		return dayOfMonth[month-1] + day - 1
	}
	return dayOfMonth[month-1] + day - 1
}

func (t *Time) DaysInMonth() int {
	switch t.Month() {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	}
	if t.IsLeapYear() {
		return 29
	}
	return 28
}

func (t *Time) WeeksOfYear() int {
	_, week := t.ISOWeek()
	return week
}

func formatToStdLayout(format string) string {
	var builder strings.Builder
	for i := 0; i < len(format); i++ {
		switch format[i] {
		case '\\':
			if i < len(format)-1 {
				builder.WriteByte(format[i+1])
				i++
			}
		default:
			if f, ok := formats[format[i]]; ok {
				switch format[i] {
				case 'j':
					builder.WriteString("2")
				case 'G':
					builder.WriteString("15")
				case 'u':
					if i > 0 && format[i-1] == '.' {
						builder.WriteString("000")
					} else {
						builder.WriteString(".000")
					}
				default:
					builder.WriteString(f)
				}
			} else {
				builder.WriteByte(format[i])
			}
		}
	}
	return builder.String()
}

func formatToRegexPattern(format string) string {
	s := regexp.QuoteMeta(formatToStdLayout(format))
	s, _ = regex.ReplaceString(`[0-9]`, `[0-9]`, s)
	s, _ = regex.ReplaceString(`[A-Za-z]`, `[A-Za-z]`, s)
	s, _ = regex.ReplaceString(`\s+`, `\s+`, s)
	return s
}

func formatMonthDaySuffixMap(day string) string {
	switch day {
	case "01", "21", "31":
		return "st"
	case "02", "22":
		return "nd"
	case "03", "23":
		return "rd"
	default:
		return "th"
	}
}

// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package times

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/focela/aid/errors"
	"github.com/focela/aid/errors/code"
)

var (
	setTimeZoneMu   sync.Mutex
	setTimeZoneName string
	zoneMap         = make(map[string]*time.Location)
	zoneMu          sync.RWMutex
)

// SetTimeZone sets the time zone for the current process.
// The parameter `zone` is a string specifying the time zone, e.g., Asia/Shanghai.
//
// Notes:
// 1. This should be called before the "time" package import.
// 2. This function should be called once.
// 3. Please refer to issue: https://github.com/golang/go/issues/34814
func SetTimeZone(zone string) error {
	setTimeZoneMu.Lock()
	defer setTimeZoneMu.Unlock()

	if setTimeZoneName != "" && !strings.EqualFold(zone, setTimeZoneName) {
		return errors.NewCodef(
			code.CodeInvalidOperation,
			`process timezone already set using "%s"`,
			setTimeZoneName,
		)
	}

	location, err := time.LoadLocation(zone)
	if err != nil {
		return errors.WrapCodef(code.CodeInvalidParameter, err, `time.LoadLocation failed for zone "%s"`, zone)
	}

	time.Local = location

	if err := os.Setenv("TZ", location.String()); err != nil {
		return errors.WrapCodef(
			code.CodeUnknown,
			err,
			`set environment failed with key "%s", value "%s"`,
			"TZ", location.String(),
		)
	}

	setTimeZoneName = zone
	return nil
}

// ToLocation converts the current time to the specified location.
func (t *Time) ToLocation(location *time.Location) *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.In(location)
	return newTime
}

// ToZone converts the current time to the specified time zone, e.g., Asia/Shanghai.
func (t *Time) ToZone(zone string) (*Time, error) {
	location, err := t.getLocationByZoneName(zone)
	if err != nil {
		return nil, err
	}
	return t.ToLocation(location), nil
}

func (t *Time) getLocationByZoneName(name string) (*time.Location, error) {
	zoneMu.RLock()
	location := zoneMap[name]
	zoneMu.RUnlock()

	if location == nil {
		var err error
		location, err = time.LoadLocation(name)
		if err != nil {
			return nil, errors.Wrapf(err, `time.LoadLocation failed for name "%s"`, name)
		}
		zoneMu.Lock()
		zoneMap[name] = location
		zoneMu.Unlock()
	}
	return location, nil
}

// Local converts the time to the local time zone.
func (t *Time) Local() *Time {
	newTime := t.Clone()
	newTime.Time = newTime.Time.Local()
	return newTime
}

// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package times

import (
	"database/sql/driver"
)

// Scan implements the Scanner interface from database/sql.
// It scans a value from the database into the Time struct.
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	*t = *New(value)
	return nil
}

// Value implements the Valuer interface from database/sql/driver.
// It returns the time value to be stored in the database.
func (t *Time) Value() (driver.Value, error) {
	if t == nil || t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

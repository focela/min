// Copyright (c) 2024 Focela Technologies. All rights reserved.
//
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package minerror

import "encoding/json"

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// Note: Using json.Marshal to safely handle escaping of special characters.
func (err Error) MarshalJSON() ([]byte, error) {
	// Use json.Marshal to handle escaping and serialization
	return json.Marshal(err.Error())
}

// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package structs

import (
	"strings"

	"github.com/focela/aid/util/tags"
)

// TagJsonName returns the `json` tag name string of the field.
func (f *Field) TagJsonName() string {
	jsonTag := f.Tag(tags.Json)
	if jsonTag != "" {
		return strings.Split(jsonTag, ",")[0]
	}
	return ""
}

// TagDefault returns the most commonly used tag `default/d` value of the field.
func (f *Field) TagDefault() string {
	if v := f.Tag(tags.Default); v != "" {
		return v
	}
	return f.Tag(tags.DefaultShort)
}

// TagParam returns the most commonly used tag `param/p` value of the field.
func (f *Field) TagParam() string {
	if v := f.Tag(tags.Param); v != "" {
		return v
	}
	return f.Tag(tags.ParamShort)
}

// TagValid returns the most commonly used tag `valid/v` value of the field.
func (f *Field) TagValid() string {
	if v := f.Tag(tags.Valid); v != "" {
		return v
	}
	return f.Tag(tags.ValidShort)
}

// TagDescription returns the most commonly used tag `description/des/dc` value of the field.
func (f *Field) TagDescription() string {
	if v := f.Tag(tags.Description); v != "" {
		return v
	}
	if v := f.Tag(tags.DescriptionShort); v != "" {
		return v
	}
	return f.Tag(tags.DescriptionShort2)
}

// TagSummary returns the most commonly used tag `summary/sum/sm` value of the field.
func (f *Field) TagSummary() string {
	if v := f.Tag(tags.Summary); v != "" {
		return v
	}
	if v := f.Tag(tags.SummaryShort); v != "" {
		return v
	}
	return f.Tag(tags.SummaryShort2)
}

// TagAdditional returns the most commonly used tag `additional/ad` value of the field.
func (f *Field) TagAdditional() string {
	if v := f.Tag(tags.Additional); v != "" {
		return v
	}
	return f.Tag(tags.AdditionalShort)
}

// TagExample returns the most commonly used tag `example/eg` value of the field.
func (f *Field) TagExample() string {
	if v := f.Tag(tags.Example); v != "" {
		return v
	}
	return f.Tag(tags.ExampleShort)
}

// TagIn returns the most commonly used tag `in` value of the field.
func (f *Field) TagIn() string {
	return f.Tag(tags.In)
}

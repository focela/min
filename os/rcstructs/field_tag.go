// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcstructs

import (
	"strings"

	"github.com/focela/ratcatcher/utils/rctag"
)

// TagJsonName returns the `json` tag name string of the field.
func (f *Field) TagJsonName() string {
	if jsonTag := f.Tag(rctag.Json); jsonTag != "" {
		return strings.Split(jsonTag, ",")[0]
	}
	return ""
}

// TagDefault returns the most commonly used tag `default/d` value of the field.
func (f *Field) TagDefault() string {
	v := f.Tag(rctag.Default)
	if v == "" {
		v = f.Tag(rctag.DefaultShort)
	}
	return v
}

// TagParam returns the most commonly used tag `param/p` value of the field.
func (f *Field) TagParam() string {
	v := f.Tag(rctag.Param)
	if v == "" {
		v = f.Tag(rctag.ParamShort)
	}
	return v
}

// TagValid returns the most commonly used tag `valid/v` value of the field.
func (f *Field) TagValid() string {
	v := f.Tag(rctag.Valid)
	if v == "" {
		v = f.Tag(rctag.ValidShort)
	}
	return v
}

// TagDescription returns the most commonly used tag `description/des/dc` value of the field.
func (f *Field) TagDescription() string {
	v := f.Tag(rctag.Description)
	if v == "" {
		v = f.Tag(rctag.DescriptionShort)
	}
	if v == "" {
		v = f.Tag(rctag.DescriptionShort2)
	}
	return v
}

// TagSummary returns the most commonly used tag `summary/sum/sm` value of the field.
func (f *Field) TagSummary() string {
	v := f.Tag(rctag.Summary)
	if v == "" {
		v = f.Tag(rctag.SummaryShort)
	}
	if v == "" {
		v = f.Tag(rctag.SummaryShort2)
	}
	return v
}

// TagAdditional returns the most commonly used tag `additional/ad` value of the field.
func (f *Field) TagAdditional() string {
	v := f.Tag(rctag.Additional)
	if v == "" {
		v = f.Tag(rctag.AdditionalShort)
	}
	return v
}

// TagExample returns the most commonly used tag `example/eg` value of the field.
func (f *Field) TagExample() string {
	v := f.Tag(rctag.Example)
	if v == "" {
		v = f.Tag(rctag.ExampleShort)
	}
	return v
}

// TagIn returns the most commonly used tag `in` value of the field.
func (f *Field) TagIn() string {
	v := f.Tag(rctag.In)
	return v
}

// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rstructs

import (
	"strings"

	"github.com/focela/ratcatcher/utils/rtag"
)

// TagJsonName returns the `json` tag name string of the field.
func (f *Field) TagJsonName() string {
	if jsonTag := f.Tag(rtag.Json); jsonTag != "" {
		return strings.Split(jsonTag, ",")[0]
	}
	return ""
}

// TagDefault returns the most commonly used tag `default/d` value of the field.
func (f *Field) TagDefault() string {
	v := f.Tag(rtag.Default)
	if v == "" {
		v = f.Tag(rtag.DefaultShort)
	}
	return v
}

// TagParam returns the most commonly used tag `param/p` value of the field.
func (f *Field) TagParam() string {
	v := f.Tag(rtag.Param)
	if v == "" {
		v = f.Tag(rtag.ParamShort)
	}
	return v
}

// TagValid returns the most commonly used tag `valid/v` value of the field.
func (f *Field) TagValid() string {
	v := f.Tag(rtag.Valid)
	if v == "" {
		v = f.Tag(rtag.ValidShort)
	}
	return v
}

// TagDescription returns the most commonly used tag `description/des/dc` value of the field.
func (f *Field) TagDescription() string {
	v := f.Tag(rtag.Description)
	if v == "" {
		v = f.Tag(rtag.DescriptionShort)
	}
	if v == "" {
		v = f.Tag(rtag.DescriptionShort2)
	}
	return v
}

// TagSummary returns the most commonly used tag `summary/sum/sm` value of the field.
func (f *Field) TagSummary() string {
	v := f.Tag(rtag.Summary)
	if v == "" {
		v = f.Tag(rtag.SummaryShort)
	}
	if v == "" {
		v = f.Tag(rtag.SummaryShort2)
	}
	return v
}

// TagAdditional returns the most commonly used tag `additional/ad` value of the field.
func (f *Field) TagAdditional() string {
	v := f.Tag(rtag.Additional)
	if v == "" {
		v = f.Tag(rtag.AdditionalShort)
	}
	return v
}

// TagExample returns the most commonly used tag `example/eg` value of the field.
func (f *Field) TagExample() string {
	v := f.Tag(rtag.Example)
	if v == "" {
		v = f.Tag(rtag.ExampleShort)
	}
	return v
}

// TagIn returns the most commonly used tag `in` value of the field.
func (f *Field) TagIn() string {
	v := f.Tag(rtag.In)
	return v
}

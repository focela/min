// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package structs

import (
	"reflect"
	"strconv"

	"github.com/focela/aid/errors"
	"github.com/focela/aid/errors/code"
	"github.com/focela/aid/util/tags"
)

// ParseTag parses a tag string into a map.
// For example:
// ParseTag(`v:"required" p:"id" d:"1"`) => map[v:required p:id d:1].
func ParseTag(tag string) map[string]string {
	data := make(map[string]string)
	for tag != "" {
		// Skip leading space.
		tag = trimLeadingSpace(tag)
		if tag == "" {
			break
		}

		// Scan to colon.
		key, remainingTag := parseTagKey(tag)
		if key == "" || remainingTag == "" {
			break
		}
		tag = remainingTag

		// Scan quoted string to find value.
		value, remainingTag, err := parseTagValue(tag)
		if err != nil {
			panic(errors.WrapCodef(code.CodeInvalidParameter, err, `error parsing tag "%s"`, tag))
		}
		data[key] = tags.Parse(value)
		tag = remainingTag
	}
	return data
}

func trimLeadingSpace(s string) string {
	for i := 0; i < len(s) && s[i] == ' '; i++ {
		s = s[i+1:]
	}
	return s
}

func parseTagKey(tag string) (string, string) {
	i := 0
	for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
		i++
	}
	if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
		return "", ""
	}
	return tag[:i], tag[i+1:]
}

func parseTagValue(tag string) (string, string, error) {
	i := 1
	for i < len(tag) && tag[i] != '"' {
		if tag[i] == '\\' {
			i++
		}
		i++
	}
	if i >= len(tag) {
		return "", "", errors.New("invalid tag format")
	}
	quotedValue := tag[:i+1]
	remainingTag := tag[i+1:]
	value, err := strconv.Unquote(quotedValue)
	if err != nil {
		return "", "", err
	}
	return value, remainingTag, nil
}

// TagFields retrieves and returns struct tags as []Field from `pointer`.
func TagFields(pointer interface{}, priority []string) ([]Field, error) {
	return getFieldValuesByTagPriority(pointer, priority, make(map[string]struct{}))
}

// TagMapName retrieves and returns struct tags as map[tag]attribute from `pointer`.
func TagMapName(pointer interface{}, priority []string) (map[string]string, error) {
	fields, err := TagFields(pointer, priority)
	if err != nil {
		return nil, err
	}
	tagMap := make(map[string]string, len(fields))
	for _, field := range fields {
		tagMap[field.TagValue] = field.Name()
	}
	return tagMap, nil
}

// TagMapField retrieves struct tags as map[tag]Field from `pointer`.
func TagMapField(object interface{}, priority []string) (map[string]Field, error) {
	fields, err := TagFields(object, priority)
	if err != nil {
		return nil, err
	}
	tagMap := make(map[string]Field, len(fields))
	for _, field := range fields {
		tagMap[field.TagValue] = field
	}
	return tagMap, nil
}

func getFieldValues(structObject interface{}) ([]Field, error) {
	reflectValue := reflect.ValueOf(structObject)
	reflectKind := reflectValue.Kind()

	for {
		switch reflectKind {
		case reflect.Ptr, reflect.Array, reflect.Slice:
			reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
			reflectKind = reflectValue.Kind()
		default:
			if reflectKind != reflect.Struct {
				return nil, errors.NewCode(
					code.CodeInvalidParameter,
					"given value should be either type of struct/*struct/[]struct/[]*struct",
				)
			}
			return extractFields(reflectValue), nil
		}
	}
}

func extractFields(reflectValue reflect.Value) []Field {
	structType := reflectValue.Type()
	length := reflectValue.NumField()
	fields := make([]Field, length)
	for i := 0; i < length; i++ {
		fields[i] = Field{
			Value: reflectValue.Field(i),
			Field: structType.Field(i),
		}
	}
	return fields
}

func getFieldValuesByTagPriority(
	pointer interface{}, priority []string, repeatedTagFilteringMap map[string]struct{},
) ([]Field, error) {
	fields, err := getFieldValues(pointer)
	if err != nil {
		return nil, err
	}
	tagFields := make([]Field, 0)
	for _, field := range fields {
		if !field.IsExported() {
			continue
		}
		tagValue := getPriorityTagValue(field, priority)
		if tagValue != "" && !isRepeatedTag(tagValue, repeatedTagFilteringMap) {
			tagField := field
			tagField.TagName = priority[0] // Assuming first priority as tag name
			tagField.TagValue = tagValue
			tagFields = append(tagFields, tagField)
		}
		if field.IsEmbedded() && field.OriginalKind() == reflect.Struct {
			subTagFields, err := getFieldValuesByTagPriority(field.Value, priority, repeatedTagFilteringMap)
			if err != nil {
				return nil, err
			}
			tagFields = append(tagFields, subTagFields...)
		}
	}
	return tagFields, nil
}

func getPriorityTagValue(field Field, priority []string) string {
	for _, p := range priority {
		if tagValue := field.Tag(p); tagValue != "" && tagValue != "-" {
			return tagValue
		}
	}
	return ""
}

func isRepeatedTag(tagValue string, repeatedTagFilteringMap map[string]struct{}) bool {
	if _, exists := repeatedTagFilteringMap[tagValue]; exists {
		return true
	}
	repeatedTagFilteringMap[tagValue] = struct{}{}
	return false
}

// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package structs provides functions for struct information retrieving.
package structs

import (
	"reflect"

	"github.com/focela/aid/errors"
)

// Type wraps reflect.Type for additional features.
type Type struct {
	reflect.Type
}

// Field contains information of a struct field.
type Field struct {
	Value reflect.Value       // The underlying value of the field.
	Field reflect.StructField // The underlying field of the field.

	// Retrieved tag name. It depends on TagValue.
	TagName string

	// Retrieved tag value.
	// There might be more than one tag in the field,
	// but only one can be retrieved according to calling function rules.
	TagValue string
}

// FieldsInput is the input parameter struct type for the function Fields.
type FieldsInput struct {
	// Pointer should be of type struct/*struct.
	Pointer interface{}

	// RecursiveOption specifies the way to retrieve the fields recursively if the attribute
	// is an embedded struct. It is RecursiveOptionNone by default.
	RecursiveOption RecursiveOption
}

// FieldMapInput is the input parameter struct type for the function FieldMap.
type FieldMapInput struct {
	// Pointer should be of type struct/*struct.
	Pointer interface{}

	// PriorityTagArray specifies the priority tag array for retrieving from high to low.
	// If it's given `nil`, it returns map[name]Field, where `name` is the attribute name.
	PriorityTagArray []string

	// RecursiveOption specifies the way to retrieve the fields recursively if the attribute
	// is an embedded struct. It is RecursiveOptionNone by default.
	RecursiveOption RecursiveOption
}

type RecursiveOption int

const (
	RecursiveOptionNone          RecursiveOption = 0 // No recursive retrieval of fields as a map if the field is an embedded struct.
	RecursiveOptionEmbedded      RecursiveOption = 1 // Recursively retrieve fields as a map if the field is an embedded struct.
	RecursiveOptionEmbeddedNoTag RecursiveOption = 2 // Recursively retrieve fields as a map if the field is an embedded struct and the field has no tag.
)

// Fields retrieves and returns the fields of `pointer` as a slice.
func Fields(in FieldsInput) ([]Field, error) {
	var (
		fieldFilterMap       = make(map[string]struct{})
		retrievedFields      = make([]Field, 0)
		currentLevelFieldMap = make(map[string]Field)
		rangeFields, err     = getFieldValues(in.Pointer)
	)
	if err != nil {
		return nil, err
	}

	for _, field := range rangeFields {
		currentLevelFieldMap[field.Name()] = field
	}

	for _, field := range rangeFields {
		if _, ok := fieldFilterMap[field.Name()]; ok {
			continue
		}
		if field.IsEmbedded() && in.RecursiveOption != RecursiveOptionNone {
			switch in.RecursiveOption {
			case RecursiveOptionEmbeddedNoTag:
				if field.TagStr() != "" {
					break
				}
				fallthrough
			case RecursiveOptionEmbedded:
				structFields, err := Fields(FieldsInput{
					Pointer:         field.Value,
					RecursiveOption: in.RecursiveOption,
				})
				if err != nil {
					return nil, err
				}
				// The current level fields can overwrite the sub-struct fields with the same name.
				for _, structField := range structFields {
					fieldName := structField.Name()
					if _, ok := fieldFilterMap[fieldName]; ok {
						continue
					}
					fieldFilterMap[fieldName] = struct{}{}
					if v, ok := currentLevelFieldMap[fieldName]; !ok {
						retrievedFields = append(retrievedFields, structField)
					} else {
						retrievedFields = append(retrievedFields, v)
					}
				}
				continue
			}
		}
		fieldFilterMap[field.Name()] = struct{}{}
		retrievedFields = append(retrievedFields, field)
	}
	return retrievedFields, nil
}

// FieldMap retrieves and returns struct fields as map[name/tag]Field from `pointer`.
func FieldMap(in FieldMapInput) (map[string]Field, error) {
	fields, err := getFieldValues(in.Pointer)
	if err != nil {
		return nil, err
	}
	mapField := make(map[string]Field)
	for _, field := range fields {
		if !field.IsExported() {
			continue
		}
		tagValue := ""
		for _, p := range in.PriorityTagArray {
			tagValue = field.Tag(p)
			if tagValue != "" && tagValue != "-" {
				break
			}
		}
		tempField := field
		tempField.TagValue = tagValue
		if tagValue != "" {
			mapField[tagValue] = tempField
		} else {
			if in.RecursiveOption != RecursiveOptionNone && field.IsEmbedded() {
				switch in.RecursiveOption {
				case RecursiveOptionEmbeddedNoTag:
					if field.TagStr() != "" {
						mapField[field.Name()] = tempField
						break
					}
					fallthrough
				case RecursiveOptionEmbedded:
					m, err := FieldMap(FieldMapInput{
						Pointer:          field.Value,
						PriorityTagArray: in.PriorityTagArray,
						RecursiveOption:  in.RecursiveOption,
					})
					if err != nil {
						return nil, err
					}
					for k, v := range m {
						if _, ok := mapField[k]; !ok {
							mapField[k] = v
						}
					}
				}
			} else {
				mapField[field.Name()] = tempField
			}
		}
	}
	return mapField, nil
}

// StructType retrieves and returns the struct Type of the specified struct/*struct.
// The parameter `object` should be either of type struct/*struct/[]struct/[]*struct.
func StructType(object interface{}) (*Type, error) {
	var (
		reflectValue reflect.Value
		reflectKind  reflect.Kind
	)
	if rv, ok := object.(reflect.Value); ok {
		reflectValue = rv
	} else {
		reflectValue = reflect.ValueOf(object)
	}
	for {
		switch reflectValue.Kind() {
		case reflect.Ptr:
			if !reflectValue.IsValid() || reflectValue.IsNil() {
				reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
			} else {
				reflectValue = reflectValue.Elem()
			}
		case reflect.Array, reflect.Slice:
			reflectValue = reflect.New(reflectValue.Type().Elem()).Elem()
		default:
			goto exitLoop
		}
	}

exitLoop:
	if reflectValue.Kind() != reflect.Struct {
		return nil, errors.Newf(
			`invalid object kind "%s", kind of "struct" is required`,
			reflectValue.Kind(),
		)
	}
	return &Type{
		Type: reflectValue.Type(),
	}, nil
}

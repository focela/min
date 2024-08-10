// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

// Package tags provides tag content storing for struct.
//
// Note that calling functions of this package is not concurrently safe,
// which means you cannot call them in runtime but in boot procedure.
package tags

const (
	Default           = "default"      // Default is the tag used to specify default values for struct fields when receiving parameters from an HTTP request.
	DefaultShort      = "d"            // Short name of Default.
	Param             = "param"        // Parameter name for converting certain parameter to specified struct field.
	ParamShort        = "p"            // Short name of Param.
	Valid             = "valid"        // Validation rule tag for struct of field.
	ValidShort        = "v"            // Short name of Valid.
	NoValidation      = "nv"           // NoValidation is used to skip validation for specified structs or fields.
	ORM               = "orm"          // ORM tag for ORM feature, which performs different features according scenarios.
	Arg               = "arg"          // Arg tag for struct, usually for command argument option.
	Brief             = "brief"        // Brief tag for struct, usually be considered as summary.
	Root              = "root"         // Root tag for struct, usually for nested commands management.
	Additional        = "additional"   // Additional tag for struct, usually for additional description of command.
	AdditionalShort   = "ad"           // Short name of Additional.
	Path              = `path`         // Route path for HTTP request.
	Method            = `method`       // Route method for HTTP request.
	Domain            = `domain`       // Route domain for HTTP request.
	Mime              = `mime`         // MIME type for HTTP request/response.
	Consumes          = `consumes`     // MIME type for HTTP request.
	Summary           = `summary`      // Summary for struct, usually for OpenAPI in request struct.
	SummaryShort      = `sm`           // Short name of Summary.
	SummaryShort2     = `sum`          // Short name of Summary.
	Description       = `description`  // Description for struct, usually for OpenAPI in request struct.
	DescriptionShort  = `dc`           // Short name of Description.
	DescriptionShort2 = `des`          // Short name of Description.
	Example           = `example`      // Example for struct, usually for OpenAPI in request struct.
	ExampleShort      = `eg`           // Short name of Example.
	Examples          = `examples`     // Examples for struct, usually for OpenAPI in request struct.
	ExamplesShort     = `egs`          // Short name of Examples.
	ExternalDocs      = `externalDocs` // ExternalDocs for struct, always for OpenAPI in request struct.
	ExternalDocsShort = `ed`           // Short name of ExternalDocs.
	Conv              = "conv"         // GConv defines the converting target name for specified struct field.
	ConvShort         = "c"            // GConv defines the converting target name for specified struct field.
	Json              = "json"         // Json tag is supported by stdlib.
	Security          = "security"     // Security defines scheme for authentication. Detail to see https://swagger.io/docs/specification/authentication/
	In                = "in"           // Swagger distinguishes between the following parameter types based on the parameter location. Detail to see https://swagger.io/docs/specification/describing-parameters/
)

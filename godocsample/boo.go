package godocsample

import (
	"fmt"
)

// Decode assigns property values to exported fields of a struct.
//
// Decode traverses v recursively and returns an error if a value cannot be
// converted to the field type or a required value is missing for a field.
//
// The following type dependent decodings are used:
//
// String, boolean, numeric fields have the value of the property key assigned.
// The property key name is the name of the field. A different key and a default
// value can be set in the field's tag. Fields without default value are
// required. If the value cannot be converted to the field type an error is
// returned.
//
// time.Duration fields have the result of time.ParseDuration() assigned.
//
// time.Time fields have the vaule of time.Parse() assigned. The default layout
// is time.RFC3339 but can be set in the field's tag.
//
// Arrays and slices of string, boolean, numeric, time.Duration and time.Time
// fields have the value interpreted as a comma separated list of values. The
// individual values are trimmed of whitespace and empty values are ignored. A
// default value can be provided as a semicolon separated list in the field's
// tag.
//
// Struct fields are decoded recursively using the field name plus "." as
// prefix. The prefix (without dot) can be overridden in the field's tag.
// Default values are not supported in the field's tag. Specify them on the
// fields of the inner struct instead.
//
// Map fields must have a key of type string and are decoded recursively by
// using the field's name plus ".' as prefix and the next element of the key
// name as map key. The prefix (without dot) can be overridden in the field's
// tag. Default values are not supported.
//
// Examples:
//
//     // Field is ignored.
//     Field int `properties:"-"`
//
//     // Field is assigned value of 'Field'.
//     Field int
//
//     // Field is assigned value of 'myName'.
//     Field int `properties:"myName"`
//
//     // Field is assigned value of key 'myName' and has a default
//     // value 15 if the key does not exist.
//     Field int `properties:"myName,default=15"`
//
//     // Field is assigned value of key 'Field' and has a default
//     // value 15 if the key does not exist.
//     Field int `properties:",default=15"`
//
//     // Field is assigned value of key 'date' and the date
//     // is in format 2006-01-02
//     Field time.Time `properties:"date,layout=2006-01-02"`
//
//     // Field is assigned the non-empty and whitespace trimmed
//     // values of key 'Field' split by commas.
//     Field []string
//
//     // Field is assigned the non-empty and whitespace trimmed
//     // values of key 'Field' split by commas and has a default
//     // value ["a", "b", "c"] if the key does not exist.
//     Field []string `properties:",default=a;b;c"`
//
//     // Field is decoded recursively with "Field." as key prefix.
//     Field SomeStruct
//
//     // Field is decoded recursively with "myName." as key prefix.
//     Field SomeStruct `properties:"myName"`
//
//     // Field is decoded recursively with "Field." as key prefix
//     // and the next dotted element of the key as map key.
//     Field map[string]string
//
//     // Field is decoded recursively with "myName." as key prefix
//     // and the next dotted element of the key as map key.
//     Field map[string]string `properties:"myName"`
func Boo() {
	fmt.Printf("hello, world\n")
}

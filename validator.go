package mc_json_validation

import (
	"errors"
	"reflect"
)

var (
	FieldCanNotBeEmpty = " cannot be empty"
)

// Validator - gets values and types of the interface
func Validator(i interface{}) error {
	// Get interface field type
	t := reflect.TypeOf(i)

	// Get interface field value
	v := reflect.ValueOf(i)

	// Iterate over fields for validation (x is the index)
	for x := 0; x < t.NumField() /* could also use v.NumField() */ ; x++ {
		err := validation(t, v, x)
		if err != nil {
			return err
		}
	}

	return nil
}

func validation(t reflect.Type, v reflect.Value, x int) error {
	val := v.Field(x).Interface()

	// Check field type
	switch t.Field(x).Type.Name() {

	// Validation for string fields
	case "string":
		if val == "" {
			return errors.New(t.Field(x).Name + FieldCanNotBeEmpty)
		}

	// Validation for numeric fields
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		if val == 0 {
			return errors.New(t.Field(x).Name + FieldCanNotBeEmpty)
		}
	}

	return nil
}

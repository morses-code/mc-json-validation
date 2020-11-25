package mc_json_validation

import (
	"errors"
	"reflect"
)

var (
	IsNotOfKindStruct     = "is not of kind struct"
	FieldCanNotBeEmpty    = " cannot be empty"
	UnableToDetermineType = "unable to determine type"
)

type Validator struct {
	Fields map[string]bool
}

// Validator - gets values and types of the interface (needs to be of struct kind)
func (n *Validator) Validate(i interface{}) error {
	// Get interface field type
	t := reflect.TypeOf(i)

	// Get interface field value
	v := reflect.ValueOf(i)

	// Check that i is a struct
	if t.Kind() != reflect.Struct {
		return errors.New(IsNotOfKindStruct)
	}

	// Iterate over fields for validation (x is the index)
	for x := 0; x < t.NumField() /* could also use v.NumField() */ ; x++ {
		err := n.validation(t, v, x)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n Validator) validation(t reflect.Type, v reflect.Value, x int) error {
	val := v.Field(x).Interface()

	for key, value := range n.Fields {
		if t.Field(x).Name == key && value {
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

			// Default return, when type doesn't match any other case.
			default:
				return errors.New(UnableToDetermineType)
			}
		}
	}



	return nil
}

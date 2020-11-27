package validator

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	IsNotOfKindStruct     = "is not of kind struct"
	UnableToDetermineType = "unable to determine type"
)

type Validator struct {
	Fields map[string]map[interface{}]string
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

	err := n.fieldIterator(t, x, val)
	if err != nil {
		return err
	}

	return nil
}

func (n Validator) fieldIterator(t reflect.Type, x int, val interface{}) error {

	// field represents the field in the struct and rules is the map for the validation logic
	for field, rules := range n.Fields {
		validate, err := strconv.ParseBool(t.Field(x).Tag.Get("validate"))
		if err != nil {
			return err
		}
		if t.Field(x).Name == field && validate {
			err := n.ruleIterator(t, x, val, rules)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (n Validator) ruleIterator(t reflect.Type, x int, val interface{}, rules map[interface{}]string) error {

	// rule is the value that the field should not be and the message is the message provided for the error
	for rule, message := range rules {

		// Check field type
		switch t.Field(x).Type.Name() {

		// Validation for string fields
		case "string":
			if val == rule {
				return errors.New(t.Field(x).Name + " " + message)
			}

		// Validation for numeric fields
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
			if val == rule {
				return errors.New(t.Field(x).Name + " " + message)
			}

		// Default return, when type doesn't match any other case.
		default:
			return errors.New(UnableToDetermineType)
		}
	}

	return nil
}

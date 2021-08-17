package sample

import (
	"log"

	. "github.com/morses-code/mc-json-validation"
)

type sampleStruct struct {
	Name string `json:"name" validate:"true"`
	ID   int	`json:"id" validate:"true"`
}

func sampleFunc() {

	s := newSampleStruct()

	v := newSampleValidator()

	err := v.Validate(s)

	if err != nil {
		log.Print(err)
	}

}

func newSampleStruct() sampleStruct {
	return sampleStruct{
		Name: "Test",
		ID:   123,
	}
}

func newSampleValidator() Validator {
	return Validator{Fields: map[string]map[interface{}]string{
		"Name": {
			"": "can't be empty",
		},
		"ID": {
			0: "can't be zero",
		},
	}}
}

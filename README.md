# mc-json-validation

### How to use

`$ go get github.com/morses-code/mc-json-validation`

Sample struct

```
type club struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
```

Sample code

```
type sampleStruct struct {
	Name string `json:"name" validate:"true"`
	ID   int	`json:"id" validate:"true"`
}

func main() {

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
```

### For the future

The library will allow the developer to define custom validation rules and pass them to the validator.
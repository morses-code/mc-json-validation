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
func sample() {
    c := club{
    	ID:   1,
    	Name: "Test",
    }

    v := Validator{}

    err := mc-json-validation.v.Validator(c)

    if err != nil {
        // TODO - handle the error
    }
}
```

### For the future

The library will allow the developer to define custom validation rules and pass them to the validator.
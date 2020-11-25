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

    // key value represents if the field should be validated --- this is due for a rework but for now works.
    m := map[string]map[bool]map[interface{}]string {
    	"Name": {
    		true: {
    			"": "can not be empty",
    		},
    	},
    	"ID": {
    		true: {
    			0: "can not be 0",
    		},
   		},
   		"IsActive": {
   			true: {
   				false: "can not be false",
   			},
   		},
   	}

    v := Validator{
        Fields: m,
    }

    err := mc-json-validation.v.Validator(c)

    if err != nil {
        // TODO - handle the error
    }
}
```

### For the future

The library will allow the developer to define custom validation rules and pass them to the validator.
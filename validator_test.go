package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidator_Validate(t *testing.T) {

	// Dummy fields map
	m := map[string]map[interface{}]string{
		"Name": {
			"": "can not be empty",
		},
		"ID": {
			0: "can not be 0",
		},
		"IsActive": {
			false: "can not be false",
		},
	}

	// Dummy validator
	v := Validator{
		Fields: m,
	}

	// Dummy fields map 2
	m2 := map[string]map[interface{}]string{
		"Name": {
			"": "can not be empty",
		},
		"ID": {
			0: "can not be 0",
		},
		"IsActive": {
			false: "can not be false",
		},
	}

	// Dummy validator 2
	v2 := Validator{
		Fields: m2,
	}

	// Dummy fields map 2
	m3 := map[string]map[interface{}]string{
		"Name": {
			"can not be this": "can not be 'can not be this'",
		},
		"ID": {
			0: "can not be 0",
		},
		"IsActive": {
			false: "can not be false",
		},
	}

	// Dummy validator 2
	v3 := Validator{
		Fields: m3,
	}

	// Dummy struct
	s := struct {
		Name string `validate:"true"`
		ID   int    `validate:"true"`
	}{}

	// Dummy struct 2
	s2 := struct {
		Name     string `validate:"true"`
		ID       int    `validate:"true"`
		IsActive bool   `validate:"true"`
	}{}

	// Dummy struct 3
	s3 := struct {
		Name string `validate:"false"`
		ID   int    `validate:"true"`
	}{}

	// Dummy struct 4
	s4 := struct {
		Name string `validate:"not a bool"`
		ID   int    `validate:"true"`
	}{}

	Convey("Given a valid struct the  validator will return no errors", t, func() {

		s.Name = "Test"
		s.ID = 123

		Convey("When Reflector is called passing the interface", func() {
			err := v.Validate(s)

			Convey("Then no errors are returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a valid struct with missing 'Name' field the validator will return an error", t, func() {

		s.Name = ""
		s.ID = 123

		Convey("When Reflector is called passing the interface", func() {
			err := v.Validate(s)

			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Name can not be empty")
			})
		})
	})

	Convey("Given a valid struct with missing 'ID' field the validator will return an error", t, func() {

		s.Name = "Test"
		s.ID = 0

		Convey("Validate Validator is called passing the interface", func() {
			err := v.Validate(s)

			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "ID can not be 0")
			})
		})
	})

	Convey("Given a valid struct a type that isn't in the validator the validator will return an error", t, func() {

		s2.Name = "Test"
		s2.ID = 123
		s2.IsActive = true

		Convey("Validate Validator is called passing the interface", func() {
			err := v.Validate(s2)

			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "unable to determine type")
			})
		})
	})

	Convey("Given an invalid type the validator will return an error", t, func() {

		i := "not a struct"

		Convey("Validate Validator is called passing the interface", func() {
			err := v.Validate(i)

			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "is not of kind struct")
			})
		})
	})

	Convey("Given a valid struct and the Name field is empty but flagged to skip validation", t, func() {

		s3.Name = ""
		s3.ID = 123

		Convey("When Reflector is called passing the interface", func() {
			err := v2.Validate(s3)

			Convey("Then no errors are returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a valid struct and the Name field has incorrect tag value for validate", t, func() {

		s4.Name = "Test"
		s4.ID = 123

		Convey("When Reflector is called passing the interface", func() {
			err := v2.Validate(s4)

			Convey("Then no errors are returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a valid struct and the Name field matches custom validation then validator returns an error", t, func() {

		s.Name = "can not be this"
		s.ID = 123

		Convey("When Reflector is called passing the interface", func() {
			err := v3.Validate(s)

			Convey("Then no errors are returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Viper should return correct config", t, func() {

		Convey("When Config.Get() is called", func() {
			cfg, err := Get()

			Convey("Then no errors are returned", func() {
				So(err, ShouldBeNil)
				So(cfg.Password, ShouldEqual, "12341234")
			})
		})
	})

	Convey("Viper should return correct config", t, func() {

		Convey("When Config.Get() is called", func() {
			cfg = &Config{
				Password:     "dmV2KZJ3mq#eZE4xm^GKMVbYVWASad73",
				AwsAuthToken: "#^n@JnmcyzwS91B%$!d2Wb#CVnZt8D3L",
			}

			cfg, err := Get()

			Convey("Then no errors are returned", func() {
				So(err, ShouldBeNil)
				So(cfg.Password, ShouldEqual, "dmV2KZJ3mq#eZE4xm^GKMVbYVWASad73")
			})
		})
	})
}

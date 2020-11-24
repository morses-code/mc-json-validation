package mc_json_validation

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReflector(t *testing.T) {

	// Dummy struct
	s := struct {
		Name string
		ID int
	}{}

	Convey("Given a valid struct the  validator will return no errors", t, func() {

		s.Name = "Test"
		s.ID = 123

		Convey("When Reflector is called passing the interface", func() {
			err := Validator(s)

			Convey("Then no errors are returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a valid struct with missing 'Name' field the validator will return an error", t, func() {

		s.Name = ""
		s.ID = 123

		Convey("When Reflector is called passing the interface", func() {
			err := Validator(s)

			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a valid struct with missing 'ID' field the validator will return an error", t, func() {

		s.Name = ""
		s.ID = 123

		Convey("When Validator is called passing the interface", func() {
			err := Validator(s)

			Convey("Then an error is returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

}

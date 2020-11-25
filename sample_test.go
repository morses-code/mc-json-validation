package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSampleValidator(t *testing.T) {

	s := newSampleStruct()

	v := newSampleValidator()

	Convey("Sample should not return an error", t, func() {
		err := v.Validate(s)
		So(err, ShouldBeNil)
	})

}
package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCanPlay(t *testing.T) {
	Convey("CanPlay false", t, func() {
		lat := 49.282223
		long := -123.116580
		canplay, dist := CanPlayWithDistance(&lat, &long, 49.281040, -123.114509)
		So(canplay, ShouldBeFalse)
		So(dist, ShouldAlmostEqual, 200.0, 2.0)
	})

	Convey("CanPlay true", t, func() {
		lat := 49.282497
		long := -123.112477
		canplay, dist := CanPlayWithDistance(&lat, &long, 49.2823396, -123.112318)
		So(canplay, ShouldBeTrue)
		So(dist, ShouldAlmostEqual, 16.0, 5.0)
	})
}

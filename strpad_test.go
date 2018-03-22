package pad

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLeftPad(t *testing.T) {
	Convey("test left padding", t, func() {
		inputStr := "abc"

		Convey("pad length has equal length with inputstr", func() {
			So(Left(inputStr, "*", 3), ShouldEqual, inputStr)
		})

		Convey("pad length less than input str", func() {
			So(Left(inputStr, "*", 2), ShouldEqual, inputStr)
		})

		Convey("pad length more than input str", func() {
			So(Left(inputStr, "*", 5), ShouldEqual, "**"+inputStr)
		})

		Convey("pad length more than input str and pad str more than one char", func() {
			So(Left(inputStr, "**", 6), ShouldEqual, "***"+inputStr)
		})
	})
}

func TestRightPad(t *testing.T) {
	Convey("test right padding", t, func() {
		inputStr := "abc"

		Convey("pad length has equal length with inputstr", func() {
			So(Right(inputStr, "*", 3), ShouldEqual, inputStr)
		})

		Convey("pad length less than input str", func() {
			So(Right(inputStr, "*", 2), ShouldEqual, inputStr)
		})

		Convey("pad length more than input str", func() {
			So(Right(inputStr, "*", 5), ShouldEqual, inputStr+"**")
		})

		Convey("pad length more than input str and pad str more than one char", func() {
			So(Right(inputStr, "**", 6), ShouldEqual, inputStr+"***")
		})
	})
}

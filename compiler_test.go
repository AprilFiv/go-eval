package eval

import (
	"context"
	"errors"
	"testing"
)

func TestDemo(t *testing.T) {

	t.Run("AddMinus", func(t *testing.T) {

		expressions := map[string]string{
			"TestAdd1":     "A + 1",
			"TestAdd2":     "A + 1 + B",
			"TestMinus1":   "C - D",
			"TestMinus2":   "C - D - E",
			"TestMinus3":   "C - D - A",
			"TestAddMinus": "A + E - E - C + D + C - ( A- D )",
		}

		descriptor, e := Parse(context.Background(), expressions)
		if e != nil {
			panic(e)
		}

		c := descriptor.Context()

		descriptor.Run(context.Background(), c, "A", float64(10), nil)
		if c.Output["TestAdd1"].Value != float64(11) {
			panic("1")
		}
		if c.Output["TestAdd2"] != nil {
			panic("2")
		}

		descriptor.Run(context.Background(), c, "B", float64(10), errors.New("test error"))
		if c.Output["TestAdd2"].Err == nil {
			panic("3")
		}

		descriptor.Run(context.Background(), c, "C", int64(10), nil)
		if c.Output["TestMinus1"] != nil {
			panic(4)
		}

		descriptor.Run(context.Background(), c, "D", int64(200), nil)
		if c.Output["TestMinus1"].Value != int64(-190) {
			panic(4)
		}
		if c.Output["TestMinus3"].Value != float64(-200) {
			panic(5)
		}
		if c.Output["TestMinus2"] != nil {
			panic(6)
		}
		descriptor.Run(context.Background(), c, "E", int64(5), nil)
		if c.Output["TestMinus2"].Value != int64(-195) {
			panic(7)
		}
		if c.Output["TestAddMinus"].Value != float64(400) {
			panic(8)
		}
	})

	t.Run("Unary", func(t *testing.T) {
		expressions := map[string]string{
			"Test1": "!A",
			"Test2": "-B",
			"Test3": "-B+C",
			"Test4": "-(B+C)",
		}

		descriptor, e := Parse(context.Background(), expressions)
		if e != nil {
			panic(e)
		}

		c := descriptor.Context()

		descriptor.Run(context.Background(), c, "A", true, nil)
		if c.Output["Test1"].Value != false {
			panic("1")
		}

		descriptor.Run(context.Background(), c, "B", int64(10), nil)
		if c.Output["Test2"].Value != int64(-10) {
			panic("2")
		}
		descriptor.Run(context.Background(), c, "C", int64(10), nil)
		if c.Output["Test3"].Value != int64(0) {
			panic(3)
		}

		if c.Output["Test4"].Value != int64(-20) {
			panic(4)
		}
	})

	t.Run("", func(t *testing.T) {
		expressions := map[string]string{
			"Test1": "!A",
			"Test2": "-B",
			"Test3": "-B+C",
			"Test4": "-(B+C)",
		}

		descriptor, e := Parse(context.Background(), expressions)
		if e != nil {
			panic(e)
		}

		c := descriptor.Context()

		descriptor.Run(context.Background(), c, "A", true, nil)
		if c.Output["Test1"].Value != false {
			panic("1")
		}

		descriptor.Run(context.Background(), c, "B", int64(10), nil)
		if c.Output["Test2"].Value != int64(-10) {
			panic("2")
		}
		descriptor.Run(context.Background(), c, "C", int64(10), nil)
		if c.Output["Test3"].Value != int64(0) {
			panic(3)
		}

		if c.Output["Test4"].Value != int64(-20) {
			panic(4)
		}
	})
}

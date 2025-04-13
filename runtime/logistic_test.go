package runtime

import (
	"fmt"
	"testing"
)

func TestAnd(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		res, err := And(true, true)
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false1", func(t *testing.T) {
		res, err := And(true, false)
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("false2", func(t *testing.T) {
		res, err := And(false, true)
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("false3", func(t *testing.T) {
		res, err := And(false, false)
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("err", func(t *testing.T) {
		res, err := And(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestOr(t *testing.T) {
	t.Run("true1", func(t *testing.T) {
		res, err := Or(true, true)
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("true2", func(t *testing.T) {
		res, err := Or(true, false)
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("true3", func(t *testing.T) {
		res, err := Or(false, true)
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false", func(t *testing.T) {
		res, err := Or(false, false)
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("err", func(t *testing.T) {
		res, err := Or(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestNot(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		res, err := Not(true)
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true", func(t *testing.T) {
		res, err := Not(false)
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("err", func(t *testing.T) {
		res, err := Not(float64(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

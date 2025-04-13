package runtime

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Run("true1", func(t *testing.T) {
		res, err := Equal(true, true)
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("true2", func(t *testing.T) {
		res, err := Equal("1", "1")
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("true3", func(t *testing.T) {
		res, err := Equal(int64(1), int64(1))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("true4", func(t *testing.T) {
		res, err := Equal(float64(1), float64(1))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("err", func(t *testing.T) {
		res, err := Equal(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})
}

func TestGreatThan(t *testing.T) {
	t.Run("true1", func(t *testing.T) {
		res, err := GreatThan("12", "11")
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false1", func(t *testing.T) {
		res, err := GreatThan("12", "2")
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true2", func(t *testing.T) {
		res, err := GreatThan(int64(10), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false2", func(t *testing.T) {
		res, err := GreatThan(int64(9), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true3", func(t *testing.T) {
		res, err := GreatThan(float64(10), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false3", func(t *testing.T) {
		res, err := GreatThan(float64(9), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true4", func(t *testing.T) {
		res, err := GreatThan(int64(10), int64(9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false4", func(t *testing.T) {
		res, err := GreatThan(int64(9), int64(10))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true5", func(t *testing.T) {
		res, err := GreatThan(float64(10), int64(9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false5", func(t *testing.T) {
		res, err := GreatThan(float64(7), int64(9))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("err", func(t *testing.T) {
		res, err := Equal(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestGreatThanOrEqual(t *testing.T) {
	t.Run("true1", func(t *testing.T) {
		res, err := GreatThanOrEqual("12", "11")
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false1", func(t *testing.T) {
		res, err := GreatThanOrEqual("12", "2")
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true2", func(t *testing.T) {
		res, err := GreatThanOrEqual(int64(10), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false2", func(t *testing.T) {
		res, err := GreatThanOrEqual(int64(9), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true3", func(t *testing.T) {
		res, err := GreatThanOrEqual(float64(10), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false3", func(t *testing.T) {
		res, err := GreatThanOrEqual(float64(9), float64(9.9))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true4", func(t *testing.T) {
		res, err := GreatThanOrEqual(int64(10), int64(9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false4", func(t *testing.T) {
		res, err := GreatThanOrEqual(int64(9), int64(10))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("true5", func(t *testing.T) {
		res, err := GreatThanOrEqual(float64(10), int64(9))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("false5", func(t *testing.T) {
		res, err := GreatThanOrEqual(float64(7), int64(9))
		if err != nil {
			panic(1)
		}
		if res != false {
			panic(2)
		}
	})

	t.Run("err", func(t *testing.T) {
		res, err := Equal(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

	t.Run("true6", func(t *testing.T) {
		res, err := GreatThanOrEqual(float64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

	t.Run("true7", func(t *testing.T) {
		res, err := GreatThanOrEqual(int64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res != true {
			panic(2)
		}
	})

}

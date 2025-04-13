package runtime

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("int64 , int64", func(t *testing.T) {
		res, err := Add(int64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(int64) != 20 {
			panic(2)
		}
	})

	t.Run("int64 , float64", func(t *testing.T) {
		res, err := Add(int64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 20 {
			panic(2)
		}
	})

	t.Run("float64 , float64", func(t *testing.T) {
		res, err := Add(float64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 20 {
			panic(2)
		}
	})

	t.Run("float64 , int64", func(t *testing.T) {
		res, err := Add(float64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 20 {
			panic(2)
		}
	})

	t.Run("float64 + int32", func(t *testing.T) {
		res, err := Add(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestMinus(t *testing.T) {
	t.Run("int64 , int64", func(t *testing.T) {
		res, err := Minus(int64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 0 {
			panic(2)
		}
	})

	t.Run("int64 , float64", func(t *testing.T) {
		res, err := Minus(int64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 0 {
			panic(2)
		}
	})

	t.Run("float64 , float64", func(t *testing.T) {
		res, err := Minus(float64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 0 {
			panic(2)
		}
	})

	t.Run("float64 , int64", func(t *testing.T) {
		res, err := Minus(float64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 0 {
			panic(2)
		}
	})

	t.Run("float64 , int32", func(t *testing.T) {
		res, err := Minus(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestMultiply(t *testing.T) {
	t.Run("int64 , int64", func(t *testing.T) {
		res, err := Multiply(int64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 100 {
			panic(2)
		}
	})

	t.Run("int64 , float64", func(t *testing.T) {
		res, err := Multiply(int64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 100 {
			panic(2)
		}
	})

	t.Run("float64 , float64", func(t *testing.T) {
		res, err := Multiply(float64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 100 {
			panic(2)
		}
	})

	t.Run("float64 , int64", func(t *testing.T) {
		res, err := Multiply(float64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 100 {
			panic(2)
		}
	})

	t.Run("float64 , int32", func(t *testing.T) {
		res, err := Multiply(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestDivide(t *testing.T) {
	t.Run("int64 , int64", func(t *testing.T) {
		res, err := Divide(int64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("int64 , float64", func(t *testing.T) {
		res, err := Divide(int64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("float64 , float64", func(t *testing.T) {
		res, err := Divide(float64(10), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("float64 , int64", func(t *testing.T) {
		res, err := Divide(float64(10), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("float64 , int32", func(t *testing.T) {
		res, err := Divide(float64(10), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestNegative(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		res, err := Negative(int64(10))
		if err != nil {
			panic(1)
		}
		if res.(int64) != -10 {
			panic(2)
		}
	})

	t.Run("float64", func(t *testing.T) {
		res, err := Negative(float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != -10 {
			panic(2)
		}
	})

	t.Run("int", func(t *testing.T) {
		res, err := Negative(10)
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

func TestMod(t *testing.T) {
	t.Run("int64 , int64", func(t *testing.T) {
		res, err := Mod(int64(11), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("int64 , float64", func(t *testing.T) {
		res, err := Mod(int64(11), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("float64 , float64", func(t *testing.T) {
		res, err := Mod(float64(11), float64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("float64 , int64", func(t *testing.T) {
		res, err := Mod(float64(11), int64(10))
		if err != nil {
			panic(1)
		}
		if res.(float64) != 1 {
			panic(2)
		}
	})

	t.Run("float64 , int32", func(t *testing.T) {
		res, err := Mod(float64(11), int32(10))
		fmt.Println(res, err)
		if err == nil {
			panic(1)
		}
	})

}

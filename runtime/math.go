package runtime

import (
	"errors"
	"fmt"
	"math"
)

func Add(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x + y, nil
		case float64:
			return float64(x) + y, nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x + float64(y), nil
		case float64:
			return x + y, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T + %T", a, b))
}

func Mod(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return int64(math.Mod(float64(x), float64(y))), nil
		case float64:
			return math.Mod(float64(x), y), nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return math.Mod(x, float64(y)), nil
		case float64:
			return math.Mod(x, y), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T %% %T", a, b))
}

func Minus(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x - y, nil
		case float64:
			return float64(x) - y, nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x - float64(y), nil
		case float64:
			return x - y, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T - %T", a, b))
}

func Multiply(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x * y, nil
		case float64:
			return float64(x) * y, nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x * float64(y), nil
		case float64:
			return x * y, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T * %T", a, b))
}

func Divide(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x / y, nil
		case float64:
			return float64(x) / y, nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x / float64(y), nil
		case float64:
			return x / y, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T / %T", a, b))
}

func Negative(a interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		return -x, nil
	case float64:
		return -x, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: -%T", a))
}

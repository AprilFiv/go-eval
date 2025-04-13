package runtime

import (
	"errors"
	"fmt"
)

func Equal(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x == y, nil
		case float64:
			return float64(x) == float64(y), nil
		case bool:
			return false, nil
		case string:
			return false, nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return float64(x) == float64(y), nil
		case float64:
			return x == y, nil
		case bool:
			return false, nil
		case string:
			return false, nil
		}
	case string:
		switch y := b.(type) {
		case string:
			return x == y, nil
		case bool:
			return false, nil
		case int64:
			return false, nil
		case float64:
			return false, nil
		}
	case bool:
		switch y := b.(type) {
		case string:
			return false, nil
		case bool:
			return x == y, nil
		case int64:
			return false, nil
		case float64:
			return false, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T == %T", a, b))
}

func NotEqual(a, b interface{}) (interface{}, error) {
	res, err := Equal(a, b)
	if err != nil {
		return nil, err
	}
	return !res.(bool), nil
}

func GreatThan(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x > y, nil
		case float64:
			return float64(x) > float64(y), nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return float64(x) > float64(y), nil
		case float64:
			return x > y, nil
		}
	case string:
		switch y := b.(type) {
		case string:
			return x > y, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("invalid operation: %T > %T", a, b))
}

func GreatThanOrEqual(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x >= y, nil
		case float64:
			return float64(x) >= float64(y), nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return float64(x) >= float64(y), nil
		case float64:
			return x >= y, nil
		}
	case string:
		switch y := b.(type) {
		case string:
			return x >= y, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("invalid operation: %T >= %T", a, b))
}

func LessThan(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x < y, nil
		case float64:
			return float64(x) < float64(y), nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return float64(x) < float64(y), nil
		case float64:
			return x < y, nil
		}
	case string:
		switch y := b.(type) {
		case string:
			return x < y, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("invalid operation: %T < %T", a, b))
}

func LessThanOrEqual(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case int64:
		switch y := b.(type) {
		case int64:
			return x <= y, nil
		case float64:
			return float64(x) <= float64(y), nil
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return float64(x) <= float64(y), nil
		case float64:
			return x <= y, nil
		}
	case string:
		switch y := b.(type) {
		case string:
			return x <= y, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("invalid operation: %T <= %T", a, b))
}

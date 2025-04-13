package runtime

import (
	"errors"
	"fmt"
)

func And(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case bool:
		switch y := b.(type) {
		case bool:
			return x && y, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T && %T", a, b))
}

func Or(a, b interface{}) (interface{}, error) {
	switch x := a.(type) {
	case bool:
		switch y := b.(type) {
		case bool:
			return x || y, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: %T || %T", a, b))
}

func Not(a interface{}) (interface{}, error) {
	switch x := a.(type) {
	case bool:
		return !x, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid operation: !%T", a))
}

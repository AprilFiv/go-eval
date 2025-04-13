package runtime

func ShouldBool(v interface{}) bool {
	_, ok := v.(bool)
	return ok
}

func ShouldInt64(v interface{}) bool {
	_, ok := v.(int64)
	return ok
}

func ShouldFloat64(v interface{}) bool {
	_, ok := v.(float64)
	return ok
}

func ShouldString(v interface{}) bool {
	_, ok := v.(string)
	return ok
}

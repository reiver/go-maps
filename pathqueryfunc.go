package maps

func PathQueryFunc[T any](fn func(any)(T,bool), data map[string]any, path ...string) (T, bool) {
	var empty T

	var value any
	var found bool

	value, found = Type(data).PathQuery(path...)
	if !found {
		return empty, false
	}

	return fn(value)
}

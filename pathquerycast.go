package maps

// PathQueryCast is similar to [PathQuery] except that it
// №1 return the result of the query as the parameterized type, and
// №2 only returns the result if it is of the correct type.
//
// For example, consider this map:
//
//	var data map[string]any = map[string]any{
//		"apple":"1",
//		"banana":2,
//		"cherry":"three",
//	}
//
// This:
//	value, found := maps.PathQueryCast[string](data, "cherry")
//
// Would return:
//
//	value == "three"
//	found == true
//
// This:
//	value, found := maps.PathQueryCast[int](data, "cherry")
//
// Would return:
//
//	value == 0
//	found == false
//
// (Since, although the map key "cherry" exists, it is of type "string" rather than type "int".)
func PathQueryCast[T any](data map[string]any, path ...string) (T, bool) {
	return PathQueryFunc(castfunc[T], data, path...)
}

func castfunc[T any](value any)(T, bool) {
	var empty T

	v, casted := value.(T)
	if !casted {
		return empty, false
	}

	return v, true
}

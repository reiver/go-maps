package maps

func PathQuery(data map[string]any, path ...string) (any, bool) {
	return Type(data).PathQuery(path...)
}

package maps

func PathQueryForString(data map[string]any, path ...string) (string, bool) {
	return Type(data).PathQueryForString(path...)
}

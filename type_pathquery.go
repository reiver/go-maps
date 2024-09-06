package maps

func (receiver Type) PathQuery(path ...string) (any, bool) {
	var empty any

	if nil == receiver {
		return empty, false
	}

	if 1 > len(path) {
		var result any = receiver
		switch casted := result.(type) {
		case Type:
			result = map[string]any(casted)
		}

		return result, true
	}

	name := path[0]

	pointer, found := receiver[name]
	if !found {
		return empty, false
	}

	if 1 == len(path) {
		return pointer, true
	}

	msi, casted := pointer.(map[string]any)
	if !casted {
		return empty, false
	}

	return Type(msi).PathQuery(path[1:]...)
}

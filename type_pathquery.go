package maps

func (receiver Type) PathQuery(path ...string) (any, bool) {
	var empty any

	if nil == receiver {
		return empty, false
	}

	var pathlength int = len(path)

	if pathlength < 1 {
		return map[string]any(receiver), true
	}

	name := path[0]

	pointer, found := receiver[name]
	if !found {
		return empty, false
	}

	if 1 == pathlength {
		return pointer, true
	}

//@TODO: Could this code to turn a map[any]any into a map[string]any be improved? Maybe solve the issue it is solving another way?
	switch casted := pointer.(type) {
	case map[any]any:
		var alt map[string]any = map[string]any{}

		for nameAny, value := range casted {
			var name string
			var nameCasted bool

			name, nameCasted = nameAny.(string)
			if !nameCasted {
				return empty, false
			}

			alt[name] = value
		}

		pointer = alt
	}

	themap, casted := pointer.(map[string]any)
	if !casted {
		return empty, false
	}

	return Type(themap).PathQuery(path[1:]...)
}

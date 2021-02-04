package mapstringinterface

func PathQuery(data map[string]interface{}, path ...string) (interface{}, bool) {
	return Type(data).PathQuery(path...)
}

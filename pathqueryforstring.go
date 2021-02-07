package mapstringinterface

func PathQueryForString(data map[string]interface{}, path ...string) (interface{}, bool) {
	return Type(data).PathQueryForString(path...)
}

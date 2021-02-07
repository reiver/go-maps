package mapstringinterface

func PathQueryForString(data map[string]interface{}, path ...string) (string, bool) {
	return Type(data).PathQueryForString(path...)
}

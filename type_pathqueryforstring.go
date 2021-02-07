package mapstringinterface

import (
	"github.com/reiver/go-cast"
)

func (receiver Type) PathQueryForString(path ...string) (string, bool) {
	i, found := receiver.PathQuery(path...)
	if !found {
		return "", found
	}

	s, err := cast.String(i)

	if nil != err {
		return "", false
	}

	return s, true
}

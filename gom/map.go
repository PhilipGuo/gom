package gom

import "errors"

type gom struct {
	m map[interface{}]*interface{} // map
}

const (
	iniItemCnt = 64
)

//
//
//
func Newgom() *gom {
	return &gom{
		m: make(map[interface{}]*interface{}, iniItemCnt),
	}
}

//
//
//
func (m *gom) Add(k interface{}, v interface{}, replace bool) (bool, error) {
	if nil == m || nil == m.m {
		return false, errors.New("m is nil")
	}

	if nil == k {
		return false, errors.New("k is nil")
	}

	if _, ok := m.m[k]; ok {
		if replace {
			m.m[k] = &v
		} else {
			return false, errors.New("k is already exists, and replace is false")
		}
	} else {
		m.m[k] = &v
	}
	return true, nil
}

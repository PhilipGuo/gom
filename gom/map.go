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

//
// remove by key
//
func (m *gom) RemoveByKey(k interface{}) (bool, error) {
	if nil == m || nil == m.m {
		return false, errors.New("m is nil")
	}

	if _, ok := m.m[k]; ok {
		delete(m.m, k)
		return true, nil
	}
	return false, errors.New("k not exists")
}

//
// remove by value
//
func (m *gom) RemoveByValue(v *interface{}) (bool, error) {
	if nil == m || nil == m.m {
		return false, errors.New("m is nil")
	}
	opted := false
	for key, val := range m.m {
		if *val == *v {
			delete(m.m, key)
			opted = true
		}
	}
	if opted {
		return true, nil
	}
	return false, errors.New("v not exists")
}

func (m *gom) GetValueByKey(k interface{}) (bool, error) {
	return false, nil
}

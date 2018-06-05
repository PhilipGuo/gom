package gom

import (
	"errors"
	"sync"
)

type gom struct {
	m map[interface{}]*interface{} // map

	sync.RWMutex
}

const (
	iniItemCnt = 64
)

// Newgom ...
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

	m.Lock()
	defer m.Unlock()
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

	m.Lock()
	defer m.Unlock()

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

	m.Lock()
	defer m.Unlock()

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

//
// get value by key
//
func (m *gom) GetValueByKey(k interface{}) (interface{}, error) {
	if nil == m || nil == m.m {
		return nil, errors.New("m is nil")
	}

	m.Lock()
	defer m.Unlock()

	if val, ok := m.m[k]; ok {
		return val, nil
	}
	return nil, errors.New("k not exists")
}

//
// get keys([]interface{}) by value
//
func (m *gom) GetKeysByValue(v *interface{}) ([]interface{}, error) {
	if nil == m || nil == m.m {
		return nil, errors.New("m is nil")
	}

	m.Lock()
	defer m.Unlock()

	var keys []interface{}
	for key, val := range m.m {
		if *val == *v {
			keys = append(keys, key)
		}
	}

	return keys, nil
}

//
// get count
//
func (m *gom) GetCount() int {
	if nil == m || nil == m.m {
		return 0
	}

	m.Lock()
	defer m.Unlock()

	return len(m.m)
}

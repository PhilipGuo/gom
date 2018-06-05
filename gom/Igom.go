package gom

type Igom interface {

	// 添加
	// key, value, replace
	Add(interface{}, *interface{}, bool) (bool, error)

	// 根据键删除
	ReomveByKey(interface{}) (bool, error)
	// 根据值删除
	RemoveByValue(interface{}) (bool, error)

	// 根据键读取
	GetValueByKey(interface{}) (interface{}, error)
	// 根据值读取
	GetKeysByValue(*interface{}) ([]interface{}, error)

	// 计数
	Count() int
}

package gom

type Igom interface {

	// 添加
	// key, value, replace
	Add(interface{}, *interface{}, bool) (bool, string)

	// 根据键删除
	ReomveByKey(interface{}) bool
	// 根据值删除
	RemoveByValue(interface{}) bool

	// 根据键读取
	GetByKey(interface{}) *interface{}
	// 根据值读取
	GetByValue(interface{}) []*interface{}

	// 计数
	Count()
}

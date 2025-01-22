package datagroup

import "github.com/spf13/cast"

// Child 实现该接口的类型可以调用group方法和tree方法
type Child[K comparable] interface {
	GetParentId() K
}

// MapListGroup map按parenKey分组
func MapListGroup(list []map[string]any, parentKey string) map[string][]map[string]any {
	res := make(map[string][]map[string]any)
	for _, item := range list {
		anyKey, ok := item[parentKey]
		if ok {
			key := cast.ToString(anyKey)
			var resItem []map[string]any
			if getResItem, ok := res[key]; ok {
				resItem = getResItem
			}
			resItem = append(resItem, item)
			res[key] = resItem
		}
	}
	return res
}

// ListGroup 结构体分组
func ListGroup[K comparable, T Child[K]](list []T) map[K][]T {
	res := make(map[K][]T)
	for _, item := range list {
		key := item.GetParentId()
		var resItem []T
		if getResItem, ok := res[key]; ok {
			resItem = getResItem
		}
		resItem = append(resItem, item)
		res[key] = resItem
	}
	return res
}

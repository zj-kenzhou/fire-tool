package datamap

import "github.com/spf13/cast"

type Id interface {
	GetId() string
}

// MapListIdMap map按key转map
func MapListIdMap(list []map[string]any, idKey string) map[string]map[string]any {
	res := make(map[string]map[string]any)
	for _, item := range list {
		anyKey, ok := item[idKey]
		if ok {
			key := cast.ToString(anyKey)
			res[key] = item
		}
	}
	return res
}

// ListIdMap 结构体转成Id为Key的map
func ListIdMap[T Id](list []T) map[string]T {
	res := make(map[string]T)
	for _, item := range list {
		key := item.GetId()
		res[key] = item
	}
	return res
}

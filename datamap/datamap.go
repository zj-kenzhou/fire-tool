package datamap

import (
	"github.com/spf13/cast"
	"github.com/zj-kenzhou/go-col/cmap"
)

type Id[K comparable] interface {
	GetId() K
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

func CMapListIdMap(list []cmap.Map[string, any], idKey string) map[string]cmap.Map[string, any] {
	res := make(map[string]cmap.Map[string, any])
	for _, item := range list {
		if item.ContainsKey(idKey) {
			anyKey := item.Get(idKey)
			key := cast.ToString(anyKey)
			res[key] = item
		}
	}
	return res
}

// ListIdMap 结构体转成Id为Key的map
func ListIdMap[K comparable, T Id[K]](list []T) map[K]T {
	res := make(map[K]T)
	for _, item := range list {
		key := item.GetId()
		res[key] = item
	}
	return res
}

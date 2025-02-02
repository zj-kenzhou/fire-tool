package datatree

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/zj-kenzhou/fire-tool/datagroup"
	"github.com/zj-kenzhou/fire-tool/datamap"
)

type Tree[K comparable] interface {
	datagroup.Child[K]
	datamap.Id[K]
	SetChild(self, t any)
}

// MapListToTree map转换成tree
func MapListToTree(list []map[string]any, parentKey, idKey, childKey, rootId string) []map[string]any {
	groupMap := datagroup.MapListGroup(list, parentKey)
	rootList := groupMap[rootId]
	toTreeForMap(&rootList, idKey, childKey, groupMap)
	return rootList
}

// ListToTree list转换成tree
func ListToTree[K comparable, T Tree[K]](list []T, rootId K) []T {
	groupMap := datagroup.ListGroup(list)
	rootList := groupMap[rootId]
	toTreeForStruct(&rootList, groupMap)
	return rootList
}

func toTreeForStruct[K comparable, T Tree[K]](list *[]T, groupMap map[K][]T) {
	for index, item := range *list {
		key := item.GetId()
		children, ok := groupMap[key]
		if ok {
			if reflect.TypeOf(item).Kind() == reflect.Ptr {
				item.SetChild(item, children)
			} else {
				item.SetChild(&item, children)
			}
			(*list)[index] = item
			toTreeForStruct(&children, groupMap)
		}
	}
}

func toTreeForMap(list *[]map[string]any, idKey, childKey string, groupMap map[string][]map[string]any) {
	for _, item := range *list {
		idAny, ok := item[idKey]
		if ok {
			key := cast.ToString(idAny)
			children, ok := groupMap[key]
			if ok {
				item[childKey] = children
				toTreeForMap(&children, idKey, childKey, groupMap)
			}
		}
	}
}

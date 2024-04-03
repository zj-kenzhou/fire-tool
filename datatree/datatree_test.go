package datatree

import (
	"encoding/json"
	"testing"
)

type childItem struct {
	Name   string
	Parent string
	Id     string
	Child  []*childItem
}

func (c childItem) GetParentId() string {
	return c.Parent
}

func (c childItem) GetId() string {
	return c.Id
}

func (c childItem) SetChild(self, item any) {
	this := self.(*childItem)
	this.Child = item.([]*childItem)
}

func TestMapListToTree(t *testing.T) {
	var list []map[string]any
	list = append(list, map[string]any{
		"name":   "1111",
		"parent": "0",
		"id":     "1",
	})
	list = append(list, map[string]any{
		"name":   "222",
		"parent": "0",
		"id":     "2",
	})
	list = append(list, map[string]any{
		"name":   "33",
		"parent": "1",
		"id":     "3",
	})
	list = append(list, map[string]any{
		"name":   "44",
		"parent": "1",
		"id":     "4",
	})
	list = append(list, map[string]any{
		"name":   "55",
		"parent": "4",
		"id":     "5",
	})
	list = append(list, map[string]any{
		"name":   "66",
		"parent": "2",
		"id":     "6",
	})
	res := MapListToTree(list, "parent", "id", "children", "0")
	t.Log(res)
}

func TestListToTree(t *testing.T) {
	var list []*childItem
	list = append(list, &childItem{
		Name:   "1111",
		Parent: "0",
		Id:     "1",
	})
	list = append(list, &childItem{
		Name:   "222",
		Parent: "0",
		Id:     "2",
	})
	list = append(list, &childItem{
		Name:   "33",
		Parent: "1",
		Id:     "3",
	})
	list = append(list, &childItem{
		Name:   "44",
		Parent: "1",
		Id:     "4",
	})
	list = append(list, &childItem{
		Name:   "55",
		Parent: "4",
		Id:     "5",
	})
	list = append(list, &childItem{
		Name:   "66",
		Parent: "2",
		Id:     "6",
	})
	res := ListToTree(list, "0")
	resByte, _ := json.Marshal(res)
	t.Log(string(resByte))
}

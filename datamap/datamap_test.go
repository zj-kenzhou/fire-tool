package datamap

import "testing"

type entity struct {
	Id   string
	Name string
}

func (e entity) GetId() string {
	return e.Id
}

func TestListIdMap(t *testing.T) {
	var list []entity
	list = append(list, entity{
		Id:   "1",
		Name: "1",
	})
	list = append(list, entity{
		Id:   "2",
		Name: "2",
	})
	list = append(list, entity{
		Id:   "3",
		Name: "3",
	})
	res := ListIdMap(list)
	if len(res) != 3 {
		t.Error("res len is not 3")
	}
	t.Log(res)
}

func TestMapListIdMap(t *testing.T) {
	var list []map[string]any
	list = append(list, map[string]any{
		"Id":   "1",
		"Name": "1",
	})
	list = append(list, map[string]any{
		"Id":   "2",
		"Name": "2",
	})
	res := MapListIdMap(list, "Id")
	if len(res) != 2 {
		t.Error("res len is not 2")
	}
	t.Log(res)
}

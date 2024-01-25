package datagroup

import "testing"

type testData struct {
	Id      string
	ParenId string
	Name    string
}

func (t testData) GetParentId() string {
	return t.ParenId
}

func TestListGroup(t *testing.T) {
	var list []testData
	list = append(list, testData{
		Id:      "101",
		ParenId: "1",
		Name:    "101",
	})
	list = append(list, testData{
		Id:      "102",
		ParenId: "1",
		Name:    "102",
	})
	list = append(list, testData{
		Id:      "201",
		ParenId: "2",
		Name:    "201",
	})
	list = append(list, testData{
		Id:      "202",
		ParenId: "2",
		Name:    "202",
	})
	res := ListGroup(list)
	if len(res) != 2 {
		t.Error("res len is not 2")
	}
	if res["1"][0].Name != "101" {
		t.Error("res 0 0 name is not 101")
	}
	t.Log(res)
}

func TestMapListGroup(t *testing.T) {
	var list []map[string]any
	list = append(list, map[string]any{
		"Id":      "101",
		"ParenId": "1",
		"Name":    "101",
	})
	list = append(list, map[string]any{
		"Id":      "102",
		"ParenId": "1",
		"Name":    "102",
	})
	list = append(list, map[string]any{
		"Id":      "201",
		"ParenId": "2",
		"Name":    "201",
	})
	list = append(list, map[string]any{
		"Id":      "202",
		"ParenId": "2",
		"Name":    "202",
	})
	res := MapListGroup(list, "ParenId")

	if len(res) != 2 {
		t.Error("res len is not 2")
	}
	if res["1"][0]["Name"] != "101" {
		t.Error("res 0 0 name is not 101")
	}
	t.Log(res)
}

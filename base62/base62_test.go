package base62

import "testing"

func TestEncode(t *testing.T) {
	res := Encode(1000002)
	if res != "49c4" {
		t.Error("res is not valid")
	}
	t.Log(res)
}

func TestDecode(t *testing.T) {
	res := Decode("49c4")
	if res != 1000002 {
		t.Error("res is not valid")
	}
	t.Log(res)
}

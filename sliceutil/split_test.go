package sliceutil

import "testing"

func TestSliceSplit(t *testing.T) {
	testList := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	res := Split(testList, 3)
	t.Log(res)
}

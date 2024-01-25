package sliceutil

func Split(list []any, splitNum int) [][]any {
	if list == nil || len(list) == 0 {
		return [][]any{}
	}
	listLen := len(list)
	if listLen <= splitNum {
		return [][]any{list}
	}
	var res [][]any
	start := 0
	for {
		end := start + splitNum
		if end > listLen {
			end = listLen
		}
		res = append(res, list[start:end])
		start = start + splitNum
		if end >= listLen {
			break
		}
	}
	return res
}

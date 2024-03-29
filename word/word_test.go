package word

import "testing"

func TestUnderLineToCamelCaseTitle(t *testing.T) {
	str := "aaa_bbb_ccc _ddd"
	res := UnderLineToCamelCaseTitle(str)
	t.Log(res)
}

func TestCamelCaseToCenterLine(t *testing.T) {
	str := "aaaBaaaCCaaa"
	res := CamelCaseToCenterLine(str)
	if "aaa-baaa-c-caaa" != res {
		t.Error("convert err")
	}
	t.Log(res)
}

func TestCenterLineToCamelCase(t *testing.T) {
	str := "aaa-baaa-c-caaa"
	res := CenterLineToCamelCase(str)
	if "aaaBaaaCCaaa" != res {
		t.Error("convert err")
	}
	t.Log(res)
}

func TestUnderLineToCamelCase(t *testing.T) {
	str := "aaa_baaa_c_caaa"
	res := UnderLineToCamelCase(str)
	if "aaaBaaaCCaaa" != res {
		t.Error("convert err")
	}
	t.Log(res)
}

func TestCamelCaseToUnderLine(t *testing.T) {
	str := "aaaBaaaCCaaa"
	res := CamelCaseToUnderLine(str)
	if "aaa_baaa_c_caaa" != res {
		t.Error("convert err")
	}
	t.Log(res)
}

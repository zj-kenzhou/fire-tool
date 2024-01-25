package word

import (
	"testing"
)

// MaskString 脱敏数据
func TestMaskString(t *testing.T) {
	input := "1769313"
	start := int32(1)
	end := int32(3)
	t.Log(MaskString(input, &start, &end))
}

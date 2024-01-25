package sha

import "testing"

func TestHmacSHA256(t *testing.T) {
	res := HmacSHA256("aaa", "fasfsafsf")
	t.Log(res)
}

func TestHmacSHA256Equal(t *testing.T) {
	param := HmacSHA256("aaa", "fasfsafsf")
	res := HmacSHA256Equal("NmE4NjZlYWE0ZDA2YWNlYTAzYjhiNDFlYWFjOTU3NDJiYmNiNzM0NjBiNzlhMmI1YmY5NzQyMDJhNjhmMmY5Mg==", param)
	t.Log(res)
}

func TestGenKey(t *testing.T) {
	str := "e123213214324134124124124312"
	res := GenKey(str)
	t.Log(len(res))
	t.Log(string(res))
}

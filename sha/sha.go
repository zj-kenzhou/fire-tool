package sha

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HmacSHA256(str string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

func HmacSHA256Equal(str string, str1 string) bool {
	return hmac.Equal([]byte(str), []byte(str1))
}

// GenKey 生成32位的key
func GenKey(src string) []byte {
	m := sha256.New()
	m.Write([]byte(src))
	return m.Sum(nil)
}

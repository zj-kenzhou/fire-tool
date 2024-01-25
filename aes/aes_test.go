package aes

import "testing"

func TestEncryptWithKey(t *testing.T) {
	res := EncryptWithKey("aaaaa", []byte("cApdfK6B3xihDtf1"))
	if res != "YUpfvD2RTwbN0BoUZuNQUw" {
		t.Error("encode err")
	}
	t.Log(res)
}

func TestDecryptWithKey(t *testing.T) {
	res := DecryptWithKey("YUpfvD2RTwbN0BoUZuNQUw", []byte("cApdfK6B3xihDtf1"))
	if res != "aaaaa" {
		t.Error("encode err")
	}
	t.Log(res)
}

func TestEncrypt(t *testing.T) {
	res := Encrypt("aaaaa")
	if res != "YUpfvD2RTwbN0BoUZuNQUw" {
		t.Error("encode err")
	}
	t.Log(res)
}

func TestDecrypt(t *testing.T) {
	res := Decrypt("YUpfvD2RTwbN0BoUZuNQUw")
	if res != "aaaaa" {
		t.Error("encode err")
	}
	t.Log(res)
}

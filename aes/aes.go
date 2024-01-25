package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const DefaultKey = `cApdfK6B3xihDtf1`

func EncryptWithKey(orig string, k []byte) string {
	// 转成字节数组
	origData := []byte(orig)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = pKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	crated := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(crated, origData)
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	return base64.RawURLEncoding.EncodeToString(crated)

}

func Encrypt(orig string) string {
	return EncryptWithKey(orig, []byte(DefaultKey))
}

func DecryptWithKey(crated string, k []byte) string {
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	cratedByte, _ := base64.RawURLEncoding.DecodeString(crated)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cratedByte))
	// 解密
	blockMode.CryptBlocks(orig, cratedByte)
	// 去补全码
	orig = pKCS7UnPadding(orig)
	return string(orig)
}

func Decrypt(crated string) string {
	return DecryptWithKey(crated, []byte(DefaultKey))
}

// pKCS7Padding 补码
func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// pKCS7UnPadding 去码
func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

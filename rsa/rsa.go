package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

const defaultPublicKey = `
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKucO9V/93kja3FbL0oaghy4MtsLeoLm
qLd8Ox0kyvOpFsPXjTD+988U/Nhr7MKDa0XSBOCVd/9GcyVRuWCvXXcCAwEAAQ==
-----END PUBLIC KEY-----
`
const defaultPrivateKey = `
-----BEGIN PRIVATE KEY-----
MIIBVgIBADANBgkqhkiG9w0BAQEFAASCAUAwggE8AgEAAkEAq5w71X/3eSNrcVsv
ShqCHLgy2wt6guaot3w7HSTK86kWw9eNMP73zxT82GvswoNrRdIE4JV3/0ZzJVG5
YK9ddwIDAQABAkEAkztYbmT6yjikPr3vxZEWnlM+doXTlykUCVARW4CsivzYfXgR
yTCgVEvf1aGnJZ1yYUHWnzvLH14117AwqF6CAQIhANKK/XaGqK73CnSP3st1G05f
iZX1E72baLGBwipSlTEBAiEA0KlhcVUEWyIwnzbkhCxq7mNBqd8ba9GXY6XWTSOf
lncCIA+sl0gOqFo2PAoHd++vrTO9exG16B1Sh43HejwKX98BAiEAu7Y0x+ywj3+R
eBm8t77xr2UxdC0WkQcrBuTwhdKvRX8CIQDBmQ9ZpjlG6RRsyLXfUV8s/2e48gXZ
fKdnZgW5Mv6jOA==
-----END PRIVATE KEY-----
`

// Encrypt 加密
func Encrypt(origData string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	res, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	return base64.StdEncoding.EncodeToString(res), err
}

func EncryptDefaultKey(origData string) (string, error) {
	return Encrypt(origData, []byte(defaultPublicKey))
}

// Decrypt rss 解密（密码用途）
func Decrypt(cipherText string, privateKey []byte) (res string, err error) {
	text, _ := base64.StdEncoding.DecodeString(cipherText)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	parseResult, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", errors.New("parse error")
	}
	privy := parseResult.(*rsa.PrivateKey)
	re, err := rsa.DecryptPKCS1v15(rand.Reader, privy, text)
	return string(re), err
}

func DecryptDefaultKey(origData string) (string, error) {
	return Decrypt(origData, []byte(defaultPrivateKey))
}

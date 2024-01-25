package rsa

import "testing"

const testPublicKey = `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANMKhZCu3bJpx0gzwZD92X/lHd/AaDFt
96w4shw/nxMZKxUvzbbJ6aWOFWjY8oSZKlUrgY8etSJDS1jYsLT4NucCAwEAAQ==
-----END PUBLIC KEY-----`
const testPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEA0wqFkK7dsmnHSDPB
kP3Zf+Ud38BoMW33rDiyHD+fExkrFS/NtsnppY4VaNjyhJkqVSuBjx61IkNLWNiw
tPg25wIDAQABAkBOn7zjCkLOFVqJK78RpYFChLl8gtJOmAmNwBGbNJivJCOBupKh
wJLy6GjN9a+Te8vsNM8vPW89PU9MdoUtQZEBAiEA/Fs5iRcs6yc2V9O0Gt6dyoDS
PjH1/Qys1XVNJdzilw0CIQDWFpXC0RGJh3NMR1HT3I1UrGoWXLD0iInqNgV8YJHI
wwIgEEbGADyZbhVaYPHLpQwEu2HxrYdFaWve3emx7yY/CyUCIF/BkFCn9IEgHXjz
uWzajDhgOb8wzFFDtPQWH5QHGY3DAiB9dwklZW3Kh4lU7uZaKeUtAumLt73xBRpG
M8U3CaLPFg==
-----END PRIVATE KEY-----`

func TestEncrypt(t *testing.T) {
	publicEncRes, err := Encrypt("aaa", []byte(testPublicKey))
	if err != nil {
		t.Error(err)
	}
	t.Log(publicEncRes)
}
func TestDecrypt(t *testing.T) {
	res, err := Decrypt("LEhp/Tb53S7K8pQSjw7NjNo+Mk1F/JW8W1pwl8BPf/EcsTFSlMin1P5y8skYShNv5i/J2exA6fy4SyFlesTVNA==", []byte(testPrivateKey))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestEncryptDefaultKey(t *testing.T) {
	publicEncRes, err := EncryptDefaultKey("aaa")
	if err != nil {
		t.Error(err)
	}
	t.Log(publicEncRes)
}

func TestDecryptDefaultKey(t *testing.T) {
	res, err := DecryptDefaultKey("Wp5wPC3n0YlEZKemvG7ngpErd5m/XGzXxmcWKokJ/8TzH+jvauuFNPnxi3y2OmFGAUohNeGUPjAh0o41Dfg28w==")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

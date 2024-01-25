package ip

import "testing"

func TestExternalIp(t *testing.T) {
	ip, err := ExternalIp()
	if err != nil {
		t.Error(err)
	}
	t.Log(ip)
}

func TestExternalIpNoErr(t *testing.T) {
	ip := ExternalIpNoErr()
	t.Log(ip)
}

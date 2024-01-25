package bcrypt

import "testing"

func TestEncodePassword(t *testing.T) {
	res := EncodePassword("admin")
	t.Log(res)
}

func TestValidatePassword(t *testing.T) {
	res := ValidatePassword("$2a$04$BS0hgxukDA265Y4uHxxRNupkr9PCKDPYt/vXr2m/GXtwUD/HWj1dW", "admin")
	t.Log(res)
}

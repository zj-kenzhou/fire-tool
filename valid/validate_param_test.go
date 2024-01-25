package valid

import "testing"

type testParam struct {
	Id string `validate:"required,min=3,max=20"`
}

func TestValidate(t *testing.T) {
	param := testParam{
		Id: "1",
	}
	err := Validate(param)
	if err != nil {
		t.Log(err)
	}
}

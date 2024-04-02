package valid

import "testing"

type testParam struct {
	Id   string `validate:"required,min=3,max=20"`
	Name string `validate:"required,tableName"`
}

func TestValidate(t *testing.T) {
	param := testParam{
		Id:   "2",
		Name: "test",
	}
	err := Validate(param)
	if err != nil {
		t.Log(err)
	}
}

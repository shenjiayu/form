package form

import (
	"testing"
)

func TestForm(t *testing.T) {
	var content = "shenjiayu"
	err := Validate("username", content)
	if err != nil {
		t.Error(err)
	}
	content = "blablabla"
	err = Validate("password", content)
	if err != nil {
		t.Error(err)
	}
	content = "sjy19930312@gmail.com"
	err = Validate("email", content)
	if err != nil {
		t.Error(err)
	}
}

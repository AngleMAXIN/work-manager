package utils

import "testing"

func TestMakePassword(t *testing.T) {
	var ok bool
	rowPassword := "admin12"
	errPassword := "admin122"
	savePassword, err := MakePassword(rowPassword)
	if err != nil {
		t.Error("password is faild")
	}
	if ok = CheckPassword(rowPassword, savePassword); !ok {
		t.Error("password is Error")
	}

	if ok = CheckPassword(errPassword, savePassword); ok {
		t.Error("password is Error")
	}

}

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	n := 33
	res := RandStringBytesMaskImprSrc(n)
	t.Log(res)
}

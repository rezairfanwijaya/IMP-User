package helper

import "testing"

func TestGetENV(t *testing.T) {
	env, err := GetENV("../.env")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(env)
}

func TestHashedPassword(t *testing.T) {
	res, err := HashingPassword("12345")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(res)
}

func TestVerifyPassword(t *testing.T) {
	if err := VerifyPassword("$2a$10$5xE4X8yW1iQr5MHUGx0KU.qwdpQEImqL.8/zNWO5KpIKrTpXMcKci", "12345"); err != nil {
		t.Fatal(err.Error())
	}
}

package helper

import "testing"

func TestGetENV(t *testing.T) {
	env, err := GetENV("../.env")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(env)
}

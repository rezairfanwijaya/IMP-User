package database

import "testing"

func TestNewConnection(t *testing.T) {
	conn, err := NewConnection("../.env")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(conn)
}

package auth

import "testing"

var service = Service{}

func TestGenerateToken(t *testing.T) {
	token, err := service.GenerateToken(2)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(token)
}

func TestVerifyToken(t *testing.T) {
	token, err := service.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.JbhU1dl-845Le3GDhy1wPusjuSkiE5BC4xDH50KJJgk")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(token)
}

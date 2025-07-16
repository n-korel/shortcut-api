package jwt_test

import (
	"testing"

	"github.com/n-korel/shortcut-api/pkg/jwt"
)

func TestJWTCreate(t *testing.T) {
	const email = "nick@d.ru"
	jwtService := jwt.NewJWT("/2+XnmJGz1j3ehIVI/5P9kl+CghrE3DcS7rnT+qar5w=")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}

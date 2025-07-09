package auth

import (
	"fmt"
	"net/http"

	"github.com/n-korel/shortcut-api/configs"
	"github.com/n-korel/shortcut-api/pkg/req"
	"github.com/n-korel/shortcut-api/pkg/res"
)

type AuthHandlerDeps struct{
	*configs.Config
}

type AuthHandler struct{
	*configs.Config
}

func NewAuthHandler(router * http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return 
		}

		fmt.Println(body)

		data := LoginResponse{
			Token: "12345678",
		}
		res.Json(w, data, 200)
		
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return 
		}

		fmt.Println(body)

		data := RegisterResponse{
			Token: "12345678",
		}
		res.Json(w, data, 200)
	}
}
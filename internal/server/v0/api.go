package v0

import (
	"net/http"

	"github.com/MulDeveloper/go-test-crud/internal/data"
	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	return r
}

package v1

import (
	"net/http"

	"github.com/KovacDR/go-mysql-api/internal/storage"
	"github.com/KovacDR/go-mysql-api/pkg/post"
	"github.com/KovacDR/go-mysql-api/pkg/user"
	"github.com/go-chi/chi"
)


func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &user.UserRepository{
			Storage: storage.New(),
		},
	}

	pr := &PostRouter{
		Repository: &post.PostRepository{
			Storage: storage.New(),
		},
	}

	r.Mount("/users", ur.Routes())
	r.Mount("/posts", pr.Routes())

	return r
}
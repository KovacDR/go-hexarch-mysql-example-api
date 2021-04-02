package v1

import (
	"net/http"

	"github.com/KovacDR/go-mysql-api/pkg/post"
	"github.com/go-chi/chi"
)


type PostRouter struct {
	Repository post.Repository
}

func (pr *PostRouter) Routes() http.Handler {
	r := chi.NewRouter()


	return r
}
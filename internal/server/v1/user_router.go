package v1

import (
	"encoding/json"
	"net/http"

	"github.com/KovacDR/go-mysql-api/pkg/response"
	"github.com/KovacDR/go-mysql-api/pkg/user"
	"github.com/go-chi/chi"
)


type UserRouter struct {
	Repository user.Repository
}

func (ur *UserRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/all", ur.GetAllHandler)
	r.Post("/", ur.CreateHanlder)

	return r
}

func (ur *UserRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, response.Map{"users": users})
}

func (ur *UserRouter) CreateHanlder(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	u.Password = ""
	response.JSON(w, r, http.StatusCreated, response.Map{"user": u})
}
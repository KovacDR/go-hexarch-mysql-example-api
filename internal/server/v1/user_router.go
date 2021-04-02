package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	r.Get("/one", ur.GetOneHanlder)
	r.Get("/one/username", ur.GetByUsernameHandler)
	r.Post("/", ur.CreateHanlder)
	r.Put("/{id}", ur.UpdateHandler)
	r.Delete("/{id}", ur.DeleteHanlder)

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

func (ur *UserRouter) GetOneHanlder(w http.ResponseWriter, r *http.Request) {
	sId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	user, err := ur.Repository.GetOne(r.Context(), id)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, response.Map{"user": user})
}

func (ur *UserRouter) GetByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	user, err := ur.Repository.GetByUsername(r.Context(), username)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, response.Map{"user": user})
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

func (ur *UserRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	sId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	var user user.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	err = ur.Repository.Update(r.Context(), id, user)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, response.Map{"updated": true, "udated Fields": user})
}

func (ur *UserRouter) DeleteHanlder(w http.ResponseWriter, r *http.Request) {
	sId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = ur.Repository.Delete(r.Context(), id)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, r, http.StatusAccepted, response.Map{"deleted": true})
}
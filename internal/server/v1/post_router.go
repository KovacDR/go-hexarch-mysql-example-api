package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KovacDR/go-mysql-api/pkg/post"
	"github.com/KovacDR/go-mysql-api/pkg/response"
	"github.com/go-chi/chi"
)


type PostRouter struct {
	Repository post.Repository
}

func (pr *PostRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/all", pr.GetAllHandler)
	r.Get("/one", pr.GetOneHandler)
	r.Get("/one/username", pr.GetByUserHandler)
	r.Post("/", pr.CreateHanlder)
	r.Put("/update/{id}", pr.UpdateHandler)
	r.Delete("/delete/{id}", pr.DeleteHandler)

	return r
}

func (pr *PostRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := pr.Repository.GetAll(r.Context())
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, posts)
}

func (pr *PostRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	sId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	post, err := pr.Repository.GetOne(r.Context(), id)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, post)
}

func (pr *PostRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("username")
	userID, err := strconv.Atoi(user)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	post, err := pr.Repository.GetByUser(r.Context(), userID)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, r, http.StatusAccepted, post)
}

func (pr *PostRouter) CreateHanlder(w http.ResponseWriter, r *http.Request) {
	var p post.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	defer r.Body.Close()

	err = pr.Repository.Create(r.Context(), &p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusCreated, p)
}

func (pr *PostRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	sId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	ctx := r.Context()

	p, err := pr.Repository.GetOne(ctx, id)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	var post post.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	defer r.Body.Close()

	if post.Body == "" {
		post.Body = p.Body
	}
	if post.Image == "" {
		post.Image = p.Image
	}

	err = pr.Repository.Update(ctx, id, post)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, response.Map{"updated": true, "updated fields": post})
}

func (pr *PostRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	sId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	err = pr.Repository.Delete(r.Context(), id)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, r, http.StatusAccepted, response.Map{"deleted": true})
}
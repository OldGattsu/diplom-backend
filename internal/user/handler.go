package user

import (
	"diplom/internal/handlers"
	"diplom/pkg/logging"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/users/{uuid}"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger,
	}
}

func (h *handler) Register(r chi.Router) {
	r.Get(usersURL, h.GetList)
	r.Get(userURL, h.GetUserByUUID)
	r.Post(usersURL, h.CreateUser)
	r.Put(userURL, h.UpdateUser)
	r.Patch(userURL, h.PartiallyUpdateUser)
	r.Delete(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("this is user by uuid"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("this is create user"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))
}

package application

import (
	"encoding/json"
	"github.com/oldgattsu/diplom2/internal/models"
	"github.com/oldgattsu/diplom2/internal/storage"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type registrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registrationResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

func (app *Application) handlerRegistration(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler registration")

	r := registrationRequest{}

	errUnmarshal := json.NewDecoder(req.Body).Decode(&r)
	if errUnmarshal != nil {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}

	uq := &storage.AddUserQuery{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		IsAdmin:  false,
	}

	u, errUser := app.store.AddUser(req.Context(), uq)
	if errUser != nil {
		// todo обработать ошибку нормально
		//if errors.Is(errUser, storage.ErrUserNotFound) {
		//	http.Error(rw, "unauthorized", http.StatusUnauthorized)
		//	return
		//}

		app.logger.Error("error query db", zap.Error(errUser))
		http.Error(rw, "error query db", http.StatusInternalServerError)
		return
	}

	token := uuid.New().String()

	errSaveToken := app.store.SaveToken(req.Context(), u.ID, token)
	if errSaveToken != nil {
		app.logger.Error("error save token", zap.Error(errSaveToken))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	resp := registrationResponse{
		Token: token,
		User:  u,
	}

	data, errMarshal := json.Marshal(resp)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errSaveToken))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}

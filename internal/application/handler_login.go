package application

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/storage"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (app *Application) handlerLogin(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler login")

	body, errReadBody := io.ReadAll(req.Body)
	if errReadBody != nil {
		app.logger.Error("error read request body", zap.Error(errReadBody))
		http.Error(rw, "error read body", http.StatusInternalServerError)
		return
	}

	r := loginRequest{}

	errUnmarshal := json.Unmarshal(body, &r)
	if errUnmarshal != nil {
		app.logger.Error("error unmarshal request body", zap.Error(errReadBody))
		http.Error(rw, "error unmarshal body", http.StatusBadRequest)
		return
	}

	u, errUser := app.store.GetUser(req.Context(), r.Email, r.Password)
	if errUser != nil {
		if errors.Is(errUser, storage.ErrUserNotFound) {
			http.Error(rw, "unauthorized", http.StatusUnauthorized)
			return
		}

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

	resp := loginResponse{Token: token}

	data, errMarshal := json.Marshal(resp)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errSaveToken))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}

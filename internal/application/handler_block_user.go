package application

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type blockUserRequest struct {
	ID        int  `json:"id"`
	IsBlocked bool `json:"is_blocked"`
}

func (app *Application) handlerBlockUser(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler block user")

	var bur *blockUserRequest

	errUnmarshal := json.NewDecoder(req.Body).Decode(&bur)
	if errUnmarshal != nil {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}

	errBlockUser := app.store.BlockUser(req.Context(), bur.ID, bur.IsBlocked)
	if errBlockUser != nil {
		app.logger.Error("error update book", zap.Error(errBlockUser))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}
}

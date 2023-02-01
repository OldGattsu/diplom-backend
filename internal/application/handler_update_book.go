package application

import (
	"encoding/json"
	"github.com/oldgattsu/diplom2/internal/storage"
	"net/http"

	"go.uber.org/zap"
)

type updateBookRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Authors     []int  `json:"authors"`
}

func (app *Application) handlerUpdateBook(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler update book")

	r := updateBookRequest{}

	errUnmarshal := json.NewDecoder(req.Body).Decode(&r)
	if errUnmarshal != nil {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}

	if r.Name == "" {
		http.Error(rw, "name is empty", http.StatusBadRequest)
		return
	}

	bq := &storage.UpdateBookQuery{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Authors:     r.Authors,
	}
	errUpdateBook := app.store.UpdateBook(req.Context(), bq)
	if errUpdateBook != nil {
		app.logger.Error("error update book", zap.Error(errUpdateBook))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}
}

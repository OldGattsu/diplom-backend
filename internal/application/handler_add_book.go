package application

import (
	"encoding/json"
	"github.com/oldgattsu/diplom2/internal/models"
	"github.com/oldgattsu/diplom2/internal/storage"
	"net/http"

	"go.uber.org/zap"
)

type addBookRequest struct {
	Name    string `json:"name"`
	Authors []int  `json:"authors"`
}

func (app *Application) handlerAddBook(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler add book")

	r := addBookRequest{}

	errUnmarshal := json.NewDecoder(req.Body).Decode(&r)
	if errUnmarshal != nil {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}

	if r.Name == "" {
		http.Error(rw, "name is empty", http.StatusBadRequest)
		return
	}

	bq := &storage.AddBookQuery{
		Name:    r.Name,
		Authors: r.Authors,
	}
	bookID, errAddBook := app.store.AddBook(req.Context(), bq)
	if errAddBook != nil {
		app.logger.Error("error add book", zap.Error(errAddBook))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	b := &models.Book{
		ID:   bookID,
		Name: r.Name,
	}
	resp, errMarshal := json.Marshal(b)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

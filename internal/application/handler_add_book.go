package application

import (
	"encoding/json"
	"net/http"

	"github.com/oldgattsu/diplom2/internal/models"

	"go.uber.org/zap"
)

type addBookRequest struct {
	Name string `json:"name"`
}

func (app *Application) handlerAddBook(rw http.ResponseWriter, req *http.Request) {
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

	b := &models.Book{
		Name: r.Name,
	}

	bookID, errAddBook := app.store.AddBook(req.Context(), b)
	if errAddBook != nil {
		app.logger.Error("error add book", zap.Error(errAddBook))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	b.ID = bookID

	resp, errMarshal := json.Marshal(b)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

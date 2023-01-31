package application

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/storage"
)

func (app *Application) handlerDeleteBook(rw http.ResponseWriter, req *http.Request) {
	idStr := chi.URLParam(req, "id")
	id, errConvert := strconv.Atoi(idStr)
	if errConvert != nil {
		http.Error(rw, "book id must be an integer", http.StatusBadRequest)
		return
	}

	errDeleteBook := app.store.DeleteBook(req.Context(), id)
	if errDeleteBook != nil {
		if errors.Is(errDeleteBook, storage.ErrBookNotFound) {
			http.Error(rw, "", http.StatusNotFound)
			return
		}
		app.logger.Error("error delete book", zap.Error(errDeleteBook))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}
}

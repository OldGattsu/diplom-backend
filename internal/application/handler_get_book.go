package application

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/storage"
)

func (app *Application) handlerGetBook(rw http.ResponseWriter, req *http.Request) {
	idStr := chi.URLParam(req, "id")
	id, errConvert := strconv.Atoi(idStr)
	if errConvert != nil {
		http.Error(rw, "book id must be an integer", http.StatusBadRequest)
		return
	}

	b, errGetBook := app.store.GetBook(req.Context(), id)
	if errGetBook != nil {
		if errors.Is(errGetBook, storage.ErrBookNotFound) {
			http.Error(rw, "", http.StatusNotFound)
			return
		}
		app.logger.Error("error get book", zap.Error(errGetBook))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(b) // todo: check error

	rw.Write(resp)
}

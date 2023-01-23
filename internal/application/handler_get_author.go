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

func (app *Application) handlerGetAuthor(rw http.ResponseWriter, req *http.Request) {
	idStr := chi.URLParam(req, "id")
	id, errConvert := strconv.Atoi(idStr)
	if errConvert != nil {
		http.Error(rw, "author id must be an integer", http.StatusBadRequest)
		return
	}

	b, errGetAuthor := app.store.GetAuthor(req.Context(), id)
	if errGetAuthor != nil {
		if errors.Is(errGetAuthor, storage.ErrAuthorNotFound) {
			http.Error(rw, "", http.StatusNotFound)
			return
		}
		app.logger.Error("error get book", zap.Error(errGetAuthor))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(b) // todo: check error

	rw.Write(resp)
}

package application

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/storage"
)

func (app *Application) handlerDeleteComment(rw http.ResponseWriter, req *http.Request) {
	idStr := chi.URLParam(req, "id")
	id, errConvert := strconv.Atoi(idStr)
	if errConvert != nil {
		http.Error(rw, "comment id must be an integer", http.StatusBadRequest)
		return
	}

	errDeleteComment := app.store.DeleteComment(req.Context(), id)
	if errDeleteComment != nil {
		if errors.Is(errDeleteComment, storage.ErrCommentNotFound) {
			http.Error(rw, "", http.StatusNotFound)
			return
		}
		app.logger.Error("error delete book", zap.Error(errDeleteComment))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}
}

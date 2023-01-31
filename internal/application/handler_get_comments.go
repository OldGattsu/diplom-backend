package application

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (app *Application) handlerGetComments(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler getComments")

	idStr := chi.URLParam(req, "id")
	id, errConvert := strconv.Atoi(idStr)
	if errConvert != nil {
		http.Error(rw, "comment id must be an integer", http.StatusBadRequest)
		return
	}

	comments, errGetComments := app.store.GetComments(req.Context(), id)
	if errGetComments != nil {
		app.logger.Error("error get comments", zap.Error(errGetComments))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	resp, errMarshal := json.Marshal(comments)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

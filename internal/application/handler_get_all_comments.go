package application

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (app *Application) handlerGetAllComments(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler get all comments")

	comments, errGetComments := app.store.GetAllComments(req.Context())
	if errGetComments != nil {
		app.logger.Error("error get books", zap.Error(errGetComments))
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

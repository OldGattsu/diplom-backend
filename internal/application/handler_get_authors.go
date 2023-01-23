package application

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (app *Application) handlerGetAuthors(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler getAuthors")

	books, errGetAuthors := app.store.GetAuthors(req.Context())
	if errGetAuthors != nil {
		app.logger.Error("error get authors", zap.Error(errGetAuthors))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	resp, errMarshal := json.Marshal(books)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

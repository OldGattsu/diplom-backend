package application

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (app *Application) handlerGetBooks(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler getBooks")

	books, errGetBooks := app.store.GetBooks(req.Context())
	if errGetBooks != nil {
		app.logger.Error("error get books", zap.Error(errGetBooks))
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

package application

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (app *Application) handlerGetUsers(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler getUsers")

	users, errGetUsers := app.store.GetUsers(req.Context())
	if errGetUsers != nil {
		app.logger.Error("error get books", zap.Error(errGetUsers))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	resp, errMarshal := json.Marshal(users)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

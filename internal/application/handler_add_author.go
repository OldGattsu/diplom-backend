package application

import (
	"encoding/json"
	"net/http"

	"github.com/oldgattsu/diplom2/internal/models"

	"go.uber.org/zap"
)

type addAuthorRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (app *Application) handlerAddAuthor(rw http.ResponseWriter, req *http.Request) {
	r := addAuthorRequest{}

	errUnmarshal := json.NewDecoder(req.Body).Decode(&r)
	if errUnmarshal != nil {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}

	if r.Name == "" {
		http.Error(rw, "name is empty", http.StatusBadRequest)
		return
	}

	a := &models.Author{
		Name:        r.Name,
		Description: r.Description,
	}

	authorID, errAddAuthor := app.store.AddAuthor(req.Context(), a)
	if errAddAuthor != nil {
		app.logger.Error("error add book", zap.Error(errAddAuthor))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	a.ID = authorID

	resp, errMarshal := json.Marshal(a)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

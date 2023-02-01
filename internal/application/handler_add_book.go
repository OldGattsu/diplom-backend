package application

import (
	"encoding/json"
	"fmt"
	"github.com/oldgattsu/diplom2/internal/imageUploader"
	"github.com/oldgattsu/diplom2/internal/models"
	"github.com/oldgattsu/diplom2/internal/storage"
	"go.uber.org/zap"
	"net/http"
)

type addBookRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"'`
	Authors     []int  `json:"authors"`
	Poster      string `json:"poster"`
}

func (app *Application) handlerAddBook(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler add book")

	r := addBookRequest{}

	errUnmarshal := json.NewDecoder(req.Body).Decode(&r)

	imageURL, uploadImageError := imageUploader.Load(r.Poster)
	if uploadImageError != nil {
		fmt.Printf("error: %v", uploadImageError)
		return
	}

	if errUnmarshal != nil {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}

	if r.Name == "" {
		http.Error(rw, "name is empty", http.StatusBadRequest)
		return
	}

	bq := &storage.AddBookQuery{
		Name:        r.Name,
		Authors:     r.Authors,
		Description: r.Description,
		Poster:      imageURL,
	}
	bookID, errAddBook := app.store.AddBook(req.Context(), bq)
	if errAddBook != nil {
		app.logger.Error("error add book", zap.Error(errAddBook))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	b := &models.Book{
		ID:          bookID,
		Name:        r.Name,
		Description: r.Description,
		Poster:      imageURL,
	}
	resp, errMarshal := json.Marshal(b)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

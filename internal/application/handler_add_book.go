package application

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/oldgattsu/diplom2/internal/config"
	"github.com/oldgattsu/diplom2/internal/models"
	"github.com/oldgattsu/diplom2/internal/storage"
	"go.uber.org/zap"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
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

	index := strings.Index(r.Poster, "base64,")
	posterBase64, _ := b64.StdEncoding.DecodeString(r.Poster[index+7:])
	reader := bytes.NewReader(posterBase64)

	image, format, imageDecodeErr := image.Decode(reader)
	if imageDecodeErr != nil {
		fmt.Printf("error: %v", imageDecodeErr)
		return
	}
	fmt.Printf("image format: %v", format)

	dir := "upload/images"
	uniqueFileName := uuid.New().String()
	fullFileName := fmt.Sprintf("%s/%s.%s", dir, uniqueFileName, format)

	f, openFilerErr := os.OpenFile(fullFileName, os.O_WRONLY|os.O_CREATE, 0777)
	if openFilerErr != nil {
		fmt.Printf("error: %v", openFilerErr)
		return
	}

	switch format {
	case "png":
		png.Encode(f, image)
	case "jpeg":
		jpeg.Encode(f, image, nil)
	}

	cfg, errCfg := config.Load()
	if errCfg != nil {
		log.Printf("error load config, %v", errCfg)
	}
	imageURL := fmt.Sprintf("%s/%s", cfg.Address, fullFileName)

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

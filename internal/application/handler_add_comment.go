package application

import (
	"encoding/json"
	"github.com/oldgattsu/diplom2/internal/models"
	"github.com/oldgattsu/diplom2/internal/storage"
	"net/http"

	"go.uber.org/zap"
)

type addCommentRequest struct {
	UserID int    `json:"user_id"`
	BookID int    `json:"book_id"`
	Text   string `json:"text"`
}

func (app *Application) handlerAddComment(rw http.ResponseWriter, req *http.Request) {
	app.logger.Debug("handler add comment")

	r := addCommentRequest{}

	errUnmarshal := json.NewDecoder(req.Body).Decode(&r)
	if errUnmarshal != nil {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}

	if r.Text == "" {
		http.Error(rw, "text is empty", http.StatusBadRequest)
		return
	}

	cq := &storage.AddCommentQuery{
		UserID: r.UserID,
		BookID: r.BookID,
		Text:   r.Text,
	}
	commentID, errAddComment := app.store.AddComment(req.Context(), cq)
	if errAddComment != nil {
		app.logger.Error("error add comment", zap.Error(errAddComment))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	c := &models.Comment{
		ID:     commentID,
		UserId: r.UserID,
		BookID: r.BookID,
		Text:   r.Text,
	}
	resp, errMarshal := json.Marshal(c)
	if errMarshal != nil {
		app.logger.Error("error marshal response", zap.Error(errMarshal))
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(resp)
}

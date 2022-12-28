package application

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/storage"
)

type contextKey string

var (
	contextKeyUser contextKey = "user"
)

func (app *Application) middlewareAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		app.logger.Debug("middlewareAuth")

		token := req.Header.Get("Authorization")
		if token == "" {
			http.Error(rw, "", http.StatusForbidden)
			return
		}

		user, errGetUser := app.store.GetUserByToken(req.Context(), token)
		if errGetUser != nil {
			if errors.Is(errGetUser, storage.ErrUserNotFound) {
				http.Error(rw, "", http.StatusForbidden)
				return
			}
			app.logger.Error("error get user", zap.Error(errGetUser))
			http.Error(rw, "", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(req.Context(), contextKeyUser, user)

		handler.ServeHTTP(rw, req.WithContext(ctx))
	})
}

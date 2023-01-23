package application

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/models"
	"github.com/oldgattsu/diplom2/internal/storage"
)

type Application struct {
	store  *storage.Storage
	logger *zap.Logger
}

func New(store *storage.Storage, logger *zap.Logger) *Application {
	app := &Application{
		store:  store,
		logger: logger,
	}

	return app
}

func (app *Application) Run(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, ln net.Listener) {
	defer wg.Done()

	r := chi.NewRouter()
	r.Use(app.middlewareResponseHeaders)

	r.Post("/login", app.handlerLogin)

	r.Group(func(r chi.Router) {
		r.Use(app.middlewareAuth)
		r.Get("/book/{id}", app.handlerGetBook)
		r.Get("/books", app.handlerGetBooks)
		r.Get("/author/{id}", app.handlerGetAuthor)
		r.Get("/authors", app.handlerGetAuthors)

		r.Group(func(r chi.Router) {
			r.Use(app.middlewareIsAdmin)
			r.Post("/book", app.handlerAddBook)
			r.Post("/author", app.handlerAddAuthor)
		})

	})

	server := http.Server{
		Handler: r,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()

		app.logger.Info("start main server", zap.String("address", ln.Addr().String()))

		errServe := server.Serve(ln)
		if errServe != nil {
			if !errors.Is(errServe, http.ErrServerClosed) {
				app.logger.Error("error serve main server", zap.Error(errServe))
			}
		}
	}()

	<-ctx.Done()

	app.logger.Info("stop main server")

	errShutdown := server.Shutdown(context.Background())
	if errShutdown != nil {
		app.logger.Error("error shutdown server", zap.Error(errShutdown))
	}
}

func getUserFromContext(ctx context.Context) *models.User {
	return ctx.Value(contextKeyUser).(*models.User)
}

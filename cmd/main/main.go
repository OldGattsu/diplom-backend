package main

import (
	"context"
	"fmt"
	"github.com/oldgattsu/diplom2/internal/storage"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"github.com/oldgattsu/diplom2/internal/application"
	"github.com/oldgattsu/diplom2/internal/config"
)

func main() {
	cfg, errCfg := config.Load()
	if errCfg != nil {
		log.Printf("error load config, %v", errCfg)
		os.Exit(1)
	}

	logger, errLogger := createLogger(cfg.Debug)
	if errLogger != nil {
		log.Printf("error create logger, %v", errLogger)
		os.Exit(1)
	}

	errRun := run(cfg, logger)
	if errRun != nil {
		log.Printf("error run application, %v", errRun)
		os.Exit(1)
	}
}

func run(cfg *config.Config, logger *zap.Logger) error {
	logger.Info("run")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db, errConnectDB := pgxpool.Connect(ctx, buildPGConnectionString(cfg.Postgres))
	if errConnectDB != nil {
		return fmt.Errorf("error connect to db, %w", errConnectDB)
	}
	defer db.Close()

	store := storage.New(db, logger)

	ln, errLn := net.Listen("tcp", cfg.Address)
	if errLn != nil {
		return fmt.Errorf("error listen address, %w", errLn)
	}
	defer ln.Close()

	wg := sync.WaitGroup{}

	app := application.New(store, logger)

	wg.Add(1)
	go app.Run(ctx, cancel, &wg, ln)

	<-ctx.Done()

	wg.Wait()

	logger.Info("stop")
	return nil
}

func buildPGConnectionString(cfg config.Postgres) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&sslrootcert=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
		cfg.SSLCertPath)
}

package app

import (
	"context"
	"currency/internal/db"
	third_api "currency/internal/third-api"
	"currency/pkg/config"
	zapLogger "currency/pkg/logger"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "currency/internal/docs"

	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type App struct {
	mysql  db.Repository
	logger *zap.SugaredLogger
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{
		logger: zapLogger.InitZap(),
	}

	err := a.initDeps(ctx)

	if err != nil {
		return nil, err
	}

	a.logger.Infoln("App inited")
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initClients,
	}
	var err error

	for _, f := range inits {
		err = f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) fetchCurrenciesDaily() {
	currencies, err := third_api.FetchCurrencyRates()
	if err != nil {
		a.logger.Fatalf("Some troubles with fetching data: %w\n Restarting APP...", err)
	}

	err = a.mysql.InsertCurrencies(currencies)

	if err != nil {
		a.logger.Fatalf("Some troubles with inserting data: %w\n Restarting APP...", err)
	}
}

func (app *App) GoFetchCurrenciesDaily(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)

	app.fetchCurrenciesDaily()
	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				app.logger.Info("Stopping fetchCurrencyRatesDaily task")
				return
			case <-ticker.C:
				app.fetchCurrenciesDaily()
			}
		}
	}()
}

func (app *App) StartServer(ctx context.Context) {
	port := os.Getenv("SERVER_PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	http.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")
		if date == "" {
			app.HandleGetAllCurrencies(w, r)
		} else {
			app.HandleGetCurrenciesByDate(w, r)
		}
	})

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	go func() {
		app.logger.Infof("server listening at %v\n", port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			app.logger.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()

	go func() {
		i := 0
		for {
			app.logger.Infof("Shutting down server... %ds ", i)
			time.Sleep(time.Second)
			i++
		}
	}()

	if err := server.Shutdown(context.Background()); err != nil {
		app.logger.Fatalf("server shutdown failed: %v", err)
	}

	app.logger.Infoln("Server gracefully stopped")
}

func (a *App) initConfig(ctx context.Context) error {
	nodeEnv := os.Getenv("NODE_ENV")
	if nodeEnv != "DOCKER" {
		err := config.Load(".env")
		if err != nil {
			return err
		}

		a.logger.Infoln("Env loaded")
		return nil
	}

	return nil
}

func (a *App) initClients(ctx context.Context) error {
	var err error
	a.mysql, err = db.NewMySQLRepository()

	if err != nil {
		return err
	}

	if a.mysql == nil {
		return fmt.Errorf("clients are not inited")
	}

	a.logger.Infoln("Clients inited")
	return nil
}

package main

import (
	"context"
	app "currency/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize APP: %v", err)
	}

	a.GoFetchCurrenciesDaily(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigs
		cancel()
	}()

	a.StartServer(ctx)
}

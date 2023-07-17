package main

import (
	"os"
	"os/signal"
	"ticker/app/ticker"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err)
	}
	projectId := os.Getenv("PROJECT_ID")
	tickerManager, err := ticker.NewTickerManager(projectId)
	if err != nil {
		log.Fatal().Err(err)
	}

	timeTicker := time.NewTicker(10 * time.Minute)
	defer timeTicker.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	tickerManager.UpdateTickerPrices()
	for {
		select {
		case <-timeTicker.C:
			tickerManager.UpdateTickerPrices()
		case <-quit:
			log.Print("Shutting down gracefully...")
			os.Exit(0)
		}
	}
}

package ticker

import (
	"ticker/app/firebase"
	"ticker/app/yahoo"
	"time"

	"github.com/rs/zerolog/log"
)

type TickerManager struct {
	firebaseClient *firebase.FirebaseClient
}

func NewTickerManager(projectId string) (*TickerManager, error) {
	firebaseClient, err := firebase.NewFirebaseClient(projectId)
	if err != nil {
		return nil, err
	}
	return &TickerManager{firebaseClient: firebaseClient}, nil
}

func (tm *TickerManager) UpdateTickerPrices() {
	log.Print("updating prices")
	tickers, err := tm.firebaseClient.GetUqTickers()
	if err != nil {
		log.Printf("Error getting tickers: %v", err)
		return
	}

	for _, ticker := range tickers {
		price, err := yahoo.GetStockOptionPrice(ticker)
		if err != nil {
			log.Printf("Error getting price for ticker %s: %v", ticker, err)
			continue
		}

		tick := &firebase.Tick{
			TickerName:   ticker,
			CurrentPrice: price,
			Timestamp:    time.Now(),
		}

		log.Printf("tick: %+v", tick)

		err = tm.firebaseClient.AddNewTick(tick)
		if err != nil {
			log.Printf("Error adding tick for ticker %s: %v", ticker, err)
		}
	}
	log.Print("finished updating prices, going to sleep")
}

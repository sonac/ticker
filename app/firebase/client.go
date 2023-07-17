package firebase

import (
	"context"
	"ticker/app/utils"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	client *firestore.Client
}

type Tick struct {
	TickerName   string
	CurrentPrice float64
	Timestamp    time.Time
}

func NewFirebaseClient(projectId string) (*FirebaseClient, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("/Users/sonac/.config/gcc/firebase.json")
	client, err := firestore.NewClient(ctx, "hilmm-62092", opt)

	if err != nil {
		return nil, err
	}

	return &FirebaseClient{client: client}, nil
}

func (fc *FirebaseClient) GetUqTickers() ([]string, error) {
	iter := fc.client.Collection("investments").Documents(context.Background())
	tickers := []string{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}
		data := doc.Data()
		tickers = append(tickers, data["ticker"].(string))
	}
	utils.RemoveDuplicates(&tickers)

	return tickers, nil
}

func (fc *FirebaseClient) AddNewTick(tick *Tick) error {
	ctx := context.Background()

	_, _, err := fc.client.Collection("ticks").Add(ctx, map[string]interface{}{
		"TickerName":   tick.TickerName,
		"CurrentPrice": tick.CurrentPrice,
		"Timestamp":    tick.Timestamp,
	})

	if err != nil {
		return err
	}

	return nil
}

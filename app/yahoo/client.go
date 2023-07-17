package yahoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	Client HttpClient
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OptionsData struct {
	OptionChain struct {
		Result []struct {
			Quote struct {
				RegularMarketPrice float64 `json:"regularMarketPrice"`
			} `json:"quote"`
		} `json:"result"`
	} `json:"optionChain"`
}

func init() {
	Client = &http.Client{}
}

func GetStockOptionPrice(tickerName string) (float64, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/options/%s", tickerName), nil)

	resp, err := Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data OptionsData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	if len(data.OptionChain.Result) == 0 {
		return 0, fmt.Errorf("no option data available for %s", tickerName)
	}

	return data.OptionChain.Result[0].Quote.RegularMarketPrice, nil
}

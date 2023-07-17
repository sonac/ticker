package yahoo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockHttpClient struct {
	ResponseBody string
	ResponseCode int
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.ResponseCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(c.ResponseBody)),
	}, nil
}

func TestGetStockOptionPrice(t *testing.T) {
	Client = &MockHttpClient{
		ResponseCode: 200,
		ResponseBody: `{
			"optionChain": {
				"result": [{
						"quote": {
							"regularMarketPrice": 100.0
						}
				}]
			}
		}`,
	}

	stockName := "AAPL"
	price, err := GetStockOptionPrice(stockName)
	if err != nil {
		t.Fatalf("Expected no error, but got: %s", err)
	}

	if price != 100.0 {
		t.Fatalf("Expected price to be 100.0, but got: %.2f", price)
	}
}

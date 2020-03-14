package wazirx

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// MarketTrade holds trade history of a perticular market
type MarketTrade struct {
	ID        int64       `json:"id"`
	Price     string      `json:"price"`
	Volume    string      `json:"volume"`
	Funds     string      `json:"funds"`
	Market    string      `json:"market"`
	CreatedAt time.Time   `json:"created_at"`
	Side      interface{} `json:"side"`
}

// MarketTrade returs trade history of a market.
// Pass any market to get the desired trade history.
func (c Client) MarketTrade(ctx context.Context, market string) (data []MarketTrade, err error) {
	if market == "" {
		return nil, errors.New("market required")
	}

	endpoint := "/api/v2/trades" + "?market=" + market

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	return data, nil
}

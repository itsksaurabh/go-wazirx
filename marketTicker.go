package wazirx

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// TickerData holds active market data with all ticker related values
type TickerData struct {
	// ticker code of base market
	BaseUnit string `json:"base_unit"`
	// ticker code of quote asset
	QuoteUnit string `json:"quote_unit"`
	// 24 hrs lowest price of base asset
	Low float64 `json:"low"`
	// 24 hrs highest price of base asset
	High float64 `json:"high"`
	// Last traded price in current market
	Last float64 `json:"last"`
	Type string  `json:"type"`
	// Market Open price 24hrs ago
	Open interface{} `json:"open"`
	// Last 24hrs traded volume
	Volume float64 `json:"volume"`
	// Top ask order price
	Sell float64 `json:"sell"`
	// Top bid order price
	Buy float64 `json:"buy"`
	// Timestamp when ticker information is fetched
	At time.Time `json:"at"`
	// Display text of market
	Name string `json:"name"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// coverts Timestamp from `int` to unix timestamp
func (t *TickerData) UnmarshalJSON(data []byte) error {
	type Alias TickerData
	aux := &struct {
		At     int64  `json:"at"`
		Low    string `json:"low,omitempty"`
		High   string `json:"high,omitempty"`
		Last   string `json:"last,omitempty"`
		Volume string `json:"volume,omitempty"`
		Sell   string `json:"sell,omitempty"`
		Buy    string `json:"buy,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.Low, _ = strconv.ParseFloat(aux.Low, 64)
	t.High, _ = strconv.ParseFloat(aux.High, 64)
	t.Last, _ = strconv.ParseFloat(aux.Last, 64)
	t.Volume, _ = strconv.ParseFloat(aux.Volume, 64)
	t.Sell, _ = strconv.ParseFloat(aux.Sell, 64)
	t.Buy, _ = strconv.ParseFloat(aux.Buy, 64)
	t.At = time.Unix(aux.At, 0)
	return nil
}

// MarketTicker returs  the latest market heart-beat for all the markets for the last 24hrs.
func (c Client) MarketTicker(ctx context.Context) (data map[string]TickerData, err error) {
	endpoint := "/api/v2/tickers"

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	return data, nil
}

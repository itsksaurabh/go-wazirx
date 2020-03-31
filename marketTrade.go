package wazirx

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// MarketTrade holds trade history of a perticular market
type MarketTrade struct {
	ID        int64       `json:"id"`
	Price     float64     `json:"price"`
	Volume    float64     `json:"volume"`
	Funds     float64     `json:"funds"`
	Market    string      `json:"market"`
	CreatedAt time.Time   `json:"created_at"`
	Side      interface{} `json:"side"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It coverts some string type numeric to numeric types.
func (t *MarketTrade) UnmarshalJSON(data []byte) error {
	type Alias MarketTrade
	aux := &struct {
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Funds  string `json:"funds"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.Price, _ = strconv.ParseFloat(aux.Price, 64)
	t.Volume, _ = strconv.ParseFloat(aux.Volume, 64)
	t.Funds, _ = strconv.ParseFloat(aux.Funds, 64)
	return nil
}

// MarketTrade returs trade history of a market.
// Pass any market to get the desired trade history.
func (c Client) MarketTrade(ctx context.Context, market string) (data []MarketTrade, err error) {
	if market == "" {
		return nil, errors.New("market required")
	}

	endpoint := "/api/v2/trades" + "?market=" + market

	if err := c.makeGetRequest(ctx, endpoint, &data); err != nil {
		return nil, err
	}
	return data, nil
}

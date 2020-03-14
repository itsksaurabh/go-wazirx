package wazirx

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// MakerTaker holds order's maker-taker
type MakerTaker struct {
	Maker float64 `json:"maker"`
	Taker float64 `json:"taker"`
}

// Fee holds bid and ask order's maker-taker fee percentage
type Fee struct {
	Bid MakerTaker `json:"bid"`
	Ask MakerTaker `json:"ask"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Since `Fee` is returned as an map so it converts it
// into struct for better access
func (f *Fee) UnmarshalJSON(data []byte) error {
	var aux map[string]interface{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	mapstructure.Decode(aux, f)
	return nil
}

// Market holds market related data
type Market struct {
	// Ticker code of base asset
	BaseMarket string `json:"baseMarket"`
	// Ticker code of quote asset
	QuoteMarket string `json:"quoteMarket"`
	// Minimum buy amount of base asset
	MinBuyAmount float64 `json:"minBuyAmount,omitempty"`
	// Minumum sell amount of base asset
	MinSellAmount float64 `json:"minSellAmount,omitempty"`
	// Maximum precision of base asset, this the decimal point.
	BasePrecision int `json:"basePrecision,omitempty"`
	// Maximum precision of quote asset
	QuotePrecision int `json:"quotePrecision,omitempty"`
	// This defines the current state of the market. This can be active or suspended
	Status string `json:"status"`
	// JSON Object consists of bid and ask order's maker-taker fee percentage
	Fee Fee `json:"fee,omitempty"`
	// 24 hrs lowest price of base asset
	Low string `json:"low,omitempty"`
	// 24 hrs highest price of base asset
	High string `json:"high,omitempty"`
	// Last traded price in current market
	Last string `json:"last,omitempty"`
	// This defines the type of market, currently we have SPOT and P2P
	Type string `json:"type"`
	// Market Open price 24hrs ago
	Open float64 `json:"open,omitempty"`
	// Last 24hrs traded volume
	Volume string `json:"volume,omitempty"`
	// Top ask order price
	Sell string `json:"sell,omitempty"`
	// Top bid order price
	Buy                string    `json:"buy,omitempty"`
	At                 time.Time `json:"at,omitempty"`
	MaxBuyAmount       int       `json:"maxBuyAmount,omitempty"`
	MinBuyVolume       int       `json:"minBuyVolume,omitempty"`
	MaxBuyVolume       int       `json:"maxBuyVolume,omitempty"`
	FeePercentOnProfit float64   `json:"feePercentOnProfit,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// coverts Timestamp from `int` to unix timestamp
func (m *Market) UnmarshalJSON(data []byte) error {
	type Alias Market
	aux := &struct {
		At int64 `json:"at,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	m.At = time.Unix(aux.At, 0)
	return nil
}

// Asset holds asset related data
type Asset struct {
	// asset code
	Type string `json:"type"`
	// Display name of asset
	Name string `json:"name"`
	// Denotes whether deposit is enabled or disabled
	Deposit string `json:"deposit"`
	// Denotes whether withdrawal is enabled or disabled
	Withdrawal  string `json:"withdrawal"`
	ListingType string `json:"listingType"`
	Category    string `json:"category"`
	// Withdrawal fee of asset
	WithdrawFee float64 `json:"withdrawFee,omitempty"`
	// Minimum withdrawal amount in a single transaction
	MinWithdrawAmount float64 `json:"minWithdrawAmount,omitempty"`
	// Maximum withdrawal amount in a single transaction
	MaxWithdrawAmount float64 `json:"maxWithdrawAmount,omitempty"`
	// This is the min Deposit amount that will be accepted as deposit
	MinDepositAmount float64 `json:"minDepositAmount,omitempty"`
	// Is the min number of block height needed to confirm a block chain deposit transaction.
	Confirmations int `json:"confirmations,omitempty"`
}

// MarketStatus holds the response from endpoint /api/v2/market-status
type MarketStatus struct {
	Markets []Market `json:"markets"`
	Assets  []Asset  `json:"assets"`
}

// MarketStatus returs overview of markets and assets
func (c Client) MarketStatus(ctx context.Context) (data MarketStatus, err error) {
	endpoint := "/api/v2/market-status"

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return MarketStatus{}, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return MarketStatus{}, errors.Wrap(err, "request failed")
	}
	return data, nil
}

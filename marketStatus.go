package wazirx

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
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
	Low float64 `json:"low,omitempty"`
	// 24 hrs highest price of base asset
	High float64 `json:"high,omitempty"`
	// Last traded price in current market
	Last float64 `json:"last,omitempty"`
	// This defines the type of market, currently we have SPOT and P2P
	Type string `json:"type"`
	// Market Open price 24hrs ago
	Open float64 `json:"open,omitempty"`
	// Last 24hrs traded volume
	Volume float64 `json:"volume,omitempty"`
	// Top ask order price
	Sell float64 `json:"sell,omitempty"`
	// Top bid order price
	Buy                float64   `json:"buy,omitempty"`
	At                 time.Time `json:"at,omitempty"`
	MaxBuyAmount       float64   `json:"maxBuyAmount,omitempty"`
	MinBuyVolume       float64   `json:"minBuyVolume,omitempty"`
	MaxBuyVolume       float64   `json:"maxBuyVolume,omitempty"`
	FeePercentOnProfit float64   `json:"feePercentOnProfit,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// coverts some string type numeric to numeric type
// coverts Timestamp from `int` to unix timestamp
func (m *Market) UnmarshalJSON(data []byte) error {
	type Alias Market
	aux := &struct {
		At     int64  `json:"at,omitempty"`
		Low    string `json:"low,omitempty"`
		High   string `json:"high,omitempty"`
		Last   string `json:"last,omitempty"`
		Volume string `json:"volume,omitempty"`
		Sell   string `json:"sell,omitempty"`
		Buy    string `json:"buy,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	m.Low, _ = strconv.ParseFloat(aux.Low, 64)
	m.High, _ = strconv.ParseFloat(aux.High, 64)
	m.Last, _ = strconv.ParseFloat(aux.Last, 64)
	m.Volume, _ = strconv.ParseFloat(aux.Volume, 64)
	m.Sell, _ = strconv.ParseFloat(aux.Sell, 64)
	m.Buy, _ = strconv.ParseFloat(aux.Buy, 64)
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

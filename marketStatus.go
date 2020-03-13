package wazirx

import (
	"context"
	"encoding/json"
	"net/http"

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
	BaseMarket         string  `json:"baseMarket"`
	QuoteMarket        string  `json:"quoteMarket"`
	MinBuyAmount       float64 `json:"minBuyAmount,omitempty"`
	MinSellAmount      float64 `json:"minSellAmount,omitempty"`
	BasePrecision      int     `json:"basePrecision,omitempty"`
	QuotePrecision     int     `json:"quotePrecision,omitempty"`
	Status             string  `json:"status"`
	Fee                Fee     `json:"fee,omitempty"`
	Low                string  `json:"low,omitempty"`
	High               string  `json:"high,omitempty"`
	Last               string  `json:"last,omitempty"`
	Type               string  `json:"type"`
	Open               float64 `json:"open,omitempty"`
	Volume             string  `json:"volume,omitempty"`
	Sell               string  `json:"sell,omitempty"`
	Buy                string  `json:"buy,omitempty"`
	At                 int     `json:"at,omitempty"`
	MaxBuyAmount       int     `json:"maxBuyAmount,omitempty"`
	MinBuyVolume       int     `json:"minBuyVolume,omitempty"`
	MaxBuyVolume       int     `json:"maxBuyVolume,omitempty"`
	FeePercentOnProfit float64 `json:"feePercentOnProfit,omitempty"`
}

// Asset holds asset related data
type Asset struct {
	Type              string  `json:"type"`
	Name              string  `json:"name"`
	Deposit           string  `json:"deposit"`
	Withdrawal        string  `json:"withdrawal"`
	ListingType       string  `json:"listingType"`
	Category          string  `json:"category"`
	WithdrawFee       float64 `json:"withdrawFee,omitempty"`
	MinWithdrawAmount float64 `json:"minWithdrawAmount,omitempty"`
	MaxWithdrawAmount float64 `json:"maxWithdrawAmount,omitempty"`
	MinDepositAmount  float64 `json:"minDepositAmount,omitempty"`
	Confirmations     int     `json:"confirmations,omitempty"`
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

package wazix

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// Bid holds maker-taker bid
type Bid struct {
	Maker float64 `json:"maker"`
	Taker float64 `json:"taker"`
}

// Ask holds maker-taker ask
type Ask struct {
	Maker float64 `json:"maker"`
	Taker float64 `json:"taker"`
}

// Fee holds bid and ask order's maker-taker fee percentage
type Fee struct {
	Bid Bid `json:"bid"`
	Ask Ask `json:"ask"`
}

// Market holds market related data
type Market struct {
	BaseMarket         string  `json:"baseMarket"`
	QuoteMarket        string  `json:"quoteMarket"`
	MinBuyAmount       int     `json:"minBuyAmount,omitempty"`
	MinSellAmount      int     `json:"minSellAmount,omitempty"`
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
	Type              string `json:"type"`
	Name              string `json:"name"`
	Deposit           string `json:"deposit"`
	Withdrawal        string `json:"withdrawal"`
	ListingType       string `json:"listingType"`
	Category          string `json:"category"`
	WithdrawFee       int    `json:"withdrawFee,omitempty"`
	MinWithdrawAmount int    `json:"minWithdrawAmount,omitempty"`
	MaxWithdrawAmount int    `json:"maxWithdrawAmount,omitempty"`
	MinDepositAmount  int    `json:"minDepositAmount,omitempty"`
	Confirmations     int    `json:"confirmations,omitempty"`
}

// MarketStatus holds the response from endpoint /api/v2/market-status
type MarketStatus struct {
	Markets []Market `json:"markets"`
	Assets  []Asset  `json:"assets"`
}

// MarketStatus returs overview of markets and assets
func (c Client) MarketStatus(ctx context.Context) (data *MarketStatus, err error) {
	endpoint := "/listings/latest"

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), data); err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	return data, nil
}

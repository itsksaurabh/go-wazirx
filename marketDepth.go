package wazirx

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// MarketDepth holds orderbook data of a perticular market
type MarketDepth struct {
	Timestamp int `json:"timestamp"`
	// list order's asks
	Asks [][]string `json:"asks"`
	// list order's bids
	Bids [][]string `json:"bids"`
}

// MarketDepth returs orderbook of any market.
// Pass any market to get the desired order book.
func (c Client) MarketDepth(ctx context.Context, market string) (data MarketDepth, err error) {
	endpoint := "/api/v2/depth" + "?market=" + market

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return MarketDepth{}, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return MarketDepth{}, errors.Wrap(err, "request failed")
	}
	return data, nil
}

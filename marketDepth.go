package wazirx

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// PriceVolume holds price and volume
type PriceVolume struct {
	PRICE  string
	VOLUME string
}

// MarketDepth holds orderbook data of a perticular market
type MarketDepth struct {
	Timestamp time.Time `json:"timestamp"`
	// list order's asks
	Asks []PriceVolume `json:"asks"`
	// list order's bids
	Bids []PriceVolume `json:"bids"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// converts slices into struct for better access and
// coverts Timestamp from `int` to unix timestamp
func (d *MarketDepth) UnmarshalJSON(data []byte) error {
	aux := struct {
		Timestamp int64      `json:"timestamp"`
		Asks      [][]string `json:"asks"`
		Bids      [][]string `json:"bids"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var tmp PriceVolume
	for _, ask := range aux.Asks {
		tmp.PRICE = ask[0]
		tmp.VOLUME = ask[1]
		d.Asks = append(d.Asks, tmp)
	}

	for _, bid := range aux.Bids {
		tmp.PRICE = bid[0]
		tmp.VOLUME = bid[1]
		d.Bids = append(d.Bids, tmp)
	}

	d.Timestamp = time.Unix(aux.Timestamp, 0)
	return nil
}

// MarketDepth returs orderbook of any market.
// Pass any market to get the desired order book.
func (c Client) MarketDepth(ctx context.Context, market string) (data MarketDepth, err error) {
	if market == "" {
		return MarketDepth{}, errors.New("market required")
	}

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

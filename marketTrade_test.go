package wazirx_test

import (
	"context"
	"testing"

	wazirx "github.com/itsksaurabh/go-wazirx"
)

func TestMarketTrade(t *testing.T) {
	c := wazirx.Client{}
	market := "btcusdt"
	_, err := c.MarketTrade(context.Background(), market)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

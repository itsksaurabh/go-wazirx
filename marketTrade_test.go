package wazirx_test

import (
	"context"
	"testing"
)

func TestMarketTrade(t *testing.T) {
	market := "btcusdt"
	_, err := testClient(t).MarketTrade(context.Background(), market)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

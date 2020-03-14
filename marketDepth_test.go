package wazirx_test

import (
	"context"
	"testing"
)

func TestMarketDepth(t *testing.T) {
	market := "btcusdt"
	_, err := client(t).MarketDepth(context.Background(), market)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

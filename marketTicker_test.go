package wazirx_test

import (
	"context"
	"testing"
)

func TestMarketTicker(t *testing.T) {
	_, err := client(t).MarketTicker(context.Background())
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

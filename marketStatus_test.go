package wazirx_test

import (
	"context"
	"testing"
)

func TestMarketStatus(t *testing.T) {
	_, err := client(t).MarketStatus(context.Background())
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

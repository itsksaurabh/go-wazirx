package wazirx_test

import (
	"context"
	"testing"
)

func TestMarketStatus(t *testing.T) {
	_, err := testClient(t).MarketStatus(context.Background())
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

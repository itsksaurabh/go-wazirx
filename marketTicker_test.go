package wazirx_test

import (
	"context"
	"testing"

	wazirx "github.com/itsksaurabh/go-wazirx"
)

func TestMarketTicker(t *testing.T) {
	c := wazirx.Client{}
	_, err := c.MarketTicker(context.Background())
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

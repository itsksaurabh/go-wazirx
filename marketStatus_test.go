package wazirx_test

import (
	"context"
	"testing"

	wazirx "github.com/itsksaurabh/go-wazirx"
)

func TestMarketStatus(t *testing.T) {
	c := wazirx.Client{}
	_, err := c.MarketStatus(context.Background())
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

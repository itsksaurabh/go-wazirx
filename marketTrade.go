package wazirx

import (
	"time"
)

// MarketTrade holds trade history of a perticular market
type MarketTrade struct {
	ID        int64       `json:"id"`
	Price     string      `json:"price"`
	Volume    string      `json:"volume"`
	Funds     string      `json:"funds"`
	Market    string      `json:"market"`
	CreatedAt time.Time   `json:"created_at"`
	Side      interface{} `json:"side"`
}

package wazirx

// TickerData holds active market data with all ticker related values
type TickerData struct {
	// ticker code of base market
	BaseUnit string `json:"base_unit"`
	// ticker code of quote asset
	QuoteUnit string `json:"quote_unit"`
	// 24 hrs lowest price of base asset
	Low string `json:"low"`
	// 24 hrs highest price of base asset
	High string `json:"high"`
	// Last traded price in current market
	Last string `json:"last"`
	Type string `json:"type"`
	// Market Open price 24hrs ago
	Open float64 `json:"open"`
	// Last 24hrs traded volume
	Volume string `json:"volume"`
	// Top ask order price
	Sell string `json:"sell"`
	// Top bid order price
	Buy string `json:"buy"`
	// Timestamp when ticker information is fetched
	At int `json:"at"`
	// Display text of market
	Name string `json:"name"`
}

package wazirx

// TickerData holds active market data with all ticker related values
type TickerData struct {
	BaseUnit  string  `json:"base_unit"`
	QuoteUnit string  `json:"quote_unit"`
	Low       string  `json:"low"`
	High      string  `json:"high"`
	Last      string  `json:"last"`
	Type      string  `json:"type"`
	Open      float64 `json:"open"`
	Volume    string  `json:"volume"`
	Sell      string  `json:"sell"`
	Buy       string  `json:"buy"`
	At        int     `json:"at"`
	Name      string  `json:"name"`
}

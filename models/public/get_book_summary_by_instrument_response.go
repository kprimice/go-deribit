package public

type GetBookSummaryByInstrumentResponse []struct {
	AskPrice          float64 `json:"ask_price" mapstructure:"ask_price"`
	BaseCurrency      string  `json:"base_currency" mapstructure:"base_currency"`
	BidPrice          float64 `json:"bid_price" mapstructure:"bid_price"`
	CreationTimestamp int64   `json:"creation_timestamp" mapstructure:"creation_timestamp"`
	High              float64 `json:"high" mapstructure:"high"`
	InstrumentName    string  `json:"instrument_name" mapstructure:"instrument_name"`
	InterestRate      float64 `json:"interest_rate" mapstructure:"interest_rate"`
	Last              float64 `json:"last" mapstructure:"last"`
	Low               float64 `json:"low" mapstructure:"low"`
	MarkPrice         float64 `json:"mark_price" mapstructure:"mark_price"`
	MidPrice          float64 `json:"mid_price" mapstructure:"mid_price"`
	OpenInterest      float64 `json:"open_interest" mapstructure:"open_interest"`
	QuoteCurrency     string  `json:"quote_currency" mapstructure:"quote_currency"`
	UnderlyingIndex   string  `json:"underlying_index" mapstructure:"underlying_index"`
	UnderlyingPrice   float64 `json:"underlying_price" mapstructure:"underlying_price"`
	Volume            float64 `json:"volume" mapstructure:"volume"`
}

package public

type GetLastSettlementsByCurrencyRequest struct {
	Count    int64  `json:"count" mapstructure:"count"`
	Currency string `json:"currency" mapstructure:"currency"`
	Type     string `json:"type" mapstructure:"type"`
}

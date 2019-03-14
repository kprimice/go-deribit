package private

type BuyRequest struct {
	Amount         int64  `json:"amount" mapstructure:"amount"`
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
	Label          string `json:"label" mapstructure:"label"`
	Type           string `json:"type" mapstructure:"type"`
}

package private

type GetOrderHistoryByInstrumentRequest struct {
	Count          int64  `json:"count" mapstructure:"count"`
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
}

package private

type ClosePositionRequest struct {
	InstrumentName string  `json:"instrument_name" mapstructure:"instrument_name"`
	Price          float64 `json:"price" mapstructure:"price"`
	Type           string  `json:"type" mapstructure:"type"`
}

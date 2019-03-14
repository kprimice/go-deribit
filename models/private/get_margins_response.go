package private

type GetMarginsResponse struct {
	Buy      float64 `json:"buy" mapstructure:"buy"`
	MaxPrice float64 `json:"max_price" mapstructure:"max_price"`
	MinPrice float64 `json:"min_price" mapstructure:"min_price"`
	Sell     int64   `json:"sell" mapstructure:"sell"`
}

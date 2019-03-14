package private

type SubmitTransferToUserRequest struct {
	Amount      float64 `json:"amount" mapstructure:"amount"`
	Currency    string  `json:"currency" mapstructure:"currency"`
	Destination string  `json:"destination" mapstructure:"destination"`
}

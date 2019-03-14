package private

type CancelWithdrawalResponse struct {
	Address            string      `json:"address" mapstructure:"address"`
	Amount             float64     `json:"amount" mapstructure:"amount"`
	ConfirmedTimestamp interface{} `json:"confirmed_timestamp" mapstructure:"confirmed_timestamp"`
	CreatedTimestamp   int64       `json:"created_timestamp" mapstructure:"created_timestamp"`
	Currency           string      `json:"currency" mapstructure:"currency"`
	Fee                float64     `json:"fee" mapstructure:"fee"`
	ID                 int64       `json:"id" mapstructure:"id"`
	Priority           float64     `json:"priority" mapstructure:"priority"`
	State              string      `json:"state" mapstructure:"state"`
	TransactionID      interface{} `json:"transaction_id" mapstructure:"transaction_id"`
	UpdatedTimestamp   int64       `json:"updated_timestamp" mapstructure:"updated_timestamp"`
}

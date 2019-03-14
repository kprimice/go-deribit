package private

type GetDepositsResponse struct {
	Count int64 `json:"count" mapstructure:"count"`
	Data  []struct {
		Address           string `json:"address" mapstructure:"address"`
		Amount            int64  `json:"amount" mapstructure:"amount"`
		Currency          string `json:"currency" mapstructure:"currency"`
		ReceivedTimestamp int64  `json:"received_timestamp" mapstructure:"received_timestamp"`
		State             string `json:"state" mapstructure:"state"`
		TransactionID     string `json:"transaction_id" mapstructure:"transaction_id"`
		UpdatedTimestamp  int64  `json:"updated_timestamp" mapstructure:"updated_timestamp"`
	} `json:"data" mapstructure:"data"`
}

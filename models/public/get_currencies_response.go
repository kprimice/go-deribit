package public

type GetCurrenciesResponse []struct {
	CoinType                       string  `json:"coin_type" mapstructure:"coin_type"`
	Currency                       string  `json:"currency" mapstructure:"currency"`
	CurrencyLong                   string  `json:"currency_long" mapstructure:"currency_long"`
	DisabledDepositAddressCreation bool    `json:"disabled_deposit_address_creation" mapstructure:"disabled_deposit_address_creation"`
	FeePrecision                   int64   `json:"fee_precision" mapstructure:"fee_precision"`
	MinConfirmations               int64   `json:"min_confirmations" mapstructure:"min_confirmations"`
	MinWithdrawalFee               float64 `json:"min_withdrawal_fee" mapstructure:"min_withdrawal_fee"`
	WithdrawalFee                  float64 `json:"withdrawal_fee" mapstructure:"withdrawal_fee"`
	WithdrawalPriorities           []struct {
		Name  string  `json:"name" mapstructure:"name"`
		Value float64 `json:"value" mapstructure:"value"`
	} `json:"withdrawal_priorities" mapstructure:"withdrawal_priorities"`
}

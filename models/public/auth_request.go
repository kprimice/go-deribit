package public

type AuthRequest struct {
	ClientID     string `json:"client_id" mapstructure:"client_id"`
	ClientSecret string `json:"client_secret" mapstructure:"client_secret"`
	GrantType    string `json:"grant_type" mapstructure:"grant_type"`
}

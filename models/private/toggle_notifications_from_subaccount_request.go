package private

type ToggleNotificationsFromSubaccountRequest struct {
	Sid   int64 `json:"sid" mapstructure:"sid"`
	State bool  `json:"state" mapstructure:"state"`
}

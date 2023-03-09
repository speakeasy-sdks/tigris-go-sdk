package shared

type DelResponse struct {
	Message *string `json:"message,omitempty"`
	Status  *string `json:"status,omitempty"`
}

package shared

type DeleteIndexResponse struct {
	Message *string `json:"message,omitempty"`
	Status  *string `json:"status,omitempty"`
}

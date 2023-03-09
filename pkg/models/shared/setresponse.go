package shared

type SetResponse struct {
	Message *string `json:"message,omitempty"`
	Status  *string `json:"status,omitempty"`
}

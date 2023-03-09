package shared

type GetSetResponse struct {
	Message  *string `json:"message,omitempty"`
	OldValue *string `json:"old_value,omitempty"`
	Status   *string `json:"status,omitempty"`
}

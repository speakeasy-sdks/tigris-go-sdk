package shared

type CreateOrUpdateCollectionResponse struct {
	Message *string `json:"message,omitempty"`
	Status  *string `json:"status,omitempty"`
}

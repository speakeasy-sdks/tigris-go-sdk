package shared

type CreateNamespaceRequest struct {
	Code *int64  `json:"code,omitempty"`
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

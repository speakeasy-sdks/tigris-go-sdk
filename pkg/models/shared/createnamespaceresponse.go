package shared

type CreateNamespaceResponse struct {
	Message   *string        `json:"message,omitempty"`
	Namespace *NamespaceInfo `json:"namespace,omitempty"`
	Status    *string        `json:"status,omitempty"`
}

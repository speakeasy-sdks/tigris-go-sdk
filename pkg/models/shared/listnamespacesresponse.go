package shared

type ListNamespacesResponse struct {
	Namespaces []NamespaceInfo `json:"namespaces,omitempty"`
}

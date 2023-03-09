package shared

type CollectionInfo struct {
	Collection *string                `json:"collection,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

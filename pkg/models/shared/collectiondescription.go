package shared

type CollectionDescription struct {
	Collection *string                `json:"collection,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Schema     map[string]interface{} `json:"schema,omitempty"`
	Size       *int64                 `json:"size,omitempty"`
}

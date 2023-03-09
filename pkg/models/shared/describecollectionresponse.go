package shared

// DescribeCollectionResponse
// A detailed description of the collection. The description returns collection metadata and the schema.
type DescribeCollectionResponse struct {
	Collection *string                `json:"collection,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Schema     map[string]interface{} `json:"schema,omitempty"`
	Size       *int64                 `json:"size,omitempty"`
}

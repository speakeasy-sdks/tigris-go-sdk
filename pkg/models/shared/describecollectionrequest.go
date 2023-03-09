package shared

type DescribeCollectionRequest struct {
	Branch       *string                `json:"branch,omitempty"`
	Collection   *string                `json:"collection,omitempty"`
	Options      map[string]interface{} `json:"options,omitempty"`
	Project      *string                `json:"project,omitempty"`
	SchemaFormat *string                `json:"schema_format,omitempty"`
}

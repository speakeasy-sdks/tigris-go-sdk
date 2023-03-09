package shared

type DescribeDatabaseRequest struct {
	Branch       *string `json:"branch,omitempty"`
	Project      *string `json:"project,omitempty"`
	SchemaFormat *string `json:"schema_format,omitempty"`
}

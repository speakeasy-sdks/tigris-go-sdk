package shared

type ProjectInfo struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Project  *string                `json:"project,omitempty"`
}

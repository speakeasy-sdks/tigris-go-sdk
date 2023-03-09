package shared

type BranchInfo struct {
	Branch   *string                `json:"branch,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

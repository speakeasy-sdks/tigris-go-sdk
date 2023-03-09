package shared

type DropCollectionRequest struct {
	Branch  *string                `json:"branch,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

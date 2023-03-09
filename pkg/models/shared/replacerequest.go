package shared

type ReplaceRequest struct {
	Branch    *string                  `json:"branch,omitempty"`
	Documents []map[string]interface{} `json:"documents,omitempty"`
	Options   *ReplaceRequestOptions   `json:"options,omitempty"`
}

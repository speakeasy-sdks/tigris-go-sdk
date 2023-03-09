package shared

type InsertRequest struct {
	Branch    *string                  `json:"branch,omitempty"`
	Documents []map[string]interface{} `json:"documents,omitempty"`
	Options   *InsertRequestOptions    `json:"options,omitempty"`
}

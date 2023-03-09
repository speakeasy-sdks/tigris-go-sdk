package shared

type DeleteRequest struct {
	Branch  *string                `json:"branch,omitempty"`
	Filter  map[string]interface{} `json:"filter,omitempty"`
	Options *DeleteRequestOptions  `json:"options,omitempty"`
}

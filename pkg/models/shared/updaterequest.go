package shared

type UpdateRequest struct {
	Branch  *string                `json:"branch,omitempty"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
	Filter  map[string]interface{} `json:"filter,omitempty"`
	Options *UpdateRequestOptions  `json:"options,omitempty"`
}

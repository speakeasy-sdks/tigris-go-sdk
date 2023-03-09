package shared

type ReadRequest struct {
	Branch  *string                `json:"branch,omitempty"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
	Filter  map[string]interface{} `json:"filter,omitempty"`
	Options *ReadRequestOptions    `json:"options,omitempty"`
	Sort    *string                `json:"sort,omitempty"`
}

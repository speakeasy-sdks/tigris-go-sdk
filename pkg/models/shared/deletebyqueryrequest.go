package shared

type DeleteByQueryRequest struct {
	Filter  *string `json:"filter,omitempty"`
	Index   *string `json:"index,omitempty"`
	Project *string `json:"project,omitempty"`
}

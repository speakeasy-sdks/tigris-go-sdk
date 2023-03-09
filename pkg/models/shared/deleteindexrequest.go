package shared

type DeleteIndexRequest struct {
	Name    *string `json:"name,omitempty"`
	Project *string `json:"project,omitempty"`
}

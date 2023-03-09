package shared

type CreateByIDRequest struct {
	Document *string `json:"document,omitempty"`
	ID       *string `json:"id,omitempty"`
	Index    *string `json:"index,omitempty"`
	Project  *string `json:"project,omitempty"`
}

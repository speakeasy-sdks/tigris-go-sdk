package shared

type CreateOrUpdateIndexRequest struct {
	Name       *string `json:"name,omitempty"`
	OnlyCreate *bool   `json:"only_create,omitempty"`
	Project    *string `json:"project,omitempty"`
	Schema     *string `json:"schema,omitempty"`
}

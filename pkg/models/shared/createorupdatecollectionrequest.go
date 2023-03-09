package shared

type CreateOrUpdateCollectionRequest struct {
	Branch     *string                `json:"branch,omitempty"`
	OnlyCreate *bool                  `json:"only_create,omitempty"`
	Options    map[string]interface{} `json:"options,omitempty"`
	Schema     map[string]interface{} `json:"schema,omitempty"`
}

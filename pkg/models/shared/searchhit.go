package shared

type SearchHit struct {
	Data     map[string]interface{} `json:"data,omitempty"`
	Metadata *SearchHitMeta         `json:"metadata,omitempty"`
}

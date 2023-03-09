package shared

type IndexDoc struct {
	Doc      *string  `json:"doc,omitempty"`
	Metadata *DocMeta `json:"metadata,omitempty"`
}

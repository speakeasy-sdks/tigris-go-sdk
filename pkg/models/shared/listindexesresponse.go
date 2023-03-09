package shared

type ListIndexesResponse struct {
	Indexes []IndexInfo `json:"indexes,omitempty"`
}

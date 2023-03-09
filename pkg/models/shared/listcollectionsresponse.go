package shared

type ListCollectionsResponse struct {
	Collections []CollectionInfo `json:"collections,omitempty"`
}

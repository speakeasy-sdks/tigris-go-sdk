package shared

// DescribeDatabaseResponse
// A detailed description of the database and all the associated collections. Description of the collection includes schema details as well.
type DescribeDatabaseResponse struct {
	Branches    []string                `json:"branches,omitempty"`
	Collections []CollectionDescription `json:"collections,omitempty"`
	Metadata    map[string]interface{}  `json:"metadata,omitempty"`
	Size        *int64                  `json:"size,omitempty"`
}

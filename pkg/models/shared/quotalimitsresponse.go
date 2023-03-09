package shared

// QuotaLimitsResponse
// Contains current quota limits
type QuotaLimitsResponse struct {
	ReadUnits   *int64 `json:"ReadUnits,omitempty"`
	StorageSize *int64 `json:"StorageSize,omitempty"`
	WriteUnits  *int64 `json:"WriteUnits,omitempty"`
}

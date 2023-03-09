package shared

// QuotaUsageResponse
// Contains current quota usage
type QuotaUsageResponse struct {
	ReadUnits            *int64 `json:"ReadUnits,omitempty"`
	ReadUnitsThrottled   *int64 `json:"ReadUnitsThrottled,omitempty"`
	StorageSize          *int64 `json:"StorageSize,omitempty"`
	StorageSizeThrottled *int64 `json:"StorageSizeThrottled,omitempty"`
	WriteUnits           *int64 `json:"WriteUnits,omitempty"`
	WriteUnitsThrottled  *int64 `json:"WriteUnitsThrottled,omitempty"`
}

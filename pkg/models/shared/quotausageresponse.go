// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// QuotaUsageResponse - Contains current quota usage
type QuotaUsageResponse struct {
	// Number of read units used per second
	ReadUnits *int64 `json:"ReadUnits,omitempty"`
	// Number of read units throttled per second. Units which was rejected with "resource exhausted error".
	ReadUnitsThrottled *int64 `json:"ReadUnitsThrottled,omitempty"`
	// Number of bytes stored
	StorageSize *int64 `json:"StorageSize,omitempty"`
	// Number of bytes throttled. Number of bytes which were attempted to write in excess of quota and were rejected.
	StorageSizeThrottled *int64 `json:"StorageSizeThrottled,omitempty"`
	// Number of write units used per second
	WriteUnits *int64 `json:"WriteUnits,omitempty"`
	// Number of write units throttled per second. Units which was rejected with "resource exhausted error".
	WriteUnitsThrottled *int64 `json:"WriteUnitsThrottled,omitempty"`
}

func (o *QuotaUsageResponse) GetReadUnits() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadUnits
}

func (o *QuotaUsageResponse) GetReadUnitsThrottled() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadUnitsThrottled
}

func (o *QuotaUsageResponse) GetStorageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.StorageSize
}

func (o *QuotaUsageResponse) GetStorageSizeThrottled() *int64 {
	if o == nil {
		return nil
	}
	return o.StorageSizeThrottled
}

func (o *QuotaUsageResponse) GetWriteUnits() *int64 {
	if o == nil {
		return nil
	}
	return o.WriteUnits
}

func (o *QuotaUsageResponse) GetWriteUnitsThrottled() *int64 {
	if o == nil {
		return nil
	}
	return o.WriteUnitsThrottled
}

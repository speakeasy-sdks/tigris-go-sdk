package shared

type ListCachesResponse struct {
	Caches []CacheMetadata `json:"caches,omitempty"`
}

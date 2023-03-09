package shared

import (
	"time"
)

// SearchHitMeta
// Contains metadata related to the search hit, has information about document created_at/updated_at as well.
type SearchHitMeta struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

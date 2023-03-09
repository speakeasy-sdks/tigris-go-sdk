package shared

import (
	"time"
)

// DocMeta
// Contains metadata related to the search hit, has information about document created_at/updated_at as well.
type DocMeta struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

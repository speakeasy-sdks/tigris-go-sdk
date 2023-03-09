package shared

import (
	"time"
)

// ResponseMetadata
// Has metadata related to the documents stored.
type ResponseMetadata struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

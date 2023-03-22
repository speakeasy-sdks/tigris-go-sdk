// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ResponseMetadata - Has metadata related to the documents stored.
type ResponseMetadata struct {
	// Time at which the document was inserted/replaced. Measured in nano-seconds since the Unix epoch.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Time at which the document was deleted. Measured in nano-seconds since the Unix epoch.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Time at which the document was updated. Measured in nano-seconds since the Unix epoch.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// InsertResponse - OK
type InsertResponse struct {
	// an array returns the value of the primary keys.
	Keys []string `json:"keys,omitempty"`
	// Has metadata related to the documents stored.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
	// An enum with value set as "inserted"
	Status *string `json:"status,omitempty"`
}

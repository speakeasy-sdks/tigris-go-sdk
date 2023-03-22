// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SearchHit struct {
	// Actual search document
	Data map[string]interface{} `json:"data,omitempty"`
	// Contains metadata related to the search hit, has information about document created_at/updated_at as well.
	Metadata *SearchHitMeta `json:"metadata,omitempty"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DescribeCollectionRequest struct {
	// Optionally specify a database branch name to perform operation on
	Branch *string `json:"branch,omitempty"`
	// Name of the collection.
	Collection *string `json:"collection,omitempty"`
	// Collection requests modifying options.
	Options map[string]interface{} `json:"options,omitempty"`
	// Project name whose db is under target to get description of its collection.
	Project *string `json:"project,omitempty"`
	// Return schema in the requested format. Format can be JSON, Go, TypeScript, Java. Default is JSON.
	SchemaFormat *string `json:"schema_format,omitempty"`
}

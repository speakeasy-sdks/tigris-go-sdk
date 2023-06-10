// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ImportRequestDocuments struct {
}

type ImportRequest struct {
	// The list of autogenerated fields of the collection
	Autogenerated []string `json:"autogenerated,omitempty"`
	// Optionally specify a database branch name to perform operation on
	Branch *string `json:"branch,omitempty"`
	// Allow to create collection if it doesn't exists
	CreateCollection *bool `json:"create_collection,omitempty"`
	// Array of documents to import. Each document is a JSON object.
	Documents []ImportRequestDocuments `json:"documents,omitempty"`
	// additional options for import requests.
	Options *ImportRequestOptions `json:"options,omitempty"`
	// List of fields which constitutes primary key of the collection If not specified and field with name 'id' is present, it's used as a primary key, further if inferred type is UUID, then it's set as autogenerated.
	PrimaryKey []string `json:"primary_key,omitempty"`
}

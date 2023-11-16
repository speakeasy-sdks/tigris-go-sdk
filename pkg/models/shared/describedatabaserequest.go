// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DescribeDatabaseRequest struct {
	// Optionally specify a database branch name to perform operation on
	Branch *string `json:"branch,omitempty"`
	// Project name whose db is under target to get description.
	Project *string `json:"project,omitempty"`
	// Return schema in the requested format. Format can be JSON, Go, TypeScript, Java. Default is JSON.
	SchemaFormat *string `json:"schema_format,omitempty"`
}

func (o *DescribeDatabaseRequest) GetBranch() *string {
	if o == nil {
		return nil
	}
	return o.Branch
}

func (o *DescribeDatabaseRequest) GetProject() *string {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *DescribeDatabaseRequest) GetSchemaFormat() *string {
	if o == nil {
		return nil
	}
	return o.SchemaFormat
}

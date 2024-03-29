// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IndexInfo struct {
	// Name of the index.
	Name *string `json:"name,omitempty"`
	// Schema of the index.
	Schema *string `json:"schema,omitempty"`
}

func (o *IndexInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IndexInfo) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

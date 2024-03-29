// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ImportRequestOptions - additional options for import requests.
type ImportRequestOptions struct {
	// Additional options to modify write requests.
	WriteOptions *WriteOptions `json:"write_options,omitempty"`
}

func (o *ImportRequestOptions) GetWriteOptions() *WriteOptions {
	if o == nil {
		return nil
	}
	return o.WriteOptions
}

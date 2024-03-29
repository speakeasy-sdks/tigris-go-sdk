// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InsertUserMetadataRequestValue struct {
}

// InsertUserMetadataRequest - Request insertion of user metadata
type InsertUserMetadataRequest struct {
	MetadataKey *string                         `json:"metadataKey,omitempty"`
	Value       *InsertUserMetadataRequestValue `json:"value,omitempty"`
}

func (o *InsertUserMetadataRequest) GetMetadataKey() *string {
	if o == nil {
		return nil
	}
	return o.MetadataKey
}

func (o *InsertUserMetadataRequest) GetValue() *InsertUserMetadataRequestValue {
	if o == nil {
		return nil
	}
	return o.Value
}

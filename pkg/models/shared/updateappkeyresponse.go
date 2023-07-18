// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateAppKeyResponse - Returns response for updating the app key description
type UpdateAppKeyResponse struct {
	// An user AppKey
	UpdatedAppKey *AppKey `json:"updated_app_key,omitempty"`
}

func (o *UpdateAppKeyResponse) GetUpdatedAppKey() *AppKey {
	if o == nil {
		return nil
	}
	return o.UpdatedAppKey
}

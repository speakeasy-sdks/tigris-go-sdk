// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ChannelMetadata struct {
	Channel *string `json:"channel,omitempty"`
}

func (o *ChannelMetadata) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}

package shared

type GetRTChannelsResponse struct {
	Channels []ChannelMetadata `json:"channels,omitempty"`
}

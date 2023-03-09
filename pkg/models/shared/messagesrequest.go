package shared

type MessagesRequest struct {
	Channel  *string   `json:"channel,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	Project  *string   `json:"project,omitempty"`
}

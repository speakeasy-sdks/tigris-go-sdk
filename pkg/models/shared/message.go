package shared

type Message struct {
	Data     *string `json:"data,omitempty"`
	ID       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Sequence *string `json:"sequence,omitempty"`
}

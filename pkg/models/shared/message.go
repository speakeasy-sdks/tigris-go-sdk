// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Message struct {
	Data *string `json:"data,omitempty"`
	// an optional id if idempotency is needed to ensure only a single time message is published during retries. If not specified then server will automatically add an id to message.
	ID       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Sequence *string `json:"sequence,omitempty"`
}

func (o *Message) GetData() *string {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Message) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Message) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Message) GetSequence() *string {
	if o == nil {
		return nil
	}
	return o.Sequence
}

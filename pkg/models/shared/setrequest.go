package shared

type SetRequest struct {
	Ex    *int64  `json:"ex,omitempty"`
	Nx    *bool   `json:"nx,omitempty"`
	Px    *int64  `json:"px,omitempty"`
	Value *string `json:"value,omitempty"`
	Xx    *bool   `json:"xx,omitempty"`
}

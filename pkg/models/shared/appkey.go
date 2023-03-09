package shared

// AppKey
// An user AppKey
type AppKey struct {
	CreatedAt   *int64  `json:"created_at,omitempty"`
	CreatedBy   *string `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Project     *string `json:"project,omitempty"`
	Secret      *string `json:"secret,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
	UpdatedBy   *string `json:"updated_by,omitempty"`
}

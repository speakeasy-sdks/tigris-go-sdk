package shared

// Collation
// A collation allows you to specify string comparison rules. Default is case-sensitive, to override it you can set this option to 'ci' that will apply to all the text fields in the filters.
type Collation struct {
	Case *string `json:"case,omitempty"`
}

package shared

type ListBranchesResponse struct {
	Branches []BranchInfo `json:"branches,omitempty"`
}

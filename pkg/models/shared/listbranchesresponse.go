// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListBranchesResponse - OK
type ListBranchesResponse struct {
	// List of all the branches in this database
	Branches []BranchInfo `json:"branches,omitempty"`
}

func (o *ListBranchesResponse) GetBranches() []BranchInfo {
	if o == nil {
		return nil
	}
	return o.Branches
}

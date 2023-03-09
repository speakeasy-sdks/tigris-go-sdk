package shared

type ListProjectsResponse struct {
	Projects []ProjectInfo `json:"projects,omitempty"`
}

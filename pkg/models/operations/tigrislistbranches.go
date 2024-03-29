// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListBranchesRequest struct {
	// List database branches in this project
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisListBranchesRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisListBranchesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ListBranchesResponse *shared.ListBranchesResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TigrisListBranchesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisListBranchesResponse) GetListBranchesResponse() *shared.ListBranchesResponse {
	if o == nil {
		return nil
	}
	return o.ListBranchesResponse
}

func (o *TigrisListBranchesResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisListBranchesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisListBranchesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

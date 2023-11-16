// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListCollectionsRequest struct {
	// Optionally specify a database branch name to perform operation on
	Branch *string `queryParam:"style=form,explode=true,name=branch"`
	// Project name whose db is under target to list collections.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisListCollectionsRequest) GetBranch() *string {
	if o == nil {
		return nil
	}
	return o.Branch
}

func (o *TigrisListCollectionsRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisListCollectionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ListCollectionsResponse *shared.ListCollectionsResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TigrisListCollectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisListCollectionsResponse) GetListCollectionsResponse() *shared.ListCollectionsResponse {
	if o == nil {
		return nil
	}
	return o.ListCollectionsResponse
}

func (o *TigrisListCollectionsResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisListCollectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisListCollectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

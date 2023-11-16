// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchCreateByIDRequest struct {
	CreateByIDRequest shared.CreateByIDRequest `request:"mediaType=application/json"`
	// document id.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// index name where to create document.
	Index string `pathParam:"style=simple,explode=false,name=index"`
	// Tigris project name.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *SearchCreateByIDRequest) GetCreateByIDRequest() shared.CreateByIDRequest {
	if o == nil {
		return shared.CreateByIDRequest{}
	}
	return o.CreateByIDRequest
}

func (o *SearchCreateByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SearchCreateByIDRequest) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *SearchCreateByIDRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type SearchCreateByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateByIDResponse *shared.CreateByIDResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SearchCreateByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchCreateByIDResponse) GetCreateByIDResponse() *shared.CreateByIDResponse {
	if o == nil {
		return nil
	}
	return o.CreateByIDResponse
}

func (o *SearchCreateByIDResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SearchCreateByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchCreateByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

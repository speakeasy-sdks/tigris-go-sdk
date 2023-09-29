// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchDeleteIndexRequest struct {
	DeleteIndexRequest shared.DeleteIndexRequest `request:"mediaType=application/json"`
	// index name.
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Tigris project name.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *SearchDeleteIndexRequest) GetDeleteIndexRequest() shared.DeleteIndexRequest {
	if o == nil {
		return shared.DeleteIndexRequest{}
	}
	return o.DeleteIndexRequest
}

func (o *SearchDeleteIndexRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SearchDeleteIndexRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type SearchDeleteIndexResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DeleteIndexResponse *shared.DeleteIndexResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SearchDeleteIndexResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchDeleteIndexResponse) GetDeleteIndexResponse() *shared.DeleteIndexResponse {
	if o == nil {
		return nil
	}
	return o.DeleteIndexResponse
}

func (o *SearchDeleteIndexResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SearchDeleteIndexResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchDeleteIndexResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

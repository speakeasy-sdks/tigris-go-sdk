// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchDeleteByQueryRequest struct {
	DeleteByQueryRequest shared.DeleteByQueryRequest `request:"mediaType=application/json"`
	// The index name of the documents that needs deletion.
	Index string `pathParam:"style=simple,explode=false,name=index"`
	// The project name.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *SearchDeleteByQueryRequest) GetDeleteByQueryRequest() shared.DeleteByQueryRequest {
	if o == nil {
		return shared.DeleteByQueryRequest{}
	}
	return o.DeleteByQueryRequest
}

func (o *SearchDeleteByQueryRequest) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *SearchDeleteByQueryRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type SearchDeleteByQueryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DeleteByQueryResponse *shared.DeleteByQueryResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SearchDeleteByQueryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchDeleteByQueryResponse) GetDeleteByQueryResponse() *shared.DeleteByQueryResponse {
	if o == nil {
		return nil
	}
	return o.DeleteByQueryResponse
}

func (o *SearchDeleteByQueryResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SearchDeleteByQueryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchDeleteByQueryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

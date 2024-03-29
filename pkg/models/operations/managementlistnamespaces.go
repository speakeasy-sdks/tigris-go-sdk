// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementListNamespacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ListNamespacesResponse *shared.ListNamespacesResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ManagementListNamespacesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ManagementListNamespacesResponse) GetListNamespacesResponse() *shared.ListNamespacesResponse {
	if o == nil {
		return nil
	}
	return o.ListNamespacesResponse
}

func (o *ManagementListNamespacesResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ManagementListNamespacesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ManagementListNamespacesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

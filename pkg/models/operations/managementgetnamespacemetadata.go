// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementGetNamespaceMetadataRequest struct {
	GetNamespaceMetadataRequest shared.GetNamespaceMetadataRequest `request:"mediaType=application/json"`
	MetadataKey                 string                             `pathParam:"style=simple,explode=false,name=metadataKey"`
}

func (o *ManagementGetNamespaceMetadataRequest) GetGetNamespaceMetadataRequest() shared.GetNamespaceMetadataRequest {
	if o == nil {
		return shared.GetNamespaceMetadataRequest{}
	}
	return o.GetNamespaceMetadataRequest
}

func (o *ManagementGetNamespaceMetadataRequest) GetMetadataKey() string {
	if o == nil {
		return ""
	}
	return o.MetadataKey
}

type ManagementGetNamespaceMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	GetNamespaceMetadataResponse *shared.GetNamespaceMetadataResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ManagementGetNamespaceMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ManagementGetNamespaceMetadataResponse) GetGetNamespaceMetadataResponse() *shared.GetNamespaceMetadataResponse {
	if o == nil {
		return nil
	}
	return o.GetNamespaceMetadataResponse
}

func (o *ManagementGetNamespaceMetadataResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ManagementGetNamespaceMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ManagementGetNamespaceMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

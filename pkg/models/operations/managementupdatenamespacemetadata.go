// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementUpdateNamespaceMetadataRequest struct {
	UpdateNamespaceMetadataRequest shared.UpdateNamespaceMetadataRequest `request:"mediaType=application/json"`
	MetadataKey                    string                                `pathParam:"style=simple,explode=false,name=metadataKey"`
}

func (o *ManagementUpdateNamespaceMetadataRequest) GetUpdateNamespaceMetadataRequest() shared.UpdateNamespaceMetadataRequest {
	if o == nil {
		return shared.UpdateNamespaceMetadataRequest{}
	}
	return o.UpdateNamespaceMetadataRequest
}

func (o *ManagementUpdateNamespaceMetadataRequest) GetMetadataKey() string {
	if o == nil {
		return ""
	}
	return o.MetadataKey
}

type ManagementUpdateNamespaceMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default error response
	Status *shared.Status
	// OK
	UpdateNamespaceMetadataResponse *shared.UpdateNamespaceMetadataResponse
}

func (o *ManagementUpdateNamespaceMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ManagementUpdateNamespaceMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ManagementUpdateNamespaceMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ManagementUpdateNamespaceMetadataResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ManagementUpdateNamespaceMetadataResponse) GetUpdateNamespaceMetadataResponse() *shared.UpdateNamespaceMetadataResponse {
	if o == nil {
		return nil
	}
	return o.UpdateNamespaceMetadataResponse
}

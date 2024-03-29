// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisCreateOrUpdateCollectionRequest struct {
	CreateOrUpdateCollectionRequest shared.CreateOrUpdateCollectionRequest `request:"mediaType=application/json"`
	// Collection name to create.
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	// Project name whose db is under target to create or update collection.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisCreateOrUpdateCollectionRequest) GetCreateOrUpdateCollectionRequest() shared.CreateOrUpdateCollectionRequest {
	if o == nil {
		return shared.CreateOrUpdateCollectionRequest{}
	}
	return o.CreateOrUpdateCollectionRequest
}

func (o *TigrisCreateOrUpdateCollectionRequest) GetCollection() string {
	if o == nil {
		return ""
	}
	return o.Collection
}

func (o *TigrisCreateOrUpdateCollectionRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisCreateOrUpdateCollectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateOrUpdateCollectionResponse *shared.CreateOrUpdateCollectionResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TigrisCreateOrUpdateCollectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisCreateOrUpdateCollectionResponse) GetCreateOrUpdateCollectionResponse() *shared.CreateOrUpdateCollectionResponse {
	if o == nil {
		return nil
	}
	return o.CreateOrUpdateCollectionResponse
}

func (o *TigrisCreateOrUpdateCollectionResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisCreateOrUpdateCollectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisCreateOrUpdateCollectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

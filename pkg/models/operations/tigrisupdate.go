// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisUpdateRequest struct {
	UpdateRequest shared.UpdateRequest `request:"mediaType=application/json"`
	// Collection name where to update documents
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	// Project name whose db is under target  to update documents
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default error response
	Status *shared.Status
	// OK
	UpdateResponse *shared.UpdateResponse
}

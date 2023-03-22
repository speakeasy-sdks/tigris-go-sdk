// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisInsertRequest struct {
	InsertRequest shared.InsertRequest `request:"mediaType=application/json"`
	// Collection name where to insert documents.
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	// Project name whose db is under target to insert documents.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisInsertResponse struct {
	ContentType string
	// OK
	InsertResponse *shared.InsertResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}

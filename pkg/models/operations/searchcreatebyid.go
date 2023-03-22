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

type SearchCreateByIDResponse struct {
	ContentType string
	// OK
	CreateByIDResponse *shared.CreateByIDResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
